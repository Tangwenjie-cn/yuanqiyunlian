/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package middleware

import (
	"bytes"
	"context"
	"fmt"
	"gin/config"
	"gin/model"
	"gin/util"
	"io"
	"slices"

	"strings"

	"github.com/gin-gonic/gin"
)

// CheckAuth 中间件，验证token和权限
func CheckAuth(c *gin.Context) {
	// 获取请求头中的token
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		util.Error(c, 2, "缺少Token")
		c.Abort()
		return
	}
	ctx := context.Background()
	username, err := config.Redis.Get(ctx, "admin"+token).Result()
	if err != nil {
		util.Error(c, 2, "Token无效")
		c.Abort()
		return
	}
	admin := model.Admin{}
	config.DB.Model(&model.Admin{}).Preload("AuthGroup").Where("username = ?", username).First(&admin)
	if admin.ID == 0 {
		util.Error(c, 2, "Token无效")
		c.Abort()
		return

	}
	if admin.Status == 0 {
		util.Error(c, 2, "账号已禁用")
		c.Abort()
		return

	}
	// 记录日志
	var adminLog model.AdminLog
	adminLog.Account = admin.Username
	adminLog.IP = c.ClientIP()
	adminLog.Uri = c.Request.RequestURI
	bodyByte, _ := c.GetRawData()
	//c.GetRawData()会清空body，所以需要重新赋值
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyByte))
	if len(bodyByte) > 500 {
		adminLog.Body = string(bodyByte[:500]) + "..."
	} else {
		adminLog.Body = string(bodyByte)
	}
	config.DB.Create(&adminLog)

	if admin.GroupId > 1 { // 不是超级管理员,需要验证权限
		auth := model.Auth{}
		config.DB.Model(&model.Auth{}).Where("url = ?", c.Request.URL.Path).Find(&auth)

		if admin.AuthGroup.Auth == "" {
			util.Error(c, 1, "没有权限访问")
			c.Abort()
			return

		}
		auths := strings.Split(admin.AuthGroup.Auth, ",")
		if admin.AuthGroup.HalfAuth != "" { // 半选权限
			// 半选权限需要单独处理，不能直接添加到权限列表中
			auths = append(auths, strings.Split(admin.AuthGroup.HalfAuth, ",")...)

		}
		authID := fmt.Sprintf("%d", auth.ID)
		// 判断权限列表中是否存在该权限
		check := slices.Contains(auths, authID)
		if !check {
			util.Error(c, 1, "没有权限访问")
			c.Abort()
			return
		}

	}

	c.Set("admin", admin)
	c.Next()
}
