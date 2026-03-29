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
		<el-table :data="list" stripe border table-layout="fixed" style="width: 100%">
		    <el-table-column align="center" prop="id" label="ID" width="80" />
			<el-table-column label="名称" prop="title" width="100" align="center" />
            <el-table-column label="地址" prop="path" align="center" />
            <el-table-column label="备注" prop="remark" align="center" />
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
			<el-table-column label="操作" fixed="right">
				<template #default="scope">
					<el-button type="primary" size="small" @click="edit(scope)">编辑</el-button>
					<el-button type="success" size="small" @click="set(scope)" v-if="scope.row.status==-1">设为首页</el-button>						
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
	<el-dialog v-model="addDialog" width="66%" title="添加">
		<el-form :model="addForm" ref="refAddForm" label-width="80px">
			<el-form-item label="名称" prop="title" required>
				<el-input v-model="addForm.title"></el-input>
			</el-form-item>
			<el-form-item label="地址" prop="path" required>
				<el-input v-model="addForm.path"></el-input>
			</el-form-item>
			<el-form-item label="备注" prop="remark">
				<el-input v-model="addForm.remark" type="textarea"></el-input>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="add">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
	<el-dialog v-model="editDialog" width="66%" title="编辑">		
		<el-form :model="editForm" ref="refEditForm" label-width="80px">
			<el-form-item label="名称" prop="title" required>
				<el-input v-model="editForm.title"></el-input>
			</el-form-item>
			<el-form-item label="地址" prop="path" required>
				<el-input v-model="editForm.path"></el-input>
			</el-form-item>
			<el-form-item label="备注" prop="remark">
				<el-input v-model="editForm.remark" type="textarea"></el-input>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="editSub">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
</template>

<script setup>
	import {ref,onMounted,inject,useTemplateRef,watch} from "vue";
	const myApi=inject('myApi');
	const ElMessage=inject('ElMessage');

	const list=ref([]);
    const searchForm=ref({});
    const page=ref(1);
	const limit=ref(10);
	const total=ref(0);

	const addDialog=ref(false);
	const addForm=ref({});
    const refAddForm=useTemplateRef('refAddForm');
	
	const editDialog=ref(false);
	const editForm=ref({});
    const refEditForm=useTemplateRef('refEditForm');

	onMounted(()=>{
		getList()
	});
	function refresh(){
        searchForm.value={}
		getList()
	}
	function getList(){
        searchForm.value.page=page.value
		searchForm.value.limit=limit.value
		myApi('PagesList',searchForm.value,'get').then((res)=>{
			list.value=res.data.data
            total.value=res.data.total
		})
	}
    function add(){
		refAddForm.value.validate((valid)=>{
			if(valid){
				myApi('PagesAdd',addForm.value,'post').then((res)=>{
					ElMessage({message:res.msg,type:'success'})
					addDialog.value=false
					getList()
				})
			}
		})
	}
	function edit(scope){
		editForm.value={...scope.row}
		editDialog.value=true
	}
    function editSub(){
		refEditForm.value.validate(valid=>{
			if(valid){
				myApi('PagesEdit',editForm.value,'post').then((res)=>{
					editDialog.value=false
					ElMessage({message:res.msg,type:'success'})
					getList()
				})
			}
		})
	}
	function del(scope){
		myApi('PagesDel',{id:scope.row.id},'post').then((res)=>{
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