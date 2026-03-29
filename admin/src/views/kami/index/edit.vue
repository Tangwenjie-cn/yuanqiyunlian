
<template>
	<el-form :model="formModel" label-width="80px" ref="refForm">
        <el-form-item label="价格" prop="price" required>
			<el-input-number v-model="formModel.price" :precision="2">
                <template #prefix>
                    <el-text>￥</el-text>
                </template>
            </el-input-number>
		</el-form-item>
        <el-form-item label="卡密" prop="key" required>
			<el-input v-model="formModel.key"></el-input>
		</el-form-item>
        <el-form-item label="状态" prop="status" required>
            <el-radio v-model="formModel.status" label="已使用" :value="1"></el-radio>
            <el-radio v-model="formModel.status" label="未使用" :value="-1"></el-radio>				
        </el-form-item>
		<el-form-item >
			<el-button type="primary" @click="save">提交</el-button>
		</el-form-item>	
	</el-form>
</template>

<script setup>
import {ref,inject} from "vue";
const myApi=inject('myApi');
const ElMessage=inject('ElMessage');
const emit = defineEmits(['submit']);
const refForm=ref()	
const formModel=defineModel()
	
function save(){	 
	refForm.value.validate(valid=>{
		if(valid){
			myApi('KamiEdit',formModel.value,'post').then((res)=>{
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