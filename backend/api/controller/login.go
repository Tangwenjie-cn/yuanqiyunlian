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
	"gin/service/cloud"
	"gin/service/task"
	"gin/service/wechat"
	"gin/util"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	data := struct {
		Account string `form:"account" json:"account" binding:"required"`
		Code    string `form:"code" json:"code" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if ok, _ := regexp.MatchString(`^1[3-9]\d{9}$`, data.Account); !ok {
		util.Error(c, 1, "手机号格式不正确")
		return
	}
	captcha := model.Captcha{}
	if err := config.DB.Where("account=? and code=? and type=1 and status=0", data.Account, data.Code).Last(&captcha).Error; err != nil {
		util.Error(c, 1, "验证码错误")
		return
	}
	now := time.Now() // 获取当前本地时间
	if now.Sub(captcha.Ctime) > 5*time.Minute {
		util.Error(c, 1, "验证码已过期")
		return
	}
	if err := config.DB.Model(&captcha).Where("id=?", captcha.ID).Update("status", 1).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	userBind := model.UserBind{}
	if err := config.DB.Where("phone=? and type=1", data.Account).Find(&userBind).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if userBind.ID == 0 {
		util.Success(c, "success", map[string]any{"phone": data.Account, "type": 1, "is_register": false})
		return
	}
	user := model.User{}
	if err := config.DB.Where("id=?", userBind.Uid).First(&user).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if user.Status == -1 {
		util.Error(c, 1, "该账号已被禁用")
		return
	}
	ctx := context.Background()
	config.Redis.Unlink(ctx, "api"+user.Token)
	user.Token = util.Sha256Hash(uuid.New().String())
	config.Redis.Set(ctx, "api"+user.Token, user.ID, time.Hour*24*180)
	config.DB.Save(&user)
	util.Success(c, "success", map[string]any{"user": user})
}
func WxMiniLogin(c *gin.Context) {
	data := struct {
		Code string `form:"code" json:"code" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	wxUser := wechat.GetMiniCode2Session(data.Code)
	if wxUser["openid"] == "" {
		util.Error(c, 1, wxUser["errmsg"].(string))
		return
	}
	userBind := model.UserBind{}
	if err := config.DB.Where("openid=? and type=2", wxUser["openid"]).Find(&userBind).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if userBind.ID == 0 {
		util.Success(c, "success", map[string]any{"openid": wxUser["openid"], "type": 2, "is_register": false})
		return
	}
	user := model.User{}
	if err := config.DB.Where("id=?", userBind.Uid).First(&user).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if user.Status == -1 {
		util.Error(c, 1, "该账号已被禁用")
		return
	}
	ctx := context.Background()
	config.Redis.Unlink(ctx, "api"+user.Token)
	user.Token = util.Sha256Hash(uuid.New().String())
	config.Redis.Set(ctx, "api"+user.Token, user.ID, time.Hour*24*180)
	config.DB.Save(&user)
	util.Success(c, "success", map[string]any{"user": user})
}
func Register(c *gin.Context) {
	data := struct {
		Type     int    `form:"type" json:"type" binding:"required"`
		Openid   string `form:"openid" json:"openid"`
		Phone    string `form:"phone" json:"phone"`
		Nickname string `form:"nickname" json:"nickname" binding:"required"`
		Pid      int    `form:"pid" json:"pid"`
		Gid      int    `form:"gid" json:"gid"`
		Platform string `json:"platform" form:"platform" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	userBind := model.UserBind{}
	query1 := config.DB.Where("type=?", data.Type)
	if data.Type == 1 {
		query1.Where("phone=?", data.Phone)
	}
	if data.Type > 1 {
		query1.Where("openid=?", data.Openid)

	}
	if err := query1.Find(&userBind).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if userBind.ID > 0 {
		util.Error(c, 1, "该账号已注册")
		return
	}
	//todo 注册
	// 开启事务
	tx := config.DB.Begin()
	user := model.User{}
	user.Nickname = data.Nickname
	user.Pid = data.Pid
	user.Status = 1
	ctx := context.Background()
	user.Token = util.Sha256Hash(uuid.New().String())
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		util.Error(c, 1, err.Error())
		return
	}
	userBind.Uid = user.ID
	userBind.Type = data.Type
	userBind.Openid = data.Openid
	userBind.Phone = data.Phone
	if err := tx.Create(&userBind).Error; err != nil {
		tx.Rollback()
		util.Error(c, 1, err.Error())
		return
	}
	tx.Commit()
	config.Redis.Set(ctx, "api"+user.Token, user.ID, time.Hour*24*180)
	util.Success(c, "success", user)
	//邀请人
	if data.Pid > 0 {
		config.DB.Model(&model.User{}).Where("id=?", data.Pid).Update("promotion", gorm.Expr("promotion + ?", 1))
		if data.Gid > 0 {
			config.DB.Model(&model.GoodsTask{}).Where("uid=? and gid=?", data.Pid, data.Gid).Update("share_nums", gorm.Expr("share_nums + ?", 1))
			task.HandleGoodsTask(data.Pid, data.Gid)
		}
	}
}
func BindUser(c *gin.Context) {
	data := struct {
		Type   int    `form:"type" json:"type" binding:"required"`
		Openid string `form:"openid" json:"openid"`
		Phone  string `form:"phone" json:"phone"`
		Uid    uint   `form:"uid" json:"uid" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	user := model.User{}
	if err := config.DB.Where("id=?", data.Uid).Find(&user).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if user.ID == 0 {
		util.Error(c, 1, "用户不存在")
		return

	}
	ctx := context.Background()
	config.Redis.Unlink(ctx, "api"+user.Token)
	user.Token = util.Sha256Hash(uuid.New().String())
	config.Redis.Set(ctx, "api"+user.Token, user.ID, time.Hour*24*180)
	config.DB.Save(&user)
	userBind := model.UserBind{}
	userBind.Uid = user.ID
	userBind.Type = data.Type
	userBind.Openid = data.Openid
	userBind.Phone = data.Phone
	if err := config.DB.Create(&userBind).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "success", user)
}
func LoginOut(c *gin.Context) {
	uid, _ := c.MustGet("uid").(int)
	token := c.Request.Header.Get("Authorization")
	ctx := context.Background()
	config.Redis.Del(ctx, "api"+token)
	config.DB.Model(&model.User{}).Where("id=?", uid).Update("token", "")
	util.Success(c, "success", nil)
}
func SendSms(c *gin.Context) {
	data := struct {
		Phone string `form:"phone" json:"phone" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var captcha model.Captcha
	if err := config.DB.Where("account=? and type=1 and status = 0 and ctime > ?", data.Phone, time.Now().Add(-time.Second*60)).First(&captcha).Error; err == nil {
		util.Error(c, 1, "请勿频繁发送短信")
		return
	}
	startTime, endTime := util.GetTodayStartEnd()
	var count int64
	config.DB.Model(&model.Captcha{}).Where("account=? and type=1 and ctime between ? and ?", data.Phone, startTime, endTime).Count(&count)
	if count >= 5 {
		util.Error(c, 1, "今日发送次数已达上限")
		return
	}
	code := rand.Intn(900000) + 100000
	captcha = model.Captcha{
		Account: data.Phone,
		Code:    strconv.Itoa(code),
		Type:    1,
	}
	if err := config.DB.Create(&captcha).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := cloud.SendSms(captcha.Account, captcha.Code); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "success", nil)
}
