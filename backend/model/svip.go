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

type Svip struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Desc       string    `json:"desc"`
	Days       int       `json:"days"`
	Price      float64   `json:"price"`
	Rebate     float64   `json:"rebate"`
	Points     int       `json:"points"`
	Status     int       `json:"status"`
	Sort       int       `json:"sort"`
	RetailOn   int       `json:"retail_on"` // 是否分销
	RetailSet  int       `json:"retail_set"`
	RetailType int       `json:"retail_type"`
	RetailYj   float64   `json:"retail_yj"`
	Ctime      time.Time `gorm:"autoCreateTime" json:"ctime"`
	Utime      time.Time `gorm:"autoUpdateTime" json:"utime"`
}
