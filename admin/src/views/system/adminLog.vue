<template>
	<!--搜索栏-->
	<div class="page-header">
		<el-form :inline="true" :model="searchForm">
			<el-form-item class="page-search-form-item">
				<el-input v-model="searchForm.keyword" clearable placeholder="账号" />
			</el-form-item>
			<el-form-item class="page-search-form-item">
				<el-date-picker v-model="searchForm.date" 
				type="datetimerange" clearable 
				value-format="YYYY-MM-DD HH:mm:ss" 
				start-placeholder="开始日期" end-placeholder="结束日期" />
			</el-form-item>
			<el-form-item class="page-search-form-item">
				<el-button type="primary" icon="search" @click="getList()" />
			</el-form-item>
		</el-form>
		<el-button type="primary" icon="refresh" @click="refresh">刷新</el-button>
	</div>
	
	<!--按钮栏-->
	<div class="page-buttons">
		
	</div>
	<!--列表栏-->
	<div class="page-table">
		<el-table :data="list" stripe border table-layout="fixed" style="width: 100%">
		    <el-table-column align="center" prop="id" label="ID" width="80" />
			<el-table-column align="center" prop="account" label="管理员" width="120" />
			<el-table-column align="center" prop="ip" label="IP" width="120" />
			<el-table-column align="center" prop="uri" label="地址" />
			<el-table-column align="center" prop="body" label="请求体" />
            <el-table-column label="创建时间" align="center">
				<template #default="scope">
					{{new Date(scope.row.ctime).toLocaleString()}}
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



</template>

<script setup>
	import {ref,onMounted,watch,inject} from "vue";
	const myApi=inject('myApi');
	const list=ref([]);
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
		myApi('AdminLogList',searchForm.value).then((res)=>{
			list.value=res.data.data
			total.value=res.data.total
		})
	}
	watch([page,limit],()=>{
		getList();
	});
</script>

<style>
	
</style>