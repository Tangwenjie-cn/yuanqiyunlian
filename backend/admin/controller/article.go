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

// Article 函数用于获取文章列表
func Article(c *gin.Context) {
	data := struct {
		Keyword string `form:"keyword" json:"keyword"`
		Status  uint8  `form:"status" json:"status"`
		Page    int    `form:"page" json:"page" binding:"required"`
		Limit   int    `form:"limit" json:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var total int64
	query := config.DB.Model(&model.Article{}).Preload("ArticleSort")
	if data.Keyword != "" {
		query = query.Where("title like ?", "%"+data.Keyword+"%")
	}
	if data.Status != 0 {
		query = query.Where("status = ?", data.Status)
	}
	query.Count(&total)
	query = query.Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Order("id desc")
	var list []model.Article
	query.Find(&list)
	res := map[string]any{
		"data":  list,
		"total": total,
	}
	util.Success(c, "获取列表成功", res)
}
func AddArticle(c *gin.Context) {
	data := struct {
		Title   string `form:"title" json:"title" binding:"required"`
		Sid     int16  `form:"sid" json:"sid" binding:"required"`
		Thumb   string `form:"thumb" json:"thumb"`
		Status  int8   `form:"status" json:"status" binding:"required"`
		Content string `form:"content" json:"content" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var article model.Article
	if err := mapstructure.Decode(data, &article); err != nil {
		util.Error(c, 1, err.Error())
		return
	}

	if len(data.Thumb) > 0 {
		article.Thumb = data.Thumb
	}
	res := config.DB.Create(&article)
	if res.Error != nil {
		util.Error(c, 1, res.Error.Error())
		return
	}
	util.Success(c, "添加成功", nil)
}
func EditArticle(c *gin.Context) {
	data := struct {
		Id      uint   `form:"id" json:"id" binding:"required"`
		Title   string `form:"title" json:"title" binding:"required"`
		Sid     int16  `form:"sid" json:"sid" binding:"required"`
		Thumb   string `form:"thumb" json:"thumb"`
		Status  int8   `form:"status" json:"status" binding:"required"`
		Content string `form:"content" json:"content" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var article map[string]any
	if err := mapstructure.Decode(data, &article); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	res := config.DB.Model(&model.Article{}).Where("id = ?", data.Id).Updates(&article)
	if res.Error != nil {
		util.Error(c, 1, res.Error.Error())
		return
	}
	util.Success(c, "修改成功", nil)
}
func DelArticle(c *gin.Context) {
	data := struct {
		Id uint `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if data.Id < 999 {
		util.Error(c, 1, "不可删除固有文章")
		return
	}
	res := config.DB.Delete(&model.Article{}, data.Id)
	if res.Error != nil {
		util.Error(c, 1, res.Error.Error())
	}
	util.Success(c, "删除成功", nil)
}
func ArticleSort(c *gin.Context) {
	data := struct {
		Page  int `form:"page" json:"page" binding:"required"`
		Limit int `form:"limit" json:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var total int64
	query := config.DB.Model(&model.ArticleSort{})
	query.Count(&total)
	query = query.Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Order("id desc")
	var list []model.ArticleSort
	query.Find(&list)
	res := map[string]any{
		"data":  list,
		"total": total,
	}
	util.Success(c, "获取列表成功", res)
}
func AddArticleSort(c *gin.Context) {
	data := struct {
		Name   string           `form:"name" json:"name" binding:"required"`
		Thumb  []map[string]any `form:"thumb" json:"thumb"`
		Status int8             `form:"status" json:"status" binding:"required"`
		Sort   uint             `form:"sort" json:"sort"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	ArticleSort := model.ArticleSort{
		Name:   data.Name,
		Status: data.Status,
		Sort:   data.Sort,
	}
	if len(data.Thumb) > 0 {
		ArticleSort.Thumb = data.Thumb[0]["url"].(string)
	}
	res := config.DB.Create(&ArticleSort)
	if res.Error != nil {
		util.Error(c, 1, res.Error.Error())
		return
	}
	util.Success(c, "添加成功", nil)
}
func EditArticleSort(c *gin.Context) {
	data := struct {
		Id     uint             `form:"id" json:"id" binding:"required"`
		Name   string           `form:"name" json:"name" binding:"required"`
		Thumb  []map[string]any `form:"thumb" json:"thumb"`
		Status int8             `form:"status" json:"status" binding:"required"`
		Sort   uint             `form:"sort" json:"sort"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	ArticleSort := model.ArticleSort{
		Name:   data.Name,
		Status: data.Status,
		Sort:   data.Sort,
	}
	if len(data.Thumb) > 0 {
		ArticleSort.Thumb = data.Thumb[0]["url"].(string)
	}
	res := config.DB.Model(&model.ArticleSort{}).Where("id = ?", data.Id).Updates(&ArticleSort)
	if res.Error != nil {
		util.Error(c, 1, res.Error.Error())
		return
	}
	util.Success(c, "修改成功", nil)
}
func DelArticleSort(c *gin.Context) {
	data := struct {
		Id uint `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	res := config.DB.Delete(&model.ArticleSort{}, data.Id)
	if res.Error != nil {
		util.Error(c, 1, res.Error.Error())
		return
	}
	util.Success(c, "删除成功", nil)
}
