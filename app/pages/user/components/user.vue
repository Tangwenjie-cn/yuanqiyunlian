<template>
	<view class="content" :style="{'background-color':`${model.backgroundColor}`}">
		<div class="user-block row-center">
			<div class="row-center">	
				<image :src="user.avatar?user.avatar:config.site_logo" class="user-avatar"></image>
				<div class="column">
					<span class="user-name" 
					:style="{'color':model.color}">{{user.nickname}}</span>				
					<span class="user-text" :style="{'color':model.color}">UID：{{user.id}}
					<span class="iconfont icon-fuzhi" style="margin-left: 5px;" @click="cpoy(user.id)"></span>
					</span>
				</div>
			</div>
			<div class="user-right">
				<uni-icons type="right" size="26" :color="model.color" @click="uni.navigateTo({
					url:'/pages/user/set'
				})"></uni-icons>
			</div>
		</div>
	</view>
	<div class="vip-block" v-if="config.v_name!='yq-free'" 
	:style="{background:`linear-gradient(to right,${model.gradientStartColor}, ${model.gradientEndColor})`}">
		<div class="vip-info center" v-if="userVip">
			<span class="iconfont icon-f-vip" :style="{color:model.iconColor}"></span>
			<span class="vip-text" :style="{'color':model.color}">{{userVipList[0].title}}</span>
			<span class="vip-time-text" :style="{'color':model.color}">到期时间:{{formatIsoTime(userVipList[0].expire_time)}}</span>
		</div>
		<div class="vip-info center" v-else>
			<span class="iconfont icon-f-vip" :style="{color:model.iconColor}"></span>
			<span class="vip-text" :style="{'color':model.color}">开通会员特权</span>
			<span class="vip-btn" :style="{'color':model.color}" @click="uni.navigateTo({url:'/pages/vip/index'})">立即开通</span>
		</div>		
	</div>
</template>

<script setup>
	import { ref } from 'vue'
	import store from '@/stores/index.js'
	import { onShow } from '@dcloudio/uni-app'
	import { request } from '/config/api.js'
	import { formatIsoTime } from '/config/function.js'
	const model=defineModel()
	const user=store().user
	const config=store().config
	const userVipList = ref([])
	const userVip=ref(false)
	const login = () => {
		uni.reLaunch({
			url: '/pages/login/index?back='+'/pages/user/index'
		})
	}
	if(!user || Object.keys(user).length==0){
		uni.reLaunch({
			url: '/pages/login/index?back='+'/pages/user/index'
		})
	}
	if(config.v_name!='yq-free'){
		request("UserVipList").then(res=>{
			userVipList.value = res.data
			if(res.data){
				if(new Date(res.data[0].expire_time).getTime() > new Date().getTime()){
					userVip.value=true
				}
			}		
		})
	}
	function cpoy(text){
		uni.setClipboardData({
			data:text,
			success:()=>{
				uni.showToast({
					title:'复制成功',
					icon:'none'
				})
			}
		})
	}
</script>

<style scoped>
	@import url("../../../static/font/iconfont.css");
	.content{
		display: flex;
		width: 100%;
		height: 120px;
		border-bottom-left-radius: 75% 25%;
		border-bottom-right-radius: 75% 25%;
		justify-content: space-between;
	}
	.vip-block{
		width: auto;
		height: 50px;
		border-radius: 10px;
		margin-top: -35px;
		margin-left: 45px;
		margin-right: 45px;
	}
	.vip-info{
		height: 100%;
	}
	.vip-text{
		font-size: 14px;
		margin-left: 5px;
	}
	.vip-time-text{
		font-size: 12px;
		margin-left: 10px;
	}
	.vip-btn{
		font-size: 14px;
		border-radius: 20px;
		padding: 3px 8px;
		margin-left: 10px;
		color: #30204a;
	}
	.user-block{
		width: 100%;
		height: 80px;
		justify-content: space-between;
	}
	.user-avatar{
		width: 60px;
		height: 60px;
		border-radius: 50%;
		margin-left: 20px;
		background-color: #ffffff;
	}
	.user-name{
		font-size: 16px;
		margin-left: 10px;
	}
	.user-text{
		font-size: 12px;
		margin-left: 10px;
	}
	.user-right{
		margin-right: 20px;
	}
</style>