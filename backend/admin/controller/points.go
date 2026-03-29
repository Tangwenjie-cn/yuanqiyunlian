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

func PointsLog(c *gin.Context) {
	data := struct {
		Page    int    `form:"page" json:"page" binding:"required"`
		Limit   int    `form:"limit" json:"limit" binding:"required"`
		Keyword string `form:"keyword" json:"keyword"`
		Date    string `form:"date" json:"date"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var total int64
	query := config.DB.Model(&model.UserPoints{}).Joins("join user on user.id=user_points.uid")
	if data.Keyword != "" {
		query = query.Where("user_points.uid = ?", data.Keyword)
	}
	if data.Date != "" {
		dateArr := strings.Split(data.Date, ",")
		query = query.Where("user_points.ctime between ? and ?", dateArr[0], dateArr[1])
	}
	query.Count(&total)
	query = query.Select("user_points.*, user.nickname, user.avatar").Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Order("user_points.id desc")
	type UserPointsInfo struct {
		model.UserPoints
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
	}
	var list []UserPointsInfo
	query.Find(&list)
	res := map[string]any{
		"data":  list,
		"total": total,
	}
	util.Success(c, "获取列表成功", res)
}
