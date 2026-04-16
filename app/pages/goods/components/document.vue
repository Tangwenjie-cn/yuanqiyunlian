<template>
	<view class="document">
		<button type="primary" @click="downloadFile">{{btnText}}</button>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	const btnText=ref('下载文档')
	const links=defineModel()
	let filePath=''
	function downloadFile(){
		// #ifdef H5
		return window.open(links.value)
		// #endif
		if(filePath){
			uni.openDocument({
				filePath: filePath,
				showMenu:true,
				success: function (res) {
					btnText.value = '打开文档'
				},
				fail: function (res) {
					btnText.value = '打开失败'
				}
			})
			return
		}
		var downloadTask = uni.downloadFile({
			url:links.value,
			success: (res) => {
				if (res.statusCode === 200) {
					filePath=res.tempFilePath
					uni.openDocument({
						filePath: res.tempFilePath,
						showMenu:true,
						success: function (res) {
							btnText.value = '打开文档'
						},
						fail: function (res) {
							btnText.value = '打开失败'
						}
					})
				}
			}
		})
		downloadTask.onProgressUpdate((res) => {
			btnText.value = '下载中' + res.progress + '%'
		})
	}
</script>

<style>
</style>