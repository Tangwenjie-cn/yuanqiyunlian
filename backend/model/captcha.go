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

type Captcha struct {
	ID      uint      `json:"id"`
	Type    int       `json:"type"`
	Status  int8      `json:"status"`
	Account string    `json:"account"`
	Code    string    `json:"code"`
	Ctime   time.Time `gorm:"autoCreateTime" json:"ctime"`
	Utime   time.Time `gorm:"autoUpdateTime" json:"utime"`
}
