<template>
	<div class="content" :style="{backgroundColor:tabBar.backgroundColor}">
		<div class="tab-bar-item" v-for="(item,index) in tabBar.list" @click="tabBarClick(index)">
			<image class="tab-bar-img" :src="tabbarSelectedIndex==index ? item.selectImgUrl :item.imgUrl"></image>
			<span class="tab-bar-text" :style="setText(index)">{{item.name}}</span>
		</div>
	</div>
</template>

<script setup>
	import { ref } from 'vue'
	import {onShow} from '@dcloudio/uni-app'
	import store from '@/stores/index.js'
	const tabbarSelectedIndex = ref(0)
	const tabBar= ref([])
	onShow(()=>{
		tabBar.value = uni.getStorageSync("theme").tabBar[0]
		tabbarSelectedIndex.value = uni.getStorageSync('tabBarSelectedIndex') || 0
	})
	const tabBarClick = (index) => {		
		if(tabbarSelectedIndex.value == index) {
			return
		}
		tabbarSelectedIndex.value = index
		uni.setStorageSync('tabBarSelectedIndex',index)
		uni.reLaunch({
			url: tabBar.value.list[index].link
		});
	}
	const setText = (index) => {
		if(tabbarSelectedIndex.value == index) {
			return `color: ${tabBar.value.selectColor};`
		}else {
			return `color: ${tabBar.value.color};`
		}
	}
	// 自动根据当前页面路径，匹配选中tab
	function  autoMatchActiveTab() {
	    // 获取当前页面路径
	    const pages = getCurrentPages()
	    if (!pages.length) return
	    
	    const currentPage = pages[pages.length - 1]
	    const currentPath = '/' + currentPage.route
	    // 找到对应 tab 的 index
	    const activeIndex = tabBar.value.list.findIndex(item => {
	      return item.link === currentPath
	    })
	
	    if (activeIndex !== -1) {
	      // 更新选中状态
	      tabbarSelectedIndex.value = activeIndex
	      uni.setStorageSync('tabBarSelectedIndex',activeIndex)
	    }
	}
	defineExpose({
		autoMatchActiveTab
	})
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