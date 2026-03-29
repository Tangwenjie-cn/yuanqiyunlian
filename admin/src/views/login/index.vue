
<template>
	<div class="body">
		<div class="login-form">
			<span class="login-title">登录控制台</span>
			<el-form :model="form" label-width="80px" class="form" @submit.native.prevent="submit">
				<el-input prefix-icon="user" v-model="form.username" />
				<el-input prefix-icon="lock" type="password" v-model="form.password" />
				<el-button style="width: 230px;" type="primary" @click="submit" native-type="submit" round>登录</el-button>				
			</el-form>
		</div>
	</div>
</template>

<script setup>
	import {ref,inject} from "vue";
	import { ElMessageBox } from 'element-plus'
	const router = inject('router');
	const store = inject('store');
	const myApi=inject('myApi');
	const ElMessage=inject('ElMessage');

	const form=ref({username:'',password:''});
	function submit(){
		if(!form.value.username){
			return ElMessage.error('请输入用户名')
		}
		if(!form.value.password){
			return ElMessage.error('请输入密码')
		}
		myApi('login',form.value,'post').then(res=>{
			if(form.value.password=='123456'){
				ElMessageBox.alert('您的密码是默认密码，请及时修改密码', '提示', {
					confirmButtonText: '确定'
				})
			}
			let user={
				menu:res.data.menu,
				userInfo:{
					token:res.data.token,
					id:res.data.id,
					group_id:res.data.group_id,
					username:res.data.username,
				},
				config:{
					logo:res.data.logo,
				}
			}
			localStorage.setItem('user',JSON.stringify(user))
			store.upUser();
			let menuActive=localStorage.getItem('menuActive')
			if (menuActive) {
				let tableTabs=JSON.parse(localStorage.getItem('tableTabs'))
				let index = tableTabs.findIndex(val => val.name==menuActive);
				if(index!=-1){
					router.push(tableTabs[index].path)
				}else{
					router.push('/home')
				}
			}else{
				router.push('/home')
			}
			
		})
	}
</script>

<style scoped>
	.body{
		display: flex;
		justify-content: center;
		align-items: center;
		position: absolute;
		top: 0;
		bottom: 0;
		left: 0;
		right: 0;
		background-image: url(../../assets/images/login-bg.jpeg);	
		background-size: cover;
		z-index: 1;	
	}
	.body::after{
		content: "";
		width: 100%;
		height: 100%;
		position: absolute;
		left: 0;
		top: 0;
		/* 从父元素继承 background 属性的设置 */
		background: inherit;
		filter: blur(20px);
		z-index: -1;
	}
	.form{
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: space-around;
		height: 160px;
	}
	.login-form{
		width: 400px;
		height: 400px;
		background-color: #FFFFFF;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		border-radius: 10px;
	}
	.login-title{
		font-size: 24px;
		margin-bottom: 20px;
		color: #909399;
	}
</style>
