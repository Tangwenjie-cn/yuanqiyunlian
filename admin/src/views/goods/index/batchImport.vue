<template>
<el-form :model="formModel" label-width="80px" ref="refForm">
    <el-form-item label="分类" prop="sort_ids" required>
        <el-cascader v-model="formModel.sort_ids" :options="sort"  :props="cascaderProps">
        </el-cascader>
    </el-form-item>
    <el-form-item label="栏目" prop="cid">
        <el-select v-model="formModel.cid" style="width: 48%;" clearable placeholder="请选择">
            <el-option v-for="item in columns" :label="item.name" :value="item.id">
            </el-option>
        </el-select>
    </el-form-item>
    <el-form-item label="价格" prop="price" required>
        <el-input v-model="formModel.price" style="width: 48%;"></el-input>
    </el-form-item>
    <el-form-item label="状态" prop="status" required>
        <el-radio v-model="formModel.status" label="上架" :value="1"></el-radio>
        <el-radio v-model="formModel.status" label="下架" :value="-1"></el-radio>				
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
    <el-form-item label="置顶">
        <el-switch
            v-model="formModel.is_top"
            inline-prompt
            active-text="是"
            inactive-text="否"
            :active-value="1"
            :inactive-value="-1"
        />			
    </el-form-item>
    <el-form-item>
        <el-alert title="文档转换异步执行，上传完成后可关闭窗口" type="warning" />
    </el-form-item>
    
    <el-form-item label="选择文档">
        <el-upload
            ref="uploadRef"
            :action="domain()+paths.UploadDocument"
            :data="formData" 
            :headers="headers"
            multiple           
            accept=".xls,.xlsx,.doc,.docx,.ppt,.pptx,.pdf,.odt,.ods,.odp,.odg"
            :auto-upload="false"
            :on-error="onError"
            >
            <template #trigger>
                <el-button>选择文件</el-button>
            </template>
        </el-upload>
    </el-form-item>
    <el-form-item>
        <el-button type="primary" @click="onSubmit">开始上传</el-button>
    </el-form-item>
</el-form>
</template>
<script setup>
import {ref,onMounted,inject} from "vue";
import {domain,paths} from "../../../config/api.js"
const myApi=inject('myApi');
const store = inject('store');
const ElMessage=inject('ElMessage');
const sort=ref()
const refForm=ref()	
const uploadRef=ref()
const columns=ref([{}])
const headers=ref({
		'Authorization':store.token
	})
const formModel=ref({sort_ids:[]})
const formData=ref({})
const cascaderProps=ref({
    multiple: true,
    checkStrictly: true,
    emitPath: false,
    value: 'id',
    label: 'name',
})
onMounted(()=>{
    myApi('goodsSort',{},'post').then((res)=>{
		sort.value=res.data
	})
    myApi('GoodsColumn').then((res)=>{
		columns.value=res.data
	})
})
function onError(err){
    if(err){
        let resp=JSON.parse(err.message)
        ElMessage({message:resp.msg,type:'error'})
        console.log(err.message)
    }
}
function onSubmit(){
    formData.value={ ...formModel.value }
    formData.value.sort_ids = [...formModel.value.sort_ids].join()
    refForm.value.validate((valid)=>{
        if(valid){
            uploadRef.value.submit()
        }
    })
}
</script>