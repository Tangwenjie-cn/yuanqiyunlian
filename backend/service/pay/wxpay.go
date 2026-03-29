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
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/go-pay/xlog"
	"github.com/mitchellh/mapstructure"
)

type wxpayConfig struct {
	AppidMini   string //小程序appid
	AppidMp     string //公众号appid
	MchId       string
	ApiV3Key    string
	SerialNo    string
	PrivateKey  string
	PublicKey   string
	PublicKeyID string
	//NotifyUrl  string
	*wechat.ClientV3
}

func WxpayInit() (*wxpayConfig, error) {
	w := wxpayConfig{}
	appidMini, _ := set.GetSet("wx_mini_appid", true)
	w.AppidMini = appidMini

	appidMp, _ := set.GetSet("wx_mp_appid", true)
	w.AppidMp = appidMp

	if appidMini == "" && appidMp == "" {
		return nil, errors.New("wx_mp_appid or wx_mini_appid not found")
	}

	mchId, _ := set.GetSet("wxpay_mchId", true)
	if mchId == "" {
		return nil, errors.New("wxpay_mchId not found")
	}
	w.MchId = mchId

	apiV3Key, _ := set.GetSet("wxpay_apiV3Key", true)
	if apiV3Key == "" {
		return nil, errors.New("wxpay_apiV3Key not found")
	}
	w.ApiV3Key = apiV3Key

	privateKey, _ := set.GetSet("wxpay_privateKey", true)
	if privateKey == "" {
		return nil, errors.New("wxpay_privateKey not found")
	}
	w.PrivateKey = privateKey
	serialNo, _ := set.GetSet("wxpay_serialNo", true)
	if serialNo == "" {
		return nil, errors.New("wxpay_serialNo not found")
	}
	w.SerialNo = serialNo
	publicKey, _ := set.GetSet("wxpay_publicKey", true)
	if publicKey == "" {
		return nil, errors.New("wxpay_publicKey not found")
	}
	w.PublicKey = publicKey
	publicKeyId, _ := set.GetSet("wxpay_publicKeyId", true)
	if publicKeyId == "" {
		return nil, errors.New("wxpay_publicKeyId not found")
	}
	w.PublicKeyID = publicKeyId
	client, err := wechat.NewClientV3(w.MchId, w.SerialNo, w.ApiV3Key, w.PrivateKey)
	if err != nil {
		xlog.Error(err)
		return nil, err
	}
	err = client.AutoVerifySignByPublicKey([]byte(w.PublicKey), w.PublicKeyID)
	if err != nil {
		xlog.Error(err)
		return nil, err
	}
	w.ClientV3 = client
	return &w, nil
}
func WxpayJsapi(description string, outTradeNo string, totalFee float64, openId string, platform string) (map[string]any, error) {
	wxpayConfig, err := WxpayInit()
	if err != nil {
		xlog.Error(err)
		return nil, err
	}
	bm := make(gopay.BodyMap)
	switch platform {
	case "mp":
		bm.Set("appid", wxpayConfig.AppidMp)
	case "mini":
		bm.Set("appid", wxpayConfig.AppidMini)
	default:
		return nil, errors.New("platform not found")
	}
	siteUrl, _ := set.GetSet("site_url", true)
	if siteUrl == "" {
		return nil, errors.New("site_url not found")
	}
	notifyUrl := siteUrl + "/pay/wxpay/notify"
	bm.Set("mchid", wxpayConfig.MchId)
	bm.Set("description", description)
	bm.Set("out_trade_no", outTradeNo)
	bm.Set("notify_url", notifyUrl)
	bm.Set("amount", gopay.BodyMap{
		"total":    int64(totalFee * 100),
		"currency": "CNY",
	})
	bm.Set("payer", gopay.BodyMap{
		"openid": openId,
	})
	client := wxpayConfig.ClientV3
	wxRsp, err := client.V3TransactionJsapi(context.Background(), bm)
	if err != nil {
		xlog.Error(err)
		return nil, err
	}
	if wxRsp.Code != 0 {
		xlog.Errorf("wxRsp:%s", wxRsp.Error)
		return nil, errors.New(wxRsp.Error)
	}
	xlog.Debugf("wxRsp: %#v", wxRsp.Response)
	var paySign map[string]any
	switch platform {
	case "mp":
		jsapi, err := client.PaySignOfJSAPI(wxpayConfig.AppidMp, wxRsp.Response.PrepayId)
		if err != nil {
			xlog.Error(err)
			return nil, err
		}
		err = mapstructure.Decode(jsapi, &paySign)
		if err != nil {
			xlog.Error(err)
			return nil, err
		}
	case "mini":
		applet, err := client.PaySignOfApplet(wxpayConfig.AppidMini, wxRsp.Response.PrepayId)
		if err != nil {
			xlog.Error(err)
			return nil, err
		}
		err = mapstructure.Decode(applet, &paySign)
		if err != nil {
			xlog.Error(err)
			return nil, err
		}
	}
	return paySign, nil
}

func WxpayTransfer(transferAmount float64, transferRemark string, openid string, platform string) (map[string]string, error) {
	wxpayConfig, err := WxpayInit()
	if err != nil {
		xlog.Error(err)
		return nil, err
	}
	bm := make(gopay.BodyMap)
	switch platform {
	case "mp":
		bm.Set("appid", wxpayConfig.AppidMp)
	case "mini":
		bm.Set("appid", wxpayConfig.AppidMini)
	default:
		return nil, errors.New("platform not found")
	}
	siteUrl, _ := set.GetSet("site_url", true)
	if siteUrl == "" {
		return nil, errors.New("site_url not found")
	}
	notifyUrl := siteUrl + "/pay/wxpay/notify"
	bm.Set("out_bill_no", util.GenerateOrderNo())
	bm.Set("transfer_scene_id", "1000")
	bm.Set("openid", openid)
	bm.Set("transfer_amount", transferAmount)
	bm.Set("transfer_remark", transferRemark)
	bm.Set("notify_url", notifyUrl)
	bm.Set("transfer_scene_report_infos", []map[string]string{
		{
			"info_type":    "活动名称",
			"info_content": "推广奖励",
		},
		{
			"info_type":    "奖励说明",
			"info_content": "推广奖励提现",
		},
	})
	client := wxpayConfig.ClientV3
	wxRsp, err := client.V3TransferBills(context.Background(), bm)
	if err != nil {
		xlog.Error(err)
		return nil, err
	}
	if wxRsp.Code != 0 {
		xlog.Errorf("wxRsp:%s", wxRsp.Error)
		return nil, errors.New(wxRsp.Error)
	}
	if wxRsp.Response.State != "WAIT_USER_CONFIRM" {
		return nil, errors.New("调用失败，建议前往商户平台查看详情，以免资金损失")
	}
	return map[string]string{
		"transfer_bill_no": wxRsp.Response.TransferBillNo,
		"package_info":     wxRsp.Response.PackageInfo,
	}, nil
}
