<template>

	<div class="search column-center">
		<div class="search-input row-center" :style="{'background-color':`${model.backgroundColor}`}">
			<uni-icons type="search" size="17" color="gray"></uni-icons>
			<input placeholder="请输入搜索内容" confirm-type="search" @confirm="search" />
		</div>
	</div>

</template>

<script setup>
	const model=defineModel()
	function search(e){
		uni.navigateTo({
			url:`/pages/goods/index?keyword=${e.detail.value}`
		})
	}
</script>

<style scoped>
	.search{
		width: 100%;
		margin-bottom: 16px;
	}
	.search-input{
		width: 70%;
		height: 36px;
		border-radius: 16px;
		padding-left: 10px;
	}
</style>