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
	"gin/service/set"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v5/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
)

type aliyunSms struct {
}

func (a *aliyunSms) createClient() (*dysmsapi.Client, error) {
	accessKeyID, err := set.GetSet("aliyun_oss_accessKeyId", true)
	if err != nil {
		return nil, err
	}
	accessKeySecret, err := set.GetSet("aliyun_oss_accessKeySecret", true)
	if err != nil {
		return nil, err
	}
	credentialsConfig := new(credentials.Config).
		SetType("access_key").
		SetAccessKeyId(accessKeyID).
		SetAccessKeySecret(accessKeySecret)
	akCredential, err := credentials.NewCredential(credentialsConfig)
	if err != nil {
		return nil, err
	}

	config := &openapi.Config{}
	config.Credential = akCredential
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	return dysmsapi.NewClient(config)
}

func (a *aliyunSms) sendSms(phone, code string) error {
	client, err := a.createClient()
	if err != nil {
		return err
	}
	signName, _ := set.GetSet("aliyun_sms_signName", true)
	templateCode, _ := set.GetSet("aliyun_sms_templateCode", true)
	request := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		SignName:      tea.String(signName),
		TemplateCode:  tea.String(templateCode),
		TemplateParam: tea.String("{\"code\":\"" + code + "\"}"),
	}
	_, err = client.SendSms(request)
	if err != nil {
		return err
	}
	return nil
}
