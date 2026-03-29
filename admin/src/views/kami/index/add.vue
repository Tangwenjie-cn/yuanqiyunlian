<template>
	<el-form :model="formModel" label-width="80px" ref="refForm">
        <el-form-item label="价格" prop="price" required>
            <el-input-number v-model="formModel.price" :precision="2">
                <template #prefix>
                    <el-text>￥</el-text>
                </template>
            </el-input-number>			
		</el-form-item>
        <el-form-item label="生成数量" prop="nums" required>
            <el-input-number v-model="formModel.nums" />
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
const formModel=ref({});
	
function save(){	 
	refForm.value.validate(valid=>{
		if(valid){

			myApi('KamiAdd',formModel.value,'post').then((res)=>{
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