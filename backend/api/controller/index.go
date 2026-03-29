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
	"gin/config"
	"gin/model"
	"gin/service/set"
	"gin/util"

	"github.com/gin-gonic/gin"
)

func Theme(c *gin.Context) {
	var index model.Theme
	if err := config.DB.Where("type=? and status=?", 1, 1).First(&index).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var sort model.Theme
	if err := config.DB.Where("type=? and status=?", 2, 1).First(&sort).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var vip model.Theme
	if err := config.DB.Where("type=? and status=?", 3, 1).First(&vip).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var my model.Theme
	if err := config.DB.Where("type=? and status=?", 4, 1).First(&my).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var tabBar model.Theme
	if err := config.DB.Where("type=? and status=?", 5, 1).First(&tabBar).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	data := map[string]any{
		"index":  index.Content,
		"sort":   sort.Content,
		"vip":    vip.Content,
		"my":     my.Content,
		"tabBar": tabBar.Content,
	}
	util.Success(c, "success", data)
}
func Config(c *gin.Context) {
	allConfig := set.GetAllSet()
	data := map[string]string{
		"site_name":               allConfig["site_name"],
		"site_logo":               allConfig["site_logo"],
		"site_url":                allConfig["site_url"],
		"telephone":               allConfig["telephone"],
		"address":                 allConfig["address"],
		"cash_on":                 allConfig["cash_on"],
		"cash_wxpay_type":         allConfig["cash_wxpay_type"],
		"cash_alipay_type":        allConfig["cash_alipay_type"],
		"cash_bank_type":          allConfig["cash_bank_type"],
		"cash_min_money":          allConfig["cash_min_money"],
		"cash_max_money":          allConfig["cash_max_money"],
		"super_on":                allConfig["super_on"],
		"super_name":              allConfig["super_name"],
		"alipay_on":               allConfig["alipay_on"],
		"wxpay_on":                allConfig["wxpay_on"],
		"wx_mini_task_templateId": allConfig["wx_mini_task_templateId"],
		"wx_mini_cash_templateId": allConfig["wx_mini_cash_templateId"],
		"phone_login":             allConfig["phone_login"],
	}
	util.Success(c, "success", data)
}
