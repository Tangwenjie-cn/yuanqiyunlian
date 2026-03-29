<template>
	<div class="title-block">
		<span style="font-weight: 500;">目录</span>
	</div>
	<scroll-view scroll-y="true" class="column-block">
		<view class="center column-item" v-for="item in list">
			<navigator :url="'/pages/goods/info?id='+item.id">
				<span class="column-title" :style="check(item.id)">{{item.title}}</span>
			</navigator>
			
		</view>
	</scroll-view>
</template>

<script setup>
	const list=defineModel()
	const props=defineProps(['gid'])
	function check(id){
		if(props.gid==id){
			return 'color: #a8a8a8;'
		}
	}
</script>

<style scoped>
	.title-block{
		border-bottom: 1px solid #00000042;;
	}
	.column-block{
		width: 100%;
		height: 100px;
	}
	.column-item{
		height: 35px;
	}
	.column-title{
		font-size: 15px;
		color: #366fc7;
	}
</style>