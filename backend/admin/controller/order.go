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
	"gin/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func OrderList(c *gin.Context) {
	data := struct {
		Keyword string `form:"keyword" json:"keyword"`
		Status  int    `form:"status" json:"status"`
		Type    int    `form:"type" json:"type"`
		PayType int    `form:"pay_type" json:"pay_type"`
		Page    int    `form:"page" json:"page" binding:"required"`
		Limit   int    `form:"limit" json:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var total int64
	query := config.DB.Model(&model.Order{})
	if data.Keyword != "" {
		query = query.Where("(uid = ? or order_no = ?)", strings.TrimSpace(data.Keyword), strings.TrimSpace(data.Keyword))
	}
	if data.Status != 0 {
		if data.Status == 2 {
			data.Status = 0
		}
		query = query.Where("status = ?", data.Status)
	}
	if data.Type != 0 {
		query = query.Where("type = ?", data.Type)
	}
	if data.PayType != 0 {
		query = query.Where("pay_type = ?", data.PayType)
	}
	query.Count(&total)
	list := []struct {
		model.Order
		Title string `json:"title" gorm:"-"`
	}{}
	if err := query.Limit(data.Limit).Offset((data.Page - 1) * data.Limit).Order("id desc").Find(&list).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	for k, v := range list {
		if v.Type == 1 {
			var goods model.Goods
			if err := config.DB.Select("title").Where("id = ?", v.Gid).Find(&goods).Error; err != nil {
				util.Error(c, 1, err.Error())
				return
			}
			list[k].Title = goods.Title
		}
		if v.Type == 2 {
			var vip model.Svip
			if err := config.DB.Select("title").Where("id = ?", v.Vid).Find(&vip).Error; err != nil {
				util.Error(c, 1, err.Error())
				return
			}
			list[k].Title = vip.Title
		}
	}
	util.Success(c, "success", map[string]any{
		"data":  list,
		"total": total,
	})
}
