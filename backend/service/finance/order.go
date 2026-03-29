/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package finance

import (
	"errors"
	"fmt"
	"gin/model"
	"gin/service/set"
	"math"
	"strconv"
	"time"

	"gorm.io/gorm"
)

// TODO: 订单处理
func HandleOrder(order *model.Order, tx *gorm.DB) error {
	// TODO: 处理订单
	if order.Type == 1 {
		// TODO: 处理商品订单
		err := handelGoods(order, tx)
		if err != nil {
			return err
		}
	}
	if order.Type == 2 {
		// TODO: 处理vip订单
		err := handelVip(order, tx)
		if err != nil {
			return err
		}
	}
	retailOrder(order, tx)
	handelOrderPoints(order, tx)

	order.Status = 1
	if err := tx.Save(order).Error; err != nil {
		return err
	}
	return nil
}

// TODO: 处理VIP订单
func handelVip(order *model.Order, tx *gorm.DB) error {
	var svip model.Svip
	if err := tx.Where("id=? and status=1", order.Vid).First(&svip).Error; err != nil {
		return errors.New("svip not found")
	}
	var userSvip model.UserSvip
	if err := tx.Where("uid=? and vid=?", order.Uid, order.Vid).Find(&userSvip).Error; err != nil {
		return err
	}
	// 获取当前时间
	now := time.Now()
	if userSvip.ID > 0 {
		if userSvip.ExpireTime.Before(now) {
			userSvip.ExpireTime = now.AddDate(0, 0, svip.Days)
		} else {
			userSvip.ExpireTime = userSvip.ExpireTime.AddDate(0, 0, svip.Days)
		}
		if err := tx.Save(&userSvip).Error; err != nil {
			return err
		}
	} else {
		userSvip.Uid = order.Uid
		userSvip.Vid = order.Vid
		userSvip.ExpireTime = now.AddDate(0, 0, svip.Days)
		if err := tx.Create(&userSvip).Error; err != nil {
			return err
		}
	}
	return nil
}

// TODO: 处理商品订单
func handelGoods(order *model.Order, tx *gorm.DB) error {
	var goods model.Goods
	if err := tx.Where("id=? and status=1", order.Gid).First(&goods).Error; err != nil {
		return err
	}
	var UserGoods model.UserGoods
	if err := tx.Where("uid=? and gid=?", order.Uid, order.Gid).Find(&UserGoods).Error; err != nil {
		return err
	}
	if UserGoods.ID > 0 {
		return errors.New("已拥有该商品")
	}
	UserGoods.Uid = order.Uid
	UserGoods.Gid = order.Gid
	UserGoods.Type = 1
	if err := tx.Create(&UserGoods).Error; err != nil {
		return err
	}
	return nil
}

// TODO: 处理用户分销
func retailOrder(order *model.Order, tx *gorm.DB) error {
	retailOn, _ := set.GetSet("retail_on", true)
	if retailOn == "0" || retailOn == "" { //未开启分销
		return nil
	}
	retail_type, _ := set.GetSet("retail_type", true)
	if retail_type == "" {
		return errors.New("retail_type not found")
	}
	retail_yj, _ := set.GetSet("retail_yj", true)
	if retail_yj == "" {
		return errors.New("retail_yj not found")
	}
	var retailType int
	var retailYj float64
	if order.Type == 1 {
		var goods model.Goods
		if err := tx.Where("id=?", order.Gid).First(&goods).Error; err != nil {
			return err
		}
		if goods.RetailOn == -1 {
			return nil
		}
		if goods.RetailSet == 1 { //公共设置
			retailType, _ = strconv.Atoi(retail_type)
			retailYj, _ = strconv.ParseFloat(retail_yj, 64)
		} else {
			retailType = goods.RetailType
			retailYj = goods.RetailYj
		}
	}
	if order.Type == 2 {
		var svip model.Svip
		if err := tx.Where("id=?", order.Vid).First(&svip).Error; err != nil {
			return err
		}
		if svip.RetailOn == -1 {
			return nil
		}
		if svip.RetailSet == 1 { //公共设置
			retailType, _ = strconv.Atoi(retail_type)
			retailYj, _ = strconv.ParseFloat(retail_yj, 64)
		} else {
			retailType = svip.RetailType
			retailYj = svip.RetailYj
		}
	}
	var user model.User
	if err := tx.Where("id=?", order.Uid).Select("pid", "nickname").First(&user).Error; err != nil {
		return err
	}
	if user.Pid > 0 {
		//推广人佣金
		if retailYj > 0 {
			if retailType == 1 {
				order.RetailPrice = retailYj
			} else {
				order.RetailPrice = (retailYj / 100) * order.PayPrice
			}
			if order.RetailPrice > 0 {
				err := model.UserBill{}.Add(user.Pid, 1, 1, order.RetailPrice, fmt.Sprintf("推广(%s ID:%d)获得佣金", user.Nickname, user.ID), int(order.ID), tx)
				if err != nil {
					return err
				}
			}

		}
		//判断特级身份
		var puser model.User
		if err := tx.Where("id=?", user.Pid).First(&puser).Error; err != nil {
			return err
		}
		if puser.IsSuper > 0 {
			var superSet int
			var superYj float64
			if puser.SuperSet == 1 { //公共设置
				super_type, _ := set.GetSet("super_type", true)
				if super_type == "" {
					return errors.New("super_type not found")
				}
				super_yj, _ := set.GetSet("super_yj", true)
				if super_yj == "" {
					return errors.New("super_yj not found")
				}
				superSet, _ = strconv.Atoi(super_type)
				superYj, _ = strconv.ParseFloat(super_yj, 64)
			} else {
				superSet = puser.SuperType
				superYj = puser.SuperYj
			}
			if superYj > 0 {
				if superSet == 1 {
					order.SuperPrice = superYj
				} else {
					order.SuperPrice = (superYj / 100) * order.PayPrice
				}
				if order.SuperPrice > 0 {
					super_name, _ := set.GetSet("super_name", true)
					err := model.UserBill{}.Add(user.Pid, 1, 2, order.SuperPrice, fmt.Sprintf("%s(%s ID:%d)获得佣金", super_name, puser.Nickname, puser.ID), int(order.ID), tx)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

// TODO: 处理订单积分发放
func handelOrderPoints(order *model.Order, tx *gorm.DB) error {
	pointsMoneyRatioStr, _ := set.GetSet("points_money_ratio", true)
	fmt.Println(pointsMoneyRatioStr)
	if pointsMoneyRatioStr == "" {
		return nil
	}
	pointsMoneyRatio, _ := strconv.ParseFloat(pointsMoneyRatioStr, 64)
	if pointsMoneyRatio <= 0 {
		return nil
	}
	pointsFloat := math.Floor(order.PayPrice / pointsMoneyRatio)
	if pointsFloat == 0 {
		return nil
	}
	points := int(pointsFloat)
	var user model.User
	if err := tx.Where("id=?", order.Uid).First(&user).Error; err != nil {
		return err
	}
	if err := tx.Create(&model.UserPoints{
		Uid:    order.Uid,
		Points: points,
		Type:   4,
		Mode:   1,
		After:  user.Points + points,
		Before: user.Points,
		Remark: "支付积分奖励",
	}).Error; err != nil {
		return err
	}
	if err := tx.Model(&user).UpdateColumn("points", gorm.Expr("points + ?", points)).Error; err != nil {
		return err
	}
	return nil
}
