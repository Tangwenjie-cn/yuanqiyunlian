/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package controller

import (
	"gin/config"
	"gin/model"
	"gin/service/pay"
	"gin/service/set"
	"gin/service/wechat"
	"gin/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func CashList(c *gin.Context) {
	data := struct {
		Keyword string `form:"keyword" json:"keyword"`
		Status  int    `form:"status" json:"status"`
		Page    int    `form:"page" json:"page" binding:"required"`
		Limit   int    `form:"limit" json:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var total int64
	query := config.DB.Model(&model.Cash{})
	if data.Keyword != "" {
		query = query.Where("cash.uid = ?", strings.TrimSpace(data.Keyword))
	}
	if data.Status != 0 {
		if data.Status == 2 {
			data.Status = 0
		}
		query = query.Where("cash.status = ?", data.Status)
	}
	query.Count(&total)
	type cashInfo struct {
		model.Cash
		Nickname string `json:"nickname"`
	}
	var list []cashInfo
	if err := query.Joins("join user on user.id=cash.uid").Select("cash.*", "user.nickname").
		Limit(data.Limit).Offset((data.Page - 1) * data.Limit).Order("cash.id desc").Find(&list).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取成功", map[string]any{
		"total": total,
		"data":  list,
	})
}
func CashCheck(c *gin.Context) {
	data := struct {
		Id     int    `form:"id" json:"id" binding:"required"`
		Status int    `form:"status" json:"status" binding:"required"`
		Remark string `form:"remark" json:"remark"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	tx := config.DB.Begin()
	var cash model.Cash
	if err := tx.Where("id = ?", data.Id).First(&cash).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if cash.Status != 0 {
		util.Error(c, 1, "该提现记录已审核")
		return
	}

	switch data.Status {
	case 1:
		cash.Status = 1
	case 2:
		//处理商家转账
		if cash.Type == 2 {
			cash.Status = 1
			thirdOrder, err := pay.AliapyTransfer(cash.Actual, "提现到账", cash.Realname, cash.Account)
			if err != nil {
				tx.Rollback()
				util.Error(c, 1, err.Error())
				return
			}
			cash.ThirdOrder = thirdOrder
		}
		if cash.Type == 1 {
			cash.Status = 2
			var userBind model.UserBind
			switch cash.Platform {
			case "mp":
				if err := config.DB.Where("uid = ? and type=1", cash.Uid).First(&userBind).Error; err != nil {
					tx.Rollback()
					util.Error(c, 1, "用户未绑定微信")
					return
				}
			case "mini":
				if err := config.DB.Where("uid = ? and type=2", cash.Uid).First(&userBind).Error; err != nil {
					tx.Rollback()
					util.Error(c, 1, "用户未绑定微信")
					return
				}
			default:
				tx.Rollback()
				util.Error(c, 1, "参数错误")
				return
			}
			transferInfo, err := pay.WxpayTransfer(cash.Actual*100, "提现到账", userBind.Openid, cash.Platform)
			if err != nil {
				tx.Rollback()
				util.Error(c, 1, err.Error())
				return
			}
			cash.ThirdOrder = transferInfo["transfer_bill_no"]
			cash.PackageInfo = transferInfo["package_info"]
		}
	case -1:
		cash.Status = -1
		cash.Remark = data.Remark
		err := model.UserBill{}.Add(cash.Uid, 1, 3, cash.Money, "提现失败退回", int(cash.ID), tx)
		if err != nil {
			tx.Rollback()
			util.Error(c, 1, err.Error())
			return
		}
	default:
		util.Error(c, 1, "参数错误")
		return
	}
	if err := tx.Save(&cash).Error; err != nil {
		tx.Rollback()
		util.Error(c, 1, err.Error())
		return
	}
	tx.Commit()
	sendCashTemplateMessage(&cash)
	util.Success(c, "操作成功", nil)
}
func sendCashTemplateMessage(cash *model.Cash) {
	templateId, _ := set.GetSet("wx_mini_cash_templateId", true)
	if templateId == "" {
		return
	}
	var userBind model.UserBind
	if err := config.DB.Where("uid = ? and type=2", cash.Uid).First(&userBind).Error; err != nil {
		return
	}
	var statusText string
	remark := "无"
	switch cash.Status {
	case 1:
		statusText = "成功"
	case -1:
		statusText = "拒绝"
		remark = cash.Remark
	case 2:
		statusText = "待确认"
		remark = "请在24小时内前往提现列表页确认收款"
	default:
		return
	}
	data := map[string]any{
		"amount1": map[string]any{
			"value": cash.Money,
		},
		"phrase2": map[string]any{
			"value": statusText,
		},
		"time4": map[string]any{
			"value": cash.Utime.Format("2006-01-02 15:04:05"),
		},
		"thing6": map[string]any{
			"value": remark,
		},
	}
	wechat.SendTemplateMessage(userBind.Openid, templateId, "pages/cash/list", data)
}
