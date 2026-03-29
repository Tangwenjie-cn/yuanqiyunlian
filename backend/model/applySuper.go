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

type ApplySuper struct {
	Id     int       `json:"id"`  //主键id
	Uid    int       `json:"uid"` //用户id
	Status int       `json:"status"`
	Remark string    `json:"remark"`
	Ctime  time.Time `gorm:"autoCreateTime" json:"ctime"`
	Utime  time.Time `gorm:"autoUpdateTime" json:"utime"`
}
