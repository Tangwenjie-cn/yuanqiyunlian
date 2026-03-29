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

type Article struct {
	ID          uint        `json:"id"`
	Title       string      `json:"title"`
	Thumb       string      `json:"thumb"`
	Sid         int16       `json:"sid"`
	Content     string      `json:"content"`
	Status      int8        `json:"status"`
	Ctime       time.Time   `gorm:"autoCreateTime" json:"ctime"`
	Utime       time.Time   `gorm:"autoUpdateTime" json:"utime"`
	UtimeText   string      `json:"utime_text" gorm:"-"`
	ArticleSort ArticleSort `gorm:"foreignKey:Sid"`
}
