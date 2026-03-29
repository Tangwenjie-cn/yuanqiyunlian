
<template>
	<el-container>
		<el-header class="header" style="padding: 0;">
			<div class="header-left">
				<img :src="store.config.logo" class="logo">
				<span style="margin-left: 5px;">控制台</span>
			</div>						
			<div class="header-right">
				<div>
					<el-icon @click="collapse" class="icon"><Operation /></el-icon>
				</div>
				<div class="header-right-right">
					
					<el-dropdown style="margin-left: 20px;">
						<span>
							{{ store.userInfo.username }}
							<el-icon class="icon"><arrow-down /></el-icon>
						</span>
						<template #dropdown>
							<el-dropdown-menu>
								<el-dropdown-item @click="refresh">刷新</el-dropdown-item>
								<el-dropdown-item @click="editPwdDialog=true">修改密码</el-dropdown-item>
								<el-dropdown-item @click="logOut">退出登录</el-dropdown-item>							
							</el-dropdown-menu>
						</template>
					</el-dropdown>
				</div>
				
			</div>
		</el-header>
		<el-container class="body">
			<el-aside class="aside" :style="asideSyle">
				<el-menu :default-active="menuActive" :collapse="isCollapse">
					<menuTree v-for="item in menu" :item="item"></menuTree>	
				</el-menu>
			</el-aside>
			<el-main class="main">
				<el-tabs
					v-model="tableTabsValue"
					type="card"
					closable
					@tab-remove="tabRemove"
					@tab-change="tabChange"
				>
					<el-tab-pane
					v-for="item in tableTabs"
					:label="item.title"
					:name="item.name"
					>
					</el-tab-pane>
				</el-tabs>
				
				<router-view v-slot="{ Component }">
					<keep-alive :max="10">
					<component :is="Component" />
					</keep-alive>
				</router-view>				
			</el-main>
		</el-container>
	</el-container>
	<!--弹出层---->
	<el-dialog v-model="editPwdDialog" title="修改密码" width="460px">
		<el-form :model="editPwdForm" ref="refEditPwdForm" label-width="80px">
			<el-form-item label="旧密码" prop="old_password" required>
				<el-input v-model="editPwdForm.old_password" type="password" autocomplete="new-password"></el-input>
			</el-form-item>
			<el-form-item label="新密码" prop="password" required>
				<el-input v-model="editPwdForm.password" type="password" autocomplete="new-password"></el-input>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="editPwd">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
</template>

