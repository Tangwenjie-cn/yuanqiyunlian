
<template>
	<!--搜索栏-->
	<div class="page-header">
		<el-form :inline="true" :model="searchForm">
			
		</el-form>
		<el-button type="primary" icon="refresh" @click="refresh">刷新</el-button>
	</div>
	
	<!--按钮栏-->
	<div class="page-buttons">
		<el-button type="primary" @click="addDialog=true">添加分类</el-button>
	</div>
	<!--列表栏-->
	<div class="page-table">
		<el-table :data="list" stripe border table-layout="fixed" style="width: 100%">
		    <el-table-column align="center" prop="id" label="ID" width="80" />
			<el-table-column label="名称" prop="name" width="300" align="center" />
			<el-table-column label="缩略图" width="120" align="center">
				<template #default="scope">
					<el-image :src="scope.row.thumb" fit="contain">
						<template #error>
							<el-icon><Picture /></el-icon>
						</template>
					</el-image>
				</template>
			</el-table-column>			
			<el-table-column label="状态" width="120" align="center">
				<template #default="scope">
					<span v-if="scope.row.status==1" style="color: #67C23A;">正常</span>
					<span v-if="scope.row.status==-1" style="color: #F56C6C;">禁用</span>
				</template>
			</el-table-column>
			<el-table-column label="创建时间" width="180" align="center">
				<template #default="scope">
					{{new Date(scope.row.ctime).toLocaleString()}}
				</template>
			</el-table-column>
			<el-table-column label="操作">
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
		v-model:page-size="limit" v-model:current-page="page" 
		:page-sizes="[10,20, 50, 100, 200]" layout="total,sizes,prev,pager,next" :total="total"/>
	</div>
	<!--弹出层-->
	<el-dialog v-model="addDialog" title="添加">
		<el-form :model="addForm" label-width="80px" ref="refAddForm">
			<el-form-item label="名称" required>
				<el-input v-model="addForm.name"></el-input>
			</el-form-item>
			<el-form-item label="封面">
				<uploadImg v-model="addForm.thumb"></uploadImg>
			</el-form-item>
			<el-form-item label="状态" required>
				<el-radio v-model="addForm.status" label="启用" :value="1" />
				<el-radio v-model="addForm.status" label="禁用" :value="-1" />				
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="add">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
	<el-dialog v-model="editDialog" title="编辑">		
		<el-form :model="editForm" label-width="80px" ref="refEditForm">
			<el-form-item label="名称" required>
				<el-input v-model="editForm.name"></el-input>
			</el-form-item>
			<el-form-item label="封面">
				<uploadImg v-model="editForm.thumb"></uploadImg>
			</el-form-item>
			<el-form-item label="状态" required>
				<el-radio v-model="editForm.status" label="启用" :value="1" />
				<el-radio v-model="editForm.status" label="禁用" :value="-1" />			
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="editSub">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
</template>

<script setup>
	import {ref,onMounted,watch,inject} from "vue";
	import uploadImg from "../../components/upload/uploadImg.vue"
	const myApi=inject('myApi');
	const ElMessage=inject('ElMessage');
	const list=ref([]);
	const addDialog=ref(false);
	const editDialog=ref(false);
	const addForm=ref({status:1});
	const refAddForm=ref(null);
	const editForm=ref({});
	const refEditForm=ref(null);
	const searchForm=ref({});
	const page=ref(1);
	const limit=ref(10);
	const total=ref(0);
	onMounted(()=>{
		getList()
	});
	function refresh(){
		searchForm.value={},
		getList()
	}
	function getList(){
		searchForm.value.page=page.value
		searchForm.value.limit=limit.value
		myApi('articleSort',searchForm.value,'post').then((res)=>{
			list.value=res.data.data
			total.value=res.data.total
		})
	}
	function add(){
		refAddForm.value.validate((valid)=>{
			if(valid){
				myApi('addArticleSort',addForm.value,'post').then((res)=>{
					addDialog.value=false
					ElMessage({message:res.msg,type:'success'})
					getList()
				})
			}
		})
		
	}
	function edit(scope){
		editDialog.value=true
		editForm.value={
			id:scope.row.id,
			status:scope.row.status,
			name:scope.row.name,
			thumb:scope.row.thumb
		}
	}
	function editSub(){
		refEditForm.value.validate((valid)=>{
			if(valid){
				myApi('editArticleSort',editForm.value,'post').then((res)=>{
					editDialog.value=false
					ElMessage({message:res.msg,type:'success'})
					getList()
				})
			}
		})
		
	}
	function del(scope){
		myApi('delArticleSort',{id:scope.row.id},'post').then((res)=>{
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