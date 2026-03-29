<template>
	<div :style="{background:`linear-gradient(${model.gradientStartColor}, ${model.gradientEndColor})`}">
		<div class="page">
			<swiper :indicator-dots="true"
			:autoplay="true" 
			:interval="2000">
				<swiper-item v-for="item in model.list">
					<image :src="item.imgUrl" @click="swiperClick(item)"></image>
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
</script>
<style scoped>
	.page{
		margin-left: 5px;
		margin-right: 5px;
		margin-bottom: 16px;
	}
</style>
