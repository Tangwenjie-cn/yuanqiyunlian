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
	"github.com/google/uuid"
)

func KamiList(c *gin.Context) {
	data := struct {
		Keyword string `form:"keyword" json:"keyword"`
		Status  int    `form:"status" json:"status"`
		Page    int    `form:"page" json:"page"`
		Limit   int    `form:"limit" json:"limit"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var list []model.Kami
	var total int64
	query := config.DB.Model(&model.Kami{})
	if data.Keyword != "" {
		query = query.Where("key = ?", data.Keyword)
	}
	if data.Status != 0 {
		query.Where("status = ?", data.Status)
	}
	query.Count(&total)
	query.Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Order("id desc")
	query.Find(&list)
	util.Success(c, "成功", map[string]any{
		"data":  list,
		"total": total,
	})
}
func AddKami(c *gin.Context) {
	data := struct {
		Price float64 `json:"price" form:"price" binding:"required"`
		Nums  int     `json:"nums" form:"nums" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var kami []model.Kami

	for i := 0; i < data.Nums; i++ {
		key, _ := uuid.NewRandom()
		kami = append(kami, model.Kami{
			Price:  data.Price,
			Status: -1,
			Key:    key.String(),
		})
	}
	if err := config.DB.Create(&kami).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "添加成功", nil)
}
func EditKami(c *gin.Context) {
	data := struct {
		Id     int     `json:"id" form:"id" binding:"required"`
		Status int     `json:"status" form:"status" binding:"required"`
		Price  float64 `json:"price" form:"price" binding:"required"`
		Key    string  `json:"key" form:"key" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Model(&model.Kami{}).Where("id = ?", data.Id).Updates(model.Kami{
		Status: data.Status,
		Price:  data.Price,
		Key:    data.Key,
	}).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "修改成功", nil)
}
func DelKami(c *gin.Context) {
	data := struct {
		Id int `json:"id" form:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Where("id = ?", data.Id).Delete(&model.Kami{}).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "删除成功", nil)
}
