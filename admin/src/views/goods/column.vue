
<template>
	<!--搜索栏-->
	<div class="page-header">
		<el-form :inline="true" :model="searchForm">
			
		</el-form>
		<el-button type="primary" icon="refresh" @click="refresh">刷新</el-button>
	</div>
	
	<!--按钮栏-->
	<div class="page-buttons">
		<el-button type="primary" @click="addDialog=true">添加</el-button>
	</div>
	<!--列表栏-->
	<div class="page-table">
		<el-table :data="list" stripe border table-layout="fixed" row-key="id" style="width: 100%">
		    <el-table-column align="center" prop="id" label="ID" width="80" />
		    <el-table-column label="名称" prop="name" width="350" align="center" />
            <el-table-column label="创建时间" width="180" align="center">
				<template #default="scope">
					{{new Date(scope.row.ctime).toLocaleString()}}
				</template>
			</el-table-column>
			<el-table-column fixed="right" width="150" label="操作">
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
	</div>
	<!--弹出层-->
	<el-dialog v-model="addDialog" title="添加">
		<el-form :model="addForm" ref="refAddForm" label-width="80px">
			<el-form-item label="名称" prop="name" required>
				<el-input v-model="addForm.name"></el-input>
			</el-form-item>
			<el-form-item label-width="80px">
				<el-button type="primary" @click="add">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
	<el-dialog v-model="editDialog" title="编辑">
		<el-form :model="editForm" ref="refEditForm" label-width="80px">
			<el-form-item label="名称" prop="name" required>
				<el-input v-model="editForm.name"></el-input>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="editSub">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
</template>

<script setup>
	import {ref,onMounted,inject} from "vue";
	const myApi=inject('myApi');
	const ElMessage=inject('ElMessage');
	const list=ref([]);
	const addDialog=ref(false);
	const editDialog=ref(false);
	const addForm=ref({});
	const editForm=ref({});
	const refAddForm=ref()
	const refEditForm=ref()
	const searchForm=ref({});
	onMounted(()=>{
		getList()
	});
	function refresh(){
		getList()
	}
	function getList(){
		myApi('GoodsColumn').then((res)=>{
			list.value=res.data
		})
	}
	function add(){
		refAddForm.value.validate(valid=>{
			if(valid){
				addForm.value.sort=Number(addForm.value.sort)
				myApi('AddGoodsColumn',addForm.value,'post').then((res)=>{
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
			name:scope.row.name,
			id:scope.row.id,
		}		
	}
	function editSub(){
		refEditForm.value.validate(valid=>{
			if(valid){
				myApi('EditGoodsColumn',editForm.value,'post').then((res)=>{
					editDialog.value=false
					ElMessage({message:res.msg,type:'success'})
					getList()
                    refEditForm.value.resetFields()
				})
			}
		})		
	}
	function del(scope){
		myApi('DelGoodsColumn',{id:scope.row.id},'post').then((res)=>{
			ElMessage({message:res.msg,type:'success'})
			getList()
		})
	}
	
</script>

<style>
	
</style>
