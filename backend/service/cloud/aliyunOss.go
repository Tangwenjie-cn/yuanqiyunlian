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
	"log"
	"net/url"
	"time"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/google/uuid"
)

type aliyunOss struct {
	bucketName string
	*oss.Client
}

func (a *aliyunOss) aliyunOssInit() (*aliyunOss, error) {
	// 请根据实际要求设置region，以实例华东1（杭州）为例，regionID为cn-hangzhou
	region, err := set.GetSet("aliyun_oss_region", true)
	if err != nil {
		return nil, err
	}
	// 填写RAM用户的Access Key ID和Access Key Secret
	accessKeyID, err := set.GetSet("aliyun_oss_accessKeyId", true)
	if err != nil {
		return nil, err
	}
	accessKeySecret, err := set.GetSet("aliyun_oss_accessKeySecret", true)
	if err != nil {
		return nil, err
	}
	bucketName, err := set.GetSet("aliyun_oss_bucket", true)
	if err != nil {
		return nil, err
	}
	a.bucketName = bucketName
	provider := credentials.NewStaticCredentialsProvider(accessKeyID, accessKeySecret)
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(provider).
		WithRegion(region)
	client := oss.NewClient(cfg)
	return &aliyunOss{bucketName: bucketName, Client: client}, nil
}

func (a *aliyunOss) upload(file io.Reader, ext string, location *string) (string, error) {
	client, err := a.aliyunOssInit()
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
	// 创建上传对象的请求
	putRequest := &oss.PutObjectRequest{
		Bucket: oss.Ptr(client.bucketName), // 存储空间名称
		Key:    oss.Ptr(objectName),        // 对象名称
		Body:   file,
	}
	// 执行上传对象的请求
	result, err := client.PutObject(context.TODO(), putRequest)
	if err != nil {
		return "", err
	}
	// 打印上传对象的结果
	log.Printf("upload object result:%#v\n", result)
	url, err := set.GetSet("aliyun_oss_url", true)
	if err != nil {
		panic(err)
	}
	return url + "/" + objectName, nil
}
func (a *aliyunOss) delete(rawURL string) error {
	client, err := a.aliyunOssInit()
	if err != nil {
		return err
	}
	// 解析URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return err
	}
	objectName := parsedURL.Path[1:] // 获取对象名称，去掉前面的斜杠
	// 创建删除对象的请求
	request := &oss.DeleteObjectRequest{
		Bucket: oss.Ptr(client.bucketName), // 存储空间名称
		Key:    oss.Ptr(objectName),        // 对象名称
	}
	// 执行删除对象的操作并处理结果
	result, err := client.DeleteObject(context.TODO(), request)
	if err != nil {
		return err
	}

	// 打印删除对象的结果
	log.Printf("delete object result:%#v\n", result)
	return nil
}
