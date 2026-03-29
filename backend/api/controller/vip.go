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
	"gin/util"

	"github.com/gin-gonic/gin"
)

func VipList(c *gin.Context) {
	list := []model.Svip{}
	if err := config.DB.Order("sort asc").Find(&list).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, "成功", list)
}
