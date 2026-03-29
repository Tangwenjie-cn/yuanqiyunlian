/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package cloud

import (
	"errors"
	"gin/service/set"
)

type mySms interface {
	sendSms(phone, code string) error
}

func SendSms(phone, code string) error {
	smsType, _ := set.GetSet("sms_type", true)
	switch smsType {
	case "1":
		var aliyunSms mySms = &aliyunSms{}
		return aliyunSms.sendSms(phone, code)
	case "2":
		var tencentSms mySms = &tencentSms{}
		return tencentSms.sendSms(phone, code)
	default:
		return errors.New("短信服务商类型错误")
	}
}
