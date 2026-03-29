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

type Cash struct {
	ID          uint      `json:"id"`
	ThirdOrder  string    `json:"third_order"`
	Uid         int       `json:"uid"`
	Type        int       `json:"type"`
	Status      int       `json:"status"`
	Money       float64   `json:"money"`
	Fee         float64   `json:"fee"`
	Actual      float64   `json:"actual"`
	Account     string    `json:"account"`
	Realname    string    `json:"realname"`
	Qrcode      string    `json:"qrcode"`
	Bank        string    `json:"bank"`
	Remark      string    `json:"remark"`
	Platform    string    `json:"platform"`
	PackageInfo string    `json:"package_info"`
	Ctime       time.Time `gorm:"autoCreateTime" json:"ctime"`
	Utime       time.Time `gorm:"autoUpdateTime" json:"utime"`
}
