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
	"encoding/json"
	"gin/config"
	"gin/model"
	"gin/util"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

func ThemeIndex(c *gin.Context) {
	var themes []model.Theme
	config.DB.Where("type = ?", 1).Find(&themes)
	util.Success(c, "success", themes)
}
func ThemeAddIndex(c *gin.Context) {
	data := struct {
		Title   string `form:"title" json:"title" binding:"required"`
		Content []any  `form:"content" json:"content" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	content, err := json.Marshal(data.Content)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Create(&model.Theme{
		Title:   data.Title,
		Content: datatypes.JSON(content),
		Type:    1,
		Status:  -1,
	}).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "success", nil)
}
func ThemeEdit(c *gin.Context) {
	data := struct {
		Id      int    `form:"id" json:"id" binding:"required"`
		Title   string `form:"title" json:"title" binding:"required"`
		Content []any  `form:"content" json:"content" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	content, err := json.Marshal(data.Content)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Model(&model.Theme{}).Where("id = ?", data.Id).Updates(model.Theme{
		Title:   data.Title,
		Content: datatypes.JSON(content),
	}).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "success", nil)
}
func ThemeDelIndex(c *gin.Context) {
	data := struct {
		Id int `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Delete(&model.Theme{}, data.Id).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "success", nil)
}
func ThemeSetIndex(c *gin.Context) {
	data := struct {
		Id int `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Model(&model.Theme{}).Where("type = ?", 1).Update("status", -1).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Model(&model.Theme{}).Where("id = ?", data.Id).Update("status", 1).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "success", nil)
}
func ThemeGet(c *gin.Context) {
	typeId := c.Query("type")
	var theme model.Theme
	config.DB.Where("type = ?", typeId).Find(&theme)
	util.Success(c, "success", theme)
}
func PagesList(c *gin.Context) {
	data := struct {
		Page  int `form:"page" json:"page" binding:"required"`
		Limit int `form:"limit" json:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var total int64
	query := config.DB.Model(&model.Pages{})
	query.Count(&total)
	query = query.Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Order("id desc")
	var list []model.Pages
	query.Find(&list)
	res := map[string]any{
		"data":  list,
		"total": total,
	}
	util.Success(c, "获取列表成功", res)
}
func PagesAdd(c *gin.Context) {
	data := struct {
		Title  string `form:"title" json:"title" binding:"required"`
		Path   string `form:"path" json:"path" binding:"required"`
		Remark string `form:"remark" json:"remark"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Create(&model.Pages{
		Title:  data.Title,
		Path:   data.Path,
		Remark: data.Remark,
	}).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "success", nil)
}
func PagesEdit(c *gin.Context) {
	data := struct {
		Id     int    `form:"id" json:"id" binding:"required"`
		Title  string `form:"title" json:"title" binding:"required"`
		Path   string `form:"path" json:"path" binding:"required"`
		Remark string `form:"remark" json:"remark"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	upData := map[string]any{
		"title":  data.Title,
		"path":   data.Path,
		"remark": data.Remark,
	}
	if err := config.DB.Model(&model.Pages{}).Where("id = ?", data.Id).Updates(&upData).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "success", nil)
}
func PagesDel(c *gin.Context) {
	data := struct {
		Id int `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Delete(&model.Pages{}, data.Id).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "success", nil)
}
