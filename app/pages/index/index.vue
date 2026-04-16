<template>
	<div class="container">
		<template v-for="item in theme">
			<template v-if="item.type==='status'">{{ setStatus(item.data) }}</template>
			<yq-search-swiper v-else-if="item.type==='search-carousel'" v-model="item.data"></yq-search-swiper>
			<yq-swiper v-else-if="item.type==='carousel'" v-model="item.data"></yq-swiper>
			<yq-search v-else-if="item.type==='search'" v-model="item.data"></yq-search>
			<yq-nav v-else-if="item.type==='classify'" v-model="item.data"></yq-nav>
			<yq-notice v-else-if="item.type==='notice'" v-model="item.data"></yq-notice>
			<yq-article v-else-if="item.type==='article'" v-model="item.data"></yq-article>
			<yq-image v-else-if="item.type==='image'" v-model="item.data"></yq-image>
			<yq-video v-else-if="item.type==='video'" v-model="item.data"></yq-video>
			<yq-goods v-else-if="item.type==='goods'" v-model="item.data"></yq-goods>
			<template v-else-if="item.type==='ad'">
				<!-- #ifdef MP-WEIXIN -->
				<ad :unit-id="4654"></ad>
				<!-- #endif -->
			</template>			
		</template>
		<yq-tab-bar ref="tabBar"></yq-tab-bar>
	</div>	
</template>

<script setup>
	import { nextTick, ref } from 'vue'
	import { onShow,onLoad,onShareAppMessage,onShareTimeline } from '@dcloudio/uni-app'
	import store from '@/stores/index.js'
	import yqSearchSwiper from '/pages/index/components/search-swiper.vue'
	import yqSwiper from '/pages/index/components/swiper.vue'
	import yqSearch from '/pages/index/components/search.vue'
	import yqNav from '/pages/index/components/nav.vue'
	import yqNotice from '/pages/index/components/notice.vue'
	import yqArticle from '/pages/index/components/article.vue'
	import yqImage from '/pages/index/components/image.vue'
	import yqVideo from '/pages/index/components/video.vue'
	import yqGoods from '/pages/index/components/goods.vue'	
	const theme = ref([])
	const tabBar = ref(null)
	onShareAppMessage((res)=>{
		if(res.from=='menu'){
			return {
				title:store().config.site_name,
				path: '/pages/index/index?pid='+store().user.id?store().user.id:0
			}
		}	
	})
	onShareTimeline(()=>{
		return {
			title:store().config.site_name,
			query: 'pid='+store().user.id?store().user.id:0
		}
	})
	onLoad((query)=>{	
		if(query.pid){
			uni.setStorageSync('pid',query.pid)
		}
		if(query.scene){
			decodeURIComponent(query.scene).split('&').forEach(item=>{
				if (!item) return;
				let [key, val] = item.split('=');
				uni.setStorageSync(key, val);
			})			
		}
	})
	onShow(()=>{
		theme.value=uni.getStorageSync('theme').index
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

<style>

</style>
