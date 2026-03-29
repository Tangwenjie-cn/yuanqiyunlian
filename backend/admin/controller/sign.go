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

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func SignAdd(c *gin.Context) {
	data := struct {
		Type   int `json:"type" form:"type" binding:"required"`
		Days   int `json:"days" form:"days" binding:"required"`
		Points int `json:"points" form:"points" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var signReward model.SignReward
	if err := mapstructure.Decode(data, &signReward); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Create(&signReward).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "ok", nil)
}
func SignEdit(c *gin.Context) {
	data := struct {
		Id     int `json:"id" form:"id" binding:"required"`
		Type   int `json:"type" form:"type" binding:"required"`
		Days   int `json:"days" form:"days" binding:"required"`
		Points int `json:"points" form:"points" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var signReward model.SignReward
	if err := mapstructure.Decode(data, &signReward); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Where("id=?", data.Id).Updates(&signReward).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "ok", nil)
}
func SignDel(c *gin.Context) {
	data := struct {
		Id int `json:"id" form:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Delete(&model.SignReward{}, data.Id).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "ok", nil)
}
func SignList(c *gin.Context) {
	data := struct {
		Page  int `form:"page" json:"page" binding:"required"`
		Limit int `form:"limit" json:"limit" binding:"required"`
		Type  int `form:"type" json:"type"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var total int64
	query := config.DB.Model(&model.SignReward{})
	if data.Type != 0 {
		query = query.Where("type = ?", data.Type)
	}
	query.Count(&total)
	query = query.Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Order("id desc")
	var list []model.SignReward
	query.Find(&list)
	res := map[string]any{
		"data":  list,
		"total": total,
	}
	util.Success(c, "获取列表成功", res)
}
