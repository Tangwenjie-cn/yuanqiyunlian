<template>
	<view class="container content">
		<image :src="store().config.site_logo" class="logo"></image>
		<span class="title">{{store().config.site_name}}</span>
		<div class="phone-block" v-if="config.phone_login==1">
			<div class="phone-line">
				<span class="iconfont icon-shouji"></span>
				<input v-model="phoneForm.account" placeholder="请输入手机号码" />
			</div>
			<div class="phone-line">
				<span class="iconfont icon-yanzhengma"></span>
				<input v-model="phoneForm.code" placeholder="请输入验证码" style="width: 55%;" />
				<span class="captcha-text" @click="sendSms">{{captchaText}}</span>
			</div>
			<button style="background-color: #1255d0;" @click="login">登录</button>
		</div>
		<div class="flex-column-center">
			<div class="title-block row-center">
				<div class="underline"></div>
				<div>登录方式</div>
				<div class="underline"></div>
			</div>
			<span class="iconfont icon-weixin" @click="wxLogin"></span>
		</div>	
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	import {onLoad} from '@dcloudio/uni-app'
	import { request } from '/config/api.js'
	import store from '@/stores/index.js'
	import {objToQuery,generateRandomString} from '@/config/function.js'
	const config=store().config
	const captchaText=ref('获取验证码')
	const countdown=ref(60)
	const phoneForm=ref({account:'',code:''})
	onLoad((query)=>{
		if(query.back!==undefined){
			store().loginBack=query.back
		}
	})
	function wxLogin(){
		uni.login({
			provider: 'weixin',
			onlyAuthorize:true,
			success: (res) => {
				wxMiniLogin(res.code)
			},
			fail: (err) => {
				uni.showToast({
					title: err.errMsg,
					icon:"none"
				});
			}
		})	
	}
	function wxMiniLogin(code){
		request('WxMiniLogin',{code:code},'POST').then(res=>{
			if(res.data.user===undefined){
				uni.showActionSheet({
					title: '请选择登录方式',
					itemList: ['创建新用户', '绑定已有用户'],
					success: function (res1) {									
						if (res1.tapIndex === 0) {
							return register(res.data)
						}else if(res1.tapIndex === 1){
							uni.showModal({
								title:'绑定用户',
								editable: true,
								placeholderText: '请输入绑定UID',
								success: (res2) => {
									if (res2.confirm) {
										if(!res2.content){
											uni.showToast({
												title: '请输入绑定UID',
												icon:"fail"
											});
											return
										}
										res.data.uid=parseInt(res2.content) 
										BindUser(res.data)
									}
								}
							})
							return
						}
					}
				});
			}else{
				uni.setStorageSync('user',res.data.user)
				store().user=res.data.user
				uni.reLaunch({
					url: store().loginBack
				})
			}
		})
	}	
	function BindUser(data){
		request('BindUser',data,'POST').then(res=>{
			uni.setStorageSync('user',res.data)
			store().user=res.data
			uni.reLaunch({
				url: store().loginBack
			})
		})
	}

	function sendSms(){
		if(countdown.value!==60){
			return
		}
		if(!phoneForm.value.account){
			uni.showToast({
				title: '请输入手机号',
				icon:"none",
			});
			return
		}
		const phoneReg = /^1[3-9]\d{9}$/;
		if(!phoneReg.test(phoneForm.value.account.trim())){
			uni.showToast({
				title: '手机格式错误',
				icon:"none",
			});
			return
		}
		request('SendSms',{phone:phoneForm.value.account},'POST').then(res=>{
			let timer = setInterval(() => {
				captchaText.value = countdown.value + 's后重试'
				countdown.value--
				if (countdown.value == 0) {
					captchaText.value = '获取验证码'
					countdown.value = 60
					clearInterval(timer);
					timer = null;
				}
			}, 1000);	
		})
		
	}
	function login(){
		if(!phoneForm.value.account || !phoneForm.value.code){
			uni.showToast({
				title: '请填写手机或验证码',
				icon:"none"
			});
			return
		}
		request('Login',phoneForm.value,'POST').then(res=>{
			if(res.data.user===undefined){
				uni.showActionSheet({
					title: '请选择登录方式',
					itemList: ['创建新用户', '绑定已有用户'],
					success: function (res1) {									
						if (res1.tapIndex === 0) {
							return register(res.data)
						}else if(res1.tapIndex === 1){
							uni.showModal({
								title:'绑定用户',
								editable: true,
								placeholderText: '请输入绑定UID',
								success: (res2) => {
									if (res2.confirm) {
										if(!res2.content){
											uni.showToast({
												title: '请输入绑定UID',
												icon:"none"
											});
											return
										}
										res.data.uid=parseInt(res2.content) 
										BindUser(res.data)
									}
								}
							})
						}
					}
				});
			}else{
				uni.setStorageSync('user',res.data.user)
				store().user=res.data.user
				uni.reLaunch({
					url: store().loginBack
				})
			}
		})
	}
	function register(params){
		let pid=uni.getStorageSync('pid')
		if(pid){
			params.pid=Number(pid)
		}
		let gid=uni.getStorageSync('gid')
		if(gid){
			params.gid=Number(gid)
		}
		params.nickname = '用户'+generateRandomString()
		request('Register',params,'POST').then(res=>{
			uni.setStorageSync('user',res.data)
			store().user=res.data
			uni.reLaunch({
				url: store().loginBack
			})
		})
	}
</script>

<style scoped>
	@import url("../../static/font/iconfont.css");
	.content{
		display: flex;
		flex-direction: column;
		align-items: center;
	}
	.logo{
		width: 99px;
		height: 99px;
		border-radius: 50%;
		background-color: #fff;
		margin-top: 39px;
	}
	.title{
		font-size: 22px;
		font-weight: bold;
		margin-top: 20px;
	}
	.phone-block{
		margin-top: 10px;
		display: flex;
		flex-direction: column;
	}
	.phone-line{
		width: 450rpx;
		margin-top: 10px;
		display: flex;
		align-items: center;
		background-color: #fff;
		border-radius: 10px;
		height: 39px;
		padding-left: 10px;
	}
	.captcha-text{
		margin-left: 2px;
		font-size: 14px;
		color: #231f18e3;
	}
	button{
		width: 160px;
		height: 28px;
		color: #fff;
		border: none;
		border-radius: 18px;		
		margin-top: 20px;
		display: flex;
		justify-content: center;
		align-items: center;
		font-size: 16px;
	}
	.iconfont{
		color: #2d2d2d;
	}
	.icon-weixin{
    font-size: 36px;
    color: #51af53;		
	}
	.title-block{
		height: 60px;
	}
	.underline{
		flex: 1;
		height: 1px;
		background-color: #a6a6a6;
		margin: 0 10px;
		width: 42px;
	}
</style>
