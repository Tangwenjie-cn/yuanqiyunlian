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
	"fmt"
	"gin/config"
	"gin/model"
	"gin/service/set"
	"gin/util"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func SetList(c *gin.Context) {
	configList := []model.Config{}
	data := struct {
		Group int8 `form:"group" json:"group" binding:"required"`
		Page  int  `form:"page" json:"page" binding:"required"`
		Limit int  `form:"limit" json:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	config.DB.Model(&model.Config{}).
		Where("`group` = ?", data.Group).
		Offset((data.Page - 1) * data.Limit).
		Limit(data.Limit).Order("sort asc").
		Find(&configList)
	var total int64
	config.DB.Model(&model.Config{}).
		Where("`group` = ?", data.Group).
		Count(&total)
	res := map[string]any{
		"data":  configList,
		"total": total,
	}
	util.Success(c, "Login successful", res)
}
func SetGroupList(c *gin.Context) {
	configModel := model.Config{}
	config.DB.Model(&model.Config{}).Where("id = 1").
		Find(&configModel)
	groups := strings.Split(configModel.Param, "\n")
	list := []map[string]any{}
	for _, v := range groups {
		value := strings.Split(strings.TrimSpace(v), "=")
		list = append(list, map[string]any{
			"name":  value[0],
			"value": value[1],
		})
	}
	util.Success(c, "success", list)
}
func AddSet(c *gin.Context) {
	data := struct {
		Title  string `form:"title" json:"title" binding:"required"`
		Key    string `form:"key" json:"key" binding:"required"`
		Param  string `form:"param" json:"param"`
		Value  string `form:"value" json:"value"`
		Group  int8   `form:"group" json:"group" binding:"required"`
		Type   int8   `form:"type" json:"type" binding:"required"`
		Sort   int16  `form:"sort" json:"sort"`
		Remark string `form:"remark" json:"remark"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	configModel := model.Config{}
	config.DB.Model(&model.Config{}).Where("`key` = ?", data.Key).
		Find(&configModel)
	if configModel.ID > 0 {
		util.Error(c, 1, "参数名已存在")
		return
	}
	var configModel1 model.Config
	if err := mapstructure.Decode(data, &configModel1); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	config.DB.Model(&model.Config{}).Create(&configModel1)
	if configModel1.ID > 0 {
		util.Success(c, "添加成功", nil)
	} else {
		util.Error(c, 1, "添加失败")
	}
}
func GetSet(c *gin.Context) {
	data := struct {
		ID uint `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	configModel := model.Config{}
	config.DB.Model(&model.Config{}).Where("id = ?", data.ID).Find(&configModel)
	util.Success(c, "success", configModel)
}
func EditSet(c *gin.Context) {
	data := struct {
		ID     uint   `form:"id" json:"id" binding:"required"`
		Title  string `form:"title" json:"title" binding:"required"`
		Key    string `form:"key" json:"key" binding:"required"`
		Param  string `form:"param" json:"param"`
		Value  string `form:"value" json:"value"`
		Group  int8   `form:"group" json:"group" binding:"required"`
		Type   int8   `form:"type" json:"type" binding:"required"`
		Sort   int16  `form:"sort" json:"sort"`
		Remark string `form:"remark" json:"remark"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	configModel := model.Config{}
	config.DB.Model(&model.Config{}).Where("`key` = ? and id != ?", data.Key, data.ID).
		Find(&configModel)
	if configModel.ID > 0 {
		util.Error(c, 1, "参数名已存在")
		return
	}
	var dataMap map[string]any
	if err := mapstructure.Decode(data, &dataMap); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	res := config.DB.Model(&model.Config{}).Where("id = ?", data.ID).Updates(&dataMap)
	if res.Error != nil {
		util.Error(c, 1, "编辑失败")
		return
	}
	util.Success(c, "编辑成功", nil)
}
func DelSet(c *gin.Context) {
	data := struct {
		ID uint `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	configModel := model.Config{}
	res := config.DB.Model(&model.Config{}).Where("id = ?", data.ID).
		Delete(&configModel)
	if res.RowsAffected == 0 {
		util.Error(c, 1, "删除失败")
		return
	}
	util.Success(c, "删除成功", nil)
}
func GetSetAllInfo(c *gin.Context) {
	list := []model.Config{}
	config.DB.Where("id > 1").Order("`group` asc, sort asc").Find(&list)
	data := map[int8]any{}
	for _, v := range list {
		item := []model.Config{}
		if _, ok := data[v.Group]; ok {
			item = data[v.Group].([]model.Config)
		} else {
			data[v.Group] = []model.Config{}
		}
		switch v.Type {
		case 3, 5, 8:
			arr := strings.Split(v.Param, "\n")
			for _, v1 := range arr {
				value := strings.Split(strings.TrimSpace(v1), "=")
				v.List = append(v.List, map[string]any{
					"name":  value[0],
					"value": value[1],
				})

			}
		}
		item = append(item, v)
		data[v.Group] = item
	}
	util.Success(c, "获取成功", data)
}
func SaveAllSet(c *gin.Context) {
	data := map[string][]model.Config{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}

	for _, v := range data {
		for _, v1 := range v {
			if err := config.DB.Model(&model.Config{}).Where("id = ?", v1.ID).Update("value", strings.TrimSpace(v1.Value)).Error; err != nil {
				util.Error(c, 1, err.Error())
				return
			}
		}

	}
	util.Success(c, "保存成功", nil)
}
func GetUpdateInfo(c *gin.Context) {

}
func Update(c *gin.Context) {

}
func WxMiniUpload(c *gin.Context) {
	data := struct {
		Version string `form:"version" json:"version" binding:"required"`
		Desc    string `form:"desc" json:"desc" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	rootPath, err := util.RootPath()
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	projectPath := filepath.Join(rootPath, "foreEnd", "mp-weixin")
	domain, _ := set.GetSet("site_url", true)
	if domain == "" {
		util.Error(c, 1, "请先配置网站域名")
		return
	}
	jsStr := `"use strict"; exports.config={"url":"` + domain + `"};`
	err = os.WriteFile(filepath.Join(projectPath, "config.js"), []byte(jsStr), 0755)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	privateKeyPath := filepath.Join(rootPath, "wxUploadPrivateKey.key")
	privateKeyContent, _ := set.GetSet("wx_mini_upload_private_key", true)
	if privateKeyContent == "" {
		util.Error(c, 1, "请先配置小程序上传私钥")
		return
	}
	err = os.WriteFile(privateKeyPath, []byte(privateKeyContent), 0755)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	appid, _ := set.GetSet("wx_mini_appid", true)
	if appid == "" {
		util.Error(c, 1, "请先配置小程序appid")
		return
	}
	args := []string{
		"exec",
		"--",
		"miniprogram-ci",
		"upload",
		"--pp",
	}
	args = append(args, projectPath)
	args = append(args, "--pkp")
	args = append(args, privateKeyPath)
	args = append(args, "--appid")
	args = append(args, appid)
	args = append(args, "--uv")
	args = append(args, data.Version)
	args = append(args, "--ud")
	args = append(args, data.Desc)
	cmd := exec.Command("npm", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "success", nil)
}
func AdminLogList(c *gin.Context) {
	data := struct {
		Keyword string `form:"keyword" json:"keyword"`
		Page    int    `form:"page" json:"page" binding:"required"`
		Limit   int    `form:"limit" json:"limit" binding:"required"`
		Date    string `form:"date" json:"date"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var total int64
	query := config.DB.Model(&model.AdminLog{})
	if data.Keyword != "" {
		query = query.Where("account = ?", strings.TrimSpace(data.Keyword))
	}
	if data.Date != "" {
		dateArr := strings.Split(data.Date, ",")
		query = query.Where("ctime between ? and ?", dateArr[0], dateArr[1])
	}
	query.Count(&total)
	query = query.Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Order("id desc")
	var list []model.AdminLog
	query.Find(&list)
	res := map[string]any{
		"data":  list,
		"total": total,
	}
	util.Success(c, "获取列表成功", res)
}
