/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package config

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	// 日志输出
	gin.DisableConsoleColor()
	// Logging to a file.
	f, _ := os.Create("log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
