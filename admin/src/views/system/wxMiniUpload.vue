<template>
    <div v-loading="loading">
        <el-form :model="formData" ref="refForm" label-width="auto" style="max-width: 500px">
            <el-form-item label="版本号" prop="version" required>
                <el-input v-model="formData.version" />
            </el-form-item>
            <el-form-item label="简介" prop="desc" required>
                <el-input v-model="formData.desc" type="textarea" />
            </el-form-item>
            <el-form-item> 
                <el-card>
                    <el-text type="warning">上传成功后请前往小程序后台，设置为预览版测试使用，确认无问题再提交审核。</el-text>
                </el-card>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm">提交</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>
<script setup>
    import { ref,useTemplateRef, inject} from 'vue'
    const myApi=inject('myApi');
    const ElMessage=inject('ElMessage');
    const formData = ref({})
    const refForm = useTemplateRef('refForm')
    const loading = ref(false)
    function submitForm(){
        refForm.value.validate((valid)=>{
            if(valid){
                loading.value = true
                myApi('WxMiniUpload',formData.value,'post').then(res=>{
                    loading.value = false
                    ElMessage({message: '上传成功，请登录微信小程序后台查看',type:'success'})
                }).catch(err=>{
                    loading.value = false
                })
            }
        })
    }
</script>