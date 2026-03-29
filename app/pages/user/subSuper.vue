<template>
	<view>
		<button type="primary" @click="subSuper()">提交申请</button>
		<uni-card v-if="Object.keys(info).length>0">
			<span style="color: #0055ff;" v-if="info.status==0">已提交申请，请等待管理审核</span>
			<span style="color: #ff5500;" v-else>拒绝原因：{{info.remark}}</span>
		</uni-card>
		<uni-card>
			<mp-html :content="desc"></mp-html>
		</uni-card>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	import { request } from '/config/api.js'
	import { onShow } from '@dcloudio/uni-app'
	const info=ref({})
	const desc=ref('')
	onShow(()=>{
		request('GetSubSuper').then(res=>{
			if(res.data){
				info.value=res.data
			}			
		})
		request('ArticleInfo',{id:5}).then((res)=>{
			desc.value=res.data.content
		})
	})
	function subSuper(){
		uni.showModal({
			title:'提示',
			content:'确认满足条件提交吗',
			confirmText:'提交',
			success:function(res){
				if(res.confirm){
					request('SubSuper',{},'post').then(res=>{						
						uni.showToast({
							title:'提交成功',
							icon:'success'
						})					
					})
				}
			}
		})
	}
</script>

<style>
	       
</style>
