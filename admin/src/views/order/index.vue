
<template>
	<!--搜索栏-->
	<div class="page-header">
		<el-form :inline="true" :model="searchForm">
			<el-form-item class="page-search-form-item page-search-select">
				<el-select v-model="searchForm.status" clearable placeholder="状态">
					<el-option :value="1" label="已支付"></el-option>
                    <el-option :value="2" label="待支付"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item class="page-search-form-item page-search-select">
				<el-select v-model="searchForm.type" clearable placeholder="订单类型">
					<el-option :value="1" label="商品订单"></el-option>
                    <el-option :value="2" label="VIP订单"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item class="page-search-form-item page-search-select">
				<el-select v-model="searchForm.pay_type" clearable placeholder="支付类型">
					<el-option :value="1" label="微信"></el-option>
                    <el-option :value="2" label="支付宝"></el-option>
                    <el-option :value="3" label="卡密"></el-option>
                    <el-option :value="4" label="余额"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item class="page-search-form-item">
				<el-input v-model="searchForm.keyword" clearable placeholder="UID|订单号" />
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
			<el-table-column align="center" prop="order_no" label="订单号" width="180" />
			<el-table-column align="center" prop="third_order" label="三方订单号" width="180" />	
			<el-table-column align="center" prop="title" label="商品名称" width="180" />
			<el-table-column align="center" label="类型" width="120">
				<template #default="scope">
                    <el-text v-if="scope.row.type==1">商品订单</el-text>
                    <el-text v-else-if="scope.row.type==2">VIP订单</el-text>
				</template>
			</el-table-column>
			<el-table-column align="center" label="支付类型" width="120">
				<template #default="scope">
                    <el-text v-if="scope.row.pay_type==1">微信</el-text>
                    <el-text v-else-if="scope.row.pay_type==2">支付宝</el-text>
					<el-text v-else-if="scope.row.pay_type==3">卡密</el-text>
					<el-text v-else-if="scope.row.pay_type==4">余额</el-text>
				</template>
			</el-table-column>			
			<el-table-column align="center" prop="price" label="价格" width="120" />
			<el-table-column align="center" prop="discount_price" label="抵扣" width="120" />
			<el-table-column align="center" label="抵扣类型" width="120">
				<template #default="scope">
                    <el-text v-if="scope.row.discount_type==1">优惠券</el-text>
                    <el-text v-else-if="scope.row.discount_type==2">会员优惠</el-text>
				</template>
			</el-table-column>
			<el-table-column align="center" prop="pay_price" label="支付金额" width="120" />
			<el-table-column align="center" prop="retail_price" label="上级分润" width="120" />
			<el-table-column align="center" prop="super_price" label="高级分润" width="120" />
			<el-table-column label="状态" width="120" align="center">
				<template #default="scope">
                    <el-text type="success" v-if="scope.row.status==1">已支付</el-text>
                    <el-text type="warning" v-else-if="scope.row.status==0">待支付</el-text>
                    <el-text type="danger" v-else-if="scope.row.status==-1">已退款</el-text>
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
		myApi('OrderList',searchForm.value).then((res)=>{
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