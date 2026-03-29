<template>
	<!--搜索栏-->
	<div class="page-header">
		<el-form :inline="true" :model="searchForm">
			<el-form-item class="page-search-form-item">
				<el-input v-model="searchForm.keyword" clearable placeholder="名称" />
			</el-form-item>
			<el-form-item class="page-search-form-item">
				<el-button type="primary" icon="search" @click="getList()" />
			</el-form-item>
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
			<el-table-column align="center" prop="name" label="特权名称" width="200" />	
			<el-table-column label="特权类型" align="center" width="120">
				<template #default="scope">
					<span v-if="scope.row.type==1">商品类型</span>
					<span v-else-if="scope.row.type==2">商品分类</span>
					<span v-else-if="scope.row.type==3">商品栏目</span>
				</template>
			</el-table-column>
			<el-table-column label="创建时间" align="center" width="180">
				<template #default="scope">
					{{new Date(scope.row.ctime).toLocaleString()}}
				</template>
			</el-table-column>
			<el-table-column fixed="right" label="操作" width="150">
				<template #default="scope">					
					<el-button type="primary" size="small" @click="edit(scope)">查看</el-button>
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
		layout="total,sizes,prev,pager,next" :total="total"/>
	</div>
	<!--弹出层-->
	<el-dialog v-model="addDialog" title="添加">
		<addPrivilege v-if="addDialog" @submit="getList" />
	</el-dialog>
	<el-dialog v-model="editDialog" title="编辑">
		<editPrivilege v-if="editDialog" v-model="editForm" @submit="getList" />
	</el-dialog>
</template>

<script setup>
	import {ref,onMounted,watch,inject} from "vue";
    import addPrivilege from "./privilege/add.vue"
    import editPrivilege from "./privilege/edit.vue"
	const myApi=inject('myApi');
    const ElMessage=inject('ElMessage');
	const list=ref([]);
	const searchForm=ref({});
	const page=ref(1);
	const limit=ref(10);
	const total=ref(0);
    const addDialog=ref(false);
	const editDialog=ref(false);
	const editForm=ref();
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
		myApi('PrivilegeList',searchForm.value,'post').then((res)=>{
			list.value=res.data.data
			total.value=res.data.total
            addDialog.value=false
			editDialog.value=false
		})
	}
	function edit(scope){
		editDialog.value=true;
		editForm.value=scope.row
	}
    function del(scope){
		myApi('DelPrivilege',{id:scope.row.id},'post').then((res)=>{
			ElMessage({message:res.msg,type:'success'})
			getList()
		})
	}
	watch([page,limit],()=>{
		getList();
	});	
</script>

<style>
	
</style>