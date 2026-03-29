/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package upload

import (
	"fmt"
	"gin/service/cloud"
	"gin/service/set"
	"gin/util"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func UploadFile(file io.Reader, ext string) (string, error) {
	upload_type, err := set.GetSet("upload_type", true)
	if err != nil {
		return "", err
	}
	switch upload_type {
	case "local": //本地存储
		filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		dst := filepath.Join("upload", time.Now().Format("20060102"), filename)
		exePath, err := util.RootPath()
		if err != nil {
			return "", err
		}
		newFilePath := filepath.Join(exePath, "public", dst)
		if err = os.MkdirAll(filepath.Dir(newFilePath), 0750); err != nil {
			return "", err
		}
		out, err := os.Create(newFilePath)
		if err != nil {
			return "", err
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			return "", err
		}
		domain, err := set.GetSet("site_url", true)
		if err != nil {
			return "", err
		}
		imageUrl := domain + "/static/" + dst

		return strings.ReplaceAll(imageUrl, "\\", "/"), nil
	case "aliyun":
		return cloud.UploadOss(file, upload_type, ext, nil)
	case "tencent":
		return cloud.UploadOss(file, upload_type, ext, nil)
	default:
		return "", fmt.Errorf("存储类型不存在")
	}
}
func DeleteFile(path string) (string, error) {
	upload_type, _ := set.GetSet("upload_type", true)
	parsedURL, err := url.Parse(path)
	if err != nil {
		return "", err
	}
	pathArr := strings.Split(parsedURL.Path, `/`)
	switch upload_type {
	case "local": //本地存储
		dir, err := util.RootPath()
		if err != nil {
			return "", err
		}
		dir = filepath.Join(dir, "public")
		for k, v := range pathArr {
			if k > 1 {
				dir = filepath.Join(dir, v)
			}
		}
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			return "", fmt.Errorf("文件不存在")
		}
		err = os.Remove(dir)
		if err != nil {
			return "", err
		}
		return "删除成功", nil
	case "aliyun":
		err := cloud.DeleteOss(path, upload_type)
		if err != nil {
			return "", err
		}
		return "删除成功", nil
	case "tencent":
		err := cloud.DeleteOss(path, upload_type)
		if err != nil {
			return "", err
		}
		return "删除成功", nil
	default:
		return "", fmt.Errorf("存储类型不存在")
	}
}
