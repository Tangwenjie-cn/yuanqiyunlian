/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package model

import (
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Pass      string    `json:"-"`
	Token     string    `json:"token"`
	Status    int       `json:"status"`
	Pid       int       `json:"pid"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	Balance   float64   `json:"balance"`
	Points    int       `json:"points"`
	SignNum   int       `json:"sign_num"`
	Promotion int       `json:"promotion"`
	IsSuper   int       `json:"is_super"`
	SuperSet  int       `json:"super_set"`
	SuperType int       `json:"super_type"`
	SuperYj   float64   `json:"super_yj"`
	Phone     string    `json:"phone" gorm:"-"`
	Ctime     time.Time `gorm:"autoCreateTime" json:"ctime"`
	Utime     time.Time `gorm:"autoUpdateTime" json:"utime"`
}
