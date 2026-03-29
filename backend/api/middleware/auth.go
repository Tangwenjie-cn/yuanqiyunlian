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
	"context"
	"gin/config"
	"gin/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CheckToken(c *gin.Context) {
	notVerify := map[string]bool{
		"/api/login/login":    true,
		"/api/login/wxMini":   true,
		"/api/login/register": true,
		"/api/login/bindUser": true,
		"/api/login/sendSms":  true,

		"/api/index/theme":  true,
		"/api/index/config": true,

		"/api/article/list": true,
		"/api/article/info": true,

		"/api/goods/list": true,
		"/api/goods/info": true,
		"/api/goods/sort": true,
	}
	path := c.Request.URL.Path
	if _, ok := notVerify[path]; !ok {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			util.Error(c, 2, "请登录")
			c.Abort()
			return
		}
		ctx := context.Background()
		rUid, _ := config.Redis.Get(ctx, "api"+token).Result()
		if rUid == "" {
			util.Error(c, 2, "token错误")
			c.Abort()
			return
		}
		uid, err := strconv.Atoi(rUid)
		if err != nil {
			util.Error(c, 2, "token错误")
			config.Redis.Del(ctx, "api"+token)
			c.Abort()
			return
		}
		c.Set("uid", uid)
	}
	c.Next()
}
