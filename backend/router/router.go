/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	corsConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowHeaders:     []string{"Authorization", "Content-Type", "Accept", "Origin", "Content-Length"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposeHeaders:    []string{"Content-Disposition"},
		MaxAge:           12 * 3600, // 12小时
	}
	corsMiddleware := cors.New(corsConfig)
	r.Use(corsMiddleware) // 使用cors中间件
	admin(r)
	api(r)
	r.Static("/static", "./public")
}
