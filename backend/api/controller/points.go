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
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserPoints(c *gin.Context) {
	data := struct {
		Page  int `json:"page" form:"page" binding:"required"`
		Limit int `json:"limit" form:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	uid, ok := c.MustGet("uid").(int)
	if !ok {
		util.Error(c, 1, "获取用户信息失败")
		return
	}
	list := []model.UserPoints{}
	if err := config.DB.Where("uid = ?", uid).Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Order("id desc").Find(&list).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取成功", list)
}
func PointsVipList(c *gin.Context) {
	data := struct {
		Page  int `json:"page" form:"page" binding:"required"`
		Limit int `json:"limit" form:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	list := []model.Svip{}
	if err := config.DB.Where("points>0 and status=1").Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Order("id desc").Find(&list).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取成功", list)
}
func PointsBuy(c *gin.Context) {
	data := struct {
		Type int `json:"type" form:"type" binding:"required"`
		Id   int `json:"id" form:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	uid, ok := c.MustGet("uid").(int)
	if !ok {
		util.Error(c, 1, "获取用户信息失败")
		return
	}
	tx := config.DB.Begin()
	var user model.User
	if err := tx.Where("id = ? and status = 1", uid).First(&user).Error; err != nil {
		util.Error(c, 1, "用户不存在")
		return
	}
	var (
		points int
		remark string
		Ptype  int
	)

	switch data.Type {
	case 1:
		//兑换会员
		var svip model.Svip
		if err := tx.Where("id = ? and status = 1", data.Id).First(&svip).Error; err != nil {
			util.Error(c, 1, "会员不存在")
			return
		}
		if user.Points < svip.Points {
			util.Error(c, 1, "积分不足")
			return
		}
		points = svip.Points
		remark = "兑换会员(" + svip.Title + ")"
		Ptype = 2
		if err := handelVip(&svip, uid, tx); err != nil {
			tx.Rollback()
			util.Error(c, 1, err.Error())
			return
		}
	case 2:
		//兑换商品
		var goods model.Goods
		if err := tx.Where("id = ? and status = 1", data.Id).First(&goods).Error; err != nil {
			util.Error(c, 1, "商品不存在")
			return
		}
		if user.Points < goods.Points {
			util.Error(c, 1, "积分不足")
			return
		}
		points = goods.Points
		remark = "兑换商品(" + goods.Title + ")"
		Ptype = 3
		var userGoods model.UserGoods
		if err := tx.Where("uid=? and gid=?", uid, goods.ID).First(&userGoods).Error; err == nil {
			util.Error(c, 1, "商品已兑换")
		}
		if err := tx.Create(&model.UserGoods{
			Uid:  uid,
			Gid:  int(goods.ID),
			Type: 3,
		}).Error; err != nil {
			tx.Rollback()
			util.Error(c, 1, err.Error())
			return
		}
	default:
		util.Error(c, 1, "参数错误")
		return
	}
	if err := tx.Create(&model.UserPoints{
		Uid:    uid,
		Type:   Ptype,
		Points: points,
		Remark: remark,
		Mode:   2,
		Before: user.Points,
		After:  user.Points - points,
	}).Error; err != nil {
		tx.Rollback()
		util.Error(c, 1, err.Error())
		return
	}
	if err := tx.Model(&model.User{}).Where("id=?", uid).Update("points", gorm.Expr("points - ?", points)).Error; err != nil {
		tx.Rollback()
		util.Error(c, 1, err.Error())
		return
	}
	tx.Commit()
	util.Success(c, "兑换成功", nil)
}
func handelVip(svip *model.Svip, uid int, tx *gorm.DB) error {
	var userSvip model.UserSvip
	if err := tx.Where("uid=? and vid=?", uid, svip.ID).Find(&userSvip).Error; err != nil {
		return err
	}
	// 获取当前时间
	now := time.Now()
	if userSvip.ID > 0 {
		if userSvip.ExpireTime.Before(now) {
			userSvip.ExpireTime = now.AddDate(0, 0, svip.Days)
		} else {
			userSvip.ExpireTime = userSvip.ExpireTime.AddDate(0, 0, svip.Days)
		}
		if err := tx.Save(&userSvip).Error; err != nil {
			return err
		}
	} else {
		userSvip.Uid = uid
		userSvip.Vid = int(svip.ID)
		userSvip.ExpireTime = now.AddDate(0, 0, svip.Days)
		if err := tx.Create(&userSvip).Error; err != nil {
			return err
		}
	}
	return nil
}
