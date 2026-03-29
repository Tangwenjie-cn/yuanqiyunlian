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

	"gorm.io/gorm"
)

type UserBill struct {
	ID     uint      `json:"id"`
	Uid    int       `json:"uid"`
	Mode   int       `json:"mode"`
	Type   int       `json:"type"`
	Before float64   `json:"before"`
	After  float64   `json:"after"`
	Money  float64   `json:"money"`
	Remark string    `json:"remark"`
	Tid    int       `json:"tid"`
	Ctime  time.Time `gorm:"autoCreateTime" json:"ctime"`
}

func (u UserBill) Add(uid int, mode int, type_ int, money float64, remark string, tid int, tx *gorm.DB) error {
	var user User
	if err := tx.Where("id = ?", uid).First(&user).Error; err != nil {
		return err
	}
	var userBill UserBill
	userBill.Before = user.Balance
	if mode == 1 {
		userBill.After = user.Balance + money
	} else {
		userBill.After = user.Balance - money
	}
	userBill.Uid = uid
	userBill.Mode = mode
	userBill.Remark = remark
	userBill.Type = type_
	userBill.Money = money
	userBill.Tid = tid
	if err := tx.Create(&userBill).Error; err != nil {
		return err
	}
	if mode == 1 {
		if err := tx.Model(&user).Update("balance", gorm.Expr("balance+?", money)).Error; err != nil {
			return err
		}
	} else {
		if err := tx.Model(&user).Update("balance", gorm.Expr("balance-?", money)).Error; err != nil {
			return err
		}
	}

	return nil
}
