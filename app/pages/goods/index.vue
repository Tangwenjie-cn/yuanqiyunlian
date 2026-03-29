<template>
	<view class="container">
		<view class="status_bar">
		    <!-- 这里是状态栏 -->
		</view>
		<div class="row-center" :style="{height:titleHeight}">
			<uni-icons type="left" size="25" color="gray" @click="uni.navigateBack()"></uni-icons>
			<div class="search-input row-center">
				<uni-icons type="search" size="17" color="gray"></uni-icons>
				<input placeholder="请输入搜索内容" v-model="searchData.keyword" confirm-type="search" @confirm="search" />
			</div>
		</div>
		<div class="row-center" style="height: 36px;">
			<span class="filter-text" :class="{'filter-selected-text':searchData.sort_id>0}" 
			@click="setSort()">{{sortName}}</span>
			<span class="filter-text" :class="{'filter-selected-text':searchData.type==1}" 
			@click="setSearch('hot')">热门</span>
			<span class="filter-text" :class="{'filter-selected-text':searchData.type==3}" 
			@click="setSearch('price')">价格</span>
			<span class="filter-text" :class="{'filter-selected-text':searchData.type==4}" 
			@click="setSearch('time')">最新</span>
		</div>
		
		<div class="column" style="width: 100%;">
			<div class="one-item"  
			v-for="item in list" @click="uni.navigateTo({url:'/pages/goods/info?id='+item.id})">
				<image class="one-img" :src="item.thumb"></image>
				<div class="one-right-block column">
					<span class="one-title">{{item.title}}</span>
					<div class="one-right-bottom row-center">
						<div class="column">
							<span class="price">{{item.price>0?'￥'+item.price:'免费'}}</span>
							<span class="price-underline" 
							v-if="item.original_price>0">￥{{item.original_price}}</span>
						</div>
						<div class="column">
							<span class="text">{{item.utime_text}}</span>
							<span class="text">
								<uni-icons type="eye-filled" size="14" color="#999"></uni-icons>
								{{item.views_text}}
							</span>
						</div>
					</div>
				</div>
			</div>
		</div>
		<uni-load-more :status="loadMore"></uni-load-more>
		<uni-drawer ref="drawerRef" mode="left" :width="320">
			<uni-section :title="sortName" titleFontSize="20px" titleColor="#fc4710" type="circle">
				<div class="sort-block">					
					<view class="btn-item" v-for="item in sortList">
						<button size="mini" type="default" @click="clickSort(item)">{{item.name}}</button>
					</view>	
				</div>
				<div style="display: flex;">
					<button @click="sortReset()">重置</button>
					<button type="primary" @click="sortSubmit()">确定</button>
				</div>
			</uni-section>			
		</uni-drawer>
	</view>
</template>

<script setup>
	import {ref} from 'vue'
	import {onShow,onReachBottom,onLoad} from '@dcloudio/uni-app'
	import { request } from '/config/api.js'
	const drawerRef=ref()
	const sortAllList=ref([])
	const sortList=ref([])
	const sortName=ref('分类')
	const list=ref([])
	const loadMore=ref('more')
	const titleHeight=ref('90px')
	const searchData=ref({
		page:1,
		limit:20,
		sort_id:0,
		keyword:'',
		type:0
	})
	onLoad((query)=>{
		if(query.type){
			searchData.value.type=query.type
		}
		if(query.keyword){
			searchData.value.keyword=query.keyword
		}
		if(query.sort_id){
			searchData.value.sort_id=query.sort_id
		}
	})
	onShow(()=>{
		searchData.value.page=1
		list.value=[]
		getList()
		getSortList()
		// #ifdef APP
		let navigationBarHeight=48
		// #endif
		// #ifndef APP
		let navigationBarHeight=44
		// #endif
		let h=uni.getSystemInfoSync().statusBarHeight+navigationBarHeight
		titleHeight.value=h+'px'		
	})
	function getSortList(){
		request('GoodsSort').then(res=>{
			sortList.value=res.data
			sortAllList.value=res.data
		})
	}
	function setSearch(type){
		searchData.value.type=type=='hot'?1:type=='price'?3:type=='time'?4:0
		searchData.value.page=1
		list.value=[]
		getList()
	}
	function getList(){
		loadMore.value='loading'
		request('GoodsList',searchData.value).then(res=>{
			if(res.data.length==0){
				loadMore.value='noMore'
				return
			}
			list.value=list.value.concat(res.data)
			loadMore.value='more'
		})
	}
	function search(e){
		searchData.value.keyword=e.detail.value
		searchData.value.page=1
		searchData.value.sort_id=0
		searchData.value.type=0
		list.value=[]
		getList()
	}
	function setSort(){
		drawerRef.value.open()
	}
	function clickSort(v){
		sortList.value=v.children
		sortName.value=sortName.value+'/'+v.name
		searchData.value.sort_id=v.id
		
	}
	function sortReset(){
		sortList.value=sortAllList.value
		sortName.value='分类'
		searchData.value.sort_id=0
	}
	function sortSubmit(){
		drawerRef.value.close()
		searchData.value.page=1
		list.value=[]
		getList()
	}
	onReachBottom(()=>{
		if(loadMore.value=='noMore') return
		loadMore.value='loading'
		searchData.value.page++
		getList()
	})
</script>

<style scoped>
	.status_bar {
	    height: var(--status-bar-height);
	    width: 100%;
	}
	.search{
		width: 100%;
		margin-bottom: 16px;
	}
	.search-input{
		margin-left: 5px;
		width: 55%;
		height: 31px;
		border-radius: 16px;
		padding-left: 10px;
		background-color: #ffffff;
	} 
	.filter-text{
		margin-left: 20px;
		font-size: 15px;
		color: #666666;
	}
	.filter-selected-text{
		color: #e56c13;
	}
	.one-item{
		display: flex;
		border-radius: 10rpx;
		padding: 10px;
		background-color: #ffffff;
		margin-bottom: 5px;
	}
	.one-item:first-child{
		margin-top: 8px;
	}
	.one-img{
		width: 160px;
		height: 120px;
	}
	.one-right-block{
		margin-left: 10px;
		height: 120px;
		justify-content: space-between;
		flex: 1;
	}
	.one-title{
		font-size: 16px;
		font-weight: 300;
		display: -webkit-box;
		-webkit-line-clamp: 2;
		-webkit-box-orient: vertical;
		overflow: hidden;
		text-overflow: ellipsis;
	}
	.one-right-bottom{
		justify-content: space-between;
	}
	.price{
		color: #e05211;
		font-size: 14px;
		font-weight: 600;
	}
	.price-underline{
		color: #999;
		font-size: 12px;
		text-decoration: line-through;
	}
	.sort-block{
		margin-top: 20px;
		padding: 20px;
		display: flex;
		flex-wrap: wrap;
	}
	.btn-item{
		margin-bottom: 10px;
		margin-right: 10px;
	}
</style>
