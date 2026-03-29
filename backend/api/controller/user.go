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
	"context"
	"gin/config"
	"gin/model"
	"gin/service/image"
	"gin/util"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Collect(c *gin.Context) {
	data := struct {
		Gid int `form:"gid" json:"gid" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	uid, _ := c.MustGet("uid").(int)
	userCollect := model.UserCollect{}
	if err := config.DB.Where("gid = ? and uid = ?", data.Gid, uid).Find(&userCollect).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if userCollect.ID > 0 {
		config.DB.Delete(&userCollect)
		util.Success(c, "取消收藏成功", nil)
		return
	} else {
		config.DB.Create(&model.UserCollect{
			Gid: data.Gid,
			Uid: uid,
		})
		util.Success(c, "收藏成功", nil)
		return
	}
}
func UserVipList(c *gin.Context) {
	uid, _ := c.MustGet("uid").(int)
	list := []struct {
		model.UserSvip
		Title string `json:"title" gorm:"column:title"`
	}{}
	if err := config.DB.Model(&model.UserSvip{}).
		Select("user_svip.*", "svip.title").
		Joins("inner join svip on svip.id=user_svip.vid").
		Where("user_svip.uid = ?", uid).
		Order("user_svip.expire_time desc").Find(&list).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取成功", list)
}
func UpdateUserInfo(c *gin.Context) {
	uid, _ := c.MustGet("uid").(int)
	data := struct {
		Avatar   string `form:"avatar" json:"avatar" `
		Nickname string `form:"nickname" json:"nickname" `
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	user := model.User{}
	if data.Avatar != "" {
		user.Avatar = data.Avatar
	}
	if data.Nickname != "" {
		user.Nickname = data.Nickname
	}
	if err := config.DB.Where("id = ?", uid).Updates(&user).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "保存成功", nil)
}
func GetUserInfo(c *gin.Context) {
	uid, _ := c.MustGet("uid").(int)
	user := model.User{}
	if err := config.DB.Where("id = ?", uid).Find(&user).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	userBind := model.UserBind{}
	config.DB.Where("uid = ? and type = 1", uid).Find(&userBind)
	if userBind.ID > 0 {
		user.Phone = userBind.Phone
	}
	util.Success(c, "获取成功", user)
}
func UserGoods(c *gin.Context) {
	data := struct {
		Page  int `json:"page" form:"page" binding:"required"`
		Limit int `json:"limit" form:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	uid, _ := c.MustGet("uid").(int)
	list := []struct {
		model.UserGoods
		Title string `json:"title" gorm:"column:title"`
		Thumb string `json:"thumb" gorm:"column:thumb"`
	}{}
	if err := config.DB.Model(&model.UserGoods{}).Where("user_goods.uid = ?", uid).Offset((data.Page-1)*data.Limit).
		Select("user_goods.*", "goods.title", "goods.thumb").
		Joins("inner join goods on goods.id=user_goods.gid").
		Limit(data.Limit).Order("id desc").Find(&list).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取成功", list)
}
func UserBill(c *gin.Context) {
	data := struct {
		Page  int `json:"page" form:"page" binding:"required"`
		Limit int `json:"limit" form:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	uid, _ := c.MustGet("uid").(int)
	list := []model.UserBill{}
	if err := config.DB.Where("uid = ?", uid).Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Order("id desc").Find(&list).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取成功", list)
}
func GetPromotionImage(c *gin.Context) {
	data := struct {
		Platform string `json:"platform" form:"platform" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	uid, _ := c.MustGet("uid").(int)
	ctx := context.Background()
	keyName := "user_promotion_" + strconv.Itoa(uid) + "_" + data.Platform
	if value, err := config.Redis.Get(ctx, keyName).Result(); err == nil {
		if value != "" {
			util.Success(c, "获取成功", value)
			return
		}
	}
	image, err := image.GetUserPromotion(int64(uid), data.Platform)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	config.Redis.Set(ctx, keyName, image, time.Hour*24)
	util.Success(c, "获取成功", image)
}
func MyPromotion(c *gin.Context) {
	data := struct {
		Page  int `json:"page" form:"page" binding:"required"`
		Limit int `json:"limit" form:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	uid, _ := c.MustGet("uid").(int)
	var list []model.User
	if err := config.DB.Where("pid = ?", uid).Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Order("id desc").Find(&list).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取成功", list)
}
func SubSuper(c *gin.Context) {
	uid, _ := c.MustGet("uid").(int)
	var applySuper model.ApplySuper
	if err := config.DB.Where("uid = ?", uid).Find(&applySuper).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if applySuper.Id == 0 {
		applySuper.Uid = uid
		if err := config.DB.Create(&applySuper).Error; err != nil {
			util.Error(c, 1, err.Error())
			return
		}
		util.Success(c, "提交成功", nil)
	} else {
		if applySuper.Status == 0 {
			util.Error(c, 1, "您已提交申请，请勿重复提交")
			return
		}
		applySuper.Status = 0
		applySuper.Remark = ""
		if err := config.DB.Save(&applySuper).Error; err != nil {
			util.Error(c, 1, err.Error())
		}
		util.Success(c, "提交成功", nil)
	}

}
func GetSubSuper(c *gin.Context) {
	uid, _ := c.MustGet("uid").(int)
	applySuper := model.ApplySuper{}
	if err := config.DB.Where("uid = ?", uid).Find(&applySuper).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if applySuper.Id == 0 {
		util.Success(c, "获取成功", nil)
		return
	}
	util.Success(c, "获取成功", applySuper)
}
func GetUserGoodsTask(c *gin.Context) {
	data := struct {
		Page  int `json:"page" form:"page" binding:"required"`
		Limit int `json:"limit" form:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	uid, _ := c.MustGet("uid").(int)
	type gt struct {
		model.GoodsTask
		Title string `json:"title" gorm:"column:title"`
	}
	list := []gt{}
	if err := config.DB.Model(&model.GoodsTask{}).Where("goods_task.uid = ?", uid).
		Select("goods_task.*", "goods.title").
		Joins("join goods on goods.id=goods_task.gid").
		Order("goods_task.id desc").Find(&list).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取成功", list)
}
func CollectList(c *gin.Context) {
	data := struct {
		Page  int `json:"page" form:"page" binding:"required"`
		Limit int `json:"limit" form:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	uid, _ := c.MustGet("uid").(int)
	type uc struct {
		model.UserCollect
		Title string `json:"title" gorm:"column:title"`
		Thumb string `json:"thumb" gorm:"column:thumb"`
	}
	list := []uc{}
	if err := config.DB.Where("user_collect.uid = ?", uid).Model(&model.UserCollect{}).
		Offset((data.Page - 1) * data.Limit).
		Joins("join goods on goods.id=user_collect.gid").
		Select("user_collect.*, goods.title,goods.thumb").
		Limit(data.Limit).Order("user_collect.id desc").Find(&list).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取成功", list)
}
