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
	"gin/service/finance"
	"gin/service/pay"
	"gin/service/set"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/go-pay/xlog"
)

func AlipayNotify(c *gin.Context) {
	notifyReq, err := alipay.ParseNotifyToBodyMap(c.Request)
	if err != nil {
		xlog.Error(err)
		return
	}
	alipayPublicCert, err := set.GetSet("alipay_alipayPublicCert", true)
	if err != nil {
		xlog.Error(err)
		return
	}
	// 支付宝异步通知验签（公钥证书模式）
	_, err = alipay.VerifySignWithCert([]byte(alipayPublicCert), notifyReq)
	if err != nil {
		xlog.Error(err)
		return
	}
	// 业务处理
	tradeStatus := notifyReq.Get("trade_status")
	if tradeStatus != "TRADE_SUCCESS" {
		return
	}
	orderNo := notifyReq.Get("out_trade_no")
	thirdOrder := notifyReq.Get("trade_no")
	totalAmount := notifyReq.Get("total_amount")
	totalAmountFloat, err := strconv.ParseFloat(totalAmount, 64)
	if err != nil {
		xlog.Error(err)
		return
	}
	tx := config.DB.Begin()
	var order model.Order
	if err := config.DB.Where("order_no=?", orderNo).First(&order).Error; err != nil {
		tx.Rollback()
		return
	}
	if order.Status == 1 {
		tx.Rollback()
		c.String(http.StatusOK, "%s", "success")
		return
	}
	if order.PayPrice != totalAmountFloat {
		tx.Rollback()
		return
	}
	order.ThirdOrder = thirdOrder
	err = finance.HandleOrder(&order, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	c.String(http.StatusOK, "%s", "success")
}
func WxpayNotify(c *gin.Context) {
	notifyReq, err := wechat.V3ParseNotify(c.Request)
	if err != nil {
		xlog.Error(err)
		c.String(http.StatusInternalServerError, "服务器内部错误：%s", err.Error())
		return
	}
	wxpayConfig, err := pay.WxpayInit()
	if err != nil {
		xlog.Error(err)
		c.String(http.StatusInternalServerError, "服务器内部错误：%s", err.Error())
		return
	}
	client := wxpayConfig.ClientV3
	// 获取微信平台证书
	certMap := client.WxPublicKeyMap()
	// 验证异步通知的签名
	err = notifyReq.VerifySignByPKMap(certMap)
	if err != nil {
		xlog.Error(err)
		c.String(http.StatusInternalServerError, "服务器内部错误：%s", err.Error())
		return
	}
	switch notifyReq.EventType {
	case "TRANSACTION.SUCCESS":
		// 处理支付成功业务
		result, err := notifyReq.DecryptPayCipherText(wxpayConfig.ApiV3Key)
		if err != nil {
			xlog.Error(err)
			c.String(http.StatusInternalServerError, "服务器内部错误：%s", err.Error())
			return
		}
		WxpayHandle(result, c)
	case "MCHTRANSFER.BILL.FINISHED":
		// 处理转账通知
		result, err := notifyReq.DecryptTransferBillsNotifyCipherText(wxpayConfig.ApiV3Key)
		if err != nil {
			xlog.Error(err)
			c.String(http.StatusInternalServerError, "服务器内部错误：%s", err.Error())
			return
		}
		WxpayTransferBillHandle(result, c)
	default:
		c.String(http.StatusInternalServerError, "服务器内部错误：%s", "未知事件类型")
	}
}

// 微信支付处理
func WxpayHandle(result *wechat.V3DecryptPayResult, c *gin.Context) {
	if result.TradeState != "SUCCESS" {
		xlog.Error(result)
		c.String(http.StatusInternalServerError, "服务器内部错误：%s", result.TradeState)
		return
	}
	orderNo := result.OutTradeNo
	thirdOrder := result.TransactionId
	totalAmount := result.Amount.Total
	totalAmountFloat := float64(totalAmount) / 100
	tx := config.DB.Begin()
	var order model.Order
	if err := config.DB.Where("order_no=?", orderNo).First(&order).Error; err != nil {
		tx.Rollback()
		c.String(http.StatusInternalServerError, "服务器内部错误：%s", err.Error())
		return
	}
	if order.Status == 1 {
		tx.Rollback()
		c.JSON(http.StatusOK, &wechat.V3NotifyRsp{Code: gopay.SUCCESS, Message: "成功"})
		return
	}
	if order.PayPrice != totalAmountFloat {
		tx.Rollback()
		c.String(http.StatusInternalServerError, "服务器内部错误：%s", "金额不对")
		return
	}
	order.ThirdOrder = thirdOrder
	err := finance.HandleOrder(&order, tx)
	if err != nil {
		tx.Rollback()
		c.String(http.StatusInternalServerError, "服务器内部错误：%s", err.Error())
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, &wechat.V3NotifyRsp{Code: gopay.SUCCESS, Message: "成功"})
}
func WxpayTransferBillHandle(result *wechat.V3DecryptTransferBillsResult, c *gin.Context) {
	thirdOrder := result.TransferBillNo
	state := result.State
	tx := config.DB.Begin()
	var cash model.Cash
	if err := tx.Where("third_order=?", thirdOrder).First(&cash).Error; err != nil {
		tx.Rollback()
		c.String(http.StatusInternalServerError, "服务器内部错误：%s", err.Error())
		return
	}
	if cash.Status == 1 {
		tx.Rollback()
		c.JSON(http.StatusOK, &wechat.V3NotifyRsp{Code: gopay.SUCCESS, Message: "成功"})
		return
	}
	if state == "SUCCESS" {
		cash.Status = 1
	} else {
		cash.Status = -1
		cash.Remark = result.FailReason
		err := model.UserBill{}.Add(cash.Uid, 1, 3, cash.Money, "提现失败退回", int(cash.ID), tx)
		if err != nil {
			tx.Rollback()
			c.String(http.StatusInternalServerError, "服务器内部错误：%s", err.Error())
			return
		}
	}
	if err := tx.Save(&cash).Error; err != nil {
		tx.Rollback()
		c.String(http.StatusInternalServerError, "服务器内部错误：%s", err.Error())
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, &wechat.V3NotifyRsp{Code: gopay.SUCCESS, Message: "成功"})
}
