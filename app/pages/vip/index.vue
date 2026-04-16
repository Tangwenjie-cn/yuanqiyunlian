<template>
	<view class="container">
		<div class="user-block row-center" :style="{backgroundColor:theme.backgroundColor}">
			<image :src="user.avatar?user.avatar:config.site_logo" class="avatar"></image>
			<div class="column" style="margin-left: 10px;">
				<div>
					<span class="user-text" :style="{color:theme.color}">{{user.nickname}}</span>
					<span class="iconfont icon-f-vip" :style="{color:theme.iconColor}" 
					v-if="userVip">
					</span>
				</div>				
				<div class="vip-text" :style="{color:theme.color}" 
				v-if="!userVip">
					<span>普通用户</span>
				</div>
			</div>
		</div>
		<div class="column user-vip-list" v-if="userVipList.length>0">
			<div class="row-center user-vip-item" v-for="item in userVipList">
				<span>{{item.title}}</span>
				<span :style="{color:Date.parse(item.expire_time)>Date.now()?'rgb(82, 155, 46)':'rgb(196, 86, 86)'}">到期时间:{{formatIsoTime(item.expire_time)}}</span>
			</div>		
		</div>
		<uni-card title="选择会员">
			<scroll-view scroll-x="true" class="vip-list">
				<view class="vip-item" v-for="item in vipList" :class="{'vip-item-active':checkVip.id===item.id}" @click="vipClick(item)">
					<span class="vip-title">{{item.title}}</span>
					<span class="vip-price">￥{{item.price}}</span>
				</view>
			</scroll-view>
		</uni-card>
		<span class="buy-btn" @click="buy">购买此套餐</span>
		<uni-card title="说明" v-show="Object.keys(checkVip).length>0">
			<div class="column">
				<text>{{checkVip.desc}}</text>
				<span>时长:{{checkVip.days}}天</span>
				<span v-if="checkVip.rebate>0"> 折扣特权:{{checkVip.rebate}}折</span>
			</div>
			
		</uni-card>
		<yq-tab-bar ref="tabBar"></yq-tab-bar>
	</view>
</template>

<script setup>
	import {ref,watch,nextTick} from 'vue'	
	import { onShow } from '@dcloudio/uni-app'
	import store from '@/stores/index.js'
	import { request } from '/config/api.js'
	import { formatIsoTime } from '/config/function.js'
	const theme = ref({})
	const tabBar = ref(null)
	const user=store().user
	const config=store().config
	const checkVip = ref({})
	const vipList = ref([])
	const userVipList = ref([])
	const userVip=ref(false)
	function vipClick(i){
		checkVip.value = i
	}

	onShow(()=>{
		theme.value=uni.getStorageSync('theme').vip[0]
		setStatus()
		nextTick(()=>{
			if (tabBar.value){
				tabBar.value.autoMatchActiveTab()
			}
		})
		request("VipList",{},'post').then(res=>{
			vipList.value = res.data
		})
		request("UserVipList").then(res=>{
			userVipList.value = res.data
			if(new Date(res.data[0].expire_time).getTime() > new Date().getTime()){
				userVip.value=true
			}
		})
	})
	function setStatus(){
		uni.setNavigationBarColor({
			frontColor: theme.value.titleColor,
			backgroundColor: theme.value.backgroundColor,
		})
	}
	function buy(){
		if(Object.keys(checkVip.value).length===0){
			uni.showToast({
				icon:'none',
				title:'请选择会员'
			})
			return
		}
		uni.navigateTo({
			url:'/pages/order/buy-vip?id='+checkVip.value.id,
		})
	}
</script>

<style scoped>
	@import url("../../static/font/iconfont.css");
	.user-block{
		border-bottom-left-radius: 75% 25%;
		border-bottom-right-radius: 75% 25%;
		padding: 10px;
		height: 99px;
	}
	.avatar{
		width: 60px;
		height: 60px;
		border-radius: 50%;
	}
	.vip-text{
		font-size: 12px;
	}
	.user-text{
		font-size: 16px;
	}
	.icon-f-vip{
		font-size: 18px;
		margin-left: 10px;
	}
	.user-vip-list{
		margin-top: 20px;
		margin-left: 10px;
		margin-right: 10px;
		background-color: #ffffff;
		border-radius: 10px;
		padding: 0 10px;
	}
	.user-vip-item{
		justify-content: space-between;
		height: 45px;
		border-bottom: 1px solid #d4d4d4;
		color: #464646;
		font-size: 14px;
	}
	.user-vip-item:last-child{
		border-bottom: none;
	}
	.vip-list{
		display: flex;
		white-space: nowrap;
		width: 100%;
		background-color: #ffffff;
	}
	.vip-item{
		padding: 20px;
		display: inline-flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		background-color: #d5d5d67a;
		margin-right: 10px;
		border-radius: 5px;
	}
	.vip-item-active{
		background-color: #f9efdd;
	}
	.vip-price{
		color: #e05211;
		font-size: 16px;
		font-weight: bold;
	}
	.vip-title{
		font-size: 18px;
		color: #464646;
		font-weight: bold;
	}
	.buy-btn{
		height: 40px;
		background-color: #040404bd;
		color: #ffffff;
		font-size: 16px;
		border-radius: 5px;
		display: flex;
		justify-content: center;
		align-items: center;
		margin: 0 20px;
	}
</style>
