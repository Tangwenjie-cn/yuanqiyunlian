<template>
	<view class="content">
		<div class="column" v-if="model.arrange===1">
			<navigator :url="'/pages/article/index?id='+item.id" v-for="item in list">
			<div class="article-one-item" 
			:style="{'background-color': model.backgroundColor}">				
				<image class="article-one-img" :src="item.thumb" mode=""></image>
				<div class="article-one-text column"> 
					<span class="article-one-title">{{item.title}}</span>
					<span class="article-time">{{item.utime_text}}</span>
				</div>
			</div>
			</navigator>								
			
		</div>
		<div class="article-two" v-if="model.arrange===2">
			<navigator :url="'/pages/article/index?id='+item.id" v-for="item in list">
			<div class="article-two-item" 
			:style="{'background-color': model.backgroundColor}" >			
				<image class="article-two-img" :src="item.thumb" mode=""></image>
				<div class="article-two-text column"> 
					<span class="article-two-title">{{item.title}}</span>
					<span class="article-time">{{item.utime_text}}</span>
				</div>
			</div>
			</navigator>
			
		</div>				
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	import {onShow} from '@dcloudio/uni-app'
	import { request } from '/config/api.js'
	const model=defineModel()
	const list=ref([])
	request('ArticleList',{'sort_id':model.value.sortId,'page':1,'limit':model.value.limit}).then(res=>{
		list.value=res.data
	})
</script>

<style scoped>
	.content{
		padding-left: 5px;
		padding-right: 5px;
		margin-bottom: 16px;
	}
	.article-one-item{
		border-radius: 10rpx;
		padding: 10px;
		display: flex;
		margin-bottom: 10px;
	}
	.article-one-img{
		width: 160px;
		height: 120px;
		object-fit: cover;
		border-radius: 10rpx;
	}
	.article-one-text{
		margin-left: 10px;
		justify-content: space-between;
		width: 60%;
	}
	.article-one-title{
		font-size: 16px;
		color: #000000cc;  
		overflow: hidden;       
		text-overflow: ellipsis; 
	}
	.article-time{
		font-size: 12px;
		color: #999999;
	}
	.article-two{
		display: grid;
		grid-template-columns: repeat(2, 1fr);
		width: 100%;
		row-gap: 10px;
		justify-items: center;
	}
	.article-two-item{
		border-radius: 10rpx;
		width: 330rpx;
	}
	.article-two-img{
		width: 100%;
		height: 156px;
		object-fit: cover;
		border-radius: 10rpx;
	}
	.article-two-text{
		width: 100%;
		padding: 10px;
	}
	.article-two-title{
		font-size: 16px;
		color: #000000cc;  
		overflow: hidden;       
		text-overflow: ellipsis;
		white-space: nowrap;
		width: 160px;
		min-width: 0; 
	}
</style>