<template>
	<div class="content">
		<uni-collapse v-if="payData.discountList">
			<uni-collapse-item :title="payData.discountList[discountIndex].name+' - 优惠￥'+payData.discountList[discountIndex].price">
				<radio-group @change="discountChange">
					<label class="row-center" style="height: 36px;" v-for="(item,index) in payData.discountList">
						<radio :value="index" style="transform:scale(0.6);" color="#fdcd3d" :checked="discountIndex==index" />
						<view>{{item.name}} - 优惠￥{{item.price}}</view>
					</label>
				</radio-group>
			</uni-collapse-item>
		</uni-collapse>
		<radio-group @change="radioChange">
			<label class="row-center label-line" v-if="config.wxpay_on==1">
				<radio value="1" checked />
				<view><span class="iconfont icon-weixinzhifu"></span>微信</view>
			</label>
			<label class="row-center label-line" v-if="checkAlipay()">
				<radio value="2" />
				<view><span class="iconfont icon-zhifubao"></span>支付宝</view>
			</label>
			<label class="row-center label-line">
				<radio value="3" />
				<view><span class="iconfont icon-qiamizhifu"></span>卡密</view>
			</label>
			<label class="row-center label-line" v-if="config.v_name!='yq-free'">
				<radio value="4" />
				<view><span class="iconfont icon-kuaijiezhifu"></span>余额  ￥{{store().user.balance}}</view>
			</label>
		</radio-group>
		<div class="pay-block">
			<view class="pay-btn" @click="pay()">立即支付</view>
			<span class="pay-text">实付￥{{getpayPrice().toFixed(2)}}</span>	
		</div>
	</div>	
</template>
<script setup>
	import { ref,watch } from 'vue'
	import store from '/stores/index.js'
	import { request } from '/config/api.js'
	const payData=defineModel()
	const config=store().config
	request('GetUserInfo').then(res=>{
		store().setUser(res.data)
	})
	let pay_type=1
	let platform
	let discountIndex=ref(0)
	function radioChange(e){
		pay_type=Number(e.detail.value) 
	}
	function discountChange(e){
		discountIndex.value=Number(e.detail.value) 		
	}
	function checkAlipay(){
		if(config.alipay_on==1){
			// #ifdef MP-WEIXIN
			return false
			// #endif
			return true
		}
		return false
	}
	function getpayPrice(){
		if(payData.value.discountList){
			return payData.value.price - payData.value.discountList[discountIndex.value].price
		}else{
			return payData.value.price
		}
	}
	function pay(){
		let data={
			id:payData.value.id,
			pay_type:pay_type,
			type:payData.value.type,
			discount_type:0,
			discount_id:0,
		}
		
		if(payData.value.discountList){
			data.discount_id=payData.value.discountList[discountIndex.value].id
			data.discount_type=payData.value.discountList[discountIndex.value].type
		}
		switch(pay_type){
			case 1:
				weixin(data)
				break
			case 2:
				zhifubao(data)
				break
			case 3:
				kami(data)
				break
			case 4:
				yue(data)
				break			
			default:
			uni.showToast({
					title:'请选择支付方式',
					icon:'none'
				})
				break
		}
	}
	function yue(data){
		if(store().user.balance<getpayPrice()){
			uni.showToast({
				title:'余额不足',
				icon:'none'
			})
			return
		}
		request('OrderCreate',data,'post').then(res=>{
			if(data.type==1){
				uni.redirectTo({
					url:'/pages/goods/info?id='+data.id
				})
			}else{
				uni.redirectTo({
					url:'/pages/vip/index'
				})
			}
		})
	}
	function kami(data){
		uni.showModal({
			title: '卡密支付',
			placeholderText:'请输入卡密',
			editable:true,
			cancelText:'联系客服',
			cancelColor:'#00a0e9',
			success: function (res) {
				if (res.confirm) {
					if(!res.content){
						uni.showToast({
							title:'请输入卡密',
							icon:'none'
						})
						return
					}
					data.key=res.content
					request('OrderCreate',data,'post').then(res=>{
						if(data.type==1){
							uni.redirectTo({
								url:'/pages/goods/info?id='+data.id
							})
						}else{
							uni.redirectTo({
								url:'/pages/vip/index'
							})
						}
					})
				} else if (res.cancel) {
					uni.redirectTo({
						url:'/pages/article/index?id=1'
					})
				}
			}
		})
	}
	function zhifubao(data){
		request('OrderCreate',data,'post').then(res=>{
			// #ifdef H5
			return window.location.href=res.data
			// #endif
		})
	}
	function weixin(data){
		request('OrderCreate',data,'post').then(res=>{
			// #ifdef MP-WEIXIN
			wxJsPay(res.data,data)
			// #endif
			// #ifdef H5
			if(navigator.userAgent.includes('MicroMessenger')){
				wxJsPay(res.data,data)
			}
			// #endif
		})
	}
	function wxJsPay(paySign,data){
		uni.requestPayment({
			timeStamp: paySign.TimeStamp,
			nonceStr: paySign.NonceStr,
			package: paySign.Package,
			signType: paySign.SignType,
			paySign: paySign.PaySign,
			
			success: function (res) {
				if(data.type==1){
					uni.redirectTo({
						url:'/pages/goods/info?id='+data.id
					})
				}else{
					uni.redirectTo({
						url:'/pages/vip/index'
					})
				}
			},
			fail: function (res) {
				console.log(res)
				uni.showToast({
						title:'支付失败',
						icon:'none'
					})
			}
		})
	}
</script>
<style scoped>
	@import url("../../../../static/font/iconfont.css");
	radio{
		transform:scale(0.8);
	}
	.label-line{
		background-color: #fff;
		border-bottom: 1px #0000004f solid;
		height: 50px;
		padding-left: 10px;
	}
	.iconfont{
		margin: 0 6px;
	}
	.icon-weixinzhifu{
		color: #15ba11;
	}
	.icon-zhifubao{
		color: #00a0e9;
		font-size: 18px;
	}
	.icon-qiamizhifu{
		color: #ff6600;
	}
	.icon-kuaijiezhifu{
		color: #ff0000;
	}
	.pay-block{
		display: flex;
		height: 42px;
		background-color: #ffffff;
		flex-direction: row-reverse;
		align-items: center;
	}
	.pay-btn{
		width: 100px;
		height: 30px;
		background-color: #ff6417;
		border-radius: 5px;
		color: #fff;
		text-align: center;
		line-height: 30px;
		margin-right: 10px;
	}
	.pay-text{
		color: #ff6417;
		font-size: 14px;
		margin-right: 10px;
		font-weight: bold;
	}
	.content{
		position: fixed;
		bottom: 0;
		left: 0;
		right: 0;
	}
</style>
