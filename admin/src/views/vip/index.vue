<template>
	<!--搜索栏-->
	<div class="page-header">
		<el-form :inline="true" :model="searchForm">
			<el-form-item class="page-search-form-item page-search-select">
				<el-select v-model="searchForm.status" clearable placeholder="状态">
					<el-option :value="1" label="正常"></el-option>
					<el-option :value="-1" label="禁用"></el-option>
				</el-select>
			</el-form-item>
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
			<el-table-column align="center" prop="title" label="会员名称" width="120" />
			<el-table-column align="center" prop="price" label="价格" width="120" />	
			<el-table-column align="center" prop="days" label="时长" width="120" />
			<el-table-column align="center" prop="rebate" label="折扣" width="120" />
			<el-table-column align="center" prop="points" label="兑换积分" width="120" />
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
		<addVip v-if="addDialog" @submit="getList" />
	</el-dialog>
	<el-dialog v-model="editDialog" title="编辑">
		<editVip v-if="editDialog" v-model="editForm" @submit="getList" />
	</el-dialog>
</template>

<script setup>
	import {ref,onMounted,watch,inject} from "vue";
	import addVip from './index/add.vue';
	import editVip from './index/edit.vue';
	const myApi=inject('myApi');
	const ElMessage=inject('ElMessage');
	const list=ref([]);
	const searchForm=ref({});
	const page=ref(1);
	const limit=ref(10);
	const total=ref(0);
	const addDialog=ref(false);
	const editDialog=ref(false);
	const editForm=ref({});
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
		myApi('VipList',searchForm.value,'post').then((res)=>{
			list.value=res.data.data
			total.value=res.data.total
			addDialog.value=false
			editDialog.value=false
		})
	}
	function edit(scope){
		editDialog.value=true;
		editForm.value=scope.row;
	}
	function del(scope){
		myApi('DelVip',{id:scope.row.id},'post').then((res)=>{
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