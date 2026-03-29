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
			<el-form-item class="page-search-form-item page-search-select">
				<el-select v-model="searchForm.status" clearable placeholder="状态">
					<el-option :value="1" label="正常"></el-option>
					<el-option :value="-1" label="禁用"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item class="page-search-form-item">
				<el-input v-model="searchForm.keyword" clearable placeholder="UID|昵称" />
			</el-form-item>
			<el-form-item class="page-search-form-item">
				<el-input v-model="searchForm.pid" clearable placeholder="用户推广查询(UID)" />
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
			<el-table-column align="center" label="推广人" width="120">
				<template #default="scope">
					<span v-if="scope.row.pid==0">无</span>
					<span v-else>{{scope.row.pname}}（UID：{{scope.row.pid}}）</span>
				</template>
			</el-table-column>
			<el-table-column label="头像" width="120" align="center">
				<template #default="scope">
					<el-avatar :size="36" :src="scope.row.avatar" icon="Avatar" />
				</template>
			</el-table-column>
			<el-table-column align="center" prop="nickname" label="昵称" width="120" />	
			<el-table-column align="center" prop="balance" label="余额" width="120" />
			<el-table-column align="center" prop="points" label="积分" width="120" />
			<el-table-column align="center" prop="promotion" label="推广人数" width="120" />
			<el-table-column align="center" label="高级推广" width="120">
				<template #default="scope">
					<span v-if="scope.row.is_super==1" style="color: #67C23A;">是</span>
					<span v-else style="color: #F56C6C;">否</span>
				</template>
			</el-table-column>
			<el-table-column align="center" label="高推设置" width="120">
				<template #default="scope">
					<span v-if="scope.row.super_set==1">公共</span>
					<span v-else-if="scope.row.super_set==2">独有</span>
					<span v-else>未设置</span>
				</template>
			</el-table-column>
			<el-table-column align="center" label="高推佣金方式" width="120">
				<template #default="scope">
					<span v-if="scope.row.super_type==1">固定额</span>
					<span v-else-if="scope.row.super_type==2">百分比</span>
					<span v-else>未设置</span>
				</template>
			</el-table-column>
			<el-table-column align="center" label="高推佣金" width="120">
				<template #default="scope">
					<span v-if="scope.row.super_type==1">{{scope.row.super_yj}}</span>
					<span v-else-if="scope.row.super_type==2">{{scope.row.super_yj}}%</span>
				</template>
			</el-table-column>
			<el-table-column label="状态" width="120" align="center">
				<template #default="scope">
					<span v-if="scope.row.status==1" style="color: #67C23A;">正常</span>
					<span v-else style="color: #F56C6C;">禁用</span>
				</template>
			</el-table-column>
			<el-table-column label="创建时间" width="180" align="center">
				<template #default="scope">
					{{new Date(scope.row.ctime).toLocaleString()}}
				</template>
			</el-table-column>
			<el-table-column fixed="right" label="操作" width="180" align="center">
				<template #default="scope">
					<el-button type="primary" size="small" @click="info(scope)">详情</el-button>
					<el-button type="primary" size="small" @click="vip(scope)">会员设置</el-button>					
					<el-button type="primary" size="small" @click="setTg(scope)">高级推广设置</el-button>					
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
	<el-dialog v-model="infoDialog" title="用户详情" width="600px" @close="dialogClose">
		<user-info :onShow="infoDialog" :id="uid"></user-info>
	</el-dialog>
	<el-dialog v-model="vipDialog" title="会员设置" width="600px" @close="dialogClose">
		<user-svip :onShow="vipDialog" :id="uid"></user-svip>
	</el-dialog>
	<el-dialog v-model="tgDialog" title="高级推广设置" width="600px" @close="dialogClose">
		<el-form :model="formData" ref="tgForm">
			<el-form-item label="高级推广" prop="is_super" required>
				<el-radio-group v-model="formData.is_super">
					<el-radio :value="1">是</el-radio>
					<el-radio :value="-1">否</el-radio>
				</el-radio-group>
			</el-form-item>
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
				<el-button type="primary" @click="submitTgForm()">提交</el-button>
			</el-form-item>
		</el-form>
	</el-dialog>
</template>

<script setup>
	import {ref,onMounted,watch,inject,useTemplateRef} from "vue";
	import userInfo from "./userInfo.vue";
	import userSvip from "./userSvip.vue";
	const myApi=inject('myApi');
	const ElMessage=inject('ElMessage');
	const store=inject('store');
	const list=ref([]);
	const searchForm=ref({});
	const page=ref(1);
	const limit=ref(10);
	const total=ref(0);
	const infoDialog=ref(false);
	const uid=ref(0);
	const tgDialog=ref(false);
	const tgForm=useTemplateRef('tgForm');
	const formData=ref({})
	const vipDialog=ref(false);
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
		myApi('userList',searchForm.value).then((res)=>{
			list.value=res.data.data
			total.value=res.data.total
		})
	}
	watch([page,limit],()=>{
		getList();
	});
	function info(scope){
		infoDialog.value=true
		uid.value=scope.row.id
	}
	function vip(scope){
		vipDialog.value=true
		uid.value=scope.row.id
	}
	function dialogClose(){
		getList()
	}
	function setTg(scope){
		tgDialog.value=true
		formData.value.is_super=scope.row.is_super
		formData.value.super_set=scope.row.super_set
		formData.value.super_type=scope.row.super_type
		formData.value.super_yj=scope.row.super_yj
		formData.value.id=scope.row.id
	}
	function submitTgForm(){
		tgForm.value.validate((valid)=>{
			if(valid){
			  	myApi('SetUserSuper',formData.value,'post').then(()=>{
					ElMessage.success('设置成功')
					tgDialog.value=false
					getList()
				})  
			}
		})
	}
</script>

<style>
	
</style>