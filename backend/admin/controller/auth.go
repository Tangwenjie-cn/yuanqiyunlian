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
	"bytes"
	"fmt"
	"gin/config"
	"gin/model"
	"gin/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// AuthList 获取权限列表
func AuthList(c *gin.Context) {
	// 获取权限列表
	auths, err := model.Auth{}.GetTree(false, 0)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取权限列表成功", auths)
}

// AddAuth 添加权限
func AddAuth(c *gin.Context) {
	data := struct {
		Name   string `form:"name" json:"name" binding:"required"`
		Pid    uint   `form:"pid" json:"pid" binding:"gte=0"`
		Url    string `form:"url" json:"url"`
		Path   string `form:"path" json:"path"`
		Icon   string `form:"icon" json:"icon"`
		IsMenu uint8  `form:"is_menu" json:"is_menu"`
		Sort   uint   `form:"sort" json:"sort"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var auth model.Auth
	if err := mapstructure.Decode(data, &auth); err != nil {
		util.Error(c, 1, err.Error())
		return
	}

	res := config.DB.Model(&model.Auth{}).Create(&auth)
	if res.Error != nil {
		util.Error(c, 1, res.Error.Error())
		return
	}
	util.Success(c, "添加权限成功", nil)
}
func GetAuth(c *gin.Context) {
	data := struct {
		ID uint `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var auth model.Auth
	config.DB.Model(&model.Auth{}).Where("id = ?", data.ID).First(&auth)
	util.Success(c, "获取权限成功", auth)
}

// EditAuth 编辑权限
func EditAuth(c *gin.Context) {
	data := struct {
		ID     uint   `form:"id" json:"id" binding:"required"`
		Name   string `form:"name" json:"name" binding:"required"`
		Pid    uint   `form:"pid" json:"pid" binding:"gte=0"`
		Url    string `form:"url" json:"url"`
		Path   string `form:"path" json:"path"`
		Icon   string `form:"icon" json:"icon"`
		IsMenu uint8  `form:"is_menu" json:"is_menu"`
		Sort   uint   `form:"sort" json:"sort"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var authMap map[string]any
	if err := mapstructure.Decode(data, &authMap); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	res := config.DB.Model(&model.Auth{}).Where("id = ?", data.ID).Updates(&authMap)
	if res.Error != nil {
		util.Error(c, 1, res.Error.Error())
		return
	}
	util.Success(c, "编辑权限成功", nil)
}

// DelAuth 删除权限
func DelAuth(c *gin.Context) {
	data := struct {
		ID uint `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var auth model.Auth
	if err := config.DB.Where("id = ?", data.ID).First(&auth).Error; err != nil {
		util.Error(c, 1, "删除权限失败，权限不存在")
		return
	}

	if err := config.DB.Where("pid = ?", data.ID).First(&model.Auth{}).Error; err == nil {
		util.Error(c, 1, "删除权限失败，该权限下有子权限")
		return
	}

	if err := config.DB.Delete(&auth).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "删除权限成功", nil)
}

// GroupList 获取权限组列表
func GroupList(c *gin.Context) {
	data := struct {
		Keyword string `form:"keyword" json:"keyword"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		fmt.Println(err.Error())
		util.Error(c, 1, err.Error())
		return
	}
	query := config.DB.Model(&model.AuthGroup{})
	if data.Keyword != "" {
		query = query.Where("name like ?", "%"+data.Keyword+"%")
	}
	query = query.Where("admin_id = ?", c.MustGet("admin").(model.Admin).ID)
	var list []model.AuthGroup
	query.Find(&list)
	util.Success(c, "获取权限列表成功", list)
}

// AddGroup 添加权限组
func AddGroup(c *gin.Context) {
	admin := c.MustGet("admin").(model.Admin)
	data := struct {
		Name string `form:"name" json:"name" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	authGroup := model.AuthGroup{
		Name:    data.Name,
		AdminId: int(admin.ID),
	}
	res := config.DB.Model(&model.AuthGroup{}).Create(&authGroup)
	if res.Error != nil {
		util.Error(c, 1, res.Error.Error())
	}
	util.Success(c, "添加权限组成功", nil)
}

// EditGroup 编辑权限组
func EditGroup(c *gin.Context) {
	data := struct {
		ID   uint   `form:"id" json:"id" binding:"required"`
		Name string `form:"name" json:"name" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	res := config.DB.Model(&model.AuthGroup{}).Where("id = ?", data.ID).Update("name", data.Name)
	if res.Error != nil {
		util.Error(c, 1, "编辑权限组失败，权限组不存在")
		return
	}
	util.Success(c, "编辑权限组成功", nil)
}

// DelGroup 删除权限组
func DelGroup(c *gin.Context) {
	data := struct {
		ID uint `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	admin := model.Admin{}
	config.DB.Where("group_id = ?", data.ID).Find(&admin)
	if admin.ID != 0 {
		util.Error(c, 1, "删除权限组失败，权限组下有管理员，请先删除管理员")
		return
	}
	res := config.DB.Model(&model.AuthGroup{}).Where("id = ?", data.ID).Delete(&model.AuthGroup{})
	if res.Error != nil {
		util.Error(c, 1, "删除权限组失败，权限组不存在")
		return
	}
	util.Success(c, "删除权限组成功", nil)
}

// GetGroupAuth 获取权限组权限
func GetGroupAuth(c *gin.Context) {
	data := struct {
		ID uint `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	authGroup := model.AuthGroup{}
	res := config.DB.Model(&model.AuthGroup{}).Where("id = ?", data.ID).First(&authGroup)
	if res.Error != nil {
		util.Error(c, 1, "获取权限组失败，权限组不存在")
		return
	}
	auths, err := model.Auth{}.GetTree(false, c.MustGet("admin").(model.Admin).GroupId)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	list := map[string]any{
		"list": auths,
	}
	if authGroup.Auth == "" {
		list["checked"] = []int{}
	} else {
		strs := strings.Split(authGroup.Auth, ",")
		var checked []int
		for _, v := range strs {
			num, _ := strconv.Atoi(v)
			checked = append(checked, num)
		}
		list["checked"] = checked
	}

	util.Success(c, "获取权限组成功", list)
}

// SetGroupAuth 设置权限组权限
func SetGroupAuth(c *gin.Context) {
	data := struct {
		ID       uint  `form:"id" json:"id" binding:"required"`
		Auth     []int `form:"auth" json:"auth" binding:"required"`
		HalfAuth []int `form:"half_auth" json:"half_auth"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var buf bytes.Buffer
	for i, v := range data.Auth {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(strconv.Itoa(v))
	}
	auth := buf.String()
	halfAuth := ""
	// 处理半选权限
	if len(data.HalfAuth) > 0 {
		buf.Reset()
		for i, v := range data.HalfAuth {
			if i > 0 {
				buf.WriteString(",")
			}
			buf.WriteString(strconv.Itoa(v))
		}
		halfAuth = buf.String()
	}
	upMap := map[string]any{
		"auth":      auth,
		"half_auth": halfAuth,
	}
	res := config.DB.Model(&model.AuthGroup{}).Where("id = ?", data.ID).Updates(upMap)
	if res.Error != nil {
		util.Error(c, 1, "设置权限组失败，权限组不存在")
		return
	}
	util.Success(c, "设置权限组成功", nil)
}

// 管理员列表
func AdminList(c *gin.Context) {
	admin := c.MustGet("admin").(model.Admin)
	data := struct {
		Keyword string `form:"keyword" json:"keyword"`
		Status  uint8  `form:"status" json:"status"`
		Page    int    `form:"page" json:"page" binding:"required"`
		Limit   int    `form:"limit" json:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var total int64
	query := config.DB.Model(&model.Admin{}).Preload("AuthGroup").Where("pid = ?", admin.ID)
	if data.Keyword != "" {
		query = query.Where("username like ?", "%"+data.Keyword+"%")
	}
	if data.Status != 0 {
		query = query.Where("status = ?", data.Status)
	}
	query.Count(&total)
	query = query.Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Order("id desc")
	var list []model.Admin
	query.Find(&list)
	res := map[string]any{
		"data":  list,
		"total": total,
	}
	util.Success(c, "获取管理员列表成功", res)
}

// 添加管理员
func AddAdmin(c *gin.Context) {
	data := struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
		GroupId  int    `form:"group_id" json:"group_id" binding:"required"`
		Status   int8   `form:"status" json:"status" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var admin model.Admin
	if err := mapstructure.Decode(data, &admin); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	admin.Password = util.Md5Hash(data.Password)
	res := config.DB.Create(&admin)
	if res.Error != nil {
		util.Error(c, 1, res.Error.Error())
		return
	}
	util.Success(c, "添加管理员成功", nil)
}

// 编辑管理员
func EditAdmin(c *gin.Context) {
	data := struct {
		ID       uint   `form:"id" json:"id" binding:"required"`
		GroupId  int    `form:"group_id" json:"group_id" binding:"required"`
		Status   int8   `form:"status" json:"status" binding:"required"`
		Password string `form:"password" json:"password"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var admin model.Admin
	if err := mapstructure.Decode(data, &admin); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if data.Password != "" {
		admin.Password = util.Md5Hash(data.Password)
	}
	res := config.DB.Model(&admin).Where("id = ?", data.ID).Updates(&admin)
	if res.Error != nil {
		util.Error(c, 1, "编辑管理员失败，管理员不存在")
		return
	}
	util.Success(c, "编辑管理员成功", nil)
}

// 删除管理员
func DelAdmin(c *gin.Context) {
	data := struct {
		ID uint `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	res := config.DB.Model(&model.Admin{}).Where("id = ?", data.ID).Delete(&model.Admin{})
	if res.Error != nil {
		util.Error(c, 1, "删除管理员失败，管理员不存在")
		return
	}
	util.Success(c, "删除管理员成功", nil)
}
