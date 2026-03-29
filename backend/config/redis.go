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
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func InitRedis(cfg *AppConfig) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port), // Redis server address
		Password: cfg.Redis.Password,                                   // No password set
		DB:       cfg.Redis.DB,                                         // use default DB
	})
	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Redis connection failed:", err)
		panic("redis连接失败")
	}
}
