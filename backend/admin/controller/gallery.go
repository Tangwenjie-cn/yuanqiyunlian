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

func ImgList(c *gin.Context) {
	data := struct {
		Sid    uint8 `form:"sid" json:"sid"`
		Status uint8 `form:"status" json:"status"`
		Page   int   `form:"page" json:"page" binding:"required"`
		Limit  int   `form:"limit" json:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var total int64
	query := config.DB.Model(&model.Img{}).Preload("ImgSort")
	if data.Sid != 0 {
		query = query.Where("sid = ?", data.Sid)
	}
	if data.Status != 0 {
		query = query.Where("status = ?", data.Status)
	}
	query.Count(&total)
	query = query.Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Order("id desc")
	var list []model.Img
	query.Find(&list)
	res := map[string]any{
		"data":  list,
		"total": total,
	}
	util.Success(c, "获取列表成功", res)
}
func AddImg(c *gin.Context) {
	data := struct {
		Sid    int16  `form:"sid" json:"sid" binding:"required"`
		Sort   uint   `form:"sort" json:"sort"`
		Status int8   `form:"status" json:"status" binding:"required"`
		Thumb  string `form:"thumb" json:"thumb"`
	}{}
	var err error
	if err = c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var img model.Img
	if err := mapstructure.Decode(data, &img); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Create(&img).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "添加成功", nil)
}
func EditImg(c *gin.Context) {
	data := struct {
		Id     uint   `form:"id" json:"id" binding:"required"`
		Sid    int16  `form:"sid" json:"sid" binding:"required"`
		Sort   uint   `form:"sort" json:"sort"`
		Status int8   `form:"status" json:"status" binding:"required"`
		Thumb  string `form:"thumb" json:"thumb"`
	}{}

	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var img model.Img
	if err := mapstructure.Decode(data, &img); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Model(&model.Img{}).Where("id = ?", data.Id).Updates(&img).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "修改成功", nil)
}
func DelImg(c *gin.Context) {
	data := struct {
		Id uint `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Where("id = ?", data.Id).Delete(&model.Img{}).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "删除成功", nil)
}
func ImgSort(c *gin.Context) {
	data := struct {
		Page  int `form:"page" json:"page" binding:"required"`
		Limit int `form:"limit" json:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var total int64
	query := config.DB.Model(&model.ImgSort{})
	query.Count(&total)
	query = query.Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Order("id desc")
	var list []model.ImgSort
	query.Find(&list)
	res := map[string]any{
		"data":  list,
		"total": total,
	}
	util.Success(c, "获取列表成功", res)
}
func AddImgSort(c *gin.Context) {
	data := struct {
		Name   string `form:"name" json:"name" binding:"required"`
		Status int8   `form:"status" json:"status" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	img := model.ImgSort{
		Name:   data.Name,
		Status: data.Status,
	}
	if err := config.DB.Create(&img).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "添加成功", nil)
}
func EditImgSort(c *gin.Context) {
	data := struct {
		Id     uint   `form:"id" json:"id" binding:"required"`
		Name   string `form:"name" json:"name" binding:"required"`
		Status int8   `form:"status" json:"status" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	img := model.ImgSort{
		Name:   data.Name,
		Status: data.Status,
	}
	if err := config.DB.Model(&model.ImgSort{}).Where("id = ?", data.Id).Updates(&img).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "修改成功", nil)
}
func DelImgSort(c *gin.Context) {
	data := struct {
		Id uint `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Where("id = ?", data.Id).Delete(&model.ImgSort{}).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "删除成功", nil)
}
