<template>
	<view class="container">
		<div class="column top">
			<span class="top-title">{{info.title}}</span>
			<span class="top-introduction">{{info.introduction}}</span>
			<div class="row-center top-info">
				<div class="column">
					<span class="top-label">时间：{{info.utime_text}}</span>
					<span class="top-label">浏览：{{info.views_text}}</span>
					<span class="top-label" v-if="info.type==1">总页数：{{info.pages}}</span>
				</div>
				<div class="column">					
					<div><span class="top-label">价格：</span><span class="price">{{info.price>0?info.price:'免费'}}</span></div>					
					<div v-if="info.original_price>0"><span class="top-label">原价：</span><span class="price-underline">{{info.original_price}}</span></div>
				</div>
			</div>
		</div>
		<div class="content" v-if="info.cid>0">
			<yq-column v-model="info.clist" :gid="goodsId"></yq-column>
		</div>
		<div class="content" v-if="info.is_auth">
			<yq-document v-if="info.type===1" v-model="info.link"></yq-document>
			<mp-html v-if="info.type===2" :content="info.goods_content.fee_content"></mp-html>
			<video v-else-if="info.type===3" :src="info.link"></video>
			<yq-audio v-else-if="info.type===4" :src="info.link"></yq-audio>
			<yq-cloud-storage v-else-if="info.type===5" v-model="info"></yq-cloud-storage>
		</div>
		<div class="content">
			<mp-html :content="info.goods_content.content"></mp-html>
		</div>
		<div class="nav-buttom">
			<uni-goods-nav :fill="true"  
			:options="options" 
			:buttonGroup="buttonGroup"  
			@click="gnClick" 
			@buttonClick="gnButtonClick" />
		</div>
		<uni-popup ref="popup">
			<yq-share :src="shareData.imgUrl" :path="shareData.path" :title="shareData.title"></yq-share>
		</uni-popup>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	import { onShow,onLoad,onShareAppMessage,onShareTimeline } from '@dcloudio/uni-app'
	import store from '@/stores/index.js'
	import { request } from '/config/api.js'
	import yqCloudStorage from './components/cloud-storage.vue'
	import yqColumn from './components/column.vue'
	import yqDocument from './components/document.vue'
	const info = ref({})
	const user=store().user
	const config=store().config
	let goodsId=0
	let pid=0
	const popup = ref()
	const shareData = ref({
		imgUrl: '',
		path: '/pages/goods/info?id='+goodsId+'&pid='+user.id?user.id:0,
		title: info.value.title
	})
	onShareAppMessage((res)=>{
		if(res.from=="menu"){
			return {
				title:info.value.title,
				path:'/pages/goods/info?id='+goodsId+'&pid='+user.id?user.id:0,
				imageUrl:info.value.thumb
			}
		}
	})
	onShareTimeline(()=>{
		return {
			title:info.value.title,
			query:'?id='+goodsId+'&pid='+user.id?user.id:0,
			imageUrl:info.value.thumb
		}
	})
	onLoad((query)=>{
		if(query.scene){
			decodeURIComponent(query.scene).split('&').forEach(item=>{
				if (!item) return;
				let [key, val] = item.split('=');
				query[key] = val
			})			
		}
		if(query.id){
			goodsId=Number(query.id) 
		}else{
			uni.showToast({
				title:'参数错误',
				icon:'none'
			})
			uni.navigateBack()
			return
		}
		if(query.pid){
			pid=query.pid
			uni.setStorageSync('pid',pid)
			uni.setStorageSync('gid',goodsId)
		}
		
	})
	onShow(()=>{
		getInfo()
	})
	function getInfo(){
		request('GoodsInfo',{id:goodsId,pid:pid}).then(res=>{
			info.value=res.data
			if(res.data.is_auth){ //判断权限
				buttonGroup.value=[]
				buttonGroup.value.push(botton3)
			}else{
				buttonGroup.value=[]
				buttonGroup.value.push(botton1)
				if(res.data.is_adv==1 || res.data.is_share==1){
					buttonGroup.value.push(botton2)
				}				
			}
			if(res.data.is_collect){ //判断是否收藏
				options.value[2].text='已收藏'
				options.value[2].icon='star-filled'					
			}else{
				options.value[2].text='收藏'
				options.value[2].icon='star'
			}
			uni.setNavigationBarTitle({
			  title: res.data.title,
			})
		})
	}
	const botton1={
		text: '立即购买',
		backgroundColor: '#ff2000',
		color: '#ffffff',
		action:'buy'
	}
	const botton2={
		text: '任务获取',
		backgroundColor: '#ffba0b',
		color: '#ffffff',
		action:'task'
	}
	const botton3={
		text: '已获取',
		backgroundColor: '#1ec304',
		color: '#ffffff',
		action:'none'
	}
	const options = ref([{
		icon: 'chat',
		text: '客服'
	}, {
		icon: 'undo',
		text: '分享',
	}, {
		icon: 'star',
		text: '收藏',
	}])
	const buttonGroup = ref([])
	const gnClick = (i) => { //图标点击
		switch(i.index){
			case 0:
				uni.navigateTo({
					url:'/pages/article/index?id=1'
				})
				break;
			case 1:
				uni.showLoading({
					title: '加载中'
				})
				request('GetGoodsPromotionImage',{gid:goodsId},'GET').then(res=>{
					uni.hideLoading()
					shareData.value.imgUrl = res.data
					popup.value.open()
				}).catch(err=>{
					uni.hideLoading()
				})
				break;
			case 2:
				request('Collect',{gid:goodsId},'POST').then(res=>{
					getInfo()
					uni.showToast({
						title:res.msg,
						icon:'none'
					})
				})
				break;
			default:
				break;
		}
	}
	const gnButtonClick = (e) => {
		switch(e.content.action){
			case 'buy':
				uni.navigateTo({
					url:'/pages/order/buy?id='+goodsId
				})
				break;
			case 'task':
				request('GetGoodsTask',{gid:goodsId}).then(res=>{
					let content='分享要求：'+res.data.share+'次，已完成：'+res.data.share_nums+'次'
					//content+='广告要求：'+res.data.adv+'次，已完成：'+res.data.adv_nums+'次\n'					
					uni.showModal({
						title: '任务要求',
						content:content,
						success:(res)=>{
							if(res.confirm){
								
							}
						}
					})
					// #ifdef MP-WEIXIN
					const targetTimestamp = new Date(res.data.ctime).getTime();
					const currentTimestamp = new Date().getTime();
					if (currentTimestamp - targetTimestamp < 10000){
						wx.requestSubscribeMessage({
							tmplIds: [config.wx_mini_task_templateId],
							success(res) {
								console.log(res)
							}
						})
					}
					// #endif
				})
				break;
			default:
				break;
		}		
	}
	onShow(()=>{
		
	})
</script>

<style scoped>
 .top{
	margin: 0 5px;
	background-color: #ffffff;
	border-radius: 10px;
	padding: 6px;
 }
 .top-info{
	justify-content: space-between;
 }
 .top-title{
	font-size: 16px;
	color: #1a1a1a;
	margin: 5px 0;
	font-weight: bold;
 }
 .top-label{
	font-size: 12px;
	color: #585858;
 }
 .price-underline{
 	color: #999;
 	font-size: 12px;
 	text-decoration: line-through;
 }
 .price{
 	color: #e05211;
 	font-size: 14px;
 	font-weight: 600;
 }
 .content{
	margin: 10px 5px 0;
	background-color: #ffffff;
	border-radius: 10px;
	padding: 6px;
 }
 .nav-buttom{
	 position: fixed;
	 bottom: 0;
	 left: 0;
	 right: 0;
 }
 video{
 	width: 100%;
 	border-radius: 10rpx;
 }
 .top-introduction{
    font-size: 14px;
    color: #000000d4;
	margin-bottom: 5px;
 }
</style>
