<template>
	<view class="container">
		<template v-for="item in theme">
			<template v-if="item.type==='status'">{{ setStatus(item.data) }}</template>
			<yq-user v-else-if="item.type==='user'" v-model="item.data"></yq-user>
			<yq-block v-else-if="item.type==='operate'" v-model="item.data"></yq-block>
			<yq-image v-else-if="item.type==='image'" v-model="item.data"></yq-image>			
		</template>
		
		<yq-tab-bar ref="tabBar"></yq-tab-bar>
	</view>
	
</template>

<script setup>
	import { ref,watch,nextTick } from 'vue'
	import { onShow } from '@dcloudio/uni-app'
	import store from '@/stores/index.js'
	import yqUser from '/pages/user/components/user.vue'
	import yqBlock from '/pages/user/components/block.vue'
	import yqImage from '/pages/index/components/image.vue'
	const theme = ref([])
	const tabBar = ref(null)
	onShow(()=>{
		theme.value=uni.getStorageSync('theme').my
		nextTick(()=>{
			if (tabBar.value){
				tabBar.value.autoMatchActiveTab()
			}
		})	
	})
	function setStatus(data){
		uni.setNavigationBarColor({
			frontColor: data.color,
			backgroundColor: data.backgroundColor,
		})
		uni.setNavigationBarTitle({
		  title: data.title,
		})
	}
</script>

<style scoped>

</style>
