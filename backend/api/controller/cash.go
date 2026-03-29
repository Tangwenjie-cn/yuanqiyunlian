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
	"gin/service/set"
	"gin/util"
	"math"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CashAdd(c *gin.Context) {
	uid, ok := c.MustGet("uid").(int)
	if !ok {
		util.Error(c, 1, "用户不存在")
		return
	}
	data := struct {
		Type     int     `json:"type" form:"type" binding:"required"`
		Money    float64 `json:"money" form:"money" binding:"required"`
		Account  string  `json:"account" form:"account"`
		Realname string  `json:"realname" form:"realname"`
		Qrcode   string  `json:"qrcode" form:"qrcode"`
		Bank     string  `json:"bank" form:"bank"`
		Platform string  `json:"platform" form:"platform" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var cash model.Cash
	if data.Type == 1 { //微信提现
		cash_wxpay_type, _ := set.GetSet("cash_wxpay_type", true)
		if cash_wxpay_type == "2" {
			switch data.Platform {
			case "h5":
				userAgent := c.Request.Header.Get("User-Agent")
				if strings.Contains(userAgent, "MicroMessenger") { //微信浏览器
					cash.Platform = "mp"
				} else {
					util.Error(c, 1, "请在微信浏览器中操作")
					return
				}
			case "weixin":
				cash.Platform = "mini"
			default:
				util.Error(c, 1, "微信提现请在微信中操作")
				return
			}
		}

	}

	var user model.User
	if err := config.DB.Where("id=?", uid).First(&user).Error; err != nil {
		util.Error(c, 1, "用户不存在")
		return
	}
	if user.Balance < data.Money {
		util.Error(c, 1, "余额不足")
		return
	}
	cashMinMoney, err := set.GetSet("cash_min_money", true)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	cashMinMoneyFloat, _ := strconv.ParseFloat(cashMinMoney, 64)
	if data.Money < cashMinMoneyFloat {
		util.Error(c, 1, "提现金额不能小于"+cashMinMoney)
		return
	}
	cashMaxMoney, err := set.GetSet("cash_max_money", true)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	cashMaxMoneyFloat, _ := strconv.ParseFloat(cashMaxMoney, 64)
	if data.Money > cashMaxMoneyFloat {
		util.Error(c, 1, "提现金额不能大于"+cashMaxMoney)
		return
	}
	cashCycle, err := set.GetSet("cash_cycle", true)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	cashCycleInt, _ := strconv.Atoi(cashCycle)
	if cashCycleInt > 0 {
		var cashs []model.Cash
		cashNums, err := set.GetSet("cash_nums", true)
		if err != nil {
			util.Error(c, 1, err.Error())
			return
		}
		cashNumsInt, _ := strconv.Atoi(cashNums)

		switch cashCycleInt {
		case 1:
			startOfDay, endOfDay := util.GetTodayStartEnd()
			if err := config.DB.Where("uid=? and ctime>=? and ctime<=? and status>=0", uid, startOfDay, endOfDay).Find(&cashs).Error; err != nil {
				util.Error(c, 1, err.Error())
				return
			}
			if len(cashs) >= cashNumsInt {
				util.Error(c, 1, "每天只能提现"+cashNums+"次")
				return
			}
		case 2:
			startOfWeek, endOfWeek := util.GetWeekStartEnd()
			if err := config.DB.Where("uid=? and ctime>=? and ctime<=? and status>=0", uid, startOfWeek, endOfWeek).Find(&cashs).Error; err != nil {
				util.Error(c, 1, err.Error())
				return
			}
			if len(cashs) >= cashNumsInt {
				util.Error(c, 1, "每周只能提现"+cashNums+"次")
				return
			}
		case 3:
			startOfMonth, endOfMonth := util.GetMonthStartEnd()
			if err := config.DB.Where("uid=? and ctime>=? and ctime<=? and status>=0", uid, startOfMonth, endOfMonth).Find(&cashs).Error; err != nil {
				util.Error(c, 1, err.Error())
				return
			}
			if len(cashs) >= cashNumsInt {
				util.Error(c, 1, "每月只能提现"+cashNums+"次")
				return
			}
		}
	}
	// 开启事务
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	cashFee, err := set.GetSet("cash_fee", true)
	if err != nil {
		tx.Rollback()
		util.Error(c, 1, err.Error())
		return
	}
	cashFeeFloat, _ := strconv.ParseFloat(cashFee, 64)
	fee := math.Round(data.Money*(cashFeeFloat/100)*100) / 100
	cash.Account = data.Account
	cash.Realname = data.Realname
	cash.Qrcode = data.Qrcode
	cash.Bank = data.Bank
	cash.Type = data.Type
	cash.Money = data.Money
	cash.Fee = fee
	cash.Actual = data.Money - fee
	cash.Uid = uid

	if err := tx.Create(&cash).Error; err != nil { // 提交事务
		tx.Rollback()
		util.Error(c, 1, err.Error())
		return
	}
	err = model.UserBill{}.Add(uid, 2, 3, cash.Money, "提现", int(cash.ID), tx)
	if err != nil {
		tx.Rollback()
		util.Error(c, 1, err.Error())
		return
	}
	tx.Commit() // 提交事务
	util.Success(c, "申请成功", nil)
}
func CashList(c *gin.Context) {
	data := struct {
		Page  int `json:"page" form:"page" binding:"required"`
		Limit int `json:"limit" form:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	uid, _ := c.MustGet("uid").(int)
	var cashs []model.Cash
	if err := config.DB.Where("uid=?", uid).Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Order("id desc").Find(&cashs).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "成功", cashs)
}
func GetWxTransfer(c *gin.Context) {
	uid, _ := c.MustGet("uid").(int)
	var cash model.Cash
	if err := config.DB.Where("uid=? and status=2 and type=1", uid).Find(&cash).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if cash.ID == 0 {
		util.Success(c, "ok", nil)
		return
	}
	mchId, _ := set.GetSet("wxpay_mchId", true)
	if mchId == "" {
		util.Error(c, 1, "wxpay_mchId not found")
		return
	}
	var appId string
	if cash.Platform == "mp" {
		appidMp, _ := set.GetSet("wx_mp_appid", true)
		if appidMp == "" {
			util.Error(c, 1, "wx_mp_appid not found")
			return
		}
		appId = appidMp
	}
	if cash.Platform == "mini" {
		appidMini, _ := set.GetSet("wx_mini_appid", true)
		if appidMini == "" {
			util.Error(c, 1, "wx_mini_appid not found")
			return
		}
		appId = appidMini
	}
	data := map[string]string{
		"mchId":   mchId,
		"appId":   appId,
		"package": cash.PackageInfo,
	}
	util.Success(c, "ok", data)
}
