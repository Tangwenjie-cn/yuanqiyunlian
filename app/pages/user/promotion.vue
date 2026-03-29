<template>
	<view class="container">
		<div class="flex-column-center block">
			<template v-if="config.super_on==1&&config.v_name!='yq-free'">
				<span v-if="store().user.is_super==1" class="super">
					<span class="iconfont icon-xunzhang"></span>{{config.super_name}}</span>
				<span v-else class="sub-super" @click="uni.navigateTo({
					url:'/pages/user/subSuper'
				})">成为{{config.super_name}}<uni-icons type="right"></uni-icons></span>
			</template>
			
			<span class="title">{{store().user.promotion}}人</span>
			<div class="row-center btn-line">
				<button class="btn" @click="uni.navigateTo({
					url: '/pages/user/myPromotion'
				})">推广记录</button>
				<button class="btn" type="primary" @click="share">推广海报</button>
			</div>
		</div>
		<uni-card title="说明">
			<mp-html :content="desc"></mp-html>
		</uni-card>
		<uni-popup ref="popup">
			<yq-share :src="shareData.imgUrl" :path="shareData.path" :title="shareData.title"></yq-share>
		</uni-popup>
	</view>
</template>

<script setup>
	import {ref} from 'vue'
	import { onShow } from '@dcloudio/uni-app'
	import { request } from '/config/api.js'
	import store from '@/stores/index.js'
	const config=store().config
	const popup = ref()
	const desc=ref('')
	const shareData = ref({
		imgUrl: '',
		path: '/pages/index/index?pid='+store().user.id,
		title: '邀请您加入'
	})
	function share(){
		uni.showLoading({
			title: '加载中'
		})
		request('UserPromotionImage').then(res=>{
			uni.hideLoading()
			shareData.value.imgUrl = res.data
			popup.value.open()
		}).catch(err=>{
			uni.hideLoading()
		})
		
	}
	onShow(()=>{
		request('GetUserInfo').then(res=>{
			store().setUser(res.data)
		})
		request('ArticleInfo',{id:4}).then((res)=>{
			desc.value=res.data.content
		})
	})
</script>

<style scoped>
	@import url("../../static/font/iconfont.css");
	.block{
		background-color: #ffffff;
		height: 200px;
		margin: 0 20px;
		border-radius: 10px;
	}
	.title{
		font-size: 36px;
		font-weight: bold;
		color: #484848;
	}
	.btn{
		width: 100px;
	}
	.btn-line{
		justify-content: space-around;
		width: 100%;
		margin-top: 20px;
	}
	.super{
		font-size: 15px;
		color: #f5bb0d;
	}
	.sub-super{
		font-size: 15px;
		color: #f17b7b;
	}
</style>
