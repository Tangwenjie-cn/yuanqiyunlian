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
	"gin/util"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	data := struct {
		Keyword string `form:"keyword" json:"keyword"`
		Status  int    `form:"status" json:"status"`
		Pid     int    `form:"pid" json:"pid"`
		Page    int    `form:"page" json:"page" binding:"required"`
		Limit   int    `form:"limit" json:"limit" binding:"required"`
		Date    string `form:"date" json:"date"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var total int64
	query := config.DB.Model(&model.User{}).Joins("left join user as puser on user.pid=puser.id")
	if data.Keyword != "" {
		query = query.Where("(user.nickname = ? or user.id = ?)", strings.TrimSpace(data.Keyword), strings.TrimSpace(data.Keyword))
	}
	if data.Status != 0 {
		query = query.Where("user.status = ?", data.Status)
	}
	if data.Pid != 0 {
		query = query.Where("user.pid = ?", data.Pid)
	}
	if data.Date != "" {
		dateArr := strings.Split(data.Date, ",")
		query = query.Where("user.ctime between ? and ?", dateArr[0], dateArr[1])
	}
	query.Count(&total)
	type userInfo struct {
		model.User
		Pname string `json:"pname"`
	}
	var list []userInfo
	query = query.Select("user.*,puser.nickname as pname").Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Order("id desc")
	if err := query.Find(&list).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	res := map[string]any{
		"data":  list,
		"total": total,
	}
	util.Success(c, "获取列表成功", res)
}
func UserInfo(c *gin.Context) {
	data := struct {
		Id int `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var user model.User
	if err := config.DB.Where("id=?", data.Id).First(&user).Error; err != nil {
		util.Error(c, 1, "用户不存在")
		return
	}
	var bind []model.UserBind
	if err := config.DB.Where("uid = ?", data.Id).Find(&bind).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	type vipInfo struct {
		model.UserSvip
		Title string `json:"title"`
	}
	var vip []vipInfo
	if err := config.DB.Model(&model.UserSvip{}).Select("user_svip.*", "svip.title").
		Joins("left join svip on svip.id = user_svip.vid").Where("user_svip.uid = ?", data.Id).Find(&vip).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取成功", map[string]any{
		"user": user,
		"bind": bind,
		"vip":  vip,
	})
}
func SetUserInfo(c *gin.Context) {
	data := struct {
		Id     int `form:"id" json:"id" binding:"required"`
		Status int `form:"status" json:"status"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var user model.User
	if data.Status != 0 {
		user.Status = data.Status
		//清除登录状态
		var user1 model.User
		if err := config.DB.Where("id = ?", data.Id).First(&user1).Error; err != nil {
			util.Error(c, 1, "用户不存在")
			return
		}
		ctx := context.Background()
		config.Redis.Del(ctx, "api"+user1.Token)
	}
	if err := config.DB.Model(&model.User{}).Where("id = ?", data.Id).Updates(&user).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "修改成功", nil)
}
func SetUserSuper(c *gin.Context) {
	data := struct {
		Id        int     `form:"id" json:"id" binding:"required"`
		IsSuper   int     `form:"is_super" json:"is_super" binding:"required"`
		SuperSet  int     `form:"super_set" json:"super_set" binding:"required"`
		SuperType int     `form:"super_type" json:"super_type"`
		SuperYj   float64 `form:"super_yj" json:"super_yj"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var user model.User
	if err := config.DB.Where("id = ?", data.Id).First(&user).Error; err != nil {
		util.Error(c, 1, "用户不存在")
		return
	}
	user.IsSuper = data.IsSuper
	user.SuperSet = data.SuperSet
	user.SuperType = data.SuperType
	user.SuperYj = data.SuperYj
	if err := config.DB.Save(&user).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "设置成功", nil)
}
func GetUserVip(c *gin.Context) {
	data := struct {
		Id int `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	type vipInfo struct {
		model.UserSvip
		Title string `json:"name"`
	}
	var userSvip []vipInfo
	if err := config.DB.Model(&model.UserSvip{}).Select("user_svip.*", "svip.title").
		Joins("left join svip on svip.id = user_svip.vid").Where("user_svip.uid = ?", data.Id).Find(&userSvip).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var vip []model.Svip
	if err := config.DB.Where("status=1").Find(&vip).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取成功", map[string]any{
		"userSvip": userSvip,
		"vip":      vip,
	})
}
func SetUserVip(c *gin.Context) {
	data := struct {
		Uid        int       `form:"uid" json:"uid" binding:"required"`
		Vid        int       `form:"vid" json:"vid" binding:"required"`
		ExpireTime time.Time `form:"expire_time" json:"expire_time" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var userSvip model.UserSvip
	if err := config.DB.Where("uid = ? and vid = ?", data.Uid, data.Vid).Find(&userSvip).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if userSvip.ID == 0 {
		userSvip.Uid = data.Uid
		userSvip.Vid = data.Vid
		userSvip.ExpireTime = data.ExpireTime
		if err := config.DB.Create(&userSvip).Error; err != nil {
			util.Error(c, 1, err.Error())
			return
		}
	} else {
		userSvip.ExpireTime = data.ExpireTime
		if err := config.DB.Save(&userSvip).Error; err != nil {
			util.Error(c, 1, err.Error())
			return
		}
	}
	util.Success(c, "设置成功", nil)
}
func GetApplySuper(c *gin.Context) {
	data := struct {
		Keyword string `form:"keyword" json:"keyword"`
		Status  int    `form:"status" json:"status"`
		Page    int    `form:"page" json:"page" binding:"required"`
		Limit   int    `form:"limit" json:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var total int64
	query := config.DB.Model(&model.ApplySuper{})
	if data.Keyword != "" {
		query = query.Where("apply_super.uid = ?", data.Keyword)
	}
	if data.Status != 0 {
		if data.Status == 2 {
			data.Status = 0
		}
		query = query.Where("apply_super.status = ?", data.Status)
	}
	query.Count(&total)
	type applySuperInfo struct {
		model.ApplySuper
		Nickname  string  `json:"nickname"`
		Balance   float64 `json:"balance"`
		Points    int     `json:"points"`
		Promotion int     `json:"promotion"`
	}
	var list []applySuperInfo
	if err := query.Joins("join user on user.id=apply_super.uid").Limit(data.Limit).
		Select("apply_super.*", "user.nickname", "user.balance", "user.points", "user.promotion").
		Offset((data.Page - 1) * data.Limit).Order("apply_super.utime desc").Find(&list).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取成功", map[string]any{
		"total": total,
		"data":  list,
	})
}
func SetApplySuper(c *gin.Context) {
	data := struct {
		Id        int     `form:"id" json:"id" binding:"required"`
		Status    int     `form:"status" json:"status" binding:"required"`
		SuperSet  int     `form:"super_set" json:"super_set"`
		SuperType int     `form:"super_type" json:"super_type"`
		SuperYj   float64 `form:"super_yj" json:"super_yj"`
		Remark    string  `form:"remark" json:"remark"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var applySuper model.ApplySuper
	if err := config.DB.Where("id = ?", data.Id).First(&applySuper).Error; err != nil {
		util.Error(c, 1, "申请不存在")
		return
	}
	if data.Status == 1 {
		if data.SuperSet == 0 {
			util.Error(c, 1, "请填写完整信息")
			return
		}
		var user model.User
		if err := config.DB.Where("id = ?", applySuper.Uid).First(&user).Error; err != nil {
			util.Error(c, 1, "用户不存在")
			return
		}
		user.IsSuper = 1
		user.SuperSet = data.SuperSet
		user.SuperType = data.SuperType
		user.SuperYj = data.SuperYj
		if err := config.DB.Save(&user).Error; err != nil {
			util.Error(c, 1, err.Error())
			return
		}

	}
	if data.Status == -1 {
		if data.Remark == "" {
			util.Error(c, 1, "请填写拒绝理由")
			return
		}
		applySuper.Remark = data.Remark
	}
	applySuper.Status = data.Status
	if err := config.DB.Save(&applySuper).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "设置成功", nil)
}
func BalanceList(c *gin.Context) {
	data := struct {
		Keyword string `form:"keyword" json:"keyword"`
		Type    int    `form:"type" json:"type"`
		Page    int    `form:"page" json:"page" binding:"required"`
		Limit   int    `form:"limit" json:"limit" binding:"required"`
		Date    string `form:"date" json:"date"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var total int64
	query := config.DB.Model(&model.UserBill{}).Joins("join user on user.id=user_bill.uid")
	if data.Keyword != "" {
		query = query.Where("user_bill.uid = ?", data.Keyword)
	}
	if data.Type != 0 {
		query = query.Where("user_bill.type = ?", data.Type)
	}
	if data.Date != "" {
		dateArr := strings.Split(data.Date, ",")
		query = query.Where("user_bill.ctime between ? and ?", dateArr[0], dateArr[1])
	}
	query.Count(&total)
	type userBillInfo struct {
		model.UserBill
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
	}
	var list []userBillInfo
	if err := query.Select("user_bill.*, user.nickname, user.avatar").Limit(data.Limit).Offset((data.Page - 1) * data.Limit).
		Order("user_bill.id desc").Find(&list).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取成功", map[string]any{
		"total": total,
		"data":  list,
	})
}
func UserGoodsList(c *gin.Context) {
	data := struct {
		Keyword string `form:"keyword" json:"keyword"`
		Type    int    `form:"type" json:"type"`
		Page    int    `form:"page" json:"page" binding:"required"`
		Limit   int    `form:"limit" json:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var total int64
	query := config.DB.Model(&model.UserGoods{}).Joins("join user on user.id=user_goods.uid").Joins("join goods on goods.id=user_goods.gid")
	if data.Keyword != "" {
		query = query.Where("user_goods.uid = ?", data.Keyword)
	}
	if data.Type != 0 {
		query = query.Where("user_goods.type = ?", data.Type)
	}
	query.Count(&total)
	type userGoodsInfo struct {
		model.UserGoods
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
		Title    string `json:"title"`
	}
	var list []userGoodsInfo
	if err := query.Select("user_goods.*, user.nickname, user.avatar, goods.title").Limit(data.Limit).
		Offset((data.Page - 1) * data.Limit).Order("user_goods.id desc").Find(&list).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取成功", map[string]any{
		"total": total,
		"data":  list,
	})
}
