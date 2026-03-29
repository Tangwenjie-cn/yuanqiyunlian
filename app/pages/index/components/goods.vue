<template>
	<view class="content">
		<div class="top row-center" :style="{'background': `linear-gradient(to right, ${model.gradientStartColor}, ${model.gradientEndColor})`}">
			<span :style="{'color':model.color,'font-size':model.fontSize}">{{model.title}}</span>
			<span :style="{'color':model.color,'font-size':model.fontSize}" @click="navigateTo">更多》</span>
		</div>
		<!-- 商品列表 -->
		<view class="goods-list">
			<div class="two" v-if="model.arrange===2">
				<div class="two-item" 
				:style="{'background-color': model.backgroundColor}"
				v-for="item in list">
				<navigator :url="'/pages/goods/info?id='+item.id">
					<image class="two-img" :src="item.thumb"></image>
					<div class="two-block column"> 
						<span class="two-title">{{item.title}}</span>
						<div class="two-bottom row-center">
							<div class="column">
								<span class="price">{{item.price>0?'￥'+item.price:'免费'}}</span>
								<span class="price-underline" v-if="item.original_price>0">￥{{item.original_price}}</span>
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
				</navigator>
					
				</div>
			</div>
			
			<div class="column" v-if="model.arrange===1">				
				<navigator v-for="item in list" :url="'/pages/goods/info?id='+item.id">
					<div class="one-item"
					:style="{'background-color': model.backgroundColor}">
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
				</navigator>	
			</div>
		</view>
	</view>
</template>

<script setup>
	import {ref} from 'vue'
	import {onShow} from '@dcloudio/uni-app'
	import { request } from '/config/api.js'
	const model=defineModel()
	const list=ref([])
	request('GoodsList',{'sort_id':model.value.sortId,'page':1,'limit':model.value.limit}).then(res=>{
		list.value=res.data
	})
	function navigateTo(){
		uni.navigateTo({
			url:`/pages/goods/index?sort_id=${model.value.sortId}&type=${model.value.type}`
		})
	}
</script>

<style scoped>
	.content{
		padding-left: 5px;
		padding-right: 5px;
		margin-bottom: 16px;
	}
	.top{
		height: 48px;
		border-top-right-radius: 10rpx;
		border-top-left-radius: 10rpx;		
		justify-content: space-between;
		padding: 0 10px;
	}
	.one-item{
		display: flex;
		border-radius: 10rpx;
		padding: 10px;
		margin-top: 5px;
	}
	.one-img{
		width: 160px;
		height: 120px;
		object-fit: cover;
	}
	.one-right-block{
		margin-left: 10px;
		width: 60%;
		justify-content: space-between;
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
	.two{
		display: grid;
		grid-template-columns: repeat(2, 1fr);
		width: 100%;
		row-gap: 10px;
		justify-items: center;
		margin-top: 5px;
	}
	.two-item{
		border-radius: 10rpx;
		width: 330rpx;
	}
	.two-img{
		width: 330rpx;
		height: 156px;
		object-fit: cover;
		border-top-right-radius: 10rpx;
		border-top-left-radius: 10rpx;
	}
	.two-block{
		padding: 5px;
	}
	.two-title{
		font-size: 16px;
		color: #000000cc;  
		overflow: hidden;       
		text-overflow: ellipsis;
		white-space: nowrap;
		width: 160px;
		min-width: 0;
		font-weight: 300;
	}
	.two-bottom{
		justify-content: space-between;
	}
	.text{
		color: #999;
		font-size: 12px;
	}
</style>