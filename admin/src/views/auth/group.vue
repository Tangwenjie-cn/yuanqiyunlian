<template>
	<!--搜索栏-->
	<div class="page-header">
		<el-form :inline="true" :model="searchForm">
			<el-form-item class="page-search-form-item">
				<el-input v-model="searchForm.keyword" placeholder="用户组" />
			</el-form-item>
			<el-form-item class="page-search-form-item">
				<el-button type="primary" icon="search" @click="getList()" />
			</el-form-item>
		</el-form>
		<el-button type="primary" icon="refresh" @click="refresh">刷新</el-button>
	</div>
	
	<!--按钮栏-->
	<div class="page-buttons">
		<el-button type="primary" @click="addDialog=true">添加用户组</el-button>
	</div>
	<!--列表栏-->
	<div class="page-table">
		<el-table :data="list" stripe border table-layout="fixed" style="width: 100%">
		    <el-table-column align="center" prop="id" label="ID" width="80" />
			<el-table-column label="用户组" prop="name" width="180" align="center" />
			<el-table-column label="创建时间" prop="ctime_text" width="180" align="center"></el-table-column>
			<el-table-column fixed="right" label="操作">
				<template #default="scope">
					<el-button type="primary" size="small" @click="edit(scope)">编辑</el-button>					
					<el-button type="warning" size="small" @click="setAuth(scope)">设置权限</el-button>					
					<el-popconfirm confirm-button-text="删除" cancel-button-text="取消" title="确定删除?" @confirm="del(scope)">
						<template #reference>
							<el-button type="danger" size="small">删除</el-button>
						</template>
					</el-popconfirm>				
			    </template>
			</el-table-column>
		</el-table>
		<!-- <el-pagination class="page-pagination" v-model:page-size="limit" v-model:current-page="page" :page-sizes="[10,20, 50, 100, 200]" layout="total,sizes,prev,pager,next" :total="total"/> -->
	</div>
	<!--弹出层-->
	<el-dialog v-model="addDialog" title="添加">
		<el-form :model="addForm" ref="refAddForm" label-width="80px">
			<el-form-item label="名称" >
				<el-input v-model="addForm.name"></el-input>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="add">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
	<el-dialog v-model="editDialog" title="编辑">
		<el-form :model="editForm" label-width="80px">
			<el-form-item label="名称">
				<el-input v-model="editForm.name"></el-input>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="editSub">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
	<el-dialog v-model="authDialog" title="权限">
		<el-tree ref="treeRef" :data="tree" 
		show-checkbox
		node-key="id"
		:default-checked-keys="checkedKeys"
		:props="defaultProps"></el-tree>
		<el-button type="primary" @click="authSub" style="margin-top: 20px;">提交</el-button>  
	</el-dialog>
</template>

<script setup>
	import {ref,onMounted,inject} from "vue";
	const myApi=inject('myApi');
	const ElMessage=inject('ElMessage');
	const list=ref([]);
	const addDialog=ref(false);
	const editDialog=ref(false);
	const authDialog=ref(false);
	const addForm=ref({});
	const refAddForm=ref();
	const editForm=ref({});
	const searchForm=ref({});
	const tree=ref([]);
	const defaultProps=ref({label: 'name'});
	const tree_groupid=ref('');
	const treeRef=ref();
	const checkedKeys=ref([]);
	onMounted(()=>{
		getList()
	});
	function refresh(){
		searchForm.value={},
		getList()
	}
	function getList(){
		//searchForm.value.page=page.value
		//searchForm.value.limit=limit.value
		myApi('groupList',searchForm.value).then((res)=>{
			
			list.value=res.data
			//total.value=res.data.total
		})
	}
	function add(){
		myApi('addGroup',addForm.value,'post').then((res)=>{
			addDialog.value=false
			ElMessage.success(res.msg)
			getList()
			refAddForm.value.resetFields()
		})
	}
	function edit(scope){
		editDialog.value=true
		editForm.value={
			name:scope.row.name,
			id:scope.row.id
		}
	}
	function editSub(){
		myApi('editGroup',editForm.value,'post').then((res)=>{
			editDialog.value=false
			ElMessage.success(res.msg)
			getList()
		})
	}
	function del(scope){
		myApi('delGroup',{id:scope.row.id},'post').then((res)=>{
			ElMessage({message:res.msg,type:'success'})
			getList()
		})
	}
	function setAuth(scope){
		tree.value=[]
		checkedKeys.value=[]
		tree_groupid.value=scope.row.id
		myApi('getGroupAuth',{id:scope.row.id},'post').then((res)=>{
			authDialog.value=true
			tree.value=res.data.list
			checkedKeys.value=res.data.checked
		})
	}
	function authSub(){

		let data={
			auth:treeRef.value.getCheckedKeys(),
			half_auth:treeRef.value.getHalfCheckedKeys(),
			id:tree_groupid.value
		}
		myApi('setGroupAuth',data,'post').then((res)=>{
			authDialog.value=false
			ElMessage({message:res.msg,type:'success'})
		})
	}
	// watch([page,limit],()=>{
	// 	getList()
	// });
</script>

<style>
	
</style>

