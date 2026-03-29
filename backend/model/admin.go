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

type Admin struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Status    int8      `json:"status"`
	Pid       int       `json:"pid"`
	LoginIp   string    `json:"login_ip"`
	Ltime     time.Time `json:"ltime"`
	Ctime     time.Time `gorm:"autoCreateTime" json:"ctime"`
	Utime     time.Time `gorm:"autoUpdateTime" json:"utime"`
	GroupId   int       `json:"group_id"`
	AuthGroup AuthGroup `gorm:"foreignKey:GroupId"`
}
