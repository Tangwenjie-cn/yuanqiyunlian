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
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignConfig(c *gin.Context) {
	uid, ok := c.MustGet("uid").(int)
	if !ok {
		util.Error(c, 1, "用户不存在")
		return
	}
	signOn, err := set.GetSet("sign_on", true)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if signOn == "0" {
		util.Error(c, 1, "签到功能已关闭")
		return
	}
	signModeStr, err := set.GetSet("sign_mode", true)
	if err != nil {
		util.Error(c, 1, "获取签到模式失败")
		return
	}
	signMode, err := strconv.Atoi(signModeStr)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var startDate, endDate time.Time
	if signMode == 2 {
		startDate, endDate = util.GetWeekStartEnd()
	} else {
		startDate, endDate = util.GetMonthStartEnd()
	}
	var dayLisy []time.Time
	dayLisy = append(dayLisy, startDate)
	for !dayLisy[len(dayLisy)-1].After(endDate) {
		dayLisy = append(dayLisy, dayLisy[len(dayLisy)-1].AddDate(0, 0, 1))
	}
	dayLisy = dayLisy[:len(dayLisy)-1]
	//获取签到记录
	var signList []model.UserPoints
	if err := config.DB.Where("ctime between ? and ? and uid=? and type=1", startDate, endDate, uid).Order("id desc").Find(&signList).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	//获取当天签到记录
	dayStartDate, dayEndDate := util.GetDayStartAndEnd(time.Now())
	var daySign model.UserPoints
	var checkSign bool
	if err := config.DB.Where("ctime between ? and ? and uid=? and type=1", dayStartDate, dayEndDate, uid).Order("id desc").First(&daySign).Error; err != nil {
		checkSign = false
	} else {
		checkSign = true
	}
	var cumulativeSignDays, signDays int64
	//获取累计签到天数
	if signMode == 1 {
		config.DB.Model(&model.UserPoints{}).Where("uid=? and type=1", uid).Count(&cumulativeSignDays)
	} else {
		config.DB.Model(&model.UserPoints{}).Where("ctime between ? and ? and uid=? and type=1", startDate, endDate, uid).Count(&cumulativeSignDays)
	}
	//$signDays累计奖励计算天数，如果今天签到，则奖励天数=累积签到天数，否则奖励天数=累积签到天数+1
	if checkSign {
		signDays = cumulativeSignDays
	} else {
		signDays = cumulativeSignDays + 1
	}
	var user model.User
	if err := config.DB.Where("id=?", uid).First(&user).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	//signNum连续签到天数，如果今天签到，则连续签到天数=连续签到天数，否则连续签到天数+1
	var signNum, continuousSignDays int
	signNum = user.SignNum
	if checkSign {
		continuousSignDays = signNum
	} else {
		//获取签到周期配置，如果周签到，每周一清空连续签到记录，如果月签到，每月一日清空连续签到记录
		switch signMode {
		case 1:
			continuousSignDays = signNum + 1
		case 2:
			if time.Now().Weekday() == time.Monday {
				continuousSignDays = 0
			}
			continuousSignDays = signNum + 1
		case 3:
			if time.Now().Day() == 1 {
				continuousSignDays = 0
			}
			continuousSignDays = signNum + 1
		default:
			util.Error(c, 1, "签到模式错误")
			return
		}
		//检测昨天是否签到，如果没有签到，连续签到清0
		yesterdayStartDate, yesterdayEndDate := util.GetDayStartAndEnd(time.Now().AddDate(0, 0, -1))
		var yesterdaySign model.UserPoints
		if err := config.DB.Where("ctime between ? and ? and uid=? and type=1", yesterdayStartDate, yesterdayEndDate, uid).
			Order("id desc").First(&yesterdaySign).Error; err != nil {
			signNum = 1
		} else {
			signNum = signNum + 1
		}
	}
	var (
		nextCumulativeSignRewardList []model.SignReward
		nextContinuousSignRewardList []model.SignReward
	)
	if err := config.DB.Where("type=2").Order("days asc").Find(&nextCumulativeSignRewardList).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if err := config.DB.Where("type=1").Order("days asc").Find(&nextContinuousSignRewardList).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	//下一次连续签到奖励还需签到几天
	nextContinuousDays := 0
	for _, v := range nextContinuousSignRewardList {
		if v.Days > continuousSignDays {
			if v.Days-continuousSignDays > 0 {
				nextContinuousDays = v.Days - continuousSignDays
			} else {
				nextContinuousDays = 1
			}
			break
		}
	}
	//下一次累计签到奖励还需签到几天
	nextCumulativeDays := 0
	for _, v := range nextCumulativeSignRewardList {
		if v.Days > int(cumulativeSignDays) {
			if v.Days-int(cumulativeSignDays) > 0 {
				nextCumulativeDays = v.Days - int(cumulativeSignDays)
			} else {
				nextCumulativeDays = 1
			}
		}
	}
	pointStr, err := set.GetSet("sign_give_points", true)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	points, err := strconv.Atoi(pointStr)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var dayLisy1 []map[string]any
	i := 0
	for _, v := range dayLisy {
		value := map[string]any{}
		if signMode == 2 {
			weeks := map[time.Weekday]string{
				1: "周一",
				2: "周二",
				3: "周三",
				4: "周四",
				5: "周五",
				6: "周六",
				0: "周日",
			}
			value["day"] = weeks[v.Weekday()]
		} else {
			value["day"] = v.Format("02日")
		}
		loc := time.Local
		now := time.Now().In(loc)
		todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
		tInLoc := v.In(loc)
		vDateStart := time.Date(tInLoc.Year(), tInLoc.Month(), tInLoc.Day(), 0, 0, 0, 0, loc)
		//判断是否是当天
		if todayStart.Equal(vDateStart) {
			value["sign_day"] = true
		} else {
			value["sign_day"] = false
		}
		//判断是否签到
		value["is_sign"] = false
		for _, v2 := range signList {
			cDateStart := time.Date(v2.Ctime.Year(), v2.Ctime.Month(), v2.Ctime.Day(), 0, 0, 0, 0, loc)
			if cDateStart.Equal(vDateStart) {
				value["is_sign"] = true
				break
			}
		}
		value["type"] = 0 //签到类型，0已签到，1连续签到，2累计签到

		value["points"] = points
		if vDateStart.Equal(todayStart) || vDateStart.After(todayStart) {
			for _, v2 := range nextCumulativeSignRewardList {
				if v2.Days-int(signDays) == i {
					value["type"] = v2.Type
					value["points"] = v2.Points
					break
				}
			}
			for _, v2 := range nextContinuousSignRewardList {
				if v2.Days-int(signNum) == i {
					value["type"] = v2.Type
					value["points"] = v2.Points
					break
				}
			}
			i++
		}
		dayLisy1 = append(dayLisy1, value)
	}
	util.Success(c, "success", map[string]any{
		"list":               dayLisy1,
		"signMode":           signMode,
		"checkSign":          checkSign,
		"cumulativeSignDays": cumulativeSignDays,
		"continuousSignDays": continuousSignDays,
		"nextContinuousDays": nextContinuousDays,
		"nextCumulativeDays": nextCumulativeDays,
	})
}
func SignAdd(c *gin.Context) {
	uid, ok := c.MustGet("uid").(int)
	if !ok {
		util.Error(c, 1, "用户不存在")
		return
	}
	signOn, err := set.GetSet("sign_on", true)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	if signOn == "0" {
		util.Error(c, 1, "签到功能已关闭")
		return
	}
	//获取当天签到记录
	dayStartDate, dayEndDate := util.GetDayStartAndEnd(time.Now())
	var daySign model.UserPoints
	if err := config.DB.Where("ctime between ? and ? and uid=? and type=1", dayStartDate, dayEndDate, uid).Order("id desc").First(&daySign).Error; err == nil {
		util.Error(c, 1, "今天已经签到过了")
		return
	}
	var user model.User
	if err := config.DB.Where("id=?", uid).First(&user).Error; err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	//检测昨天是否签到，如果没有签到，连续签到清0
	yesterdayStartDate, yesterdayEndDate := util.GetDayStartAndEnd(time.Now().AddDate(0, 0, -1))
	var yesterdaySign model.UserPoints
	if err := config.DB.Where("ctime between ? and ? and uid=? and type=1", yesterdayStartDate, yesterdayEndDate, uid).
		Order("id desc").First(&yesterdaySign).Error; err != nil {
		user.SignNum = 0
	}
	signModeStr, err := set.GetSet("sign_mode", true)
	if err != nil {
		util.Error(c, 1, "获取签到模式失败")
		return
	}
	signMode, err := strconv.Atoi(signModeStr)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	switch signMode {
	case 2:
		if time.Now().Weekday() == time.Monday {
			user.SignNum = 0
		}
	case 3:
		if time.Now().Day() == 1 {
			user.SignNum = 0
		}
	}
	user.SignNum++
	continuousDays := user.SignNum
	var startDate, endDate time.Time
	if signMode == 2 {
		startDate, endDate = util.GetWeekStartEnd()
	} else {
		startDate, endDate = util.GetMonthStartEnd()
	}
	var cumulativeDays int64
	//获取累计签到天数
	if signMode == 1 {
		config.DB.Model(&model.UserPoints{}).Where("uid=? and type=1", uid).Count(&cumulativeDays)
	} else {
		config.DB.Model(&model.UserPoints{}).Where("ctime between ? and ? and uid=? and type=1", startDate, endDate, uid).Count(&cumulativeDays)
	}
	cumulativeDays++
	pointStr, err := set.GetSet("sign_give_points", true)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	points, err := strconv.Atoi(pointStr)
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	var signReward model.SignReward
	var continuousRewardPoint, cumulativeRewardPoint int
	if err := config.DB.Where("type=1 and days=?", continuousDays).First(&signReward).Error; err == nil {
		continuousRewardPoint = signReward.Points
	}
	if err := config.DB.Where("type=2 and days=?", cumulativeDays).First(&signReward).Error; err == nil {
		cumulativeRewardPoint = signReward.Points
	}
	if cumulativeRewardPoint > 0 {
		points = cumulativeRewardPoint
	}
	if continuousRewardPoint > 0 {
		points = continuousRewardPoint
	}

	tx := config.DB.Begin()
	if err := tx.Create(&model.UserPoints{
		Uid:    uid,
		Type:   1,
		Points: points,
		Mode:   1,
		Before: user.Points,
		After:  user.Points + points,
		Remark: "签到奖励",
	}).Error; err != nil {
		tx.Rollback()
		util.Error(c, 1, err.Error())
		return
	}
	if err := tx.Model(&model.User{}).Where("id=?", uid).Updates(map[string]any{
		"sign_num": user.SignNum,
		"points":   gorm.Expr("points + ?", points),
	}).Error; err != nil {
		tx.Rollback()
		util.Error(c, 1, err.Error())
		return
	}
	tx.Commit()
	util.Success(c, fmt.Sprintf("签到成功,获得%d积分", points), nil)
}
