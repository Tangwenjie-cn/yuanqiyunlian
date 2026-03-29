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
	"database/sql"
	"gin/config"
	"gin/model"
	"gin/service/set"
	"gin/util"
	"time"

	"github.com/gin-gonic/gin"
)

func Refresh(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	admin := c.MustGet("admin").(model.Admin)
	auths, err := model.Auth{}.GetTree(true, admin.GroupId)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	logo, _ := set.GetSet("site_logo", true)
	data := map[string]any{
		"token":    token,
		"id":       admin.ID,
		"group_id": admin.GroupId,
		"username": admin.Username,
		"menu":     auths,
		"logo":     logo,
	}
	util.Success(c, "Login successful", data)
}

func EditPwd(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	admin := c.MustGet("admin").(model.Admin)
	data := struct {
		OldPassword string `form:"old_password" json:"old_password" binding:"required"`
		Password    string `form:"password" json:"password" binding:"required"`
	}{}
	if err := c.ShouldBind(&data); err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if util.Sha256Hash(data.OldPassword) != admin.Password {
		util.Error(c, 1, "旧密码错误")
		return
	}
	admin.Password = util.Sha256Hash(data.Password)
	ctx := context.Background()
	config.Redis.Del(ctx, "admin"+token)
	res := config.DB.Save(&admin)
	if res.Error != nil {
		util.Error(c, 1, res.Error.Error())
		return
	}
	util.Success(c, "修改成功", nil)
}
func Home(c *gin.Context) {
	data := map[string]any{}
	//走线图数据
	dayList := []string{}
	now := time.Now()
	t := []time.Time{}
	t = append(t, now)
	t = append(t, now.AddDate(0, 0, -1))
	t = append(t, now.AddDate(0, 0, -2))
	t = append(t, now.AddDate(0, 0, -3))
	t = append(t, now.AddDate(0, 0, -4))
	goodsSalesAmount := []float64{}
	vipSalesAmount := []float64{}
	for _, v := range t {
		dayList = append(dayList, v.Format("2006-01-02"))
		start, end := util.GetDayStartAndEnd(v)
		var gsa, vsa sql.NullFloat64
		config.DB.Model(&model.Order{}).Where("ctime between ? and ? and status = 1 and type = 1", start, end).Select("sum(price)").Scan(&gsa)
		goodsSalesAmount = append(goodsSalesAmount, gsa.Float64)
		config.DB.Model(&model.Order{}).Where("ctime between ? and ? and status = 1 and type = 2", start, end).Select("sum(price)").Scan(&vsa)
		vipSalesAmount = append(vipSalesAmount, vsa.Float64)

	}
	data["lineChart"] = map[string]any{
		"dayList":          dayList,
		"goodsSalesAmount": goodsSalesAmount,
		"vipSalesAmount":   vipSalesAmount,
	}
	//统计数据
	var goodsSalesAmountTotal, vipSalesAmountTotal sql.NullFloat64
	var userTotal, vipTotal, superTotal int64
	config.DB.Model(&model.Order{}).Where("status = 1 and type = 1").Select("sum(price)").Scan(&goodsSalesAmountTotal)
	config.DB.Model(&model.Order{}).Where("status = 1 and type = 2").Select("sum(price)").Scan(&vipSalesAmountTotal)
	config.DB.Model(&model.User{}).Count(&userTotal)
	config.DB.Model(&model.UserSvip{}).Where("expire_time >= ?", now).Distinct("uid").Count(&vipTotal)
	config.DB.Model(&model.User{}).Where("is_super > 0").Count(&superTotal)
	data["statistics"] = map[string]any{
		"goodsSalesAmountTotal": goodsSalesAmountTotal.Float64,
		"vipSalesAmountTotal":   vipSalesAmountTotal.Float64,
		"userTotal":             userTotal,
		"vipTotal":              vipTotal,
		"superTotal":            superTotal,
	}
	//待处理事项
	var cash, applySuper int64
	config.DB.Model(&model.Cash{}).Where("status = 0").Count(&cash)
	config.DB.Model(&model.ApplySuper{}).Where("status = 0").Count(&applySuper)

	data["pending"] = map[string]any{
		"cash":  cash,
		"super": applySuper,
	}
	util.Success(c, "ok", data)
}