<script setup>
	import {ref,onBeforeMount,inject,watch,onMounted} from "vue";
	import menuTree from "./menuTree.vue"
	import { ElMessageBox } from 'element-plus'
	const router = inject('router');
	const store = inject('store');
	const myApi=inject('myApi');
	//const ElMessage=inject('ElMessage');
	const menu=ref([]);	
	const tableTabsValue = ref(0)
	const tableTabs = ref([])
	const menuActive=ref(0)
	const pathname=ref('')
	const isCollapse=ref(false)
	const asideSyle=ref({width:'220px'})
	const editPwdDialog=ref(false)
	const editPwdForm=ref({
		old_password:'',
		password:'',
	})
	const refEditPwdForm=ref()
	onBeforeMount(()=>{
		store.upUser()	
		
		menu.value=store.menu
	})
	onMounted(()=>{
		
		let tableTabsStorage=localStorage.getItem('tableTabs')
		if(tableTabsStorage){
			tableTabs.value=JSON.parse(tableTabsStorage)
			menuActive.value=Number(localStorage.getItem('menuActive'))
			pathname.value=localStorage.getItem('pathname')
			tableTabsValue.value=Number(localStorage.getItem('tableTabsValue'))
			router.push(pathname.value)
		}else{
			let item=menu.value[0]
		    tableTabs.value.push({
				title:item.name,
				name:item.id,
				path:item.path,
			});
			tableTabsValue.value=item.id
			menuActive.value=item.id
			pathname.value=item.path
			router.push(item.path)
			localStorage.setItem('tableTabs',JSON.stringify(tableTabs.value))
			localStorage.setItem('menuActive',menuActive.value)
			localStorage.setItem('tableTabsValue',tableTabsValue.value)
			localStorage.setItem('pathname',pathname.value)
		}
	})
	watch(
		()=>store.treeMenuClick,
		()=>{
			if(store.treeMenuClick){
				menuClick(store.treeMenuClick)
				store.treeMenuClick=null
			}
			
		}
	)
	function menuClick(item){ //点击菜单
		let index = tableTabs.value.findIndex(val => val.name==item.id);
		if(index==-1){
			tableTabs.value.push({
				title:item.name,
				name:item.id,
				path:item.path,
			});				
		}
		tableTabsValue.value=item.id
		menuActive.value=item.id
		pathname.value=item.path		
		router.push(item.path)
		localStorage.setItem('tableTabs',JSON.stringify(tableTabs.value))
		localStorage.setItem('tableTabsValue',tableTabsValue.value)
		localStorage.setItem('menuActive',menuActive.value)
		localStorage.setItem('pathname',pathname.value)
	}
	function tabRemove(el){ //tab移除
		if(tableTabs.value.length==1){
			return ElMessageBox({
				title: '提示',
				message: '至少保留一个标签页',
				type: 'warning',
			})
		}
		let index = tableTabs.value.findIndex(val => val.name==el);
		
		tableTabs.value.splice(index,1)
		let index1=tableTabs.value.length-1
		el=tableTabs.value[index1].name
		menuActive.value=el
		tableTabsValue.value=el
		pathname.value=tableTabs.value[index1].path
		router.push(tableTabs.value[index1].path)
		localStorage.setItem('tableTabs',JSON.stringify(tableTabs.value))
		localStorage.setItem('menuActive',menuActive.value)
		localStorage.setItem('tableTabsValue',tableTabsValue.value)
		localStorage.setItem('pathname',pathname.value)
	}
	function tabChange(el){
		menuActive.value=el
		tableTabsValue.value=el
		let index = tableTabs.value.findIndex(val => val.name==el);
		router.push(tableTabs.value[index].path)
		pathname.value=tableTabs.value[index].path
		localStorage.setItem('pathname',pathname.value)
		localStorage.setItem('menuActive',menuActive.value)
		localStorage.setItem('tableTabsValue',tableTabsValue.value)
	}
	function collapse(){
		if(isCollapse.value){
			isCollapse.value=false
			asideSyle.value.width='220px'
		}else{
			isCollapse.value=true
			asideSyle.value.width='65px'
		}
	}
	function logOut(){ //退出登录
		myApi('logout').then(()=>{
			localStorage.removeItem('user')
			localStorage.removeItem('tableTabs')
			localStorage.removeItem('menuActive')
			localStorage.removeItem('tableTabsValue')
			localStorage.removeItem('pathname')
			store.upUser()
			router.push('/login')
		})		
		
	}
	function refresh(){ //刷新菜单
		myApi('refresh').then((res)=>{
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
					v_name:res.data.v_name,
				}
			}
			localStorage.setItem('user',JSON.stringify(user))
			store.upUser()
			menu.value=res.data.menu
		})
	}
	function editPwd(){
		refEditPwdForm.value.validate((valid) => {
			if (valid) {
				myApi('editPwd',editPwdForm.value,'post').then(()=>{
					ElMessageBox.alert('修改成功，请重新登录', '提示', {
						confirmButtonText: '确定',
						type: 'success'
					}).then(() => {
						logOut()
					})
				})
			} else {
				return false;
			}
		});
	}
</script>

<style>
	.logo{
		width: 53px;
		height: 53px;
		border-radius: 50%;
	}
	.header-right{
		display: flex;
		align-items: center;
		width: 100%;
		justify-content: space-between;
	}
	.header-right-right{
		display: flex;
		align-items: center;
		padding-right: 20px;
	}
	.header-left{
		width: 225px;
		display: flex;
		align-items: center;
		justify-content: center;
	}
	.header{
		display: flex;
		border-bottom: 1px solid var(--el-border-color);
		flex-direction: row;
		align-items: center;
	}
	.body{
		position: absolute;
		top: 68px;
		bottom: 0;
		left: 0;
		right: 0;
	}
	.aside{
		border-right: 1px solid var(--el-border-color);
	}
	.main{
		position: relative;
	}
	.page-header{
		display: flex;
		justify-content: space-between;
		align-items: center;
	}
	.page-search-form-item{
		margin-top: 0;
		margin-bottom: 0;
		padding-top: 10px;
	}
	.page-search-select{
		width: 150px;
	}
	.page-buttons{
		display: flex;
		align-items: center;
		margin-top: 15px;
	}
	.page-table{
		margin-top: 20px;
	}
	.page-pagination{
		margin-top: 20px;
	}
	.icon:hover{
		color: #409EFF;
	}
</style>
