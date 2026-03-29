
<template>
	<!--搜索栏-->
	<div class="page-header">
		<el-form :inline="true" :model="searchForm">
		</el-form>
		<el-button type="primary" icon="refresh" @click="refresh">刷新</el-button>
	</div>
	
	<!--按钮栏-->
	<div class="page-buttons">
		<el-button type="primary" @click="addDialog=true">新增装修</el-button>
	</div>
	<!--列表栏-->
	<div class="page-table">
		<el-table :data="list" stripe border table-layout="fixed" style="width: 100%">
		    <el-table-column align="center" prop="id" label="ID" width="80" />
			<el-table-column label="名称" prop="title" width="150" align="center" />
            <el-table-column label="状态" width="180" align="center">
				<template #default="scope">
					<span v-if="scope.row.status==1" style="color: #67C23A;">启用</span>
					<span v-if="scope.row.status==-1" style="color: #F56C6C;">禁用</span>
				</template>
			</el-table-column>
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
	</div>
	<!--弹出层-->
	<el-dialog v-model="addDialog" width="66%" title="添加">
		<add-index></add-index>
	</el-dialog>
	<el-dialog v-model="editDialog" width="66%" title="编辑">		
		<edit-index v-model="editForm"></edit-index>
	</el-dialog>
</template>

<script setup>
	import {ref,onMounted,inject} from "vue";
    import addIndex from './index/add.vue'
	import editIndex from './index/edit.vue'
	const myApi=inject('myApi');
	const ElMessage=inject('ElMessage');
	const list=ref([]);
	const addDialog=ref(false);
	const editDialog=ref(false);
	const editForm=ref();
	onMounted(()=>{
		getList()
	});
	function refresh(){
		getList()
	}
	function getList(){
		myApi('ThemeIndex').then((res)=>{
			list.value=res.data
		})
	}
	function edit(scope){
		editForm.value=scope.row
		editDialog.value=true
	}
	function del(scope){
		myApi('ThemeDelIndex',{id:scope.row.id},'post').then((res)=>{
			ElMessage({message:res.msg,type:'success'})
			getList()
		})
	}
	function set(scope){
		myApi('ThemeSetIndex',{id:scope.row.id},'post').then((res)=>{
			ElMessage({message:res.msg,type:'success'})
			getList()
		})
	}
</script>

<style>
	
</style>