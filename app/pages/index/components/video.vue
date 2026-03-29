<template>
	<view class="content">
		<video :src="model.link" controls :style="model.height?'height:'+model.height:''"></video>
	</view>
</template>

<script setup>
	const model=defineModel()
</script>

<style>
	.content{
		padding-left: 5px;
		padding-right: 5px;
		margin-bottom: 16px;
	}
	video{
		width: 100%;
		height: 200px;
		border-radius: 10rpx;
	}
</style>