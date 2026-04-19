
<template>
	<el-divider>统计</el-divider>
	<el-row :gutter="30">		
		<el-col :span="4">
			<el-card>
				<template #default>
					<div class="card-block">
					<el-text type="primary" class="icon-size"><el-icon><Avatar /></el-icon></el-text>
					<el-text>用户：{{ statistics.userTotal }}</el-text>
					</div>									
				</template>
			</el-card>
		</el-col>
		<el-col :span="4">
			<el-card>
				<template #default>
					<div class="card-block">
					<el-text type="danger" class="icon-size"><el-icon><TrendCharts /></el-icon></el-text>
					<el-text>商品销售：￥{{ statistics.goodsSalesAmountTotal }}</el-text>
					</div>									
				</template>
			</el-card>
		</el-col>
		<el-col :span="4">
			<el-card>
				<template #default>
					<div class="card-block">
					<el-text type="success" class="icon-size"><el-icon><Ticket /></el-icon></el-text>
					<el-text>会员销售：￥{{ statistics.vipSalesAmountTotal }}</el-text>
					</div>									
				</template>
			</el-card>
		</el-col>
		<el-col :span="4">
			<el-card>
				<template #default>
					<div class="card-block">
					<el-text type="warning" class="icon-size"><el-icon><Medal /></el-icon></el-text>
					<el-text>会员：{{ statistics.vipTotal }}</el-text>
					</div>									
				</template>
			</el-card>
		</el-col>
		<el-col :span="4">
			<el-card>
				<template #default>
					<div class="card-block">
					<el-text type="warning" class="icon-size"><el-icon><Avatar /></el-icon></el-text>
					<el-text>高级推广：{{ statistics.superTotal }}</el-text>
					</div>									
				</template>
			</el-card>
		</el-col>
	</el-row>
	<div style="margin: 20px;display: flex;"></div>
	<el-divider>功能与走势</el-divider>
	<el-row>
		<el-col :span="14">
			<div id="echarts" style="width: 100%;height: 450px;"></div>
		</el-col>
		<el-col :span="10">
			<el-card header="信息提示">
				<el-alert :title="'有'+pending.cash+'个提现申请待处理'" v-if="pending.cash>0" type="warning" show-icon />
				<el-alert :title="'有'+pending.super+'个高级推广申请待处理'" v-if="pending.super>0" type="warning" show-icon />	
			</el-card>
			<el-collapse v-if="pending.update" style="margin-top: 10px;">
				<el-collapse-item title="有新的版本更新">
					<el-descriptions column="2">
						<el-descriptions-item label="当前版本">{{ pending.update.version }}</el-descriptions-item>
						<el-descriptions-item label="最新版本">{{ pending.update.tag_name}}</el-descriptions-item>
						<el-descriptions-item label="版本">{{ pending.update.prerelease?'预览版':'正式版' }}</el-descriptions-item>
						<el-descriptions-item label="更新内容">{{ pending.update.body }}</el-descriptions-item>
						<el-descriptions-item label="下载地址">
							<a href="https://gitee.com/chengdu-yuanqi/yuanqi-yunlian/releases" target="_blank" rel="noopener noreferrer">点击前往</a>
						</el-descriptions-item>
					</el-descriptions>
				</el-collapse-item>
			</el-collapse>
		</el-col>		
	</el-row>
	
</template>

<script setup>
	import {onBeforeMount,inject, ref} from 'vue';
	import * as echarts from 'echarts';
	const myApi=inject('myApi');
	const statistics=ref({})
	const pending=ref({})
	onBeforeMount(()=>{
		myApi('home').then((res)=>{
			setEcharts(res.data.lineChart)
			statistics.value=res.data.statistics
			pending.value=res.data.pending
		})
	})
	function setEcharts(data){
	    const myChart = echarts.init(document.getElementById('echarts'));
		myChart.setOption({
			legend: {
				type: 'scroll',
			},
			tooltip: {
				trigger: 'axis',
				axisPointer: {
					type: 'cross'
				}
			},
			xAxis: {
				type:'category',
				name:'日期',
				data: data.dayList
			},
			yAxis: {
				type:'value',
				name:'金额(元)'
			},
			series: [
				{
				data: data.goodsSalesAmount,
				type: 'bar',
				name: '商品销量'
				},
				{
				data: data.vipSalesAmount,
				type: 'bar',
				name: '会员销量'
				}
			]
		});
	}	
</script>

<style>
.el-alert {
  margin: 20px 0 0;
}
.card-block{
	display: flex;
	flex-direction: column;
}
.icon-size{
	font-size: 32px;
}
</style>
