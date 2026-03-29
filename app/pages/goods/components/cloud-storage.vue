<template>
	<div>
		<span class="title">地址：</span>
		<span class="text">{{model.link}}</span>
		<span class="copy-text iconfont icon-fuzhi" @click="cpoy(model.link)"></span>
	</div>
	<div>
		<span class="title">提取码：</span>
		<span class="text">{{model.code}}</span>
		<span class="copy-text iconfont icon-fuzhi" @click="cpoy(model.code)"></span>
	</div>
</template>	
<script setup>
	const model=defineModel()
	function cpoy(text){
		uni.setClipboardData({
			data:text,
			success:()=>{
				uni.showToast({
					title:'复制成功',
					icon:'none'
				})
			}
		})
	}
</script>

<style scoped>
	@import url("../../../static/font/iconfont.css");
	.title{
		font-size: 14px;
		color: #999;
	}
	.text{
		font-size: 14px;
		color: #333;
	}
	.copy-text{
		font-size: 14px;
		color: #007aff;
		margin-left: 10px;
	}
</style>