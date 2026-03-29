<template>
	<div class="center">
		<div class="controls" @click="play">
			<image :src="props.poster" style="width: 60px;height: 60px;"></image>
			<span class="iconfont icon-zanting icon" v-if="isPlay"></span>
			<span class="iconfont icon-bofang icon" v-else ></span>
		</div>
		<div style="margin: 0 3px;">{{formatTime(currentTime)}}</div>
		<div class="progress-container" @click="progressClick">
		    <div
		      class="progress-bar"
		      :style="{ width: `${progress}%` }"
		    />
		</div>
		<div style="margin-left: 3px;">{{formatTime(duration)}}</div>
	</div>
</template>
<script setup>
	import { ref } from 'vue'
	import { onUnload } from '@dcloudio/uni-app'
	import {formatTime} from '@/config/function.js'
	const innerAudioContext = uni.createInnerAudioContext();	
	const props = defineProps({
		src: {
			type: String,
			default: ''
		},
		poster: {
			type: String,
			default: ''
		}
	})
	const isPlay = ref(false)
	const duration=ref(0)
	const progress = ref(0)
	const currentTime=ref(0)
	let timer = null
	innerAudioContext.src =props.src
	innerAudioContext.onCanplay(()=>{
		duration.value=Math.round(innerAudioContext.duration)
	})
	onUnload(()=>{
		innerAudioContext.destroy()
	})
	function progressClick(e) {
		const percent = e.offsetX / e.target.offsetWidth
		innerAudioContext.seek(duration.value * percent)
		progress.value = percent * 100
	}
	function play() {
		console.log(111)
		if (isPlay.value) {
			innerAudioContext.pause()
		} else {
			innerAudioContext.play()
		}		
	}
	innerAudioContext.onPlay(()=>{
		isPlay.value=true
		timer = setInterval(() => {
			progress.value = Math.round(innerAudioContext.currentTime / duration.value * 100)
			currentTime.value=parseInt(innerAudioContext.currentTime)
		}, 1000)
	})
	innerAudioContext.onPause(() => {
		isPlay.value=false
		clearInterval(timer)
	})	
	innerAudioContext.onEnded(() => {
		isPlay.value=false
		progress.value = 0
		currentTime.value=0
		clearInterval(timer)
	})
</script>
<style scoped>
	@import url("../../../../static/font/iconfont.css");
	.progress-container {
	  height: 6px;
	  width: 100%;
	  background: #f1f5f9;
	  border-radius: 3px;
	  overflow: hidden;
	}
	.progress-bar {
	  height: 100%;
	  background: #2563eb;
	  transition: width 0.1s linear;
	}
	.controls{
		width: 60px;
		height: 60px;
		background-color: #3d3d3d;
		border-radius: 5px;
		position: relative;
	}
	.icon{
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%,-50%);
		font-size: 40px;
		color: #fff;
	}
</style>
