<template>
	<view>
		<div class="list-item" v-for="(item,index) in list" >
			<image :src="item.thumb" class="thumb" @click="uni.navigateTo({
			url:'/pages/goods/info?id='+item.gid
		})"></image>
			<div class="column list-block">
				<span class="list-title" @click="uni.navigateTo({
			url:'/pages/goods/info?id='+item.gid
		})">{{item.title}}</span>				
				<div class="list-info row-center">
					<span class="list-text">{{formatIsoTime(item.ctime)}}</span>
					<span style="font-size: 14px;color: #F56C6C;" @click="delCollect(item.gid,index)">取消收藏</span>
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
		request('CollectList',searchData.value).then(res=>{
			if(res.data.length==0){
				loadMore.value='noMore'
				return
			}
			list.value=list.value.concat(res.data)
			loadMore.value='more'
		})
	}
	function delCollect(id,index){
		request('Collect',{gid:id},'POST').then(res=>{
			uni.showToast({
				title:res.msg,
				icon:'none'
			})
			list.value.splice(index,1)
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
	.thumb{
		width: 89px;
		height: 89px;
		margin: 10px;
		border-radius: 10px;
	}
	.list-block{
		flex: 1;
		margin: 10px;
		justify-content: space-between;
	}
	.list-info{ 
		justify-content: space-between;
	}
	.list-title{
		display: -webkit-box;
		-webkit-line-clamp: 2;
		-webkit-box-orient: vertical;
		overflow: hidden;
		text-overflow: ellipsis;
		font-size: 15px;
	}
	.list-text{
		font-size: 13px;
		color: #b5b5b5;
	}
</style>