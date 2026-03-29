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
	"os"

	"github.com/pelletier/go-toml/v2"
)

type ServerConfig struct {
	Host  string   `toml:"host"`
	Port  int      `toml:"port"`
	Mode  string   `toml:"mode"`
	Proxy []string `toml:"proxy"`
}

type DatabaseConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Name     string `toml:"name"`
	// Prefix   string `toml:"prefix"`
}
type RedisConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}
type AppConfig struct {
	Server   ServerConfig   `toml:"server"`
	Database DatabaseConfig `toml:"database"`
	Redis    RedisConfig    `toml:"redis"`
}

func LoadConfig(path string) (cfg *AppConfig, err error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return
	}
	if err = toml.Unmarshal(file, &cfg); err != nil {
		return
	}
	return
}
