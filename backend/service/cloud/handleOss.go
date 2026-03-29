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
	"io"
)

type myoss interface {
	// 上传文件
	upload(file io.Reader, fileName string, location *string) (string, error)
	delete(url string) error
}

func UploadOss(file io.Reader, upload_type string, ext string, location *string) (string, error) {
	switch upload_type {
	case "aliyun":
		var aliyunOss myoss = &aliyunOss{}
		return aliyunOss.upload(file, ext, location)
	case "tencent":
		var tencentOss myoss = &tencentOss{}
		return tencentOss.upload(file, ext, location)
	default:
		return "", nil
	}
}
func DeleteOss(url string, upload_type string) error {
	switch upload_type {
	case "aliyun":
		var aliyunOss myoss = &aliyunOss{}
		return aliyunOss.delete(url)
	case "tencent":
		var tencentOss myoss = &tencentOss{}
		return tencentOss.delete(url)
	default:
		return nil
	}
}
