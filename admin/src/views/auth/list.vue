
<template>
	<!--搜索栏-->
	<div class="page-header">
		<el-form :inline="true" :model="searchForm">			
		</el-form>
		<el-button type="primary" icon="refresh" @click="refresh">刷新</el-button>
	</div>
	
	<!--按钮栏-->
	<div class="page-buttons">
		<el-button type="primary" @click="addDialog=true">添加权限</el-button>
	</div>
	<!--列表栏-->
	<div class="page-table">
		<el-table :data="list" stripe border table-layout="fixed" row-key="id" style="width: 100%">
			<el-table-column align="center" width="80" />
		    <el-table-column align="center" prop="id" label="ID" width="80" />
		    <el-table-column label="权限名" prop="name" align="center" />
			<el-table-column label="图标" width="80" align="center">
				<template #default="scope">
					<el-icon v-if="scope.row.icon">
						<component :is="scope.row.icon" />
					</el-icon>
				</template>
			</el-table-column>
			<el-table-column align="center" prop="url" label="接口地址" />
			<el-table-column align="center" prop="path" label="前端路由" />
			<el-table-column label="菜单" width="80" align="center">
				<template #default="scope">
					<el-text v-if="scope.row.is_menu" type="primary">是</el-text>
					<el-text v-else>否</el-text>
				</template>
			</el-table-column>
			<el-table-column label="排序" prop="sort" width="80" align="center" />
			<el-table-column width="150" label="操作">
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
			<el-form-item label="上级权限" prop="pid" required>
				<el-cascader v-model="addForm.pid" :options="list" :props="props">
				</el-cascader>
			</el-form-item>
			<el-form-item label="权限名" prop="name" required>
				<el-input v-model="addForm.name"></el-input>
			</el-form-item>
			<el-form-item label="权限地址">
				<el-input v-model="addForm.url" placeholder="控制器/方法"></el-input>
			</el-form-item>
			<el-form-item label="前端路由">
				<el-input v-model="addForm.path"></el-input>
			</el-form-item>
			<el-form-item label="图标">
				<el-input v-model="addForm.icon"></el-input>
			</el-form-item>
			<el-form-item label="菜单" prop="is_menu" required>
				<el-switch
				v-model="addForm.is_menu"
				:active-value="1"
				:inactive-value="0"
				/>
			</el-form-item>
			<el-form-item label="排序">
				<el-input type="number" v-model="addForm.sort"></el-input>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="add">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
	<el-dialog v-model="editDialog" title="编辑">
		<el-form :model="editForm" ref="refEditForm" label-width="80px">
			<el-form-item label="上级权限" prop="pid" required>
				<el-cascader v-model="editForm.pid" :options="list"  :props="props">
				</el-cascader>
			</el-form-item>
			<el-form-item label="权限名" prop="name" required>
				<el-input v-model="editForm.name"></el-input>
			</el-form-item>
			<el-form-item label="权限地址">
				<el-input v-model="editForm.url" placeholder="控制器/方法"></el-input>
			</el-form-item>
			<el-form-item label="前端路由">
				<el-input v-model="editForm.path"></el-input>
			</el-form-item>
			<el-form-item label="图标">
				<el-input v-model="editForm.icon"></el-input>
			</el-form-item>
			<el-form-item label="菜单" prop="is_menu" required>
				<el-switch
				v-model="editForm.is_menu"
				:active-value="1"
				:inactive-value="0"
				/>
			</el-form-item>
			<el-form-item label="排序">
				<el-input type="number" v-model="editForm.sort"></el-input>
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
	const addForm=ref({pid:0,is_menu:1,sort:0});
	const editForm=ref({});
	const refAddForm=ref()
	const refEditForm=ref()
	const props=ref({
		checkStrictly: true,
		label: 'name',
		value: 'id',
		emitPath: false,
	});
	onMounted(()=>{
		getList()
	});
	function refresh(){
		getList()
	}
	function getList(){
		myApi('authList',null,'post').then((res)=>{
			list.value=res.data
		})
	}
	function add(){
		refAddForm.value.validate(valid=>{
			if(valid){
				addForm.value.sort=Number(addForm.value.sort)
				myApi('addAuth',addForm.value,'post').then((res)=>{
					addDialog.value=false
					ElMessage({message:res.msg,type:'success'})
					getList()
					refAddForm.value.resetFields()
				})
			}
		})	
	}
	function edit(scope){
		myApi('getAuth',{id:scope.row.id}).then((res)=>{
			editForm.value=res.data
			editDialog.value=true
		})	
	}
	function editSub(){
		refEditForm.value.validate(valid=>{
			if(valid){
				editForm.value.sort=Number(editForm.value.sort)
				myApi('editAuth',editForm.value,'post').then((res)=>{
					editDialog.value=false
					ElMessage({message:res.msg,type:'success'})
					getList()
				})
			}
		})		
	}
	function del(scope){
		myApi('delAuth',{id:scope.row.id},'post').then((res)=>{
			ElMessage({message:res.msg,type:'success'})
			getList()
		})
	}
	
</script>

<style>
	
</style>
