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

type Img struct {
	ID      uint      `json:"id"`
	Thumb   string    `json:"thumb"`
	Sid     int16     `json:"sid"`
	Status  int8      `json:"status"`
	Sort    uint      `json:"sort"`
	Ctime   time.Time `gorm:"autoCreateTime" json:"ctime"`
	Utime   time.Time `gorm:"autoUpdateTime" json:"utime"`
	ImgSort ImgSort   `gorm:"foreignKey:Sid"`
}
