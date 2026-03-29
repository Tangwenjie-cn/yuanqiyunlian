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
	"gin/service/task"
	"gin/util"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func GoodsList(c *gin.Context) {
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
	query := config.DB.Model(&model.Goods{})
	if data.Keyword != "" {
		query = query.Where("title like ?", "%"+data.Keyword+"%")
	}
	if data.Status != 0 {
		query = query.Where("status = ?", data.Status)
	}
	query.Count(&total)
	query = query.Preload("GsCorrelation").Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Order("id desc")
	var list []model.Goods
	query.Find(&list)
	res := map[string]any{
		"data":  list,
		"total": total,
	}
	util.Success(c, "获取列表成功", res)
}
func AddGoods(c *gin.Context) {
	data := struct {
		Type          int     `form:"type" json:"type" binding:"required"`
		FeeContent    string  `form:"fee_content" json:"fee_content"`
		Code          string  `form:"code" json:"code"`
		Thumb         string  `form:"thumb" json:"thumb" binding:"required"`
		Title         string  `form:"title" json:"title" binding:"required"`
		Link          string  `form:"link" json:"link"`
		SortIds       []int   `form:"sort_ids" json:"sort_ids" binding:"required"`
		Introduction  string  `form:"introduction" json:"introduction"`
		CostPrice     float64 `form:"cost_price" json:"cost_price"`
		Price         float64 `form:"price" json:"price"`
		Points        int     `form:"points" json:"points"`
		OriginalPrice float64 `form:"original_price" json:"original_price"`
		Status        int     `form:"status" json:"status"`
		Content       string  `form:"content" json:"content" binding:"required"`
		Pages         int     `form:"pages" json:"pages"`
		IsTop         int     `form:"is_top" json:"is_top" binding:"required"`
		Views         int     `form:"views" json:"views"`
		RetailOn      int     `form:"retail_on" json:"retail_on" binding:"required"`
		RetailSet     int     `form:"retail_set" json:"retail_set"`
		RetailType    int     `form:"retail_type" json:"retail_type"`
		RetailYj      float64 `form:"retail_yj" json:"retail_yj"`
		//IsAdv         int     `form:"is_adv" json:"is_adv" binding:"required"`
		IsShare   int `form:"is_share" json:"is_share" binding:"required"`
		ShareNums int `form:"share_nums" json:"share_nums"`
		Cid       int `form:"cid" json:"cid"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	// 开启事务
	tx := config.DB.Begin()
	var goods model.Goods
	if err := mapstructure.Decode(data, &goods); err != nil {
		util.Error(c, 1, err.Error())
		return
	}

	if err := tx.Create(&goods).Error; err != nil {
		tx.Rollback() // 回滚事务
		util.Error(c, 1, err.Error())
		return
	}
	sorts := make([]model.GsCorrelation, 0)
	for _, v := range data.SortIds {
		sorts = append(sorts, model.GsCorrelation{
			Gid: int(goods.ID),
			Sid: v,
		})
	}
	if err := tx.Create(&sorts).Error; err != nil {
		tx.Rollback() // 回滚事务
		util.Error(c, 1, err.Error())
		return
	}
	content := model.GoodsContent{
		Gid:        int(goods.ID),
		Content:    data.Content,
		FeeContent: data.FeeContent,
	}
	if err := tx.Create(&content).Error; err != nil {
		tx.Rollback() // 回滚事务
		util.Error(c, 1, err.Error())
		return
	}
	tx.Commit() // 提交事务
	util.Success(c, "添加成功", nil)
}
func GetGoods(c *gin.Context) {
	data := struct {
		ID uint `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	goods := model.Goods{}
	if err := config.DB.Preload("GsCorrelation").Preload("GoodsContent").Where("id = ?", data.ID).First(&goods).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取成功", goods)
}
func EditGoods(c *gin.Context) {
	data := struct {
		Type          int     `form:"type" json:"type" binding:"required"`
		FeeContent    string  `form:"fee_content" json:"fee_content"`
		Code          string  `form:"code" json:"code"`
		ID            uint    `form:"id" json:"id" binding:"required"`
		Thumb         string  `form:"thumb" json:"thumb" binding:"required"`
		Title         string  `form:"title" json:"title" binding:"required"`
		Link          string  `form:"link" json:"link"`
		SortIds       []int   `form:"sort_ids" json:"sort_ids" binding:"required"`
		Introduction  string  `form:"introduction" json:"introduction"`
		CostPrice     float64 `form:"cost_price" json:"cost_price"`
		Price         float64 `form:"price" json:"price"`
		Points        int     `form:"points" json:"points"`
		OriginalPrice float64 `form:"original_price" json:"original_price"`
		Status        int     `form:"status" json:"status"`
		Content       string  `form:"content" json:"content" binding:"required"`
		Pages         int     `form:"pages" json:"pages"`
		RetailOn      int     `form:"retail_on" json:"retail_on" binding:"required"`
		IsTop         int     `form:"is_top" json:"is_top" binding:"required"`
		Views         int     `form:"views" json:"views"`
		RetailSet     int     `form:"retail_set" json:"retail_set"`
		RetailType    int     `form:"retail_type" json:"retail_type"`
		RetailYj      float64 `form:"retail_yj" json:"retail_yj"`
		//IsAdv         int     `form:"is_adv" json:"is_adv" binding:"required"`
		IsShare   int `form:"is_share" json:"is_share"`
		ShareNums int `form:"share_nums" json:"share_nums"`
		Cid       int `form:"cid" json:"cid"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	// 开启事务
	tx := config.DB.Begin()
	goodsMap := map[string]any{
		"type":           data.Type,
		"code":           data.Code,
		"thumb":          data.Thumb,
		"title":          data.Title,
		"link":           data.Link,
		"introduction":   data.Introduction,
		"cost_price":     data.CostPrice,
		"price":          data.Price,
		"points":         data.Points,
		"original_price": data.OriginalPrice,
		"status":         data.Status,
		"pages":          data.Pages,
		"retail_on":      data.RetailOn,
		"is_top":         data.IsTop,
		"views":          data.Views,
		"retail_set":     data.RetailSet,
		"retail_type":    data.RetailType,
		"retail_yj":      data.RetailYj,
		//"is_adv":         data.IsAdv,
		"is_share":   data.IsShare,
		"share_nums": data.ShareNums,
		"cid":        data.Cid,
	}
	if err := tx.Model(&model.Goods{}).Where("id = ?", data.ID).Updates(goodsMap).Error; err != nil {
		tx.Rollback() // 回滚事务
		util.Error(c, 1, err.Error())
		return
	}
	context := model.GoodsContent{
		Gid:        int(data.ID),
		Content:    data.Content,
		FeeContent: data.FeeContent,
	}
	if err := tx.Model(&model.GoodsContent{}).Where("gid = ?", data.ID).Updates(context).Error; err != nil {
		tx.Rollback() // 回滚事务
		util.Error(c, 1, err.Error())
	}
	//处理分类更新
	var sorts []int
	tx.Model(&model.GsCorrelation{}).Where("gid = ?", data.ID).Pluck("sid", &sorts)
	sortsMap := make(map[int]int)
	for _, v := range sorts {
		sortsMap[v] = v
	}
	for _, v := range data.SortIds {
		if _, ok := sortsMap[v]; ok {
			delete(sortsMap, v)
		} else {
			if err := tx.Create(&model.GsCorrelation{
				Gid: int(data.ID),
				Sid: v,
			}).Error; err != nil {
				tx.Rollback() // 回滚事务
				util.Error(c, 1, err.Error())
				return
			}
		}
	}
	if len(sortsMap) > 0 {
		for k := range sortsMap {
			if err := tx.Delete(&model.GsCorrelation{}, "gid = ? and sid = ?", data.ID, k).Error; err != nil {
				tx.Rollback() // 回滚事务
				util.Error(c, 1, err.Error())
				return
			}
		}

	}
	tx.Commit() // 提交事务
	util.Success(c, "成功", nil)
}
func DelGoods(c *gin.Context) {
	data := struct {
		Id uint `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Delete(&model.Goods{}, "id = ?", data.Id).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Delete(&model.GsCorrelation{}, "gid = ?", data.Id).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Delete(&model.GoodsContent{}, "gid = ?", data.Id).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "删除成功", nil)
}
func GoodsSort(c *gin.Context) {
	data, err := model.GoodsSort{}.GetTree()
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取权限列表成功", data)
}
func AddGoodsSort(c *gin.Context) {
	data := struct {
		Name   string `form:"name" json:"name" binding:"required"`
		Status int    `form:"status" json:"status" binding:"required"`
		Thumb  string `form:"thumb" json:"thumb"`
		Pid    int    `form:"pid" json:"pid"`
		Sort   int    `form:"sort" json:"sort"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var goodsSort model.GoodsSort
	if err := mapstructure.Decode(data, &goodsSort); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Create(&goodsSort).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "添加成功", nil)
}
func EditGoodsSort(c *gin.Context) {
	data := struct {
		Id     int    `form:"id" json:"id" binding:"required"`
		Name   string `form:"name" json:"name" binding:"required"`
		Thumb  string `form:"thumb" json:"thumb"`
		Status int    `form:"status" json:"status" binding:"required"`
		Pid    int    `form:"pid" json:"pid"`
		Sort   int    `form:"sort" json:"sort"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var goodsSort model.GoodsSort
	if err := mapstructure.Decode(data, &goodsSort); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Model(&model.GoodsSort{}).Where("id = ?", data.Id).Updates(goodsSort).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "编辑成功", nil)
}
func DelGoodsSort(c *gin.Context) {
	data := struct {
		Id uint `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Delete(&model.GoodsSort{}, "id = ?", data.Id).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "删除成功", nil)
}
func GoodsColumn(c *gin.Context) {
	var list []model.GoodsColumn
	if err := config.DB.Find(&list).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取成功", list)
}
func AddGoodsColumn(c *gin.Context) {
	data := struct {
		Name string `form:"name" json:"name" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Create(&model.GoodsColumn{Name: data.Name}).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "添加成功", nil)
}
func EditGoodsColumn(c *gin.Context) {
	data := struct {
		Id   uint   `form:"id" json:"id" binding:"required"`
		Name string `form:"name" json:"name" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Model(&model.GoodsColumn{}).Where("id = ?", data.Id).Update("name", data.Name).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "编辑成功", nil)
}
func DelGoodsColumn(c *gin.Context) {
	data := struct {
		Id uint `form:"id" json:"id" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Delete(&model.GoodsColumn{}, "id = ?", data.Id).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "删除成功", nil)
}
func DocumentAuto(c *gin.Context) {
	data := struct {
		SortIds    string  `form:"sort_ids" json:"sort_ids" binding:"required"`
		Cid        int     `form:"cid" json:"cid"`
		Price      float64 `form:"price" json:"price"`
		Status     int     `form:"status" json:"status"`
		IsTop      int     `form:"is_top" json:"is_top" binding:"required"`
		RetailOn   int     `form:"retail_on" json:"retail_on" binding:"required"`
		RetailSet  int     `form:"retail_set" json:"retail_set"`
		RetailType int     `form:"retail_type" json:"retail_type"`
		RetailYj   float64 `form:"retail_yj" json:"retail_yj"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	file, _ := c.FormFile("file")
	basePath, err := util.RootPath()
	if err != nil {
		fmt.Println("获取根目录失败")
		return
	}
	dst := filepath.Join(basePath, "docFile", file.Filename)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var goods model.Goods
	if err := mapstructure.Decode(data, &goods); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	go task.HandleDocTask(dst, goods, data.SortIds)
	util.Success(c, "上传成功", nil)
}
func GoodsTask(c *gin.Context) {
	data := struct {
		Keyword string `form:"keyword" json:"keyword"`
		Page    int    `form:"page" json:"page" binding:"required"`
		Limit   int    `form:"limit" json:"limit" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var total int64
	query := config.DB.Model(&model.GoodsTask{}).Joins("join user on user.id=goods_task.uid").Joins("join goods on goods.id=goods_task.gid")
	if data.Keyword != "" {
		query = query.Where("goods_task.uid = ?", data.Keyword)
	}
	query.Count(&total)
	type GoodsTaskInfo struct {
		model.GoodsTask
		Title    string `json:"title"`
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
	}
	var list []GoodsTaskInfo
	if err := query.Select("goods_task.*,goods.title,user.nickname,user.avatar").Limit(data.Limit).Offset((data.Page - 1) * data.Limit).
		Order("goods_task.id desc").Find(&list).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取成功", map[string]any{
		"total": total,
		"data":  list,
	})
}
