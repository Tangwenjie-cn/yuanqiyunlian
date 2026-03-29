<template>
	<view>
		<div class="list-item" v-for="item in list">
			<div class="column list-block">
				<div class="row-center">
					<image class="list-img" :src="item.avatar"></image>
					<span class="list-title">{{item.nickname}}</span>
				</div>								
				<div class="list-info">
					<span class="list-text">UID：{{item.id}}</span>
					<span class="list-text">{{formatIsoTime(item.ctime)}}</span>
				</div>
			</div>
		</div>
		<uni-load-more :status="loadMore"></uni-load-more>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	import { onShow,onReachBottom } from '@dcloudio/uni-app'
	import { request } from '/config/api.js'
	import { formatIsoTime } from '/config/function.js'
	const searchData=ref({
		page:1,
		limit:10
	})
	const loadMore=ref('more')
	const list=ref([])
	onShow(()=>{
		searchData.value.page=1
		list.value=[]
		getList()
	})
	function getList(){
		loadMore.value='loading'
		request('MyPromotion',searchData.value).then(res=>{
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

<style scoped>
	.list-item{
		width: 100%;
		display: flex;
		border-bottom: 1px solid #a1a1a1;
	}
	.list-item:last-child{
		border-bottom: none;
	}
	.list-block{
		flex: 1;
		margin: 10px;
		justify-content: space-between;
	}
	.list-info{
		display: flex; 
		justify-content: space-between;
		margin-top: 10px;
	}
	.list-title{
		display: -webkit-box;
		-webkit-line-clamp: 2;
		-webkit-box-orient: vertical;
		overflow: hidden;
		text-overflow: ellipsis;
		font-size: 18px;
	}
	.list-text{
		font-size: 13px;
		color: #b5b5b5;
	}
	.list-img{
		width: 32px;
		height: 32px;
		border-radius: 50%;
		margin-right: 10px;
	}
</style>
