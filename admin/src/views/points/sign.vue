<template>
	<!--搜索栏-->
	<div class="page-header">
		<el-form :inline="true" :model="searchForm">
			<el-form-item class="page-search-form-item page-search-select">
				<el-select v-model="searchForm.type" clearable placeholder="奖励类型">
					<el-option :value="1" label="连续签到"></el-option>
					<el-option :value="2" label="累积签到"></el-option>
				</el-select>
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
			<el-table-column align="center" prop="days" label="签到天数" width="120" />	
			<el-table-column align="center" prop="points" label="奖励积分" width="120" />
			<el-table-column label="类型" width="120" align="center">
				<template #default="scope">
					<el-text v-if="scope.row.type==1">连续签到</el-text>
					<el-text v-else-if="scope.row.type==2">累积签到</el-text>
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
			<el-table-column fixed="right" label="操作" width="180" align="center">
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
		layout="total,sizes,prev,pager,next" :total="total"/>
	</div>
	<!--弹出层-->
	<el-dialog v-model="addDialog" title="添加" @close="dialogClose">
		<el-form :model="addForm" label-width="80px" ref="refAddForm">
			<el-form-item label="签到类型" prop="type" required>
				<el-radio v-model="addForm.type" label="连续签到" :value="1" />
				<el-radio v-model="addForm.type" label="累积签到" :value="2" />				
			</el-form-item>
			<el-form-item label="签到天数" prop="days" required>
				<el-input-number v-model="addForm.days">
                    <template #suffix>
                        <span>天</span>
                    </template>
                </el-input-number>
			</el-form-item>
            <el-form-item label="奖励积分" prop="points" required>
				<el-input-number v-model="addForm.points"></el-input-number>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="add">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
	<el-dialog v-model="editDialog" title="编辑" @close="dialogClose">		
		<el-form :model="editForm" ref="refEditForm" label-width="80px">
			<el-form-item label="签到类型" prop="type" required>
				<el-radio v-model="editForm.type" label="连续签到" :value="1" />
				<el-radio v-model="editForm.type" label="累积签到" :value="2" />				
			</el-form-item>
			<el-form-item label="签到天数" prop="days" required>
				<el-input-number v-model="editForm.days">
                    <template #suffix>
                        <span>天</span>
                    </template>
                </el-input-number>
			</el-form-item>
            <el-form-item label="奖励积分" prop="points" required>
				<el-input-number v-model="editForm.points"></el-input-number>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="editSub">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
</template>

<script setup>
	import {ref,onMounted,watch,inject,useTemplateRef} from "vue";
	const myApi=inject('myApi');
	const ElMessage=inject('ElMessage');
	const list=ref([]);
	const searchForm=ref({});
	const page=ref(1);
	const limit=ref(10);
	const total=ref(0);
    const addDialog=ref(false);
	const editDialog=ref(false);
	const addForm=ref({status:1,sort:0});
	const editForm=ref({});
	const refAddForm=useTemplateRef('refAddForm');
	const refEditForm=useTemplateRef('refEditForm');
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
		if(typeof searchForm.value.type!="number"){
			delete searchForm.value.type
		}
		myApi('SignList',searchForm.value).then(res=>{		
			list.value=res.data.data
			total.value=res.data.total
		})
	}
	watch([page,limit],()=>{
		getList();
	});
	function dialogClose(){
		getList()
	}
	function add(){
		refAddForm.value.validate((valid)=>{
			if(valid){
				myApi('SignAdd',addForm.value,"post").then(res=>{					
					ElMessage.success(res.msg)
					addDialog.value=false					
				})
			}
		})
	}
	function editSub(){
		refEditForm.value.validate((valid)=>{
			if(valid){
				myApi('SignEdit',editForm.value,"post").then(res=>{
					ElMessage.success(res.msg)
					editDialog.value=false					
				})
			}
		})
	}
	function edit(scope){
		editForm.value={...scope.row}
		editDialog.value=true
	}
	function del(scope){
		myApi('SignDel',{id:scope.row.id},"post").then(res=>{			
			ElMessage.success(res.msg)
			getList()

		})
	}
</script>

<style>	
	
</style>