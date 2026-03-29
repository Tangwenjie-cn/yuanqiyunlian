<template>
	<el-form :model="formModel" label-width="80px" ref="refForm">
		<el-form-item label="特权名称" prop="name" required>
			<el-input v-model="formModel.name"></el-input>
		</el-form-item>
		<el-form-item label="特权类型" prop="type" required>
			<el-select v-model="formModel.type" placeholder="请选择特权类型">
				<el-option label="商品类型" :value="1" />
				<el-option label="商品分类" :value="2" />
				<el-option label="商品栏目" :value="3" />
			</el-select>
		</el-form-item>
		<el-form-item label="特权" prop="content" required>
			<template v-if="formModel.type==1">
                <el-cascader v-model="formModel.content" :options="goodsType"  :props="cascaderProps">
                </el-cascader>
            </template>
			<template v-else-if="formModel.type==2">
				<el-cascader v-model="formModel.content" :options="sort"  :props="cascaderProps">
                </el-cascader>
			</template>
			<template v-else-if="formModel.type==3">
				<el-cascader v-model="formModel.content" :options="columns"  :props="cascaderProps">
                </el-cascader>
			</template>
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
const sort=ref()
const columns=ref()
const refForm=ref()	
const formModel=defineModel()
const cascaderProps=ref({
    multiple: true,
    checkStrictly: true,
    emitPath: false,
    value: 'id',
    label: 'name',
})
const goodsType=ref([
	{name:'文档',id:1},
	{name:'文章',id:2},
	{name:'视频',id:3},
	{name:'音频',id:4},
	{name:'网盘',id:5},
	{name:'课程',id:6},
])
onMounted(()=>{
	myApi('goodsSort',{},'post').then((res)=>{
		sort.value=res.data
	})
	myApi('GoodsColumn').then((res)=>{
		columns.value=res.data
	})
})	
function save(){	 
	refForm.value.validate(valid=>{
		if(valid){
			myApi('EditPrivilege',formModel.value,'post').then((res)=>{
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