<template>
	<view class="container">
		<div class="row-center cash-line">
			<span class="label-text">余额</span>
			<div class="row-center cash-line-right">
				<span>{{store().user.balance}}</span>
			</div>
			
		</div>
		<div class="row-center cash-line">
			<span class="label-text">提现金额</span>
			<div class="row-center cash-line-right">
				<input type="digit" v-model="cashData.money" placeholder="请输入提现金额" class="cash-input" />
			</div>
		</div>
		<div class="row-center cash-line">
			<span class="label-text">提现方式</span>
			<div class="row-center cash-line-right">
				<radio-group @change="radioChange">
					<label class="radio-label" v-if="config.cash_wxpay_type>0"><radio :value="1" :checked="cashData.type==1" />微信</label>
					<label class="radio-label" v-if="config.cash_alipay_type>0"><radio :value="2" :checked="cashData.type==2" />支付宝</label>
					<label class="radio-label" v-if="config.cash_bank_type>0"><radio :value="3" :checked="cashData.type==3" />银行卡</label>
				</radio-group>				
			</div>
		</div>
		<div class="row-center cash-line" v-if="checkInput('account')">
			<span class="label-text">收款账号</span>
			<div class="row-center cash-line-right">
				<input placeholder="请输入收款账号" v-model="cashData.account" class="cash-input" />
			</div>
		</div>
		<div class="row-center cash-line" v-if="checkInput('realname')">
			<span class="label-text">收款人</span>
			<div class="row-center cash-line-right">
				<input placeholder="请输入收款人姓名" v-model="cashData.realname" class="cash-input" />
			</div>
		</div>
		<div class="row-center cash-line" v-if="checkInput('bank')">
			<span class="label-text">银行名称</span>
			<div class="row-center cash-line-right">
				<input placeholder="请输入收款银行名称" v-model="cashData.bank" class="cash-input" />
			</div>
		</div>
		<div class="row-center cash-line" v-if="checkInput('qrcode')">
			<span class="label-text">收款码</span>
			<div class="row-center cash-line-right">
				<div class="center cash-img-block" @click="upQrCode()">
					<img :src="cashData.qrcode" v-if="cashData.qrcode" style="height: 48px;width: 48px;" />
					<uni-icons type="plusempty" size="30" v-else></uni-icons>
				</div>
			</div>
		</div>
		<div class="row-center cash-line">
			<button type="primary" size="mini" @click="sub()">提交</button>
			<button size="mini" @click="uni.navigateTo({
				url:'/pages/cash/list'
			})">记录</button>
		</div>
		<uni-card title="说明">
			<mp-html :content="desc"></mp-html>
		</uni-card>
	</view>
</template>

<script setup>
	import {ref} from 'vue'
	import { onShow } from '@dcloudio/uni-app'
	import store from '@/stores/index.js'
	import { request,upload } from '/config/api.js'
	const config=store().config
	const cashData=ref({type:1})
	const desc=ref('')
	function checkInput(type){
		switch (type) {
			case 'account':
				if(cashData.value.type==3 && config.cash_bank_type==1){
					return true
				}
				if(cashData.value.type==2 && config.cash_alipay_type==2){
					return true
				}
				return false
			case 'realname':
				if(cashData.value.type==3 && config.cash_bank_type==1){
					return true
				}
				if(cashData.value.type==1 && config.cash_wxpay_type==1){
					return true
				}
				if(cashData.value.type==2){
					return true
				}
				return false
			case 'bank':
				if(cashData.value.type==3 && config.cash_bank_type==1){
					return true
				}
				return false
			case 'qrcode':
				if(cashData.value.type==1 && config.cash_wxpay_type==1){
					return true
				}
				if(cashData.value.type==2 && config.cash_alipay_type==1){
					return true
				}
				return false
			default:
				return false
		}
	}
	function radioChange(e){
		cashData.value.type=Number(e.detail.value)
	}
	function upQrCode(){
		uni.chooseImage({
			count:1,
			success(chooseImageRes){
				let tempFilePaths = chooseImageRes.tempFilePaths;
				upload('UploadFile',tempFilePaths[0]).then(res=>{
					cashData.value.qrcode=res.data					
				})
			}
		})
	}
	onShow(()=>{
		request('GetUserInfo').then(res=>{
			store().setUser(res.data)
		})
		request('ArticleInfo',{id:3}).then((res)=>{
			desc.value=res.data.content
		})
	})
	function sub(){
		if(!cashData.value.money){
			uni.showToast({
				title:'请输入提现金额',
				icon:'none'
			})
			return
		}
		if(cashData.value.money>store().user.balance){
			uni.showToast({
				title:'余额不足',
				icon:'none'
			})
			return
		}
		cashData.value.money=parseFloat(cashData.value.money)
		request('CashAdd',cashData.value,'post').then(res=>{
			uni.showToast({
				title:res.msg
			})
			// #ifdef MP-WEIXIN
			wx.requestSubscribeMessage({
				tmplIds: [config.wx_mini_cash_templateId],
				success(res) {
					console.log(res)
				}
			})
			// #endif
		})
	}
</script>

<style scoped>
	.cash-line{
		margin-bottom: 1px;
		background-color: #ffffff;
		height: 50px;
		padding: 0 10px;
	}
	.label-text{
		width: 80px;
	}
	.cash-line-right{
		margin-left: 10px;
	}
	.cash-input{
		border: 1px solid #0000009e;
		border-radius: 5px;
		height: 29px;
	}
	.radio-label{
		margin-right: 10px;
	}
	radio{
		transform:scale(0.8)
	}
	.cash-img-block{
		border: 1px dashed #838383;
		width: 50px;
		height: 48px;
	}
</style>
