/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package router

import (
	"gin/admin/controller"
	"gin/admin/middleware"

	"github.com/gin-gonic/gin"
)

func admin(r *gin.Engine) {
	// Admin routes
	admin := r.Group("/admin", middleware.CheckAuth)
	{
		// 无需权限验证请求

		// 登录
		r.POST("/admin/login/login", controller.Login)
		r.GET("/admin/login/logout", controller.Logout)

	}
	{
		index := admin.Group("/index")
		index.GET("/home", controller.Home)
		// 刷新
		index.GET("/refresh", controller.Refresh)
		//修改密码
		index.POST("/editPwd", controller.EditPwd)
	}
	{
		// 权限管理
		auth := admin.Group("/auth")
		auth.POST("/authList", controller.AuthList)
		auth.POST("/addAuth", controller.AddAuth)
		auth.GET("/getAuth", controller.GetAuth)
		auth.POST("/editAuth", controller.EditAuth)
		auth.POST("/delAuth", controller.DelAuth)
		auth.GET("/groupList", controller.GroupList)
		auth.POST("/addGroup", controller.AddGroup)
		auth.POST("/editGroup", controller.EditGroup)
		auth.POST("/delGroup", controller.DelGroup)
		auth.POST("/getGroupAuth", controller.GetGroupAuth)
		auth.POST("/setGroupAuth", controller.SetGroupAuth)
		auth.POST("/adminList", controller.AdminList)
		auth.POST("/addAdmin", controller.AddAdmin)
		auth.POST("/editAdmin", controller.EditAdmin)
		auth.POST("/delAdmin", controller.DelAdmin)
	}
	{
		system := admin.Group("/system")
		//参数列表
		system.POST("/setList", controller.SetList)
		system.GET("/setGroupList", controller.SetGroupList)
		system.POST("/addSet", controller.AddSet)
		system.GET("/getSet", controller.GetSet)
		system.POST("/editSet", controller.EditSet)
		system.POST("/delSet", controller.DelSet)
		system.GET("/getSetAllInfo", controller.GetSetAllInfo)
		system.POST("/saveAllSet", controller.SaveAllSet)
		system.GET("/getUpdateInfo", controller.GetUpdateInfo)
		system.GET("/update", controller.Update)
		system.POST("/wxMiniUpload", controller.WxMiniUpload)
		system.GET("/adminLogList", controller.AdminLogList)
	}
	{
		upload := admin.Group("/upload")
		upload.POST("/uploadFile", controller.UploadFile)
		upload.POST("/delFile", controller.DelFile)
	}
	{
		article := admin.Group("/article")
		article.POST("/list", controller.Article)
		article.POST("/save", controller.AddArticle)
		article.POST("/update", controller.EditArticle)
		article.POST("/del", controller.DelArticle)
		article.POST("/sort", controller.ArticleSort)
		article.POST("/saveSort", controller.AddArticleSort)
		article.POST("/updateSort", controller.EditArticleSort)
		article.POST("/delSort", controller.DelArticleSort)
	}
	{
		gallery := admin.Group("/gallery")
		gallery.POST("/index", controller.ImgList)
		gallery.POST("/save", controller.AddImg)
		gallery.POST("/update", controller.EditImg)
		gallery.POST("/del", controller.DelImg)
		gallery.GET("/sort", controller.ImgSort)
		gallery.POST("/saveSort", controller.AddImgSort)
		gallery.POST("/updateSort", controller.EditImgSort)
		gallery.POST("/delSort", controller.DelImgSort)
	}
	{
		user := admin.Group("/user")
		user.GET("/list", controller.UserList)
		user.GET("/info", controller.UserInfo)
		user.POST("/setUser", controller.SetUserInfo)
		user.GET("/getUserVip", controller.GetUserVip)
		user.POST("/setUserSuper", controller.SetUserSuper)
		user.POST("/setUserVip", controller.SetUserVip)
		user.GET("/applySuper", controller.GetApplySuper)
		user.POST("/setApplySuper", controller.SetApplySuper)
		user.GET("/balanceList", controller.BalanceList)
		user.GET("/goodsList", controller.UserGoodsList)
	}
	{
		goods := admin.Group("/goods")
		goods.POST("/index", controller.GoodsList)
		goods.POST("/save", controller.AddGoods)
		goods.GET("/getGoods", controller.GetGoods)
		goods.POST("/update", controller.EditGoods)
		goods.POST("/del", controller.DelGoods)
		goods.POST("/sort", controller.GoodsSort)
		goods.POST("/addSort", controller.AddGoodsSort)
		goods.POST("/upSort", controller.EditGoodsSort)
		goods.POST("/delSort", controller.DelGoodsSort)
		goods.GET("/column", controller.GoodsColumn)
		goods.POST("/column/save", controller.AddGoodsColumn)
		goods.POST("/column/edit", controller.EditGoodsColumn)
		goods.POST("/column/del", controller.DelGoodsColumn)
		goods.POST("/doc/upload", controller.DocumentAuto)
		goods.GET("/task", controller.GoodsTask)
	}
	{
		theme := admin.Group("/theme")
		theme.GET("/index", controller.ThemeIndex)
		theme.POST("/index/save", controller.ThemeAddIndex)
		theme.POST("/index/update", controller.ThemeEdit)
		theme.POST("/index/del", controller.ThemeDelIndex)
		theme.POST("/index/set", controller.ThemeSetIndex)
		theme.GET("/index/get", controller.ThemeGet)
		theme.GET("/pages/list", controller.PagesList)
		theme.POST("/pages/save", controller.PagesAdd)
		theme.POST("/pages/update", controller.PagesEdit)
		theme.POST("/pages/del", controller.PagesDel)
	}
	{
		vip := admin.Group("/vip")
		vip.POST("/addPrivilege", controller.AddPrivilege)
		vip.POST("/privilegeList", controller.PrivilegeList)
		vip.POST("/delPrivilege", controller.DelPrivilege)
		vip.POST("/editPrivilege", controller.EditPrivilege)
		vip.POST("/save", controller.AddVip)
		vip.POST("/list", controller.VipList)
		vip.POST("/del", controller.DelVip)
		vip.POST("/edit", controller.EditVip)
	}
	{
		kami := admin.Group("/kami")
		kami.POST("/list", controller.KamiList)
		kami.POST("/save", controller.AddKami)
		kami.POST("/del", controller.DelKami)
		kami.POST("/edit", controller.EditKami)
	}
	{
		cash := admin.Group("/cash")
		cash.GET("/list", controller.CashList)
		cash.POST("/check", controller.CashCheck)
	}
	{
		order := admin.Group("/order")
		order.GET("/list", controller.OrderList)
	}
	{
		sign := admin.Group("/sign")
		sign.GET("/list", controller.SignList)
		sign.POST("/save", controller.SignAdd)
		sign.POST("/del", controller.SignDel)
		sign.POST("/edit", controller.SignEdit)
	}
	{
		points := admin.Group("/points")
		points.GET("/log", controller.PointsLog)
	}
}
