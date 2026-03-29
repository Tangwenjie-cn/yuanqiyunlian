/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
import store from '@/stores/index.js'

let apiUrl=null
let yxtConfig=null

// #ifndef H5
import config from '@/config.js'
yxtConfig=config
apiUrl=yxtConfig.url+'/api'
// #endif

// #ifdef H5
fetch('/config.json').then(res=>{
	res.json().then(res1=>{
		apiUrl=res1.url+'/api'
		yxtConfig=res1
	})
})
// #endif
export const paths={
	Theme:"/index/theme",
	Config:"/index/config",
	
	ArticleList:"/article/list",
	ArticleInfo:'/article/info',
	
	GoodsList:"/goods/list",
	GoodsInfo:"/goods/info",
	GoodsSort:'/goods/sort',
	GetGoodsPromotionImage:'/goods/promotionImage',
	GetGoodsTask:'/goods/task',
	
	WxMiniLogin:"/login/wxMini",
	Register:"/login/register",
	BindUser:"/login/bindUser",
	Login:"/login/login",
	Logout:"/login/out",
	SendSms:"/login/sendSms",
	
	Collect:'/user/collect',
	UserVipList:'/user/vipList',
	GetUserInfo:'/user/info',
	UpdateUserInfo:'/user/update',
	UserGoods:'/user/goods',
	UserBill:'/user/bill',
	UserPromotionImage:'/user/promotionImage',
	MyPromotion:'/user/myPromotion',
	SubSuper:'/user/subSuper',
	GetSubSuper:'/user/getSubSuper',
	GetUserGoodsTask:'/user/goodsTask',
	CollectList:'/user/collectList',
	
	VipList:'/vip/list',
	
	InitOrder:'/order/init',
	OrderList:'/order/list',
	OrderInfo:'/order/info',
	OrderCreate:'/order/create',	
	
	UploadFile:'/upload/file',
	
	CashAdd:'/cash/add',
	CashList:'/cash/list',
	GetWxTransfer:'/cash/getWxTransfer',
	
	SignConfig:'/sign/config',
	SignAdd:'/sign/add',
	
	UserPoints:'/points/list',
	PointsVipList:'/points/vip',
	PointsBuy:'/points/buy',
}
function delay(time) {
  return new Promise(resolve => setTimeout(resolve, time));
}
export async function request(path,data={},method='GET',headers={}){
	if(!apiUrl){
		await delay(200);
	}
	if(store().user){
		headers['Authorization']=store().user.token
	}
	// #ifdef H5
	data.platform="h5"
	// #endif
	// #ifdef MP-WEIXIN
	data.platform="weixin"
	// #endif
	// #ifdef APP-PLUS
	data.platform="app"
	// #endif
	try{
		let res=await uni.request({
			url:apiUrl+paths[path],
			method:method,
			data:data,
			header:headers
		})
		if(res.statusCode!=200){
			if(Object.hasOwn(res.data,'msg')){
				throw new Error(res.data.msg)
			}
			throw new Error('服务器错误')			
		}
		if(res.data.code==0){
			return res.data
		}
		if(res.data.code==1){
			throw new Error(res.data.msg)
		}
		if(res.data.code==2){
			uni.navigateTo({
				url:'/pages/login/index'
			})
			throw new Error(res.data.msg)
		}
	}catch(err){
		uni.showToast({
			title: err.message,
			icon:"none",
			duration:3000
		});
		throw new Error(err.message)
	}
}
export async function upload(path,filePath){
	try{
		let res=await uni.uploadFile({
			url:apiUrl+paths[path],
			filePath:filePath,
			name:'file',
			header:{
				'Authorization':store().user.token
			}
		})
		if(typeof res.data!='object'){
			res.data=JSON.parse(res.data)
		}
		if(res.data.code==0){
			return res.data
		}
		if(res.data.code==1){
			throw new Error(res.data.msg)
		}
		if(res.data.code==2){
			uni.navigateTo({
				url:'/pages/login/index'
			})
			throw new Error(res.data.msg)
		}
	}catch(err){
		uni.showToast({
			title: err.message,
			icon:"none",
			duration:3000
		});
		throw new Error(err.message)
	}
}