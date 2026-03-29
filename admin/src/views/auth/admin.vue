
<template>
	<!--搜索栏-->
	<div class="page-header">
		<el-form :inline="true" :model="searchForm">
			<el-form-item class="page-search-form-item page-search-select">
				<el-select v-model="searchForm.status" placeholder="状态">
					<el-option value="0" label="全部"></el-option>
					<el-option value="1" label="正常"></el-option>
					<el-option value="-1" label="禁用"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item class="page-search-form-item">
				<el-input v-model="searchForm.keyword" placeholder="账号" />
			</el-form-item>
			<el-form-item class="page-search-form-item">
				<el-button type="primary" icon="search" @click="getList()" />
			</el-form-item>
		</el-form>
		<el-button type="primary" icon="refresh" @click="refresh">刷新</el-button>
	</div>
	
	<!--按钮栏-->
	<div class="page-buttons">
		<el-button type="primary" @click="addDialog=true">添加用户</el-button>
	</div>
	<!--列表栏-->
	<div class="page-table">
		<el-table :data="list" stripe border table-layout="fixed" style="width: 100%">
		    <el-table-column align="center" prop="id" label="ID" width="80" />
			<el-table-column label="账号" prop="username" width="150" align="center" />
			<el-table-column label="状态" width="180" align="center">
				<template #default="scope">
					<span v-if="scope.row.status==1" style="color: #67C23A;">正常</span>
					<span v-if="scope.row.status==-1" style="color: #F56C6C;">禁用</span>
				</template>
			</el-table-column>
			<el-table-column label="用户组" width="150" align="center">
				<template #default="scope">
					{{scope.row.AuthGroup.name}}
				</template>
			</el-table-column>
			<el-table-column label="登录时间" width="180" align="center">
				<template #default="scope">
					{{new Date(scope.row.ltime).toLocaleString()}}
				</template>
			</el-table-column>
			<el-table-column label="登录IP" prop="login_ip" width="150" align="center" />
			<el-table-column label="创建时间" width="180" align="center">
				<template #default="scope">
					{{new Date(scope.row.ctime).toLocaleString()}}
				</template>
			</el-table-column>
			<el-table-column label="更新时间" width="180" align="center">
				<template #default="scope">
					{{new Date(scope.row.utime).toLocaleString()}}
				</template>
			</el-table-column>
			<el-table-column label="操作" width="150" fixed="right">
				<template #default="scope">
					<el-button type="primary" size="small" @click="edit(scope)">编辑</el-button>						
					<el-popconfirm confirm-button-text="删除" cancel-button-text="取消" title="确定删除?" @confirm="del(scope)">
						<template #reference>
							<el-button type="danger" size="small">删除</el-button>
						</template>
					</el-popconfirm>
			    </template>
			</el-table-column>
		</el-table>
		<el-pagination class="page-pagination" 
		v-model:page-size="limit" 
		v-model:current-page="page" 
		:page-sizes="[10,20, 50, 100, 200]" 
		layout="total,sizes,prev,pager,next" 
		:total="total"/>
	</div>
	<!--弹出层-->
	<el-dialog v-model="addDialog" title="添加">
		<el-form :model="addForm" ref="refAddForm" label-width="80px">
			<el-form-item label="用户组" prop="group_id" required>
				<el-select v-model="addForm.group_id" placeholder="请选择" >
					<el-option v-for="(item,index) in groupList" :value="item.id" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item label="账号" prop="username" required>
				<el-input v-model="addForm.username"></el-input>
			</el-form-item>
			<el-form-item label="密码" prop="password" required>
				<el-input v-model="addForm.password" type="password"></el-input>
			</el-form-item>
			<el-form-item label="状态" >
				<el-radio v-model="addForm.status" :value="1">启用</el-radio>
				<el-radio v-model="addForm.status" :value="-1">禁用</el-radio>				
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="add">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
	<el-dialog v-model="editDialog" title="编辑">		
		<el-form :model="editForm" ref="refEditForm" label-width="80px">
			<el-form-item label="用户组" prop="group_id" required>
				<el-select v-model="editForm.group_id" placeholder="请选择">
					<el-option v-for="(item,index) in groupList" :value="item.id" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item label="账号">
				<el-input v-model="editForm.username" disabled></el-input>
			</el-form-item>
			<el-form-item label="修改密码">
				<el-input v-model="editForm.password" placeholder="修改时填写" type="password"></el-input>
			</el-form-item>
			<el-form-item label="状态">
				<el-radio v-model="editForm.status" :value="1">启用</el-radio>
				<el-radio v-model="editForm.status" :value="-1">禁用</el-radio>				
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="editSub">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
</template>

<script setup>
	import {ref,onMounted,inject,watch} from "vue";
	const myApi=inject('myApi');
	const ElMessage=inject('ElMessage');
	const list=ref([]);
	const groupList=ref([]);
	const addDialog=ref(false);
	const editDialog=ref(false);
	const addForm=ref({status:'1'});
	const editForm=ref({});
	const refAddForm=ref()
	const refEditForm=ref()
	const searchForm=ref({});
	const page=ref(1);
	const limit=ref(10);	
	const total=ref(0);
	onMounted(()=>{
		getList()
		getGroupList()
	});
	function refresh(){
		searchForm.value={}
		getList()
		getGroupList()
	}
	function getGroupList(){
		myApi('groupList').then(res=>{
			groupList.value=res.data
		})
	}
	function getList(){
		searchForm.value.page=page.value
		searchForm.value.limit=limit.value
		myApi('adminList',searchForm.value,'post').then((res)=>{
			list.value=res.data.data
			//groupList.value=res.data.group
			total.value=res.data.total
		})
	}
	function add(){
		refAddForm.value.validate(valid=>{
			if(valid){
				myApi('addAdmin',addForm.value,'post').then((res)=>{
					addDialog.value=false
					ElMessage({message:res.msg,type:'success'})
					getList()
					refAddForm.value.resetFields()
				})
			}
		})
		
	}
	function edit(scope){
		editDialog.value=true
		editForm.value={
			username:scope.row.username,
			id:scope.row.id,
			status:scope.row.status,
			group_id:scope.row.group_id
		}
	}
	function editSub(){
		refEditForm.value.validate(valid=>{
			if(valid){
				myApi('editAdmin',editForm.value,'post').then((res)=>{
					editDialog.value=false
					ElMessage({message:res.msg,type:'success'})
					getList()
				})
			}
		})
	}
	function del(scope){
		myApi('delAdmin',{id:scope.row.id},'post').then((res)=>{
			ElMessage({message:res.msg,type:'success'})
			getList()
		})
	}
	watch([page,limit],()=>{
		getList()
	});	
</script>

<style>
	
</style>

