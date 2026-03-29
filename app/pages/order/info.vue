<template>
	<view>
		<div class="flex-column-center top-block">
			<image :src="info.thumb" v-if="info.type==1" class="thumb"></image>
			<span style="margin-top: 10px;">{{info.title}}</span>
			<span style="margin-top: 10px;color: #155ee9;font-size: 13px;" @click="uni.redirectTo({
				url:'/pages/article/index?id=1'
			})">联系客服</span>
		</div>
		<div class="column" style="padding-left: 20px;">
			<div>
				<span class="label-text">交易时间</span>
				<span class="info-text">{{info.ctime_text}}</span>
			</div>
			<div>
				<span class="label-text">订单状态</span>
				<span class="info-text">{{info.status==1?'已支付':'未支付'}}</span>
			</div>
			<div>
				<span class="label-text">支付方式</span>
				<span class="info-text" v-if="info.pay_type==1">微信</span>
				<span class="info-text" v-else-if="info.pay_type==2">支付宝</span>
				<span class="info-text" v-else-if="info.pay_type==3">卡密</span>
				<span class="info-text" v-else-if="info.pay_type==4">余额</span>
			</div>
			<div>
				<span class="label-text">支付金额</span>
				<span class="info-text" style="color: #fb2929;">{{info.pay_price}}</span>
			</div>
			<div>
				<span class="label-text">优惠金额</span>
				<span class="info-text" style="color: #e38e21;">{{info.discount_price}}</span>
			</div>
			<div>
				<span class="label-text">订单金额</span>
				<span class="info-text">{{info.price}}</span>
			</div>
			<div>
				<span class="label-text">订单编号</span>
				<span class="info-text">{{info.order_no}}</span>
			</div>
		</div>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	import { onLoad } from '@dcloudio/uni-app'
	import store from '@/stores/index.js'
	import { request } from '/config/api.js'
	import { formatIsoTime } from '/config/function.js'
	const info = ref({})
	onLoad((query) => {
		request('OrderInfo',{id:query.id}).then(res=>{
			info.value = res.data
		})
	})
</script>

<style scoped>
	.top-block{
		border-bottom: 1px solid #00000069;
		margin: 20px;
		padding-bottom: 20px;
	}
	.thumb{
		width: 79px;
		height: 79px;
		border-radius: 50%;
	}
	.label-text{
		color: #999;
		font-size: 13px;
		margin-right: 10px;
	}
	.info-text{
		color: #333;
		font-size: 13px;
	}
</style>
