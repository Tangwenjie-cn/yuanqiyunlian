<template>
	<view class="container">
		<div :style="`background: linear-gradient(${themeBackgroundColor},#ffffff);`" class="top-block">
			<span style="color: #ffffff;font-size: 26px;font-weight: 600;padding-left: 10px;padding-top: 20px;">每日签到活动</span>
			<span style="color: #ffffff;padding-top: 23px;padding-right: 5px;" @click="uni.navigateTo({
				url: '/pages/points/log'
			})">明细</span>
		</div>
		<div class="sign-block">
			<span>已连续签到{{signData.continuousSignDays}}天</span>			
			<div style="display: grid;grid-template-columns:repeat(7,1fr)">
				<div class="flex-column-center" v-for="item in signData.list" style="margin-top: 5px;">
					<div class="flex-column-center day-block" :style="item.sign_day?`border: 1px solid ${themeBackgroundColor};`:''">
						<span>+{{item.points}}</span>
						<img src='/static/images/sign-suc.png' style="width: 22px;height: 22px;" v-if="item.is_sign"  />
						<img v-else :src="item.type>0?'/static/images/sign-icon-3.png':'/static/images/sign-icon-1.png'" 
						style="width: 22px;height: 22px;" />
					</div>
					<span style="color: #6f6f6f;">{{item.is_sign?'已领':item.day}}</span>
				</div>				
			</div>
			<div class="column-center">
				<div class="center sign-btn" :style="{'background-color':signData.checkSign?'#00000057':themeBackgroundColor}" 
				@click="sign">{{signData.checkSign?'今日已签到，明日再来吧':'立即签到'}}</div>
			</div>
			<div class="center" v-if="signData.nextContinuousDays>0" style="margin: 10px 0;">
				<img src="/static/images/sign-icon-4.png" style="width: 14px;height: 14px;margin-right: 5px;" />
				<span style="color: #6f6f6f;font-size: 14px;">再连续签到{{signData.nextContinuousDays}}天，可额外获得惊喜礼包</span>
			</div>
			
		</div>
		<div class="lock"></div>
		<div class="sign-block1">
			<div class="column-center" style="margin-top: 10px;">
				<span style="color: #6f6f6f;font-size: 24px;">已累积签到</span>
			</div>
			<div class="center" style="margin-top: 10px;">
				<view class="cumulative-day center" :style="{backgroundColor:themeBackgroundColor}">{{ signCount[0] || 0 }}</view>
				<view class="cumulative-day center" :style="{backgroundColor:themeBackgroundColor}">{{ signCount[1] || 0 }}</view>
				<view class="cumulative-day center" :style="{backgroundColor:themeBackgroundColor}">{{ signCount[2] || 0 }}</view>
				<view class="cumulative-day center" :style="{backgroundColor:themeBackgroundColor}">{{ signCount[3] || 0 }}</view>
				<span>天</span>
			</div>		
			<div class="center" v-if="signData.nextCumulativeDays>0" style="margin-top: 10px;">
				<img src="/static/images/sign-icon-4.png" style="width: 14px;height: 14px;margin-right: 5px;" />
				<span style="color: #6f6f6f;font-size: 14px;">再累积签到{{signData.nextCumulativeDays}}天，可额外获得惊喜礼包</span>
			</div>
		</div>
	</view>
</template>

<script setup>
	import {ref,watch} from 'vue'
	import { onShow } from '@dcloudio/uni-app'
	import store from '@/stores/index.js'
	import { request } from '/config/api.js'
	const signData = ref({})
	const signCount=ref([])
	const themeFrontColor=ref('#fff')
	const themeBackgroundColor=ref('#f56b59')
	const theme=ref({})
	onShow(()=>{
		getSignConfig()
		theme.value=uni.getStorageSync('theme').vip[0]
		themeFrontColor.value=theme.value.titleColor
		themeBackgroundColor.value=theme.value.backgroundColor		
		setStatus()
	})
	function getSignConfig(){
		request('SignConfig').then(res=>{
			signData.value = res.data
			signCount.value=PrefixInteger(res.data.cumulativeSignDays,4) 
		})
	}
	function PrefixInteger(num, length) {
		return (Array(length).join('0') + num).slice(-length).split('');
	}
	function sign(){
		if(!signData.value.checkSign){
			request('SignAdd',{},'post').then(res=>{
				uni.showToast({
					title:res.msg,
					icon:'none'
				})
				getSignConfig()
			})
		}
	}
	function setStatus(){
		uni.setNavigationBarColor({
			frontColor: themeFrontColor.value,
			backgroundColor: themeBackgroundColor.value,
		})
	}
</script>

<style scoped>
	.top-block{
		height: 160px;
		display: flex;
		justify-content: space-between;
	}   
	.sign-block{
		border-radius: 10px;
		margin: -80px 10px 0 10px;
		padding: 10px;
		background-color: #fff;
	}
	.sign-btn{
		width: 210px;
		height: 45px;
		border-radius: 29px;
		color: #ffffff;
		margin-top: 15px;
	}
	.day-block{
		background-color: #ebebeb;
		border-radius: 5px;
		padding: 10px;		
	}
	.lock {
		background-image: url('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAi4AAABECAYAAACmur7KAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAA3ZpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTQyIDc5LjE2MDkyNCwgMjAxNy8wNy8xMy0wMTowNjozOSAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wTU09Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9tbS8iIHhtbG5zOnN0UmVmPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvc1R5cGUvUmVzb3VyY2VSZWYjIiB4bWxuczp4bXA9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8iIHhtcE1NOk9yaWdpbmFsRG9jdW1lbnRJRD0ieG1wLmRpZDowYWFmYjU3Mi03MGJhLTRiNDctOTI2Yi0zOThlZDkzZDkxMDkiIHhtcE1NOkRvY3VtZW50SUQ9InhtcC5kaWQ6MkY0OEQxQjdEMDFDMTFFODgwMTlGMzZDRjQ3RTkwMTgiIHhtcE1NOkluc3RhbmNlSUQ9InhtcC5paWQ6MkY0OEQxQjZEMDFDMTFFODgwMTlGMzZDRjQ3RTkwMTgiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENDIDIwMTggKFdpbmRvd3MpIj4gPHhtcE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6ZmRjNTM0MmUtNmFkOC1iMDRhLThjZTEtMjk2YWYzM2FkMmUxIiBzdFJlZjpkb2N1bWVudElEPSJ4bXAuZGlkOjBhYWZiNTcyLTcwYmEtNGI0Ny05MjZiLTM5OGVkOTNkOTEwOSIv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');
		background-repeat: no-repeat;
		background-size: 100% 100%;
		width: 558rpx;
		height: 68rpx;
		position: absolute;
		left: 50%;
		transform: translateX(-50%);
		z-index: 9;
		margin-top: -14px;
	}
	.sign-block1{
		border-radius: 10px;
		margin: 10px 10px 0 10px;
		padding: 10px;
		background-color: #fff;
	}
	.cumulative-day{
		height: 49px;
		border-radius: 5px;	
		padding: 8px 12px;
		color: #fff;
		font-size: 28px;
		margin-right: 5px;
	}
</style>
