
<template>
    <el-space direction="vertical" fill wrap style="width: 100%;" v-if="info">
        <el-descriptions title="基本信息">
            <el-descriptions-item label="UID">{{ info.user.id }}</el-descriptions-item>
            <el-descriptions-item label="昵称">{{ info.user.nickname }}</el-descriptions-item>
            <el-descriptions-item label="头像"><el-avatar :size="26" :src="info.user.avatar" /></el-descriptions-item>
            <el-descriptions-item label="余额">{{ info.user.balance }}</el-descriptions-item>
            <el-descriptions-item label="推广人数">{{ info.user.promotion }}</el-descriptions-item>
            <el-descriptions-item label="积分">{{ info.user.points }}</el-descriptions-item>
            <el-descriptions-item label="高级推广">
                <el-text type="success" v-if="info.user.is_super==1">是</el-text>
                <el-text v-else>否</el-text>
            </el-descriptions-item>
            <el-descriptions-item label="高推设置">
                <el-text v-if="info.user.super_set==1" type="primary">公共</el-text>
                <el-text v-else-if="info.user.super_set==2" type="primary">独有</el-text>
                <el-text v-else>无</el-text>
            </el-descriptions-item>
            <el-descriptions-item label="高推佣金方式">
                <el-text v-if="info.user.super_type==1" type="primary">固定</el-text>
                <el-text v-else-if="info.user.super_type==2" type="primary">百分比</el-text>
                <el-text v-else>无</el-text>
            </el-descriptions-item>
            <el-descriptions-item label="高推佣金">
                <el-text v-if="info.user.super_type==1">{{ info.user.super_yj }}</el-text>
                <el-text v-else-if="info.user.super_type==2">{{ info.user.super_yj }}%</el-text>
            </el-descriptions-item>
            <el-descriptions-item label="状态">
                <el-text type="success" v-if="info.user.status==1">正常</el-text>
                <el-text v-else type="danger">禁用</el-text>
                <div style="margin-left: 5px;display: inline-block;">
                    <el-button type="danger" size="small" v-if="info.user.status==1" @click="setStatus(-1)">禁用</el-button>
                    <el-button type="success" size="small" v-else @click="setStatus(1)">启用</el-button>
                </div>
                
            </el-descriptions-item>
            <el-descriptions-item label="注册时间">{{ new Date(info.user.ctime).toLocaleString() }}</el-descriptions-item>
        </el-descriptions>
        <el-card header="绑定登录">
            <el-space>
                <template v-for="item in info.bind">
                    <el-tag type="primary" v-if="item.type==1">手机：{{ item.phone }}</el-tag>
                    <el-tag type="primary" v-if="item.type==2">微信小程序</el-tag>
                </template>
            </el-space> 
        </el-card>
        <el-card header="会员信息">
            <el-table :data="info.vip" style="width: 100%">
                <el-table-column prop="name" label="名称" />
                <el-table-column label="到期时间">
                    <template #default="scope">
                        <el-text>{{ new Date(scope.row.expire_time).toLocaleString() }}</el-text>
                    </template>
                </el-table-column>
            </el-table>
        </el-card>
    </el-space>
    
</template>
<script setup>
import {ref,onMounted,inject,watch} from "vue";
const myApi=inject('myApi');
const props = defineProps(['id','onShow']);
const info=ref()
onMounted(()=>{
    onShow()
})
watch(props,()=>{
    if(props.onShow){
        onShow()
    }
})
function onShow(){
    myApi('UserInfo',{id:props.id}).then(res=>{
        info.value=res.data
    })
}
function setStatus(s){
    myApi('SetUserInfo',{id:props.id,status:s},'post').then(()=>{
        onShow()
    })
}
</script>