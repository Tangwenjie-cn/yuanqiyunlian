
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
			<el-table-column align="center" prop="type" label="类型" width="120">
				<template #default="scope">
                    <el-text v-if="scope.row.type==1">微信</el-text>
                    <el-text v-else-if="scope.row.type==2">支付宝</el-text>
                    <el-text v-else-if="scope.row.type==3">银行卡</el-text>
				</template>
			</el-table-column>			
			<el-table-column align="center" prop="money" label="金额" width="120" />
			<el-table-column align="center" prop="fee" label="手续费" width="120" />
			<el-table-column align="center" prop="actual" label="实际到账" width="120" />
            <el-table-column align="center" prop="account" label="账号" width="120" />
			<el-table-column align="center" prop="realname" label="姓名" width="120" />
			<el-table-column align="center" prop="bank" label="银行" width="120" />
			<el-table-column align="center" label="收款码" width="120">
                <template #default="scope">
					<el-image :src="scope.row.qrcode" fit="contain" 
					style="width: 59px;height: 59px;" 
					:preview-src-list="[scope.row.qrcode]" :preview-teleported="true">
						<template #error>
							<el-icon><Picture /></el-icon>
						</template>
					</el-image>
				</template>
			</el-table-column>
			<el-table-column label="状态" width="120" align="center">
				<template #default="scope">
                    <el-text type="success" v-if="scope.row.status==1">通过</el-text>
                    <el-text type="warning" v-else-if="scope.row.status==0">待审核</el-text>
                    <el-text type="danger" v-else-if="scope.row.status==-1">拒绝</el-text>
                    <el-text type="primary" v-else-if="scope.row.status==2">微信转账待确认</el-text>
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
                    <template v-if="scope.row.status==0 || scope.row.status==2">
					<el-button type="primary" size="small" @click="check(scope,1)">已打款</el-button>
					<el-button type="warning" size="small" @click="check(scope,2)">商家转账</el-button>
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
</template>

<script setup>
	import {ref,onMounted,watch,inject} from "vue";
	import { ElMessageBox } from 'element-plus'
	const myApi=inject('myApi');
	const ElMessage=inject('ElMessage');
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
		myApi('CashList',searchForm.value).then((res)=>{
			list.value=res.data.data
			total.value=res.data.total
		})
	}
	watch([page,limit],()=>{
		getList();
	});
	function check(scope,status){
        switch(status){
            case 1:
                ElMessageBox.confirm('确定已打款吗？', '提示', {
                    type: 'warning',
                }).then(()=>{
                    myApi('CashCheck',{id:scope.row.id,status:1},'post').then(()=>{
                        ElMessage.success('操作成功')
                        getList()
                    })
                })
                break;
            case 2:
                ElMessageBox.confirm('确定使用商家转账接口吗？', '提示', {
                    type: 'warning',
                }).then(()=>{
                    myApi('CashCheck',{id:scope.row.id,status:2},'post').then(()=>{
                        ElMessage.success('操作成功')
                        getList()
                    })
                })
                break;
            case -1:
                ElMessageBox.confirm('', '', {
					showInput: true,
					inputPlaceholder: '拒绝理由',
					inputType: 'textarea',
                    type: 'error',
                }).then(({value})=>{
					if(!value){
						ElMessage.error('请输入拒绝理由')
						return
					}
                    myApi('CashCheck',{id:scope.row.id,status:-1,remark:value},'post').then(()=>{
                        ElMessage.success('操作成功')
                        getList()
                    })
				})
                break;
        }

	}
</script>

<style>
	
</style>