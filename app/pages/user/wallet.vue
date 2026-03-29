<template>
	<view class="container">
		<div class="flex-column-center block">
			<span class="price">{{store().user.balance}}</span>
			<div class="row-center btn-line">
				<button class="btn" @click="uni.navigateTo({
					url: '/pages/user/bill'
				})">记录</button>
				<button class="btn" type="primary" @click="uni.navigateTo({
					url:'/pages/cash/index'
				})">提现</button>
			</div>
		</div>
		<uni-card title="说明">
			<mp-html :content="desc"></mp-html>
		</uni-card>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	import { onShow } from '@dcloudio/uni-app'
	import store from '@/stores/index.js'
	import { request } from '/config/api.js'
	const desc=ref('')
	onShow(()=>{
		request('GetUserInfo').then(res=>{
			store().setUser(res.data)
		})
		request('ArticleInfo',{id:2}).then((res)=>{
			desc.value=res.data.content
		})
	})
</script>

<style scoped>
	.block{
		background-color: #ffffff;
		height: 200px;
		margin: 0 20px;
		border-radius: 10px;
	}
	.price{
		font-size: 36px;
		font-weight: bold;
		color: #ef5c5c;
	}
	.btn{
		width: 100px;
	}
	.btn-line{
		justify-content: space-around;
		width: 100%;
		margin-top: 20px;
	}
</style>
