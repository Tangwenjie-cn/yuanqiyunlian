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

type GoodsTask struct {
	ID        int64     `json:"id"`
	Uid       int64     `json:"uid"`
	Gid       int64     `json:"gid"`
	Adv       int64     `json:"adv"`
	AdvNums   int64     `json:"adv_nums"`
	Share     int64     `json:"share"`
	ShareNums int64     `json:"share_nums"`
	Ctime     time.Time `gorm:"autoCreateTime" json:"ctime"`
	Utime     time.Time `gorm:"autoUpdateTime" json:"utime"`
}
