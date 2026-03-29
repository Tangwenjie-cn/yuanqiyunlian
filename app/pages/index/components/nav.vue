<template>
	<div class="content">
		<div class="nav-block" 
		:style="{'background-color':model.backgroundColor,'grid-template-columns':`repeat(${model.arrange}, 1fr)`}">
			<div class="nav-item column-row-center" v-for="item in model.list">
				<navigator :url="item.link">
					<image class="nav-image" :src="item.imgUrl" mode=""></image>
				</navigator>
				<span :style="{'color':model.color,'font-size':model.fontSize}">{{item.name}}</span>
			</div>

		</div>
	</div>
</template>

<script setup>
	import {onLoad} from '@dcloudio/uni-app'
	const model=defineModel()
</script>

<style scoped>
	.content{
		padding-left: 5px;
		padding-right: 5px;
		margin-bottom: 16px;
	}
	.nav-block{
		border-radius: 10rpx;
		margin-top: 20rpx;
		margin-bottom: 20rpx;
		display: grid;
		justify-items: center;
		padding: 10px;
	}
	.nav-item{
		margin-top: 6px;
	}
	.nav-image{
		width: 55px;
		height: 55px;
		border-radius: 10px;
		object-fit: cover;
	} 
</style>