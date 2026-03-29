/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright  © 2026 成都源启科技有限公司
 * @license    授权协议（Apache2.0）
 */
import { reactive } from 'vue'
const store = reactive({
	menu:null,
	userInfo:null,
	token:null,
	config:null,
	upUser(){
		let user=JSON.parse(localStorage.getItem('user'));
		if(user){
			this.menu=user.menu;
			this.userInfo=user.userInfo;
			this.token=user.userInfo.token;
			this.config=user.config;
		}else{
			this.menu=null;
			this.userInfo=null;
			this.token=null;
			this.config=null;
		}
	},
	treeMenuClick:null, //无限级菜单点击状态
});
export default store;
