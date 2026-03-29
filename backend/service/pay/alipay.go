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
	"context"
	"errors"
	"gin/service/set"
	"gin/util"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/xlog"
)

type alipayConfig struct {
	AppId      string
	PrivateKey string
	//NotifyUrl        string
	AlipayRootCert   string
	AlipayPublicCert string
	AppPublicCert    string
	*alipay.Client
}

func alipayInit() (*alipayConfig, error) {
	var a alipayConfig
	appid, _ := set.GetSet("alipay_appid", true)
	if appid == "" {
		return nil, errors.New("alipay_appid not found")
	}
	a.AppId = appid
	privateKey, _ := set.GetSet("alipay_privateKey", true)
	if privateKey == "" {
		return nil, errors.New("alipay_privateKey not found")
	}
	a.PrivateKey = privateKey
	// 传入证书内容
	appPublicCert, _ := set.GetSet("alipay_appPublicCert", true)
	if appPublicCert == "" {
		return nil, errors.New("alipay_appPublicCert not found")
	}
	a.AppPublicCert = appPublicCert
	alipayRootCert, _ := set.GetSet("alipay_alipayRootCert", true)
	if alipayRootCert == "" {
		return nil, errors.New("alipay_alipayRootCert not found")
	}
	a.AlipayRootCert = alipayRootCert
	alipayPublicCert, _ := set.GetSet("alipay_alipayPublicCert", true)
	if alipayPublicCert == "" {
		return nil, errors.New("alipay_alipayPublicCert not found")
	}
	a.AlipayPublicCert = alipayPublicCert
	client, err := alipay.NewClient(a.AppId, a.PrivateKey, true)
	if err != nil {
		xlog.Error(err)
		return nil, err
	}
	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	err = client.SetCertSnByContent([]byte(a.AppPublicCert), []byte(a.AlipayRootCert), []byte(a.AlipayPublicCert))
	if err != nil {
		xlog.Error(err)
		return nil, err
	}
	a.Client = client
	return &a, nil
}
func AliapyWapPay(subject string, totalAmount float64, orderNo string) (string, error) {
	alipayConfig, err := alipayInit()
	if err != nil {
		return "", err
	}
	siteUrl, _ := set.GetSet("site_url", true)
	if siteUrl == "" {
		return "", errors.New("site_url not found")
	}
	client := alipayConfig.Client
	client.NotifyUrl = siteUrl + "/pay/alipay/notify"
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", orderNo)     // 商户订单号
	bm.Set("total_amount", totalAmount) // 订单总金额，单位为元，精确到小数点后两位
	bm.Set("subject", subject)          // 订单标题
	payUrl, err := client.TradeWapPay(context.Background(), bm)
	return payUrl, err
}

// 商家转账
func AliapyTransfer(transAmount float64, title string, realName string, account string) (string, error) {
	alipayConfig, err := alipayInit()
	if err != nil {
		return "", err
	}
	client := alipayConfig.Client
	bm := make(gopay.BodyMap)
	bm.Set("out_biz_no", util.GenerateOrderNo())
	bm.Set("trans_amount", transAmount)
	bm.Set("product_code", "TRANS_ACCOUNT_NO_PWD")
	bm.Set("biz_scene", "DIRECT_TRANSFER")
	bm.Set("order_title", title)
	bm.Set("payee_info", map[string]any{
		"identity_type": "ALIPAY_USER_ID",
		"identity":      realName,
		"account":       account,
	})
	aliRsp, err := client.FundTransUniTransfer(context.Background(), bm)
	if err != nil {
		return "", err
	}
	if aliRsp.Response.Status != "SUCCESS" {
		return "", errors.New(aliRsp.Response.SubMsg)
	}
	return aliRsp.Response.OrderId, nil
}
