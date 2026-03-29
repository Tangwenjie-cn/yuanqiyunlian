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
	"fmt"
	"gin/config"
	"gin/model"
	"gin/service/set"
	"gin/util"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	admin := model.Admin{}
	data := struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	config.DB.Preload("AuthGroup").Where("username = ?", data.Username).Find(&admin)
	if admin.ID == 0 {
		util.Error(c, 1, "账号不存在")
		return
	}

	if util.Sha256Hash(data.Password) != admin.Password {
		util.Error(c, 1, "密码错误")
		return
	}
	admin.Ltime = time.Now()
	admin.LoginIp = c.ClientIP()
	token := util.Sha256Hash(fmt.Sprintf("%d%d", time.Now().Unix(), admin.ID))
	ctx := context.Background()
	config.Redis.Set(ctx, "admin"+token, admin.Username, time.Hour*24*7)
	config.Redis.HSet(ctx, "yqyl_admin", "admin"+token, admin.Username)
	res := config.DB.Save(&admin)
	if res.Error != nil {
		util.Error(c, 1, res.Error.Error())
		return
	}
	auths, err := model.Auth{}.GetTree(true, admin.GroupId)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	logo, _ := set.GetSet("site_logo", true)
	list := map[string]any{
		"token":    token,
		"id":       admin.ID,
		"group_id": admin.GroupId,
		"username": admin.Username,
		"menu":     auths,
		"logo":     logo,
	}
	util.Success(c, "Login successful", list)
}
func Logout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		util.Error(c, 1, "缺少Token")
		return
	}
	ctx := context.Background()
	config.Redis.Del(ctx, "admin"+token)
	config.Redis.HDel(ctx, "yqyl_admin", "admin"+token)
	util.Success(c, "Logout successful", nil)
}
