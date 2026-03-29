/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
import axios from "axios";
import store from '../store/index.js';
import { ElMessage } from 'element-plus';
import router from '../router/index.js';
let url=''; //接口域名
fetch('/config.json').then(res=>{
	res.json().then(res1=>{
		url=res1.url+'/admin/'
	})
})
export const domain=()=>url
export const paths={ //接口地址
	login:'login/login',
	logout:'login/logout',

	refresh:'index/refresh',
	editPwd:'index/editPwd',
	home:'index/home',

	authList:'auth/authList',
	addAuth:'auth/addAuth',
	getAuth:'auth/getAuth',
	editAuth:'auth/editAuth',
	delAuth:'auth/delAuth',			
	groupList:'auth/groupList',
	addGroup:'auth/addGroup',
	editGroup:'auth/editGroup',
	delGroup:'auth/delGroup',
	getGroupAuth:'auth/getGroupAuth',
	setGroupAuth:'auth/setGroupAuth',
	adminList:'auth/adminList',
	addAdmin:'auth/addAdmin',
	editAdmin:'auth/editAdmin',
	delAdmin:'auth/delAdmin',
	
	setList:'system/setList',
	setGroupList:'system/setGroupList',
	addSet:'system/addSet',
	getSet:'system/getSet',
	editSet:'system/editSet',
	delSet:'system/delSet',
	getSetAllInfo:'system/getSetAllInfo',
	saveAllSet:'system/saveAllSet',
	GetUpdateInfo:'system/getUpdateInfo',
	Update:'system/update',
	WxMiniUpload:'system/wxMiniUpload',
	AdminLogList:'system/adminLogList',
	
	delFile:'upload/delFile',
	uploadFile:'upload/uploadFile',

	article:'article/list',
	addArticle:'article/save',
	editArticle:'article/update',
	delArticle:'article/del',
	articleSort:'article/sort',
	addArticleSort:'article/saveSort',
	editArticleSort:'article/updateSort',
	delArticleSort:'article/delSort',
	
	imgList:'gallery/index',
	addImg:'gallery/save',
	editImg:'gallery/update',
	delImg:'gallery/del',
	imgSort:'gallery/sort',
	addSort:'gallery/saveSort',
	upSort:'gallery/updateSort',
	delSort:'gallery/delSort',
	
	userList:'user/list',
	UserInfo:'user/info',
	SetUserInfo:'user/setUser',
	GetUserVip:'user/getUserVip',
	SetUserVip:'user/setUserVip',
	GetApplySuper:'user/applySuper',
	SetApplySuper:'user/setApplySuper',
	SetUserSuper:'user/setUserSuper',
	BalanceList:'user/balanceList',
	UserGoodsList:'user/goodsList',

	CashList:'cash/list',
	CashCheck:'cash/check',

	OrderList:'order/list',

	goodsList:'goods/index',
	addGoods:'goods/save',
	getGoods:'goods/getGoods',
	editGoods:'goods/update',
	delGoods:'goods/del',
	goodsSort:'goods/sort',
	addGoodsSort:'goods/addSort',
	upGoodsSort:'goods/upSort',
	delGoodsSort:'goods/delSort',
	GoodsColumn:'goods/column',
	AddGoodsColumn:'goods/column/save',
	EditGoodsColumn:'goods/column/edit',
	DelGoodsColumn:'goods/column/del',
	UploadDocument:'goods/doc/upload',
	GoodsTask:'goods/task',

	ThemeIndex:'theme/index',
	ThemeAddIndex:'theme/index/save',
	ThemeEdit:'theme/index/update',
	ThemeDelIndex:'theme/index/del',
	ThemeSetIndex:'theme/index/set',
	ThemeGet:'theme/index/get',
	PagesList:'theme/pages/list',
	PagesAdd:'theme/pages/save',
	PagesEdit:'theme/pages/update',
	PagesDel:'theme/pages/del',

	AddPrivilege:'vip/addPrivilege',
	PrivilegeList:'vip/privilegeList',
	EditPrivilege:'vip/editPrivilege',
	DelPrivilege:'vip/delPrivilege',
	AddVip:'vip/save',
	VipList:'vip/list',
	EditVip:'vip/edit',
	DelVip:'vip/del',

	KamiList:'kami/list',
	KamiAdd:'kami/save',
	KamiDel:'kami/del',
	KamiEdit:'kami/edit',

	SignList:'sign/list',
	SignAdd:'sign/save',
	SignEdit:'sign/edit',
	SignDel:'sign/del',

	PointsLog:'points/log'
};
export async function request(path,data={},method='get',headers={}){
	headers.Authorization=store.token;
	let config={
		'baseURL':url,
		'url':paths[path],
		'method':method,
		//timeout:10000,
		'headers':headers,
	};
	if(method.toLowerCase()==='post'){
		//headers["content-type"]="multipart/form-data"		
		config.data=data
	}else{
		config.params=new URLSearchParams(data)
	}
	try{
		let response=await axios(config)
		if(response.status!=200){
			throw new Error('请求错误')
		}			
		if(response.data.code===undefined){
			throw new Error('返回数据格式错误')
		}else if(response.data.code===0){
			return response.data;
		}else if(response.data.code===1){ //错误信息
			throw new Error(response.data.msg)
		}else if(response.data.code===2){ //用户验证失败
			localStorage.removeItem('user');
			localStorage.removeItem('menuActive');
			localStorage.removeItem('tableTabs');
			localStorage.removeItem('tableTabsValue');
			store.upUser();
			router.push('/login');			
		}
		
	}catch(e){
		// 第一步：判断是否是Axios的HTTP错误（含500）
		let errorMsg = e.message; // 默认使用原有错误信息
		if (e.response) {
			// 存在response，说明是服务端返回的错误（如500、404等）
			const { status, data } = e.response;
			// 处理500状态码的特殊逻辑
			if (status === 500) {
				// 提取500的返回内容，兼容不同的后端返回格式
				errorMsg = data?.msg || `服务器内部错误(${status})：${JSON.stringify(data)}`;
			} else {
				// 其他HTTP错误（如404、403），也可自定义提示
				errorMsg = `请求失败(${status})：${data?.msg || '未知错误'}`;
			}
		} else if (e.request) {
			// 请求已发送但无响应（网络错误）
			errorMsg = '网络异常，请检查网络连接';
		}

		// 第二步：统一提示错误信息
		ElMessage.error(errorMsg);

		// 第三步：抛出包含完整信息的错误，供业务代码捕获（可选）
		throw new Error(errorMsg);
	}
}