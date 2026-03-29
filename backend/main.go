/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package main

import (
	"fmt"
	"gin/config"
	"gin/router"
	"gin/util"

	"github.com/gin-gonic/gin"
)

func main() {

	configPath, err := util.GetConfigPath() //获取配置文件路径
	if err != nil {
		fmt.Println("错误:", err)
		return
	}
	cfg, err := config.LoadConfig(configPath) //加载配置文件
	if err != nil {
		fmt.Println("错误:", err)
		return
	}
	config.InitDatabase(cfg) //初始化数据库连接
	config.InitRedis(cfg)    //初始化redis连接

	//gin设置
	gin.SetMode(cfg.Server.Mode) //设置运行模式
	r := gin.Default()           //创建引擎
	if len(cfg.Server.Proxy) > 0 {
		// 明确指定可信认代理IP（推荐）
		r.SetTrustedProxies(cfg.Server.Proxy)
	}
	router.Init(r)                                                // 初始化路由
	r.Run(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)) // 监听地址
}
