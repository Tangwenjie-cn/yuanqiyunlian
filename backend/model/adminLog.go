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

type AdminLog struct {
	ID      uint      `json:"id"`
	Account string    `json:"account"`
	Body    string    `json:"body"`
	IP      string    `json:"ip"`
	Uri     string    `json:"uri"`
	Ctime   time.Time `gorm:"autoCreateTime" json:"ctime"`
}
