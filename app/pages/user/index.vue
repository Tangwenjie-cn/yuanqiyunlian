<template>
	<view class="container">
		<template v-for="item in theme">
			<template v-if="item.type==='status'">{{ setStatus(item.data) }}</template>
			<yq-user v-else-if="item.type==='user'" v-model="item.data"></yq-user>
			<yq-block v-else-if="item.type==='operate'" v-model="item.data"></yq-block>
			<yq-image v-else-if="item.type==='image'" v-model="item.data"></yq-image>			
		</template>
		
		<yq-tab-bar></yq-tab-bar>
	</view>
	
</template>

<script setup>
	import { ref,watch } from 'vue'
	import { onShow } from '@dcloudio/uni-app'
	import store from '@/stores/index.js'
	import yqUser from '/pages/user/components/user.vue'
	import yqBlock from '/pages/user/components/block.vue'
	import yqImage from '/pages/index/components/image.vue'
	const theme = ref([])
	if(Object.keys(store().theme).length>0){
		theme.value=store().theme.my
	}
	watch(store(),(newVal)=>{
		theme.value=newVal.theme.my
		let tabBar=newVal.theme.tabBar[0]
		for(let i=0;i<tabBar.list.length;i++){			
			if(tabBar.list[i].link==='/pages/user/index'){
				store().tabBarSelectedIndex=i
				break
			}
		}
	})
	onShow(()=>{
		if(Object.keys(store().theme).length>0){
			let tabBar=store().theme.tabBar[0]
			for(let i=0;i<tabBar.list.length;i++){			
				if(tabBar.list[i].link==='/pages/user/index'){
					store().tabBarSelectedIndex=i
					break
				}
			}
		}	
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
