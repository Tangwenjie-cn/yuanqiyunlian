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

	"gorm.io/datatypes"
)

type Theme struct {
	ID      int            `json:"id"`
	Title   string         `json:"title"`
	Type    int8           `json:"type"`
	Content datatypes.JSON `json:"content" gorm:"type:json"`
	Status  int8           `json:"status"`
	Ctime   time.Time      `gorm:"autoCreateTime" json:"ctime"`
	Utime   time.Time      `gorm:"autoUpdateTime" json:"utime"`
}
