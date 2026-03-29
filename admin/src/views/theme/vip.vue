<template>
    <!--搜索栏-->
	<div class="page-header">
		<el-form :inline="true" :model="searchForm">
		</el-form>
		<el-button type="primary" icon="refresh" @click="refresh">刷新</el-button>
	</div> 
<div class="container">
    <div class="diy-block">

        <div class="layout-container">
            <!-- 状态栏 -->
            <div class="block">
                <div class="diy-block-status" :style="{'background-color': theme.backgroundColor}">
                <div class="diy-block-status-left">
                    <span>8:30</span>
                </div>
                <div class="diy-block-status-right">
                    <span class="iconfont icon-xinhao2 icon-status"></span>
                    <span class="iconfont icon-dianchi-man icon-status"></span>
                </div>
                </div>
                <div class="diy-block-title" :style="{'background-color': theme.backgroundColor}">
                <span :style="{'color': theme.titleColor}">{{ theme.title }}</span>
                </div>
                <div class="user-block" :style="{'background-color': theme.backgroundColor}">
                    <img src="/src/assets/images/avatar.jpeg" class="user-avatar" />
                    <div class="column" style="margin-left: 10px;">
                        <div>
                            <span class="user-text" :style="{'color': theme.color}">昵称</span>
                            <span class="iconfont icon-f-vip" :style="{'color': theme.iconColor}"></span>
                        </div>				
                    </div>
                </div>
            </div>
        </div>
        
    </div>
    <div class="component-set">
        <el-form>
            <el-form-item label="背景色">
                <el-color-picker v-model="theme.backgroundColor" />
                <el-input style="width: auto;margin-left: 5px;" v-model="theme.backgroundColor" />
            </el-form-item>
            <el-form-item label="标题">
                <el-input style="width: auto;margin-left: 5px;" v-model="theme.title" />
            </el-form-item>
            <el-form-item label="标题文字">
                <el-radio-group v-model="theme.titleColor">
                    <el-radio value="#ffffff">白色</el-radio>
                    <el-radio value="#000000">黑色</el-radio>
                </el-radio-group>
            </el-form-item>
            <el-form-item label="文字颜色">
                <el-color-picker v-model="theme.color" />
                <el-input style="width: auto;margin-left: 5px;" v-model="theme.color" />
            </el-form-item>
            <el-form-item label="图标颜色">
                <el-color-picker v-model="theme.iconColor" />
                <el-input style="width: auto;margin-left: 5px;" v-model="theme.iconColor" />
            </el-form-item>
            <el-form-item label="">
                <el-button type="primary" @click="save">保存</el-button>
            </el-form-item>
        </el-form>  
    </div>
</div>  
</template>

<script setup>
import { ref,inject, onMounted } from 'vue'
const myApi=inject('myApi');
const ElMessage=inject('ElMessage');
const theme=ref({
    backgroundColor: '#f56b59',
    title: '会员中心',
    titleColor: '#ffffff',
    color: '#ffffff'
})
const themeId=ref(0)
onMounted(()=>{
    getList()
})
function getList(){
    myApi('ThemeGet',{type:3}).then(res=>{
        themeId.value=res.data.id
        theme.value=res.data.content[0]
    })
}
function refresh(){
	getList()
}
function save(){
    myApi('ThemeEdit',{id:themeId.value,title:'会员中心',content:[theme.value]},'post').then(()=>{       
        ElMessage.success('成功')
    })
}
</script>

<style scoped>
.layout-container{
    height: 480px;
    overflow: auto;
    background-color: #eaeaea;
}
.container {
    display: flex;
}
.block{
    position: relative;
    margin-bottom: 5px;
}
.diy-block {
    width: 360px;
    height: 550px;
}
.diy-block-status {
    display: flex;
    height: 20px;
    width: 100%; 
    justify-content: space-between;
    align-items: center;
    border-top-left-radius: 5px;
    border-top-right-radius: 5px;
}
.diy-block-title{
    display: flex;
    justify-content: center;
    align-items: center;
    height: 40px;
}
.diy-block-status-left{
    margin-left: 2px;
}
.icon-status{
    margin-right: 5px;
}
.user-block{
    border-bottom-left-radius: 10px;
    border-bottom-right-radius: 10px;
    padding: 10px;
    display: flex;
    align-items: center;
}
.vip-text{
    font-size: 12px;
}
.user-text{
    font-size: 16px;
}
.user-avatar{
    width: 60px;
    height: 60px;
    border-radius: 50%;
    margin-left: 20px;
    background-color: #ffffff;
}
.icon-f-vip{
    color: #e8cb38;
    font-size: 18px;
    margin-left: 10px;
}
.component-set{
    margin-left: 20px;
}
</style>