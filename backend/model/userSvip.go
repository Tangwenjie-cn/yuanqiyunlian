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
	"encoding/json"
	"gin/config"
	"gin/util"
	"slices"
	"time"
)

type UserSvip struct {
	ID         uint      `json:"id"`
	Uid        int       `json:"uid"`
	Vid        int       `json:"vid"`
	ExpireTime time.Time `json:"expire_time"`
	Ctime      time.Time `gorm:"autoCreateTime" json:"ctime"`
	Utime      time.Time `gorm:"autoUpdateTime" json:"utime"`
}

// 检查会员用户是否有该商品的权限
func (u UserSvip) CheckGoodsAuth(uid int, gid int) bool {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	var us []UserSvip
	if err := config.DB.Where("uid = ? and expire_time > ?", uid, nowTime).Find(&us).Error; err != nil {
		return false
	}
	if len(us) == 0 {
		return false
	}
	var goods Goods
	if err := config.DB.Where("id = ?", gid).Preload("GsCorrelation").Find(&goods).Error; err != nil {
		return false
	}
	gsList := []int{}
	for _, v := range goods.GsCorrelation {
		gsList = append(gsList, v.Sid)
	}
	for _, v := range us {
		var sp []SvipPrivilege
		config.DB.Where("vid = ?", v.Vid).Preload("Privilege").Find(&sp)
		if len(sp) == 0 {
			continue
		}
		for _, v2 := range sp {
			var content []int
			json.Unmarshal(v2.Privilege.Content, &content)
			switch v2.Privilege.Type {
			case 1:
				if slices.Contains(content, goods.Type) {
					return true
				}
			case 2:
				if util.HasIntersection(content, gsList) {
					return true
				}
			case 3:
				if slices.Contains(content, goods.Cid) {
					return true
				}
			}
		}
	}
	return false
}
