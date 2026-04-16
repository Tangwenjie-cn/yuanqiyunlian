<template>
	<view class="container-sort">
		<yq-search></yq-search>
		<div class="sort-block">
			<scroll-view scroll-y="true" class="one-sort column" :enable-flex="true">
				<div v-for="(item,index) in list"
				class="one-sort-item center" 
				:class="{'one-sort-selected':selectedIndex===index}" 
				@click="clickOne(index,item.id)">{{item.name}}</div>
			</scroll-view>
			<div class="children-sort column" v-if="theme.type==1">
				<scroll-view scroll-y="true" class="sort-y-scroll" :enable-flex="true">
					<div v-for="(item1,index) in list[selectedIndex].children">
						<div class="two-title-block row-center">
							<div class="two-line"></div>
							<div @click="uni.navigateTo({
								url:'/pages/goods/index?sort_id='+item1.id
							})">{{item1.name}}</div>
							<div class="two-line"></div>
						</div>
						<div class="three-block">
							<div v-for="item2 in item1.children" 
							@click="uni.navigateTo({
								url:'/pages/goods/index?sort_id='+item2.id
							})" class="column-row-center">
								<image :src="item2.thumb" class="three-img"></image>
								<div class="three-title">{{item2.name}}</div>
							</div>
						</div>
					</div>
				</scroll-view>
			</div>
			<div class="children-sort column" v-if="theme.type==2">
				<scroll-view scroll-x="true" class="sort-x-scroll" :enable-flex="true">
					<div style="display: flex;">
						<div v-for="(item,index) in list[selectedIndex].children" 
						class="sort-btn" 
						:style="item.id==searchData.sort_id?'background:#000000':''"
						@click="getSortGoods(item.id)">
							{{item.name}}
						</div>
					</div>					
				</scroll-view>
				<scroll-view scroll-y="true" class="sort-y-scroll" @scrolltolower="scrolltolower" :enable-flex="true">
					<div class="column" style="width: 100%;">
						<div class="one-item"  
						v-for="item in goodsList" @click="uni.navigateTo({url:'/pages/goods/info?id='+item.id})">
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
				</scroll-view>				
			</div>
		</div>
		
		<yq-tab-bar ref="tabBar"></yq-tab-bar>
	</view>
</template>

<script setup>
	import {ref,nextTick} from 'vue'
	import { onShow } from '@dcloudio/uni-app'
	import store from '@/stores/index.js'
	import { request } from '/config/api.js'
	const theme = ref({})
	const tabBar = ref(null)
	onShow(()=>{
		theme.value=uni.getStorageSync('theme').sort[0]
		nextTick(()=>{
			if (tabBar.value){
				tabBar.value.autoMatchActiveTab()
			}
		})
		request('GoodsSort').then(res=>{
			list.value=res.data
			if(Object.hasOwn(res.data[0],'children')){
				searchData.sort_id=res.data[0].children[0].id
			}else{
				searchData.sort_id=res.data[0].id
			}
		})
		//要判断
		if(theme.value.type===2){
			getGoodsList()
		}
	})
	const loadMore=ref('more')
	const list = ref([])
	const goodsList=ref([])
	const selectedIndex = ref(0)
	const searchData=ref({
		page:1,
		limit:20,
		sort_id:0
	})
	function getSortGoods(i){
		searchData.value.page=1
		searchData.value.sort_id=i
		goodsList.value=[]
		getGoodsList()
	}
	function getGoodsList(){
		request('GoodsList',searchData.value).then(res=>{
			if(res.data.length==0){
				loadMore.value='noMore'
				return
			}
			goodsList.value=goodsList.value.concat(res.data)
		})
	}
	function clickOne(index,i){
		selectedIndex.value = index
		//要判断
		searchData.value.page=1
		searchData.value.sort_id=i
		goodsList.value=[]
		getGoodsList()
	}
	function scrolltolower(){
		searchData.value.page++
		getGoodsList()
	}
</script>

<style scoped>
	.container-sort{
		height: 100vh;
		background-color: #F8F8F8;
	}
	.sort-block{
		display: flex;
		height: calc(100vh - 102px);
	}
	.one-sort{
		width: 80px;
		height: 100%;
		overflow: auto;
		flex-shrink: 0;
	}
	.one-sort-item{
		width: 80px;
		height: 40px;
		text-overflow: ellipsis;
		overflow: hidden;
		white-space: nowrap;
	}
	.one-sort-selected{
		background-color: #ffffff;
		color: #f5a468;
	}
	.children-sort{
		min-width: 0;
		width: 100%;
		background-color: #ffffff;
	}
	.two-title-block{
		width: 100%;
		margin-top: 10px;
		margin-bottom: 10px;
	}
	.two-line{
		flex: 1;
		height: 1px;
		background-color: #a6a6a6;
		margin: 0 20px;
	}
	.three-block{
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		justify-items: center;
		margin-top: 10px;
		margin-bottom: 10px;
		row-gap: 10px;
	}
	.three-title{
		margin-top: 6px;
		font-size: 12px;
	}
	.three-img{
		width: 65px;
		height: 65px;
		border-radius: 5px;
		object-fit: cover;
	}
	.sort-x-scroll{
		white-space: nowrap;
		width: 100%;
	}
	.sort-y-scroll{
		height: 100%;
		overflow: auto; 
	}
	.sort-btn{
		border-radius: 15px;
		color: #fff;
		background-color: #a2a2a2;
		margin: 10px;
		padding: 5px 10px;
		font-size: 12px;
	}
	
	.one-item{
		display: flex;
		padding: 10px;
		background-color: #ffffff;
		border-top: 1px solid #00000059;
	}
	.one-item:first-child{
		margin-top: 8px;
	}
	.one-item:last-child{
		border-bottom: 1px solid #00000059;
	}
	.one-img{
		width: 120px;
		height: 120px;
		object-fit: cover;
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
	
</style>
