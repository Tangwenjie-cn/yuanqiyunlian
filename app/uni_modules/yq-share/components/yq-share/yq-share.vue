<template>
	<div class="column">
		<img :src="props.src" style="width: 300px;height: 400px;">
		<div class="center btn-line">
			<!-- #ifdef H5 -->
			<span style="color: #e52b2b;">请长按保存后分享</span>
			<!-- #endif -->
			<!-- #ifdef MP -->
			<button size="mini" open-type="share">分享</button>
			<!-- #endif -->
			<!-- #ifndef H5 -->			
			<button size="mini" @click="save">保存</button>
			<!-- #endif -->
			
		</div>
	</div>
</template>
<script setup>
	import {ref} from 'vue'
	import { onShareAppMessage  } from '@dcloudio/uni-app'
	const props = defineProps(['src','path','title'])
	onShareAppMessage((res)=>{
		return{
			title: props.title,
			path: props.path,
			imageUrl: props.src
		}
	})
	function save(){
		uni.downloadFile({
			url:props.src,
			success: (res) => {
				if (res.statusCode === 200) {
					uni.saveImageToPhotosAlbum({
						filePath: res.tempFilePath,
						success: function () {
							uni.showToast({
								title: '保存成功',
								icon: 'success',
								duration: 2000
							})
						}
					})
				}
			}
		})
	}
</script>
<style scoped>
	.btn-line{
		height: 45px;
	}
</style>
