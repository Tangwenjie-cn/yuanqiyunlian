<template>
	<view class="container">
		<div class="top">
			<span class="top-title">{{info.title}}</span>
			<span class="top-time">{{ formatIsoTime(info.utime) }}</span>
		</div>		
		<mp-html :content="info.content" style="display: flex;padding: 0 5px;width: 100%;box-sizing: border-box;"></mp-html>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	import {onLoad,onShow} from '@dcloudio/uni-app'
	import { request } from '@/config/api.js'
	import store from '@/stores/index.js'
	import {formatIsoTime} from '@/config/function.js'
	const info=ref({})
	onShow(()=>{
		// store().tabBar.list.forEach((item,index)=>{
		// 	if(item.path==='/pages/article/index'){
		// 		store().tabBar.selectedIndex=index
		// 	}
		// })
	})
	onLoad((option)=>{
		request('ArticleInfo',{id:option.id}).then((res)=>{
			info.value=res.data
			uni.setNavigationBarTitle({
			  title: res.data.title,
			})
		})
	})
</script>

<style scoped>
	.top{
		display: flex;
		flex-direction: column;
		margin-bottom: 10px;
		align-items: center;
	}
	.top-title{
		font-weight: bold;
	}
	.top-time{
		font-size: 14px;
		color: #999;
	}
</style>
