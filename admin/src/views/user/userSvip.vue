<template>
    <el-space direction="vertical" fill wrap style="width: 100%;" v-if="info">
        <el-card header="会员信息">
            <el-table :data="info.userSvip" style="width: 100%">
                <el-table-column prop="name" label="名称" />
                <el-table-column label="到期时间">
                    <template #default="scope">
                        <el-text>{{ new Date(scope.row.expire_time).toLocaleString() }}</el-text>
                    </template>
                </el-table-column>
            </el-table>
        </el-card>
        <el-card header="会员设置">
            <el-form>
                <el-form-item label="会员选择">
                    <el-select v-model="formData.vid" clearable>
                        <el-option v-for="item in info.vip" :label="item.title" :value="item.id" />
                    </el-select>
                </el-form-item>
                <el-form-item label="到期时间">
                    <el-date-picker
                        v-model="formData.expire_time"
                        type="datetime"
                        placeholder="选择会员到期时间"
                    />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="onSubmit()">提交</el-button>
                </el-form-item>
            </el-form>
        </el-card>
    </el-space>
</template>
<script setup>
import {ref,onMounted,inject,watch} from "vue";
const myApi=inject('myApi');
const props = defineProps(['id','onShow']);
const info=ref()
const formData=ref({
    uid:props.id,
})
const ElMessage=inject('ElMessage');
onMounted(()=>{
    onShow()
})
watch(props,()=>{
    if(props.onShow){
        onShow()
    }
})
function onShow(){
    myApi('GetUserVip',{id:props.id}).then(res=>{
        info.value=res.data
    })
}
function onSubmit(){
    myApi('SetUserVip',formData.value,'post').then(res=>{
        ElMessage.success('设置成功')
        onShow()
    })
}
</script>