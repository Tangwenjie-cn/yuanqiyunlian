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

func HandleKami(order *model.Order, tx *gorm.DB) (err error) {
	var kami model.Kami
	if err := tx.Where("`key` = ?", order.ThirdOrder).First(&kami).Error; err != nil {
		return err
	}
	if kami.Status == 1 {
		// 卡密已使用
		return errors.New("卡密已使用")
	}
	if order.PayPrice > kami.Price {
		// 卡密金额与订单金额不匹配
		return errors.New("卡密金额与订单金额不匹配")
	}
	if err := tx.Model(&kami).Update("status", 1).Error; err != nil {
		return err
	}
	err = finance.HandleOrder(order, tx)
	if err != nil {
		return err
	}
	return nil
}
