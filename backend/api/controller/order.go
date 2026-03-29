/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package controller

import (
	"gin/config"
	"gin/model"
	"gin/service/pay"
	"gin/util"
	"math"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func InitOrder(c *gin.Context) {
	data := struct {
		Id   int `json:"id" form:"id" binding:"required"`
		Type int `json:"type" form:"type" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	uid, _ := c.MustGet("uid").(int)
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	list := make(map[string]any)
	type discount struct {
		Price  float64 `json:"price"`
		Rebate float64 `json:"rebate"`
		Type   int     `json:"type"`
		Name   string  `json:"name"`
		Id     int     `json:"id"`
	}
	var discountList []discount
	switch data.Type {
	case 1:
		//商品订单
		var goods model.Goods
		if err := config.DB.Where("id=? and status=1", data.Id).First(&goods).Error; err != nil {
			util.Error(c, 1, "商品不存在")
			return
		}
		goods.Link = ""
		list["info"] = goods
		list["price"] = goods.Price
		list["id"] = goods.ID
	case 2:
		//会员订单
		var svip model.Svip
		if err := config.DB.Where("id=? and status=1", data.Id).First(&svip).Error; err != nil {
			util.Error(c, 1, "会员不存在")
			return
		}
		list["info"] = svip
		list["price"] = svip.Price
		list["id"] = svip.ID
	default:
		util.Error(c, 1, "类型错误")
		return
	}
	type userSvip struct {
		model.UserSvip
		Rebate float64 `json:"rebate" gorm:"column:rebate"`
		Title  string  `json:"title" gorm:"column:title"`
	}
	// 查询用户会员
	var vipList []userSvip
	if err := config.DB.Model(&model.UserSvip{}).
		Select("user_svip.*", "svip.title", "svip.rebate", "svip.price").
		Joins("inner join svip on svip.id=user_svip.vid").
		Where("user_svip.uid = ? and user_svip.expire_time > ? and svip.rebate>0", uid, nowTime).
		Order("user_svip.expire_time desc").Find(&vipList).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	for _, vip := range vipList {
		roughPrice := (1 - vip.Rebate/10) * list["price"].(float64)
		price := math.Round(roughPrice*100) / 100 //四舍五入保留两位小数
		discountList = append(discountList, discount{
			Price:  price,
			Rebate: vip.Rebate,
			Type:   2,
			Name:   vip.Title,
			Id:     int(vip.ID),
		})
	}
	list["discountList"] = discountList
	list["type"] = data.Type
	util.Success(c, "成功", list)
}
func CreateOrder(c *gin.Context) {
	data := struct {
		Id           int    `json:"id" form:"id" binding:"required"`
		Type         int    `json:"type" form:"type" binding:"required"`
		PayType      int    `json:"pay_type" form:"pay_type" binding:"required"`
		Platform     string `json:"platform" form:"platform" binding:"required"`
		Key          string `json:"key" form:"key"`
		DiscountType int    `json:"discount_type" form:"discount_type"`
		DiscountId   int    `json:"discount_id" form:"discount_id"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	uid, _ := c.MustGet("uid").(int)
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	var order model.Order
	order.Type = data.Type
	order.PayType = data.PayType
	order.Uid = uid
	order.OrderNo = util.GenerateOrderNo() //生成订单号
	// 开启事务
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var title string
	if order.Type == 1 {
		//商品订单
		var goods model.Goods
		if err := tx.Where("id=? and status=1", data.Id).First(&goods).Error; err != nil {
			util.Error(c, 1, "商品不存在")
			return
		}
		order.Gid = data.Id
		order.Price = goods.Price
		title = goods.Title
	}
	if order.Type == 2 {
		//会员订单
		var vip model.Svip
		if err := tx.Where("id=? and status=1", data.Id).First(&vip).Error; err != nil {
			util.Error(c, 1, "会员不存在")
			return
		}
		order.Vid = data.Id
		order.Price = vip.Price
		title = vip.Title
	}
	switch data.DiscountType {
	case 1:
		//优惠券
	case 2:
		//会员折扣
		var userSvip model.UserSvip
		if err := tx.Where("vid=? and uid=? and expire_time > ?", data.DiscountId, uid, nowTime).First(&userSvip).Error; err != nil {
			util.Error(c, 1, "会员折扣不存在")
			return
		}
		var svip model.Svip
		if err := tx.Where("id=? and status=1", userSvip.Vid).First(&svip).Error; err != nil {
			util.Error(c, 1, "会员不存在")
			return
		}
		roughPrice := (1 - svip.Rebate/10) * order.Price
		order.DiscountPrice = math.Round(roughPrice*100) / 100
		order.DiscountType = 2
	default:
		order.DiscountPrice = 0
		order.DiscountType = 0
	}
	order.PayPrice = order.Price - order.DiscountPrice
	if order.PayPrice <= 0 {
		util.Error(c, 1, "订单金额不能小于等于0")
		return
	}
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		util.Error(c, 1, err.Error())
		return
	}
	switch order.PayType {
	case 1:
		//微信支付
		if data.Platform == "h5" {
			userAgent := c.Request.Header.Get("User-Agent")
			if strings.Contains(userAgent, "MicroMessenger") {
				var userBind model.UserBind
				if err := tx.Where("uid=? and type=?", uid, 3).First(&userBind).Error; err != nil {
					tx.Rollback()
					util.Error(c, 1, "请先绑定微信")
					return
				}
				paySign, err := pay.WxpayJsapi(title, order.OrderNo, order.PayPrice, userBind.Openid, "mp")
				if err != nil {
					tx.Rollback()
					util.Error(c, 1, err.Error())
					return
				}
				tx.Commit()
				util.Success(c, "成功", paySign)
				return
			}
		}
		if data.Platform == "weixin" {
			var userBind model.UserBind
			if err := tx.Where("uid=? and type=?", uid, 2).First(&userBind).Error; err != nil {
				tx.Rollback()
				util.Error(c, 1, "请先绑定微信")
				return
			}
			paySign, err := pay.WxpayJsapi(title, order.OrderNo, order.PayPrice, userBind.Openid, "mini")
			if err != nil {
				tx.Rollback()
				util.Error(c, 1, err.Error())
				return
			}
			tx.Commit()
			util.Success(c, "成功", paySign)
		}
	case 2:
		//支付宝支付
		if data.Platform == "h5" {
			//h5支付
			payUrl, err := pay.AliapyWapPay(title, order.PayPrice, order.OrderNo)
			if err != nil {
				tx.Rollback()
				util.Error(c, 1, err.Error())
				return
			}
			tx.Commit()
			util.Success(c, "成功", payUrl)
		}
	case 3:
		//卡密支付
		if data.Key == "" {
			tx.Rollback()
			util.Error(c, 1, "卡密不能为空")
			return
		}
		order.ThirdOrder = strings.TrimSpace(data.Key)
		err := pay.HandleKami(&order, tx)
		if err != nil {
			tx.Rollback()
			util.Error(c, 1, err.Error())
			return
		}
		tx.Commit()
		util.Success(c, "支付成功", nil)
	case 4:
		//余额支付
		err := pay.HandleBalance(&order, tx)
		if err != nil {
			tx.Rollback()
			util.Error(c, 1, err.Error())
			return
		}
		tx.Commit()
		util.Success(c, "支付成功", nil)
	default:
		tx.Rollback()
		util.Error(c, 1, "支付方式错误")
		return
	}
}
func OrderList(c *gin.Context) {
	data := struct {
		Type  int `json:"type" form:"type" binding:"required"`
		Page  int `json:"page" form:"page" binding:"required"`
		Limit int `json:"limit" form:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	uid, _ := c.MustGet("uid").(int)
	list := []struct {
		model.Order
		Title string `json:"title" gorm:"column:title"`
		Thumb string `json:"thumb" gorm:"column:thumb"`
	}{}
	query := config.DB.Model(&model.Order{}).Where("`order`.uid=? and `order`.type=?", uid, data.Type)
	if data.Type == 1 {
		query = query.Joins("inner join goods on goods.id=`order`.gid").Select("`order`.*", "goods.title", "goods.thumb")
	}
	if data.Type == 2 {
		query = query.Joins("inner join svip on svip.id=`order`.vid").Select("`order`.*", "svip.title")
	}
	if err := query.Order("id desc").Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Find(&list).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "成功", list)
}
func OrderInfo(c *gin.Context) {
	data := struct {
		Id int `json:"id" form:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	uid, _ := c.MustGet("uid").(int)
	var order model.Order
	if err := config.DB.Where("id=? and uid=?", data.Id, uid).First(&order).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var m map[string]any
	decoderConfig := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &m,
		TagName:  "json", // 使用 json 标签
	}

	decoder, _ := mapstructure.NewDecoder(decoderConfig)
	decoder.Decode(&order)
	if order.Type == 1 {
		var goods model.Goods
		if err := config.DB.Where("id=?", order.Gid).First(&goods).Error; err != nil {
			util.Error(c, 1, err.Error())
			return
		}
		m["title"] = goods.Title
		m["thumb"] = goods.Thumb
	}
	if order.Type == 2 {
		var vip model.Svip
		if err := config.DB.Where("id=?", order.Vid).First(&vip).Error; err != nil {
			util.Error(c, 1, err.Error())
			return
		}
		m["title"] = vip.Title
	}
	m["ctime_text"] = order.Ctime.Format("2006-01-02 15:04:05")
	util.Success(c, "成功", m)
}
