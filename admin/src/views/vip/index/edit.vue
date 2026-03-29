<template>
	<el-form :model="formModel" label-width="80px" ref="refForm">
		<el-form-item label="会员名称" prop="title" required>
			<el-input v-model="formModel.title"></el-input>
		</el-form-item>
        <el-form-item label="会员价格" prop="price" required>
			<el-input-number v-model="formModel.price" :precision="2">
				<template #prefix>
					<el-text>￥</el-text>
				</template>
			</el-input-number>
		</el-form-item>
        <el-form-item label="时长" prop="days" required>
			<el-input-number v-model="formModel.days">
			</el-input-number>
			<el-text type="info" style="margin-left: 10px;">会员时长（天）</el-text>
		</el-form-item>
        <el-form-item label="折扣" prop="rebate" required>
			<el-input-number v-model="formModel.rebate" :precision="2">
			</el-input-number>
            <el-text type="info" style="margin-left: 10px;">购买商品折扣，不打折填0,如:九五折,填写9.5</el-text>
		</el-form-item>
		<el-form-item label="积分" prop="points" required>
			<el-input-number v-model="formModel.points">
			</el-input-number>
            <el-text type="info" style="margin-left: 10px;">0不允许积分兑换</el-text>
		</el-form-item>
        <el-form-item label="状态" prop="status" required>
            <el-radio v-model="formModel.status" label="启用" :value="1"></el-radio>
            <el-radio v-model="formModel.status" label="禁用" :value="-1"></el-radio>				
        </el-form-item>
		<el-form-item label="排序" prop="sort" required>
			<el-input-number v-model="formModel.sort">
			</el-input-number>
		</el-form-item>
		<el-form-item label="分销">
			<el-switch
				v-model="formModel.retail_on"
				inline-prompt
				active-text="开启"
				inactive-text="关闭"
				:active-value="1"
				:inactive-value="-1"
			/>
			<el-text type="info" style="margin-left: 10px;">开启后，购买此会员将参与分销</el-text>				
		</el-form-item>
		
		<el-form-item label="佣金设置" v-if="formModel.retail_on==1">
			<el-radio-group v-model="formModel.retail_set">
				<el-radio :value="1">公用设置</el-radio>
				<el-radio :value="2">单独设置</el-radio>
			</el-radio-group>			
		</el-form-item>
		<el-form-item label="分配方式" v-if="formModel.retail_on==1">
			<el-radio-group v-model="formModel.retail_type">
				<el-radio :value="1">固定金额</el-radio>
				<el-radio :value="2">百分比</el-radio>
			</el-radio-group>			
		</el-form-item>
		<el-form-item label="佣金" v-if="formModel.retail_on==1">
			<el-input-number v-model="formModel.retail_yj"></el-input-number>
			<el-text type="info" style="margin-left: 10px;">百分比填写百分比值，如5%填5</el-text>
		</el-form-item>
		<el-form-item label="描述" prop="desc">
			<el-input type="textarea" v-model="formModel.desc"></el-input>
		</el-form-item>
		<el-form-item label="特权" prop="content" required>
            <el-cascader v-model="formModel.content" :options="privilegeList"  :props="cascaderProps" placeholder="请选择特权">
            </el-cascader>
		</el-form-item>
		<el-form-item >
			<el-button type="primary" @click="save">提交</el-button>
		</el-form-item>	
	</el-form>
</template>

<script setup>
import {ref,onMounted,inject} from "vue";
const myApi=inject('myApi');
const ElMessage=inject('ElMessage');
const emit = defineEmits(['submit']);
const privilegeList=ref()
const refForm=ref()	
const formModel=defineModel()
const cascaderProps=ref({
    multiple: true,
    checkStrictly: true,
    emitPath: false,
    value: 'id',
    label: 'name',
})
onMounted(()=>{
	myApi('PrivilegeList',{all:true},'post').then((res)=>{
		privilegeList.value=res.data
	})
})	
function save(){	 
	refForm.value.validate(valid=>{
		if(valid){
			myApi('EditVip',formModel.value,'post').then((res)=>{
				ElMessage({message:res.msg,type:'success'})
				refForm.value.resetFields()
				emit('submit')
			})
		}
	})	
}	
</script>

<style>
</style>