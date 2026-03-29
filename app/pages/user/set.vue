<template>
	<view class="container">
		<div class="column block">
			<div class="block-item row-center">
				<span>头像</span>
				<!-- #ifdef MP-WEIXIN -->
				<div class="row-center">
					<button style="width: 30px;height: 30px;border-radius: 50%;padding: 0;"  
					open-type="chooseAvatar" @chooseavatar="onChooseAvatar">
					  <image class="avatar" :src="user.avatar?user.avatar:config.site_logo"></image>
					</button>
					<uni-icons type="right" size="25" color="#0000007d"></uni-icons>
				</div>
				<!-- #endif -->
				<!-- #ifndef MP-WEIXIN -->
				<div class="row-center" @click="editAvatar">
					<image :src="user.avatar?user.avatar:config.site_logo" class="avatar"></image>
					<uni-icons type="right" size="25" color="#0000007d"></uni-icons>
				</div>
				<!-- #endif -->				
			</div>
			<div class="block-item row-center">
				<span>昵称</span>
				<div class="row-center" @click="editNick">
					<span>{{user.nickname}}</span>
					<uni-icons type="right" size="25" color="#0000007d"></uni-icons>
				</div>
			</div>
			<!-- <div class="block-item row-center">
				<span>手机</span>
				<div class="row-center">
					<span>{{user.phone}}</span>
					<uni-icons type="right" size="25" color="#0000007d"></uni-icons>
				</div>
			</div> -->
			<div class="block-item row-center">
				<span>ID</span>
				<span>{{user.id}}</span>
			</div>
		</div>
		<div class="column block" style="margin-top: 20px;">
			<div class="block-item row-center" @click="uni.navigateTo({
				url:'/pages/article/index?id=1'
			})">
				<span>联系我们</span>
				<uni-icons type="right" size="25" color="#0000007d"></uni-icons>
			</div>
			<div class="block-item row-center" @click="uni.navigateTo({
				url:'/pages/article/list?sid=1'
			})">
				<span>帮助中心</span>
				<uni-icons type="right" size="25" color="#0000007d"></uni-icons>
			</div>
		</div>
		<button type="warn" style="margin-top: 20px;" @click="logout">退出登录</button>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	import { onShow,onLoad } from '@dcloudio/uni-app'
	import store from '@/stores/index.js'
	import { request,upload } from '/config/api.js'
	const user=store().user
	const config=store().config
	function logout(){
		request('Logout').then(res=>{
			uni.removeStorageSync('user')
			store().user=null
			uni.reLaunch({
				url:'/pages/index/index'
			})
		})
	}
	function onChooseAvatar(e){
		upload('UploadFile',e.detail.avatarUrl).then(res=>{
			request('UpdateUserInfo',{avatar:res.data},'post').then(res1=>{
				uni.showToast({
					title:'修改成功',
					icon:'success'
				})
				user.avatar=res.data
				store().setUser(user)
			})
		})
	}
	function editAvatar(){
		uni.chooseImage({
			success: (chooseImageRes) => {
				let tempFilePaths = chooseImageRes.tempFilePaths;
				upload('UploadFile',tempFilePaths[0]).then(res=>{
					request('UpdateUserInfo',{avatar:res.data},'post').then(res1=>{
						uni.showToast({
							title:'修改成功',
							icon:'success'
						})
						user.avatar=res.data
						store().setUser(user)
					})
				})
			}
		});
	}
	function editNick(){
		uni.showModal({
			title:'修改名称',
			editable: true,
			placeholderText: '请输入新名称',
			success: (res) => {
				if (res.confirm) {
					if(!res.content){
						uni.showToast({
							title: '请输入新名称',
							icon:"none"
						});
						return
					}
					request('UpdateUserInfo',{nickname:res.content},'post').then(res1=>{
						uni.showToast({
							title:'修改成功',
							icon:'success'
						})
						user.nickname=res.content
						store().setUser(user)
					})
				}
			}
		})
	}
</script>

<style scoped>
	.block{
		background-color: #ffffff;
	}
	.block-item{
		padding: 10px;
		height: 30px;
		border-bottom: 1px solid #c1c1c1;
		justify-content: space-between;
	}
	.block-item:last-child{
		border-bottom: none;
	}
	.avatar{
		width: 30px;
		height: 30px;
		border-radius: 50%;
	}
</style>
