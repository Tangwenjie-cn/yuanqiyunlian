/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin/service/set"
	"gin/util"
	"net/url"
)

type mini struct {
	AppId     string
	AppSecret string
}

func miniInit() *mini {
	appid, err := set.GetSet("wx_mini_appid", true)
	if err != nil {
		panic(err)
	}
	appsecret, err := set.GetSet("wx_mini_secret", true)
	if err != nil {
		panic(err)
	}
	options := mini{
		AppId:     appid,
		AppSecret: appsecret,
	}
	return &options
}
func getMiniStableAccessToken() string {
	options := miniInit()
	data := map[string]string{
		"grant_type": "client_credential",
		"appid":      options.AppId,
		"secret":     options.AppSecret,
	}

	reqData := make(map[string]any, 0)
	reqData["url"] = "https://api.weixin.qq.com/cgi-bin/stable_token"
	reqData["data"] = data
	reqData["method"] = "POST"
	resp, err := util.Request(reqData)
	if err != nil {
		panic(err)
	}
	resData := make(map[string]any)
	if err := json.Unmarshal(resp, &resData); err != nil {
		panic(err)
	}
	if resData["access_token"] == nil {
		panic(resData["errmsg"])
	}
	return resData["access_token"].(string)
}
func GetMiniCode2Session(code string) map[string]any {
	options := miniInit()
	params := url.Values{}
	params.Add("appid", options.AppId)
	params.Add("secret", options.AppSecret)
	params.Add("js_code", code)
	params.Add("grant_type", "authorization_code")
	reqData := make(map[string]any, 0)
	reqData["url"] = "https://api.weixin.qq.com/sns/jscode2session?" + params.Encode()
	reqData["method"] = "GET"
	resp, err := util.Request(reqData)
	if err != nil {
		panic(err)
	}
	resData := make(map[string]any)
	if err := json.Unmarshal(resp, &resData); err != nil {
		panic(err)
	}
	return resData
}
func GetMiniQRCode(page string, scene string) *bytes.Buffer {
	access_token := getMiniStableAccessToken()
	data := map[string]any{
		"page":        page,
		"scene":       scene,
		"check_path":  false,
		"env_version": "release", //正式版为 "release"，体验版为 "trial"，开发版为 "develop"
	}
	reqData := make(map[string]any, 0)
	reqData["url"] = "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=" + access_token
	reqData["data"] = data
	reqData["method"] = "POST"
	resp, err := util.Request(reqData)
	if err != nil {
		panic(fmt.Sprintf("请求失败：%v\n", err))
	}
	return bytes.NewBuffer(resp)
}
func SendTemplateMessage(openid string, templateId string, page string, data map[string]any) {
	access_token := getMiniStableAccessToken()
	bodyData := map[string]any{
		"touser":            openid,
		"template_id":       templateId,
		"page":              page,
		"data":              data,
		"miniprogram_state": "formal", // developer为开发版；trial为体验版；formal为正式版，默认为正式版
	}
	reqData := make(map[string]any, 0)
	reqData["url"] = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=" + access_token
	reqData["data"] = bodyData
	reqData["method"] = "POST"
	res, err := util.Request(reqData)
	fmt.Println(string(res))
	if err != nil {
		fmt.Println(err)
		return
	}
}
