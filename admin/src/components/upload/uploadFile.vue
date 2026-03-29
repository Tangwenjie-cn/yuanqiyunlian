<template>
    <div style="display: flex;" v-loading="loading">
        <el-upload
            class="uploader"
            :action="action"
            :headers="headers"
            :show-file-list="false"
            :on-success="uploadSuccess" :before-upload="beforeUpload" :on-error="onError">
            <img v-if="fileUrl" :src="imageUrl" class="image" />
            <el-icon v-else class="uploader-icon"><Plus /></el-icon>
        </el-upload>
        <el-popconfirm v-if="fileUrl" title="确定删除吗？" @confirm="remove">
            <template #reference>
                <el-text type="danger" style="margin-left: 10px;">删除</el-text>
            </template>
        </el-popconfirm>
    </div>   
    <el-input v-model="fileUrl" style="margin-top: 2px;" /> 
</template>
<script setup>
import {ref,inject} from "vue"
import {domain,paths} from "../../config/api.js"
import imageUrl from "../../assets/images/file.png"
const fileUrl = defineModel()
const action=ref(domain()+paths.uploadFile)
const myApi=inject('myApi');
const store = inject('store');
const ElMessage=inject('ElMessage');
const loading = ref(false)
const headers=ref({
		'Authorization':store.token
	})
function remove(){
    myApi('delFile',{"path":fileUrl.value},'post').then(res=>{
        if(res.code===0){
            fileUrl.value=''
            ElMessage(res.msg)
            return true;
        }else{
            ElMessage(res.msg)
        }					
	})
}
function uploadSuccess(res){
    loading.value=false
    if(res.code===0){
        fileUrl.value=res.data
        return
    }else{
        return ElMessage(res.msg)
    }
    	
}
function beforeUpload(file){
    loading.value=true
}
function onError(){
    loading.value=false
}
</script>
<style scoped>
.uploader .image {
  width: 80px;
  height: 80px;
  display: block;
}
</style>

<style>
.uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 80px;
  height: 80px;
  text-align: center;
}
</style>