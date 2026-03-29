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

type Order struct {
	ID            uint      `json:"id"`
	OrderNo       string    `json:"order_no"`
	ThirdOrder    string    `json:"third_order"`
	Type          int       `json:"type"`
	Uid           int       `json:"uid"`
	Gid           int       `json:"gid"`
	Vid           int       `json:"vid"`
	PayType       int       `json:"pay_type"`
	Price         float64   `json:"price"`
	DiscountPrice float64   `json:"discount_price"`
	DiscountType  int       `json:"discount_type"`
	PayPrice      float64   `json:"pay_price"`
	RetailPrice   float64   `json:"retail_price"`
	SuperPrice    float64   `json:"super_price"`
	Status        int       `json:"status"`
	Ctime         time.Time `gorm:"autoCreateTime" json:"ctime"`
	Utime         time.Time `gorm:"autoUpdateTime" json:"utime"`
}
