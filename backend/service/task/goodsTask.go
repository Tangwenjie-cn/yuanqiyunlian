/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package task

import (
	"gin/config"
	"gin/model"
	"gin/service/set"
	"gin/service/wechat"
	"strconv"
)

func HandleGoodsTask(uid int, gid int) {
	var g model.GoodsTask
	if err := config.DB.Where("uid = ? and gid = ?", uid, gid).First(&g).Error; err != nil {
		return
	}
	if g.AdvNums >= g.Adv && g.ShareNums >= g.Share {
		//完成
		var userGoods model.UserGoods
		if err := config.DB.Where("uid = ? and gid = ?", uid, gid).Find(&userGoods).Error; err != nil {
			return
		}
		if userGoods.ID > 0 {
			return
		}
		userGoods.Uid = uid
		userGoods.Gid = gid
		userGoods.Type = 2
		config.DB.Create(&userGoods)
		//发送模板消息

		templateId, err := set.GetSet("wx_mini_task_templateId", true)
		if err != nil {
			return
		}
		var userBind model.UserBind
		if err := config.DB.Where("uid = ? and type=2", uid).First(&userBind).Error; err != nil {
			return
		}
		var goods model.Goods
		if err := config.DB.Where("id = ?", gid).First(&goods).Error; err != nil {
			return
		}
		data := map[string]any{
			"thing3": map[string]any{
				"value": "获取资源任务",
			},
			"phrase12": map[string]any{
				"value": "已完成",
			},
			"thing4": map[string]any{
				"value": goods.Title,
			},
		}
		wechat.SendTemplateMessage(userBind.Openid, templateId, "pages/goods/info?id="+strconv.Itoa(gid), data)

	}
}
