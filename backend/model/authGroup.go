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

import "time"

type AuthGroup struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	Auth     string    `json:"auth"`
	HalfAuth string    `json:"half_auth"`
	AdminId  int       `json:"admin_id"`
	Ctime    time.Time `gorm:"autoCreateTime" json:"ctime"`
	Utime    time.Time `gorm:"autoUpdateTime" json:"utime"`
}
