<template>
	<!--搜索栏-->
	<div class="page-header">
		<el-form :inline="true" :model="searchForm">
			<el-form-item class="page-search-form-item page-search-select">
				<el-select v-model="searchForm.status" clearable placeholder="状态">
					<el-option value="1" label="上架"></el-option>
					<el-option value="-1" label="下架"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item class="page-search-form-item">
				<el-input v-model="searchForm.title" clearable placeholder="标题" />
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
		<el-button type="primary" @click="batchDialog=true" icon="Promotion">自动导入文档</el-button>
	</div>
	<!--列表栏-->
	<div class="page-table">
		<el-table :data="list" stripe border table-layout="fixed" style="width: 100%">
		    <el-table-column align="center" prop="id" label="ID" width="80" />
			<el-table-column label="标题" prop="title" width="300" align="center" />
			<el-table-column label="缩略图" width="120" align="center">
				<template #default="scope">
					<el-image :src="scope.row.thumb" fit="contain" style="width: 59px;height: 59px;">
						<template #error>
							<el-icon><Picture /></el-icon>
						</template>
					</el-image>
				</template>
			</el-table-column>
			<el-table-column label="价格" prop="price" width="120" align="center" />
			<el-table-column label="积分" prop="points" width="120" align="center" />
			<el-table-column label="浏览量" prop="views" width="120" align="center" />
			<el-table-column label="分销" width="120" align="center">
				<template #default="scope">
					<el-text v-if="scope.row.retail_on==1" type="primary">开启</el-text>
					<el-text v-if="scope.row.retail_on==-1">关闭</el-text>
				</template>
			</el-table-column>
			<el-table-column label="类型" width="120" align="center">
				<template #default="scope">
					<el-text>{{ typeList[scope.row.type] }}</el-text>
				</template>
			</el-table-column>			
			<el-table-column label="状态" width="120" align="center">
				<template #default="scope">
					<span v-if="scope.row.status==1" style="color: #67C23A;">上架</span>
					<span v-if="scope.row.status==-1" style="color: #F56C6C;">下架</span>
				</template>
			</el-table-column>
			<el-table-column label="置顶" width="120" align="center">
				<template #default="scope">
					<span v-if="scope.row.is_top==1" style="color: #67C23A;">是</span>
					<span v-if="scope.row.is_top==-1" style="color: #F56C6C;">否</span>
				</template>
			</el-table-column>
			<el-table-column label="创建时间" width="180" align="center">
				<template #default="scope">
					{{new Date(scope.row.ctime).toLocaleString()}}
				</template>
			</el-table-column>
			<el-table-column fixed="right" label="操作" width="160">
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
		:total="total" 
		layout="total,sizes,prev,pager,next" />
	</div>
	<!--弹出层-->
	<el-dialog v-model="addDialog" width="66%" title="添加">
		<addGoods v-if="addDialog" @submit="getList"></addGoods>
	</el-dialog>
	<el-dialog v-model="editDialog" width="66%" title="编辑">		
		<editGoods v-if="editDialog" :id="editId" @submit="getList"></editGoods>
	</el-dialog>
	<el-dialog v-model="batchDialog" title="批量导入">		
		<batchImport v-if="batchDialog" />
	</el-dialog>
</template>

<script setup>
	import {ref,onMounted,inject,watch} from "vue";
	import addGoods from "./index/add.vue";
	import editGoods from "./index/edit.vue";
	import batchImport from "./index/batchImport.vue";
	const myApi=inject('myApi');
	const ElMessage=inject('ElMessage');
	const list=ref([]);
	const addDialog=ref(false);
	const editDialog=ref(false);
	const batchDialog=ref(false);
	const searchForm=ref({});
	const page=ref(1);
	const limit=ref(10);
	const total=ref(0);
	const editId=ref();
	const typeList=ref({
		1:"文档",
		2:"文章",
		3:"视频",
		4:"音频",
		5:"网盘"
	})
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
		myApi('goodsList',searchForm.value,'post').then((res)=>{
			list.value=res.data.data
			total.value=res.data.total
			addDialog.value=false
			editDialog.value=false
		});
	}	
	function edit(scope){
		editDialog.value=true
		editId.value=scope.row.id
	}	
	function del(scope){
		myApi('delGoods',{id:scope.row.id},'post').then((res)=>{
			ElMessage({message:res.msg,type:'success'})
			getList()
		})
	}
	watch([page,limit],()=>{
		getList()
	})	
</script>

<style>
	
</style>