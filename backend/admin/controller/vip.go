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
	"encoding/json"
	"gin/config"
	"gin/model"
	"gin/util"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"gorm.io/datatypes"
)

func AddPrivilege(c *gin.Context) {
	data := struct {
		Name    string `form:"name" json:"name" binding:"required"`
		Type    int8   `form:"type" json:"type" binding:"required"`
		Content []int  `form:"content" json:"content" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	content, err := json.Marshal(data.Content)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	privilege := model.Privilege{
		Name:    data.Name,
		Type:    data.Type,
		Content: datatypes.JSON(content), // 将切片转换为JSON格式
	}
	if err := config.DB.Create(&privilege).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "添加成功", nil)
}
func PrivilegeList(c *gin.Context) {
	data := struct {
		Keyword string `form:"keyword" json:"keyword"`
		Page    int    `form:"page" json:"page"`
		Limit   int    `form:"limit" json:"limit"`
		All     bool   `form:"all" json:"all"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var list []model.Privilege
	if data.All { // 如果需要获取所有数据，则不进行分页查询
		config.DB.Find(&list)
		util.Success(c, "获取列表成功", list)
		return
	}
	var total int64
	query := config.DB.Model(&model.Privilege{})
	if data.Keyword != "" {
		query.Where("name like ?", "%"+data.Keyword+"%")
	}
	query.Count(&total)
	query.Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Order("id desc")

	query.Find(&list)
	res := map[string]any{
		"data":  list,
		"total": total,
	}
	util.Success(c, "获取列表成功", res)
}
func EditPrivilege(c *gin.Context) {
	data := struct {
		Id      int    `form:"id" json:"id" binding:"required"`
		Name    string `form:"name" json:"name" binding:"required"`
		Type    int8   `form:"type" json:"type" binding:"required"`
		Content []int  `form:"content" json:"content" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	content, err := json.Marshal(data.Content)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	privilege := model.Privilege{
		Name:    data.Name,
		Type:    data.Type,
		Content: datatypes.JSON(content), // 将切片转换为JSON格式
	}
	if err := config.DB.Where("id = ?", data.Id).Updates(&privilege).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "编辑成功", nil)
}
func DelPrivilege(c *gin.Context) {
	data := struct {
		Id int `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Where("id = ?", data.Id).Delete(&model.Privilege{}).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "删除成功", nil)
}
func AddVip(c *gin.Context) {
	data := struct {
		Title      string  `form:"title" json:"title" binding:"required"`
		Price      float64 `form:"price" json:"price" binding:"required"`
		Days       int     `form:"days" json:"days" binding:"required"`
		Rebate     float64 `form:"rebate" json:"rebate"`
		Points     int     `form:"points" json:"points"`
		Content    []int   `form:"content" json:"content" binding:"required"`
		Status     int     `form:"status" json:"status" binding:"required"`
		Desc       string  `form:"desc" json:"desc"`
		Sort       int     `form:"sort" json:"sort"`
		RetailOn   int     `form:"retail_on" json:"retail_on" binding:"required"`
		RetailSet  int     `form:"retail_set" json:"retail_set"`
		RetailType int     `form:"retail_type" json:"retail_type"`
		RetailYj   float64 `form:"retail_yj" json:"retail_yj"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var svip model.Svip
	if err := mapstructure.Decode(data, &svip); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Create(&svip).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	for _, v := range data.Content {
		config.DB.Create(&model.SvipPrivilege{
			Vid: int(svip.ID),
			Pid: v,
		})
	}
	util.Success(c, "添加成功", nil)
}
func VipList(c *gin.Context) {
	data := struct {
		Keyword string `form:"keyword" json:"keyword"`
		Status  int8   `form:"status" json:"status"`
		Page    int    `form:"page" json:"page"`
		Limit   int    `form:"limit" json:"limit"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var list []map[string]any
	var total int64
	query := config.DB.Model(&model.Svip{})
	if data.Keyword != "" {
		query.Where("name like ?", "%"+data.Keyword+"%")
	}
	if data.Status != 0 {
		query.Where("status = ?", data.Status)
	}
	query.Count(&total)
	query.Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Order("sort asc")
	query.Find(&list)
	for i, v := range list {
		var content []int
		config.DB.Model(&model.SvipPrivilege{}).Where("vid = ?", v["id"]).Pluck("pid", &content)
		list[i]["content"] = content
	}
	res := map[string]any{
		"data":  list,
		"total": total,
	}
	util.Success(c, "获取列表成功", res)
}
func EditVip(c *gin.Context) {
	data := struct {
		Id         int     `form:"id" json:"id" binding:"required"`
		Title      string  `form:"title" json:"title" binding:"required"`
		Price      float64 `form:"price" json:"price" binding:"required"`
		Days       int     `form:"days" json:"days" binding:"required"`
		Rebate     float64 `form:"rebate" json:"rebate"`
		Points     int     `form:"points" json:"points"`
		Content    []int   `form:"content" json:"content" binding:"required"`
		Status     int     `form:"status" json:"status" binding:"required"`
		Desc       string  `form:"desc" json:"desc"`
		Sort       int     `form:"sort" json:"sort"`
		RetailOn   int     `form:"retail_on" json:"retail_on" binding:"required"`
		RetailSet  int     `form:"retail_set" json:"retail_set"`
		RetailType int     `form:"retail_type" json:"retail_type"`
		RetailYj   float64 `form:"retail_yj" json:"retail_yj"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	svipMap := map[string]any{
		"title":       data.Title,
		"price":       data.Price,
		"days":        data.Days,
		"rebate":      data.Rebate,
		"points":      data.Points,
		"status":      data.Status,
		"desc":        data.Desc,
		"sort":        data.Sort,
		"retail_on":   data.RetailOn,
		"retail_set":  data.RetailSet,
		"retail_type": data.RetailType,
		"retail_yj":   data.RetailYj,
	}
	if err := config.DB.Model(&model.Svip{}).Where("id = ?", data.Id).Updates(svipMap).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Where("vid = ?", data.Id).Delete(&model.SvipPrivilege{}).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	for _, v := range data.Content {
		config.DB.Create(&model.SvipPrivilege{
			Vid: int(data.Id),
			Pid: v,
		})
	}
	util.Success(c, "编辑成功", nil)
}
func DelVip(c *gin.Context) {
	data := struct {
		Id int `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Where("id = ?", data.Id).Delete(&model.Svip{}).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Where("sid = ?", data.Id).Delete(&model.SvipPrivilege{}).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "删除成功", nil)
}
