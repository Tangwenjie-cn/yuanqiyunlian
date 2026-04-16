
<template>
	<!--搜索栏-->
	<div class="page-header">
		<el-form :inline="true" :model="searchForm">
			<el-form-item class="page-search-form-item">
                <el-date-picker
                    v-model="searchForm.date"
                    type="datetimerange"
                    start-placeholder="开始时间"
                    end-placeholder="结束时间"
                    value-format="YYYY-MM-DD HH:mm:ss"
                />
            </el-form-item>
			<el-form-item class="page-search-form-item">
				<el-input v-model="searchForm.keyword" clearable placeholder="UID" />
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
			<el-table-column align="center" prop="uid" label="UID" width="80" />
			<el-table-column label="头像" width="120" align="center">
				<template #default="scope">
					<el-avatar :size="36" :src="scope.row.avatar" icon="Avatar" />
				</template>
			</el-table-column>
			<el-table-column align="center" prop="nickname" label="昵称" width="120" />	
			<el-table-column align="center" prop="before" label="变动前" width="120" />
			<el-table-column align="center" prop="after" label="变动后" width="120" />
			<el-table-column align="center" label="积分" width="120" >
                <template #default="scope">
                    <el-text v-if="scope.row.mode==1" type="primary">+{{scope.row.points}}</el-text>
					<el-text v-else type="danger">-{{scope.row.points}}</el-text>
				</template>
			</el-table-column>
			<el-table-column align="center" label="类型" width="120">
				<template #default="scope">
					<span v-if="scope.row.type==1">签到</span>
					<span v-else-if="scope.row.type==2">兑换会员</span>
					<span v-else-if="scope.row.type==3">兑换商品</span>
				</template>
			</el-table-column>
            <el-table-column label="创建时间" width="180" align="center">
				<template #default="scope">
					{{new Date(scope.row.ctime).toLocaleString()}}
				</template>
			</el-table-column>
			<el-table-column align="center" prop="remark" label="备注" />
			
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
		myApi('PointsLog',searchForm.value).then((res)=>{
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