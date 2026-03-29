<template>
	<view class="container">
		<div class="block">
			<div class="column">
				<span style="margin-bottom: 6px;">{{orderData.info.title}}</span>
				<span class="text-msg">{{orderData.info.desc}}</span>
				<span class="text-msg">时长:{{orderData.info.days}}天</span>
				<span class="text-msg" v-if="orderData.info.rebate>0"> 购买折扣:{{orderData.info.rebate}}折</span>
				<span class="price">价格:￥{{orderData.price}}</span>
			</div>
		</div>
		<yq-pay v-model="orderData"></yq-pay>
	</view>
</template>

<script setup>
	import { onShow,onLoad } from '@dcloudio/uni-app'
	import {ref} from 'vue'
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
		request('InitOrder',{id:id,type:2},'post').then(res=>{
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
	.price{
		color: #e05211;
		font-size: 14px;
		font-weight: 600;
	}
	.text-msg{
		font-size: 12px;
		color: #999;
		margin-bottom: 6px;
	}
</style>
