<template>
	<view class="container">
		<div class="block nav-block">
			<div class="nav-item column-row-center">
				<image class="nav-image" src="/static/images/sign-icon-1.png" mode=""></image>
				<span>{{store().user.points}}积分</span>
			</div>
			<div class="nav-item column-row-center">
				<navigator url="/pages/points/sign">
					<image class="nav-image" src="/static/images/sign-in.png" mode=""></image>
				</navigator>
				<span>每日签到</span>
			</div>
			<div class="nav-item column-row-center">
				<navigator url="/pages/points/log">
					<image class="nav-image" src="/static/images/exchange.png" mode=""></image>
				</navigator>
				<span>积分记录</span>
			</div>
		</div>
		<div class="block content" style="margin-top: 20px;" v-if="vipList.length>0">
			<div class="top row-center">
				<span>兑换会员</span>
				<span @click="uni.navigateTo({
					url: '/pages/points/vipList'
				})">更多》</span>
			</div>
			<div class="vip-block column">
				<div class="vip-item column" v-for="item in vipList">
					<span style="font-size: 18px;font-weight: 500;">{{item.title}}</span>
					<span class="points-text">{{item.points}}积分</span>
					<span style="font-size: 14px;color: #333333cf;">{{item.desc}}</span>
					<span style="color: #fff;border-radius: 10px;background-color: #eb8415;height: 36px;margin: 7px 0;" 
					class="center" @click="buyVip(item.id)">兑换</span>
				</div>
			</div>
		</div>
		<div class="block content" style="margin-top: 20px;" v-if="goodsList.length>0">
			<div class="top row-center">
				<span>兑换商品</span>
				<span @click="uni.navigateTo({
					url: '/pages/points/goodsList'
				})">更多》</span>
			</div>
			<div class="column">
				<div class="goods-item" v-for="item in goodsList" @click="uni.navigateTo({
					url:'/pages/goods/info?id='+item.id
				})">
					<image class="one-img" :src="item.thumb"></image>
					<div class="one-right-block column">
						<span class="one-title">{{item.title}}</span>
						<div class="one-right-bottom row-center">
							<span class="points-text">{{item.points}}积分</span>
							<span style="color: #fff;border-radius: 10px;background-color: #eb8415;height: 32px;width: 80px;" 
					class="center" @click="buyGoods(item.id)">兑换</span>
						</div>
					</div>
				</div>
			</div>
		</div>
	</view>
</template>

<script setup>
	import {ref} from 'vue'
	import { onShow } from '@dcloudio/uni-app'
	import store from '@/stores/index.js'
	import { request } from '/config/api.js'
	const vipList = ref([])
	const goodsList=ref([])
	onShow(()=>{
		request('GetUserInfo').then(res=>{
			store().setUser(res.data)
		})
		request('PointsVipList',{page:1,limit:2}).then(res=>{
			vipList.value = res.data
		})
		request('GoodsList',{page:1,limit:2,type:5}).then(res=>{
			goodsList.value = res.data
		})
	})
	function buyVip(id){
		request('PointsBuy',{type:1,id:id},'post').then(res=>{
			uni.showToast({
				title: res.msg,
				icon: 'none'
			})
			uni.reLaunch({
				url: '/pages/vip/index'
			})
		})
	}
	function buyGoods(id){
		request('PointsBuy',{type:2,id:id},'post').then(res=>{
			uni.showToast({
				title: res.msg,
				icon: 'none'
			})
			uni.navigateTo({
				url: '/pages/goods/info?id='+id
			})
		})
	}
</script>

<style>
	.block{
		background-color: #ffffff;
		margin: 0 10px;
		border-radius: 10px;		
	}
	.nav-block{
		display: grid;
		grid-template-columns:repeat(3,1fr);
		padding: 10px;
	}
	.nav-item{
		margin-top: 6px;
	}
   .nav-image{
		width: 50px;
		height: 50px;
		border-radius: 10px;
		object-fit: cover;
   }
   .top{
		height: 48px;
		border-top-right-radius: 10rpx;
		border-top-left-radius: 10rpx;		
		justify-content: space-between;
		padding: 0 10px;
   }
   .points-text{
	   font-size: 16px;
	   color: #eb8415;
   }
   .vip-block{
		padding: 10px;
   }
   .vip-item{
	   border-top:1px solid #00000073;
	   padding-top: 6px;
   }
   .one-img{
		width: 160px;
		height: 120px;
   }
   .one-right-block{
		margin-left: 10px;
		height: 120px;
		justify-content: space-between;
		flex: 1;
   }
   .one-right-bottom{
   	justify-content: space-between;
   }
   .one-title{
   	font-size: 16px;
   	font-weight: 300;
   	display: -webkit-box;
   	-webkit-line-clamp: 2;
   	-webkit-box-orient: vertical;
   	overflow: hidden;
   	text-overflow: ellipsis;
   }
   .goods-item{
	   display: flex;
	   border-top:1px solid #00000073;
	   padding: 6px;
   }
</style>
