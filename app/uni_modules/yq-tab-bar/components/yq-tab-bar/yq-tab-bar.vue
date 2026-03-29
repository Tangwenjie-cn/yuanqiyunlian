<template>
	<div class="content" :style="{backgroundColor:store().theme.tabBar[0].backgroundColor}">
		<div class="tab-bar-item" v-for="(item,index) in store().theme.tabBar[0].list" @click="tabBarClick(index)">
			<image class="tab-bar-img" :src="store().tabBarSelectedIndex==index ? item.selectImgUrl :item.imgUrl"></image>
			<span class="tab-bar-text" :style="setText(index)">{{item.name}}</span>
		</div>
	</div>
</template>

<script setup>
	import { ref,watch } from 'vue'
	import store from '@/stores/index.js'

	const tabBarClick = (index) => {
		store().tabBarSelectedIndex = index
		uni.reLaunch({
			url: store().theme.tabBar[0].list[index].link
		});
	}
	const setText = (index) => {
		if(store().tabBarSelectedIndex == index) {
			return `color: ${store().theme.tabBar[0].selectColor};`
		}else {
			return `color: ${store().theme.tabBar[0].color};`
		}
	}
</script>

<style scoped>
	.content{
		display: flex;
		height: 48px;
		width: 100%;
		position: fixed;
		bottom: 0;
		left: 0;
		right: 0;
		z-index: 999;
		justify-content: space-around;
		align-items: center;
	}
	.tab-bar-item{
		display: flex;
		flex-direction: column;
		align-items: center;
	}
	.tab-bar-img{
		width: 24px;
		height: 24px;
	}
	.tab-bar-text{
		font-size: 12px;
	}
</style>