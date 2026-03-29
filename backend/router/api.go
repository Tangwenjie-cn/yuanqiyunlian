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
	"gin/api/controller"
	"gin/api/middleware"

	"github.com/gin-gonic/gin"
)

func api(r *gin.Engine) {
	{
		pay := r.Group("/pay")
		pay.POST("/alipay/notify", controller.AlipayNotify)
		pay.POST("/wxpay/notify", controller.WxpayNotify)
	}
	api := r.Group("/api", middleware.CheckToken)
	{
		index := api.Group("/index")
		index.GET("/theme", controller.Theme)
		index.GET("/config", controller.Config)
	}
	{
		article := api.Group("/article")
		article.GET("/list", controller.ArticleList)
		article.GET("/info", controller.ArticleInfo)
	}
	{
		goods := api.Group("/goods")
		goods.GET("/list", controller.GoodsList)
		goods.GET("/info", controller.GoodsInfo)
		goods.GET("/sort", controller.GoodsSort)
		goods.GET("/promotionImage", controller.GetGoodsPromotionImage)
		goods.GET("/task", controller.GetGoodsTask)
	}
	{
		login := api.Group("/login")
		login.POST("/login", controller.Login)
		login.POST("/wxMini", controller.WxMiniLogin)
		login.POST("/register", controller.Register)
		login.POST("/bindUser", controller.BindUser)
		login.GET("/out", controller.LoginOut)
		login.POST("/sendSms", controller.SendSms)
	}
	{
		user := api.Group("/user")
		user.POST("/collect", controller.Collect)
		user.GET("/vipList", controller.UserVipList)
		user.GET("/info", controller.GetUserInfo)
		user.POST("/update", controller.UpdateUserInfo)
		user.GET("/goods", controller.UserGoods)
		user.GET("/bill", controller.UserBill)
		user.GET("/promotionImage", controller.GetPromotionImage)
		user.GET("/myPromotion", controller.MyPromotion)
		user.POST("/subSuper", controller.SubSuper)
		user.GET("/getSubSuper", controller.GetSubSuper)
		user.GET("/goodsTask", controller.GetUserGoodsTask)
		user.GET("/collectList", controller.CollectList)
	}
	{
		vip := api.Group("/vip")
		vip.POST("/list", controller.VipList)
	}
	{
		order := api.Group("/order")
		order.POST("/create", controller.CreateOrder)
		order.POST("/init", controller.InitOrder)
		order.GET("/list", controller.OrderList)
		order.GET("/info", controller.OrderInfo)
	}
	{
		upload := api.Group("/upload")
		upload.POST("/file", controller.UploadFile)
	}
	{
		cash := api.Group("/cash")
		cash.POST("/add", controller.CashAdd)
		cash.GET("/list", controller.CashList)
		cash.GET("/getWxTransfer", controller.GetWxTransfer)
	}
	{
		sign := api.Group("/sign")
		sign.GET("/config", controller.SignConfig)
		sign.POST("/add", controller.SignAdd)
	}
	{
		points := api.Group("/points")
		points.GET("/list", controller.UserPoints)
		points.GET("/vip", controller.PointsVipList)
		points.POST("/buy", controller.PointsBuy)
	}
}
