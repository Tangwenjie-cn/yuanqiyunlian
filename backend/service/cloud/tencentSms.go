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
	"fmt"
	"gin/service/set"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type tencentSms struct {
}

func (t *tencentSms) sendSms(phone, code string) error {
	secretId, err := set.GetSet("tencent_oss_secretId", true)
	if err != nil {
		return err
	}
	secretKey, err := set.GetSet("tencent_oss_secretKey", true)
	if err != nil {
		return err
	}
	appid, err := set.GetSet("tencent_sms_appid", true)
	if err != nil {
		return err
	}
	signName, err := set.GetSet("tencent_sms_signName", true)
	if err != nil {
		return err
	}
	templateId, err := set.GetSet("tencent_sms_templateId", true)
	if err != nil {
		return err
	}
	credential := common.NewCredential(secretId, secretKey)
	cpf := profile.NewClientProfile()
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)
	request := sms.NewSendSmsRequest()
	request.SmsSdkAppId = common.StringPtr(appid)
	request.SignName = common.StringPtr(signName)
	request.TemplateId = common.StringPtr(templateId)
	request.TemplateParamSet = common.StringPtrs([]string{code})
	request.PhoneNumberSet = common.StringPtrs([]string{"+86" + phone})
	response, err := client.SendSms(request)
	if err != nil {
		return err
	}
	fmt.Println(response.ToJsonString())
	return nil
}
