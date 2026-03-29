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
	"context"
	"fmt"
	"gin/config"
	"gin/model"
	"gin/service/image"
	"gin/service/set"
	"gin/service/task"
	"gin/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

func GoodsList(c *gin.Context) {
	data := struct {
		SortId  int    `form:"sort_id" json:"sort_id"`
		Page    int    `form:"page" json:"page" binding:"required"`
		Limit   int    `form:"limit" json:"limit" binding:"required"`
		Keyword string `form:"keyword" json:"keyword"`
		Type    int    `form:"type" json:"type"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	query := config.DB.Model(&model.Goods{}).Where("goods.status=1")
	if data.Keyword != "" {
		query.Where("goods.title like ?", "%"+data.Keyword+"%")
	}
	if data.Type != 0 {
		switch data.Type {
		case 1:
			// 查询热门
			query.Order("goods.views desc")
		case 2:
			// 查询置顶
			query.Where("goods.is_top = ?", 1)
		case 3:
			// 查询价格
			query.Order("goods.price asc")
		case 4:
			// 查询最新
			query.Order("goods.utime desc")
		case 5:
			// 查询积分
			query.Where("goods.points>0")
		default:
			query.Order("goods.utime desc")
		}
	}
	if data.SortId != 0 {
		query.Joins("inner join gs_correlation on gs_correlation.gid = goods.id").Where("gs_correlation.sid = ?", data.SortId)
	}
	list := make([]model.Goods, 0)
	query.Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Find(&list)
	for k, v := range list {
		v.UtimeText = util.TimeAgo(v.Utime)
		v.ViewsText = util.FormatNumber(float64(v.Views))
		list[k] = v
	}
	util.Success(c, "success", list)
}
func GoodsInfo(c *gin.Context) {
	data := struct {
		Id       int    `form:"id" json:"id" binding:"required"`
		Pid      int    `form:"pid" json:"pid"`
		Platform string `form:"platform" json:"platform"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	id := data.Id
	goods := model.Goods{}
	config.DB.Where("goods.id = ? and goods.status = 1", id).Preload("GoodsContent").First(&goods)
	if goods.ID == 0 {
		util.Error(c, 1, "商品不存在")
		return
	}
	config.DB.Model(&model.Goods{}).Where("goods.id = ?", id).UpdateColumn("views", gorm.Expr("views + ?", 1))
	is_auth := false
	is_collect := false
	if goods.Price == 0 {
		is_auth = true
	}
	if data.Pid > 0 { //推荐人
		goodsTgType, _ := set.GetSet("goods_tg_type", true)
		if goodsTgType == "2" {
			config.DB.Model(&model.GoodsTask{}).Where("uid=? and gid=?", data.Pid, data.Id).Update("share_nums", gorm.Expr("share_nums + ?", 1))
			task.HandleGoodsTask(data.Pid, data.Id)
		}
	}
	token := c.Request.Header.Get("Authorization")
	if token != "" { // 如果有token
		user := model.User{}
		config.DB.Where("token = ?", token).Find(&user)
		if user.ID != 0 {
			//判断收藏
			userCollect := model.UserCollect{}
			config.DB.Where("uid = ? and gid = ?", user.ID, id).Find(&userCollect)
			if userCollect.ID != 0 {
				is_collect = true
			}
			//判断购买
			if !is_auth {
				userGoods := model.UserGoods{}
				config.DB.Where("uid = ? and gid = ?", user.ID, id).Find(&userGoods)
				if userGoods.ID != 0 {
					is_auth = true
				}
			}
			//判断会员权限
			if !is_auth {
				userSvip := model.UserSvip{}
				if userSvip.CheckGoodsAuth(int(user.ID), int(goods.ID)) {
					is_auth = true
				}
			}
		}

	}

	var m map[string]any
	decoderConfig := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &m,
		TagName:  "json", // 使用 json 标签
	}

	decoder, _ := mapstructure.NewDecoder(decoderConfig)
	decoder.Decode(goods)
	m["is_auth"] = is_auth
	if !is_auth {
		m["link"] = ""
		goods_content := m["goods_content"].(map[string]any)
		goods_content["fee_content"] = ""
		m["goods_content"] = goods_content
	}
	//栏目判断
	if m["cid"].(int) > 0 {
		var clist []model.Goods
		if err := config.DB.Select("id", "title", "price").
			Where("cid = ?", m["cid"]).Order("id asc").Find(&clist).Error; err != nil {
			util.Error(c, 1, err.Error())
			return
		}
		m["clist"] = clist
	}
	m["views_text"] = util.FormatNumber(float64(goods.Views))
	m["utime_text"] = util.TimeAgo(goods.Utime)
	m["is_collect"] = is_collect
	util.Success(c, "success", m)
}
func GoodsSort(c *gin.Context) {
	data, err := model.GoodsSort{}.GetTree()
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "获取列表成功", data)
}
func GetGoodsPromotionImage(c *gin.Context) {
	data := struct {
		Gid      int    `json:"gid" form:"gid" binding:"required"`
		Platform string `json:"platform" form:"platform" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	uid, _ := c.MustGet("uid").(int)
	ctx := context.Background()
	keyName := fmt.Sprintf("goods_promotion_%d_%d_%s", uid, data.Gid, data.Platform)
	if value, err := config.Redis.Get(ctx, keyName).Result(); err == nil {
		if value != "" {
			util.Success(c, "获取成功", value)
			return
		}
	}
	image, err := image.GetGoodsPromotion(int64(uid), int64(data.Gid), data.Platform)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	config.Redis.Set(ctx, keyName, image, 24*time.Hour)
	util.Success(c, "获取成功", image)
}
func GetGoodsTask(c *gin.Context) {
	data := struct {
		Gid int `json:"gid" form:"gid" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	uid, _ := c.MustGet("uid").(int)
	var goods model.Goods
	if err := config.DB.Where("id = ?", data.Gid).First(&goods).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if goods.IsAdv == -1 && goods.IsShare == -1 {
		util.Error(c, 1, "商品不支持任务获取")
		return
	}
	task := model.GoodsTask{}
	if err := config.DB.Where("gid = ? and uid = ?", data.Gid, uid).Find(&task).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if task.ID > 0 {
		util.Success(c, "success", task)
		return
	}
	task = model.GoodsTask{
		Gid:   int64(goods.ID),
		Uid:   int64(uid),
		Share: int64(goods.ShareNums),
		Adv:   int64(goods.AdvNums),
	}
	if err := config.DB.Create(&task).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "success", task)
}
