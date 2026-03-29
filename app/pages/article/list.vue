<template>
	<view>
		<div class="list-item" v-for="item in list">
			<div class="list-block row-center" @click="uni.navigateTo({
				url:'/pages/article/index?id='+item.id
			})">
				<span>{{item.title}}</span>
				<uni-icons type="right" size="22"></uni-icons>
			</div>		
		</div>
		<uni-load-more :status="loadMore"></uni-load-more>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	import { onReachBottom,onLoad } from '@dcloudio/uni-app'
	import { request } from '/config/api.js'
	const searchData=ref({
		page:1,
		limit:10,
		sort_id:0
	})
	const loadMore=ref('more')
	const list=ref([])
	onLoad((query)=>{
		if(query.sid){
			searchData.value.sort_id=query.sid
			searchData.value.page=1
			list.value=[]
			getList()
		}else{
			return uni.showToast({
				title:'参数错误',
				icon:'none'
			})
		}
		
	})
	function getList(){
		loadMore.value='loading'
		request('ArticleList',searchData.value).then(res=>{
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
</style>
