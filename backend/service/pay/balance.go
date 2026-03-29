/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package pay

import (
	"errors"
	"gin/model"
	"gin/service/finance"

	"gorm.io/gorm"
)

func HandleBalance(order *model.Order, tx *gorm.DB) (err error) {
	//TODO:处理余额支付
	var user model.User
	if err = tx.First(&user, "id = ?", order.Uid).Error; err != nil {
		return errors.New("用户不存在")
	}
	if user.Balance < order.PayPrice {
		return errors.New("余额不足")
	}
	var userBill model.UserBill
	var remark string
	if order.Type == 1 {
		var goods model.Goods
		if err = tx.Select("title").First(&goods, "id = ?", order.Gid).Error; err != nil {
			return errors.New("商品不存在")
		}
		remark = "购买" + goods.Title
	}
	if order.Type == 2 {
		var vip model.Svip
		if err = tx.Select("title").First(&vip, "id = ?", order.Vid).Error; err != nil {
			return errors.New("会员不存在")
		}
		remark = "购买" + vip.Title
	}
	err = userBill.Add(order.Uid, 2, 4, order.PayPrice, remark, int(order.ID), tx)
	if err != nil {
		return err
	}
	err = finance.HandleOrder(order, tx)
	if err != nil {
		return err
	}
	return nil
}
