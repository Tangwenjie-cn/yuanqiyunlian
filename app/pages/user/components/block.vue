<template>
	<div class="content">
		<div class="nav-block" 
		v-if="model.arrange>1"
		:style="{'background-color':`${model.backgroundColor}`,'grid-template-columns':`repeat(${model.arrange}, 1fr)`}">
			<div class="nav-item column-row-center" v-for="item in model.list">
				<navigator :url="item.link">
					<image class="nav-image" :src="item.imgUrl" mode=""></image>
				</navigator>
				<span :style="{'color':model.color,'font-size':model.fontSize}">{{item.name}}</span>
			</div>
		</div>
		<div class="nav-list column" v-else :style="{'background-color':`${model.backgroundColor}`}">
			<navigator v-for="item in model.list" :url="item.link" class="navigator-list">
			<div class="nav-list-item row-center" >
				<div class="row-center">
					<image class="nav-list-image" :src="item.imgUrl" mode=""></image>
					<span class="nav-list-title" :style="{'color':model.color,'font-size':model.fontSize}">{{item.name}}</span>
				</div>
				<uni-icons type="right" size="22"></uni-icons>
			</div>				
			</navigator>
		</div>
	</div>
</template>

<script setup>
	const model=defineModel()
</script>

<style scoped>
	.content{
		margin-top: 20px;
		padding-left: 16px;
		padding-right: 16px;
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
		margin-top: 8px;
	}
	.nav-image{
		width: 36px;
		height: 36px;
		border-radius: 10px;
		object-fit: cover;
	}
	.nav-list{
		border-radius: 10rpx;
	}
	.nav-list-item{
		justify-content: space-between;
		height: 46px;		
		margin: 0 10px;
	}
	.navigator-list{
		border-bottom: 1px solid #00000021;
	}
	.navigator-list:last-child{
		border-bottom: none;
	}
	.nav-list-item-left{
		display: flex;
		align-items: center;
	}
	.nav-list-title{
		margin-left: 10px;
	}
	.nav-list-image{
		height: 22px;
		width: 22px;
		object-fit: cover;
	}
</style>