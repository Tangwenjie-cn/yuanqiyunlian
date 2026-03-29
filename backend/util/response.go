/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Code int    `json:"code"` // 状态码
	Msg  string `json:"msg"`  // 消息
	Data any    `json:"data"` // 数据
}

func Success(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, JsonResponse{
		Code: 0,
		Msg:  msg,
		Data: data,
	})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, JsonResponse{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
