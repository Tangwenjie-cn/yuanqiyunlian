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
	"context"
	"gin/service/set"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type tencentOss struct {
	*cos.Client
	bucketUrl string
}

func (t *tencentOss) tencentOssInit() (tencentOss, error) {
	bucketUrl, err := set.GetSet("tencent_oss_backetUrl", true)
	if err != nil {
		return tencentOss{}, err
	}
	secretId, err := set.GetSet("tencent_oss_secretId", true)
	if err != nil {
		return tencentOss{}, err
	}
	secretKey, err := set.GetSet("tencent_oss_secretKey", true)
	if err != nil {
		return tencentOss{}, err
	}
	// 初始化
	u, _ := url.Parse(bucketUrl)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretId,
			SecretKey: secretKey,
		},
	})
	return tencentOss{c, bucketUrl}, nil
}
func (a *tencentOss) upload(file io.Reader, ext string, location *string) (string, error) {

	oss, err := a.tencentOssInit()
	if err != nil {
		return "", err
	}
	var objectName string
	now := time.Now()
	id := uuid.New()
	if location != nil {
		objectName = *location
	} else {
		objectName = now.Format("20060102") + "/" + id.String() + ext
	}
	_, err = oss.Object.Put(context.Background(), objectName, file, nil)
	if err != nil {
		return "", err
	}
	domain, _ := set.GetSet("tencent_oss_url", true)

	if domain == "" {
		return oss.bucketUrl + "/" + objectName, nil
	}
	return domain + "/" + objectName, nil
}
func (a *tencentOss) delete(rawURL string) error {
	oss, err := a.tencentOssInit()
	if err != nil {
		return err
	}
	// 解析URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return err
	}
	objectName := parsedURL.Path[1:] // 获取对象名称，去掉前面的斜杠
	_, err = oss.Object.Delete(context.Background(), objectName)
	return err
}
