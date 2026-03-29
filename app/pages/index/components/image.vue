<template>
	<div class="content">
		<navigator :url="model.link">
			<image :src="model.imgUrl" :style="model.height?'height:'+model.height:''"></image>
		</navigator>
		
	</div>
	
</template>

<script setup>
	const model=defineModel()
</script>

<style scoped>
	.content{
		padding-left: 5px;
		padding-right: 5px;
		margin-bottom: 16px;
	}
	image{
		width: 100%;
		height: 160px;
		border-radius: 10rpx;
		object-fit: fill;
	}
</style>