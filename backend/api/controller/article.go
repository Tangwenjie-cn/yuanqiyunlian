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
)

func ArticleList(c *gin.Context) {
	data := struct {
		SortId int `form:"sort_id" json:"sort_id" binding:"required"`
		Page   int `form:"page" json:"page" binding:"required"`
		Limit  int `form:"limit" json:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	list := []model.Article{}
	if err := config.DB.Where("sid = ? and status = 1", data.SortId).
		Select("id", "title", "utime", "thumb").Order("utime desc").
		Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Find(&list).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	for k, v := range list {
		v.UtimeText = util.TimeAgo(v.Utime)
		list[k] = v
	}
	util.Success(c, "success", list)
}
func ArticleInfo(c *gin.Context) {
	id := c.Query("id")
	info := model.Article{}
	if err := config.DB.Where("id = ? and status = 1", id).First(&info).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "success", info)
}
