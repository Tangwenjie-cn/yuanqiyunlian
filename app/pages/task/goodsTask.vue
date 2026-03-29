<template>
	<view>
		<div class="list-item" v-for="item in list">
			<div class="column list-block" @click="uni.navigateTo({
				url: '/pages/goods/info?id='+item.gid
			})">
				<span class="list-title">{{item.title}}</span>				
				<div class="list-info">
					<div class="flex-column-center">
						<span class="list-text">分享要求：{{item.share}}次，已完成：{{item.share_nums}}次</span>
						<span class="list-text"></span>
					</div>
					<div class="flex-column-center">
						<template v-if="checkTask(item)">
							<span class="list-text-s">已完成</span>
						</template>
						<template v-else>
							<span class="list-text-e">未完成</span>
						</template>
						<span class="list-text">{{formatIsoTime(item.ctime)}}</span>
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
		request('GetUserGoodsTask',searchData.value).then(res=>{
			if(res.data.length==0){
				loadMore.value='noMore'
				return
			}
			list.value=list.value.concat(res.data)
			loadMore.value='more'
		})
	}
	function checkTask(v){
		if (v.adv<=v.adv_nums && v.share<=v.share_nums){
			return true
		}
		return false
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
	.list-text-s{
		font-size: 16px;
		color: #00aa00;
	}
	.list-text-e{
		font-size: 16px;
		color: #ff0000;
	}
</style>
