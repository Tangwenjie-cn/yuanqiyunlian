
<template>
	<!--搜索栏-->
	<div class="page-header">
		<el-form :inline="true" :model="searchForm">
			<el-form-item class="page-search-form-item page-search-select">
				<el-select v-model="searchForm.status" clearable placeholder="状态">
					<el-option :value="1" label="通过"></el-option>
					<el-option :value="-1" label="拒绝"></el-option>
                    <el-option :value="2" label="待审核"></el-option>
				</el-select>
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
			<el-table-column align="center" prop="nickname" label="昵称" width="120" />	
			<el-table-column align="center" prop="balance" label="余额" width="120" />
			<el-table-column align="center" prop="points" label="积分" width="120" />
			<el-table-column align="center" prop="promotion" label="推广人数" width="120" />
			<el-table-column label="状态" width="120" align="center">
				<template #default="scope">
                    <el-text type="success" v-if="scope.row.status==1">通过</el-text>
                    <el-text type="warning" v-else-if="scope.row.status==0">待审核</el-text>
                    <el-text type="danger" v-else>拒绝</el-text>
				</template>
			</el-table-column>
            <el-table-column align="center" prop="remark" label="备注" />
			<el-table-column label="创建时间" width="180" align="center">
				<template #default="scope">
					{{new Date(scope.row.ctime).toLocaleString()}}
				</template>
			</el-table-column>
			<el-table-column fixed="right" label="操作" width="180" align="center">
				<template #default="scope">
                    <template v-if="scope.row.status==0">
					<el-button type="primary" size="small" @click="check(scope,1)">通过</el-button>
					<el-button type="danger" size="small" @click="check(scope,-1)">拒绝</el-button>
                    </template>									
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
	<el-dialog v-model="dialog" title="通过推广设置" width="600px" @close="dialogClose">
		<el-form :model="formData" ref="refForm">
			<el-form-item label="高推设置" prop="super_set" required>
				<el-radio-group v-model="formData.super_set">
					<el-radio :value="1">公共</el-radio>
					<el-radio :value="2">独有</el-radio>
				</el-radio-group>
			</el-form-item>
			<el-form-item label="高推佣金方式">
				<el-radio-group v-model="formData.super_type">
					<el-radio :value="1">固定</el-radio>
					<el-radio :value="2">百分比</el-radio>
				</el-radio-group>
			</el-form-item>
			<el-form-item label="高推佣金">
				<el-input-number v-model="formData.super_yj" :precision="2" />
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="submitForm()">提交</el-button>
			</el-form-item>
		</el-form>
	</el-dialog>
    <el-dialog v-model="dialog1" title="拒绝" width="600px" @close="dialogClose">
        <el-form :model="formData1" ref="refForm1">
            <el-form-item label="拒绝理由" prop="remark" required>
				<el-input
                    v-model="formData1.remark"
                    type="textarea"
                    placeholder="拒绝理由"
                />
			</el-form-item>
            <el-form-item>
				<el-button type="primary" @click="submitForm1()">提交</el-button>
			</el-form-item>
        </el-form>
    </el-dialog>
</template>

<script setup>
	import {ref,onMounted,watch,inject} from "vue";
	const myApi=inject('myApi');
	const ElMessage=inject('ElMessage');
	const list=ref([]);
	const searchForm=ref({});
	const page=ref(1);
	const limit=ref(10);
	const total=ref(0);
	const dialog=ref(false);
	const refForm=ref();
	const formData=ref({})
    const dialog1=ref(false);
	const refForm1=ref();
	const formData1=ref({})
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
		myApi('GetApplySuper',searchForm.value).then((res)=>{
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
	function check(scope,status){
        if(status==1){
			dialog.value=true
            formData.value.super_set=scope.row.super_set
            formData.value.super_type=scope.row.super_type
            formData.value.super_yj=scope.row.super_yj
            formData.value.id=scope.row.id
            formData.value.status=status
		}else{
            dialog1.value=true
            formData1.value.id=scope.row.id
            formData1.value.status=status
        }

	}
	function submitForm(){
		refForm.value.validate((valid)=>{
			if(valid){
			  	myApi('SetApplySuper',formData.value,'post').then(()=>{
					ElMessage.success('设置成功')
					dialog.value=false
					getList()
				})  
			}
		})
	}
    function submitForm1(){
		refForm1.value.validate((valid)=>{
			if(valid){
			  	myApi('SetApplySuper',formData1.value,'post').then(()=>{
					ElMessage.success('设置成功')
					dialog1.value=false
					getList()
				})  
			}
		})
	}
</script>

<style>
	
</style>