<template>
	<view class="container">
		<div class="block">
			<image class="goods-img" :src="orderData.info.thumb"></image>
			<div class="column goods-info">
				<span>{{orderData.info.title}}</span>
				<div><span class="price">价格：￥{{orderData.price}}</span></div>
			</div>
		</div>
		<yq-pay v-model="orderData"></yq-pay>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	import { onLoad } from '@dcloudio/uni-app'
	import store from '@/stores/index.js'
	import { request } from '/config/api.js'
	const orderData=ref({})
	onLoad((query)=>{
		let id=Number(query.id)
		if(!id){
			uni.showToast({title:'参数错误',icon:'none'})
			setTimeout(()=>{
				uni.navigateBack()
			},1000)
		}
		request('InitOrder',{id:id,type:1},'post').then(res=>{
			orderData.value=res.data
		})
	})
</script>

<style scoped>
	.block{
		margin: 0 5px;
		background-color: #fff;
		display: flex;
		padding: 6px;
		border-radius: 10px;
	}
	.goods-img{
		width: 99px;
		height: 99px;
	}
	.goods-info{
		margin-left: 10px;
		justify-content: space-between;
	}
	.price{
		color: #e05211;
		font-size: 14px;
		font-weight: 600;
	}
</style>
