<template>
	<view class="content">
		<uni-notice-bar class="notice" :text="notice.title" 
		:color="model.color" :showIcon="true" 
		:scrollable="true" :showGetMore="true" 
		:background-color="model.backgroundColor"
		@click="noticeClick"></uni-notice-bar>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	import {onShow} from '@dcloudio/uni-app'
	import { request } from '/config/api.js'
	const model=defineModel()
	const notice=ref({})
	request('ArticleList',{'sort_id':2,'page':1,'limit':1}).then((res)=>{
		notice.value=res.data[0]
	})
	function noticeClick(){
		uni.navigateTo({
			url:'/pages/article/index?id='+notice.value.id
		})
	}
</script>

<style scoped>
	.content{
		padding-left: 5px;
		padding-right: 5px;
		margin-bottom: 16px;
	}
	.notice{
		border-radius: 10rpx;
	}
</style>