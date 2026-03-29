/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package controller

import (
	"gin/service/set"
	"gin/service/upload"
	"gin/util"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func DelFile(c *gin.Context) {
	data := struct {
		Path string `json:"path" binding:"required" form:"path"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if data.Path == "" {
		util.Error(c, 1, "参数错误")
		return
	}
	msg, err := upload.DeleteFile(data.Path)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, msg, nil)
}
func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	if file == nil {
		util.Error(c, 1, "请选择文件")
		return
	}
	ext := filepath.Ext(file.Filename)
	upload_file_type, _ := set.GetSet("upload_file_type", true)
	if !strings.Contains(upload_file_type, ext) {
		util.Error(c, 1, "文件类型不正确")
		return
	}
	fileIo, err := file.Open()
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	defer fileIo.Close()
	filename, err := upload.UploadFile(fileIo, ext)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "上传成功", filename)
}
