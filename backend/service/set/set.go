/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package set

import (
	"context"
	"encoding/json"
	"errors"
	"gin/config"
	"gin/model"
	"time"
)

// 获取配置
func GetSet(name string, cache bool) (string, error) {
	ctx := context.Background()
	set := model.Config{}
	if cache {
		if config.Redis == nil {
			return "", errors.New("redis未初始化")
		}
		if value, err := config.Redis.Get(ctx, name).Result(); err == nil {
			return value, nil
		}

	}
	if err := config.DB.Select("`key`,value").Where("`key` = ?", name).Find(&set).Error; err != nil {
		return "", err
	}
	if set.Value == "" {
		return "", errors.New("未找到配置")
	}
	config.Redis.Set(ctx, name, set.Value, time.Second*30)
	return set.Value, nil
}

// 获取所有配置信息
func GetAllSet() map[string]string {
	set := []model.Config{}
	ctx := context.Background()
	if value, err := config.Redis.Get(ctx, "all_set").Result(); err == nil {
		var resultMap map[string]string
		json.Unmarshal([]byte(value), &resultMap)
		return resultMap
	}
	config.DB.Select("`key`,value").Find(&set)
	m := make(map[string]string)
	for _, v := range set {
		m[v.Key] = v.Value
	}
	jsonData, err := json.Marshal(m)
	if err == nil {
		config.Redis.Set(ctx, "all_set", jsonData, time.Second*30)
	}
	return m
}
