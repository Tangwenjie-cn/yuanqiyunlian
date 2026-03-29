<template>
	<view>
		<div class="list-item" v-for="item in list">
			<div class="row-center list-block">
				<div class="column">
					<span class="price">{{item.money}}</span>
					<span class="text-msg">手续费：{{item.fee}}</span>
					<span class="text-msg">实收：{{item.actual}}</span>
				</div>
				<div class="column">
					<span class="text-warning" v-if="item.status==0">待审核</span>
					<span class="text-success" v-if="item.status==1">完成</span>
					<span class="text-danger" v-if="item.status==-1">拒绝({{item.remark}})</span>
					<span class="text-msg" v-if="item.type==1">微信</span>
					<span class="text-msg" v-if="item.type==2">支付宝</span>
					<span class="text-msg" v-if="item.type==3">银行卡</span>
					<span class="text-msg">{{formatIsoTime(item.ctime)}}</span>
				</div>
			</div>
		</div>
		<uni-load-more :status="loadMore"></uni-load-more>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	import { onShow,onReachBottom,onLoad } from '@dcloudio/uni-app'
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
	onLoad(()=>{
		// #ifdef H5 || MP-WEIXIN
		checkWxTransfer()
		// #endif
	})
	function getList(){
		loadMore.value='loading'
		request('CashList',searchData.value).then(res=>{
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
	function checkWxTransfer(){
		request('GetWxTransfer').then(res=>{
			if(res.data){
				// #ifdef H5
				if(navigator.userAgent.includes('MicroMessenger')){
					wx.ready(function () {
						wx.checkJsApi({
							jsApiList: ['requestMerchantTransfer'],
							success: function (res1) {
								if (res1.checkResult['requestMerchantTransfer']) {
									WeixinJSBridge.invoke('requestMerchantTransfer', {
										mchId: res.data.mchId,
										appId: res.data.appId,
										package: res.data.package,
									},function (res2) {
										if (res2.err_msg === 'requestMerchantTransfer:ok') {
										    uni.showModal({
										    	content: '操作成功，如有疑问请联系客服',
												showCancel:false
										    })
										}
									});
								} else {
									alert('你的微信版本过低，请更新至最新版本。');
								}
							}
						});
					});
				}
				// #endif
				
				// #ifdef MP-WEIXIN
				if (wx.canIUse('requestMerchantTransfer')) {
				  wx.requestMerchantTransfer({
				    mchId: res.data.mchId,
				    appId: res.data.appId,
				    package: res.data.package,
				    success: (res1) => {
				        uni.showModal({
							content: '操作成功，如有疑问请联系客服',
							showCancel:false
						})
				    },
				    fail: (res2) => {
				      console.log('fail:', res2);
				    },
				  });
				} else {
				  wx.showModal({
					content: '你的微信版本过低，请更新至最新版本。',
					showCancel: false,
				  });
				}
				// #endif
			}
		})
	}
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
	.price{
		font-size: 16px;
		color: #ff2f14;
		font-weight: bold;
	}
	.text-msg{
		font-size: 14px;
		color: #a1a1a1;
	}
	.text-warning{
		font-size: 16px;
		color:#E6A23C;
	}
	.text-success{
		font-size: 16px;
		color:#67C23A;
	}
	.text-danger{
		font-size: 16px;
		color:#F56C6C;
	}
</style>
