/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
import { defineStore } from 'pinia';

export default defineStore('store', {
	state: () => {
		return {
			user:{},
			theme:{},
			config:{},
			loginBack:"/pages/index/index",
			tabBarSelectedIndex:0
		};
	},
	actions: {
		setUser(user){
			this.user = user;
			uni.setStorageSync('user', user);
		}
	}
});