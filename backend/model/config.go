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

type Config struct {
	ID     uint             `json:"id"`
	Title  string           `json:"title"`
	Key    string           `json:"key"`
	Param  string           `json:"param"`
	Value  string           `json:"value"`
	Group  int8             `json:"group"`
	Type   int8             `json:"type"`
	Sort   int16            `json:"sort"`
	Remark string           `json:"remark"`
	List   []map[string]any `json:"list" gorm:"-"` // 额外字段
	Ctime  time.Time        `gorm:"autoCreateTime" json:"ctime"`
	Utime  time.Time        `gorm:"autoUpdateTime" json:"utime"`
}
