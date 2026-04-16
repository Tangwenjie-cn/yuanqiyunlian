<script setup>
	import { onLaunch, onShow, onHide } from '@dcloudio/uni-app'
	import { request } from '@/config/api.js'
	import store from '@/stores/index.js'
	onLaunch(()=>{
		console.log('App Launch')
		request('Theme').then(res=>{
			uni.setStorageSync('theme',res.data)
		})
		request('Config').then(res=>{
			store().config=res.data
		})
		let user=uni.getStorageSync('user')
		if(user){
			store().user=user
		}
	})
	onShow(()=> {
		console.log('App Show')
	})
	onHide(()=> {
		console.log('App Hide')
	})
</script>

<style>
	/*每个页面公共css */
	image{will-change: transform}
	.container{
		min-height: 100vh;
		padding-bottom: 48px;
		background-color: #F8F8F8;
	}
	.center{
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.row-center{
		display: flex;
		align-items: center;
	}
	.column-center{
		display: flex;
		justify-content: center;
	}
	.column{
		display: flex;
		flex-direction: column;
	}
	.flex-column-center{
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}
	.column-row-center{
		display: flex;
		flex-direction: column;
		align-items: center;
	}
	.column-column-center{
		display: flex;
		flex-direction: column;
		justify-content: center;
	}

</style>
