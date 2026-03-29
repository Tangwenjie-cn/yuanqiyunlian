<template>
	<div :style="{background:`linear-gradient(${model.gradientStartColor}, ${model.gradientEndColor})`}">
		<div class="search column-center">
			<div class="search-input row-center">
				<uni-icons type="search" size="17" color="gray"></uni-icons>
				<input placeholder="请输入搜索内容" confirm-type="search" @confirm="search" />
			</div>
		</div>
		<div class="page">
			<swiper :indicator-dots="true" 
			:autoplay="true" 
			:interval="2000">
				<swiper-item v-for="item in model.list">
					<image :src="item.imgUrl" style="height:150px;width: 100%;" @click="swiperClick(item)"></image>
				</swiper-item>
			</swiper>
		</div>
	</div>
</template>
<script setup>
	import { ref } from 'vue'
	import {onShow} from '@dcloudio/uni-app'
	const model=defineModel()
	function swiperClick(item){
		uni.navigateTo({
			url:item.link
		})
	}
	function search(e){
		uni.navigateTo({
			url:`/pages/goods/index?keyword=${e.detail.value}`
		})
	}
</script>
<style scoped>
	.page{
		margin-left: 5px;
		margin-right: 5px;
		margin-bottom: 16px;
	}
	.search{
		width: 100%;
		margin-bottom: 10px;
	}
	.search-input{
		width: 70%;
		background-color: #fff;
		height: 36px;
		border-radius: 16px;
		padding-left: 10px;
	}
</style>
