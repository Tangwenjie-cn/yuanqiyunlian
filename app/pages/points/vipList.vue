<template>
	<view>
		<div class="vip-item column" v-for="item in list">
			<span style="font-size: 18px;font-weight: 500;">{{item.title}}</span>
			<span style="font-size: 16px;color: #eb8415;">{{item.points}}积分</span>
			<span style="font-size: 14px;color: #333333cf;">{{item.desc}}</span>
			<span style="color: #fff;border-radius: 10px;background-color: #eb8415;height: 36px;margin: 7px 0;" 
			class="center" @click="buyVip(item.id)">兑换</span>
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
		limit:10
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
		request('PointsVipList',searchData.value).then(res=>{
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
</script>

<style>
	.vip-item{
	   border-bottom:1px solid #00000073;
	   padding: 6px 10px;
	}       
</style>
