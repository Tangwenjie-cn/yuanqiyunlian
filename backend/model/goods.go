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

type Goods struct {
	ID            uint            `json:"id"`
	Type          int             `json:"type"`
	Code          string          `json:"code"`
	Title         string          `json:"title"`
	Thumb         string          `json:"thumb"`
	Link          string          `json:"link"`
	Introduction  string          `json:"introduction"`   // 简介
	OriginalPrice float64         `json:"original_price"` // 原价
	Price         float64         `json:"price"`
	Points        int             `json:"points"`
	CostPrice     float64         `json:"cost_price"` // 成本价
	Pages         int             `json:"pages"`      // 页数
	Status        int             `json:"status"`
	IsTop         int             `json:"is_top"` // 是否置顶
	Views         uint            `json:"views"`  // 浏览量
	ViewsText     string          `json:"views_text" gorm:"-"`
	RetailOn      int             `json:"retail_on"` // 是否分销
	RetailSet     int             `json:"retail_set"`
	RetailType    int             `json:"retail_type"`
	RetailYj      float64         `json:"retail_yj"`
	IsAdv         int             `json:"is_adv"`     // 是否开启广告
	AdvNums       int             `json:"adv_nums"`   // 广告次数
	IsShare       int             `json:"is_share"`   // 是否开启分享
	ShareNums     int             `json:"share_nums"` // 分享次数
	Cid           int             `json:"cid"`        // 栏目id
	Ctime         time.Time       `gorm:"autoCreateTime" json:"ctime"`
	Utime         time.Time       `gorm:"autoUpdateTime" json:"utime"`
	UtimeText     string          `json:"utime_text" gorm:"-"`
	GsCorrelation []GsCorrelation `gorm:"foreignKey:Gid;" json:"gs_correlation"`
	GoodsContent  GoodsContent    `gorm:"foreignKey:Gid;" json:"goods_content"`
}
