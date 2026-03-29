<template>
	<view class="container">
		<div class="column" style="width: 100%;">
			<div class="one-item"  
			v-for="item in list" @click="uni.navigateTo({url:'/pages/goods/info?id='+item.id})">
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
		<uni-load-more :status="loadMore"></uni-load-more>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	import { onShow,onReachBottom } from '@dcloudio/uni-app'
	import { request } from '/config/api.js'
	const searchData=ref({
		page:1,
		limit:10,
		type:5
	})
	const loadMore=ref('more')
	const list=ref([])
	onShow(()=>{
		searchData.value.page=1
		list.value=[]
		getList()
	})
	function getList(){
		loadMore.value='loading'
		request('GoodsList',searchData.value).then(res=>{
			if(res.data.length==0){
				loadMore.value='noMore'
				return
			}
			list.value=list.value.concat(res.data)
			loadMore.value='more'
		})
	}
	onReachBottom(()=>{
		if(loadMore.value=='noMore') return
		loadMore.value='loading'
		searchData.value.page++
		getList()
	})
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
	.points-text{
		   font-size: 16px;
		   color: #eb8415;
	}
	.one-item{
		display: flex;
		border-radius: 10rpx;
		padding: 10px;
		background-color: #ffffff;
		margin-bottom: 5px;
	}
	.one-item:first-child{
		margin-top: 8px;
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
</style>
