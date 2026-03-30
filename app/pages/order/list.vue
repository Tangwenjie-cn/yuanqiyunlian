<template>
	<view>
		<div class="row-center tab-block">
			<div class="tab-item center" 
			@click="tabClick(1)"
			:class="{'tab-item-active':searchData.type==1}">普通订单</div>
			<div class="tab-item center" 
			@click="tabClick(2)"
			:class="{'tab-item-active':searchData.type==2}">会员订单</div>
		</div>
		<div class="order-item" v-for="item in list" @click="uni.navigateTo({
			url:'/pages/order/info?id='+item.id
		})">
			<image v-if="item.type==1" :src="item.thumb" class="order-thumb"></image>
			<div class="column order-block">
				<span class="order-title">{{item.title}}</span>
				<span class="order-text">订单号：{{item.order_no}}</span>
				<div class="order-info">
					<span class="order-text">{{formatIsoTime(item.ctime)}}</span>
					<div class="column">
						<span class="order-success" v-if="item.status==1">已支付</span>
						<span class="order-wait" v-else-if="item.status==0">待支付</span>
					</div>
				</div>
				
			</div>
		</div>
		<uni-load-more :status="loadMore"></uni-load-more>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	import { onShow,onReachBottom } from '@dcloudio/uni-app'
	import store from '@/stores/index.js'
	import { request } from '/config/api.js'
	import { formatIsoTime } from '/config/function.js'
	const config=store().config
	const searchData=ref({
		type:1,
		page:1,
		limit:10
	})
	const loadMore=ref('more')
	const list=ref([])
	function tabClick(type){
		searchData.value.type=type
		searchData.value.page=1
		list.value=[]
		getList()
	}
	onShow(()=>{
		searchData.value.page=1
		list.value=[]
		getList()
	})
	function getList(){
		loadMore.value='loading'
		request('OrderList',searchData.value).then(res=>{
			if(res.data.length==0){
				loadMore.value='noMore'
				return
			}
			list.value=list.value.concat(res.data)
			loadMore.value='more'
		})
	}
	onReachBottom(()=>{
		if(loadMore.value=='noMore') return
		loadMore.value='loading'
		searchData.value.page++
		getList()
	})
</script>

<style>
	.tab-block{
		position: relative;
		top: 0;
		left: 0;
		right: 0;
		z-index: 999;
		
	}
   .tab-item{
		width: 50%;
		border-right: 1px solid #b5b5b5;
		height: 30px;
		margin: 4px;
   }
   .tab-item:last-child{
		border-right: none;
	}
	.tab-item-active{
		color: #007aff;
	}
	.order-item{
		width: 100%;
		display: flex;
		border-bottom: 1px solid #a1a1a1;
	}
	.order-item:last-child{
		border-bottom: none;
	}
	.order-thumb{
		width: 79px;
		height: 79px;
		margin: 5px;
		border-radius: 10px;
	}
	.order-block{
		flex: 1;
		margin: 5px;
		justify-content: space-between;
	}
	.order-info{
		display: flex; 
		justify-content: space-between;
	}
	.order-title{
		display: -webkit-box;
		-webkit-line-clamp: 2;
		-webkit-box-orient: vertical;
		overflow: hidden;
		text-overflow: ellipsis;
		font-size: 15px;
	}
	.order-text{
		font-size: 13px;
		color: #b5b5b5;
	}
	.order-success{
		font-size: 13px;
		color: #0aa50a;
	}
	.order-wait{
		font-size: 13px;
		color: #ffb010;
	}
</style>
