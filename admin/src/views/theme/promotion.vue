
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
            <img :src="theme.imgUrl" width="300" height="400"/>
            <div style="display: flex;position: absolute;top: 10px;left: 50px;">
                <img src="/src/assets/images/avatar.jpeg" width="40" height="40" style="border-radius: 50%;"/>
                <el-text style="margin-left: 10px;">用户001</el-text>
            </div>
            <span class="tg-title" :style="{'color': theme.color, 'font-size': theme.fontSize}">{{ theme.title }}</span>
            <img src="/src/assets/images/qrcode.jpg" width="200" height="200" style="position: absolute;bottom: 30px;">
        </div>        
    </div>
    <div class="component-set">
        <el-form>
            <el-form-item label="推广语">
                <el-input style="width: auto;margin-left: 5px;" v-model="theme.title" />
            </el-form-item>
            <el-form-item label="文字颜色">
                <el-color-picker v-model="theme.color" />
                <el-input style="width: auto;margin-left: 5px;" v-model="theme.color" />
            </el-form-item>
            <el-form-item label="文字大小">
                <el-input style="width: auto;margin-left: 5px;" v-model="theme.fontSize" />
            </el-form-item>
            <el-form-item label="推广图">
                <uploadImg v-model="theme.imgUrl"></uploadImg>
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
import uploadImg from "../../components/upload/uploadImg.vue"
const myApi=inject('myApi');
const ElMessage=inject('ElMessage');
const theme=ref({
    title: '这是推广语',
    color: '#ffffff',
    fontSize: '20px',
    imgUrl: ''
})
const themeId=ref(0)
onMounted(()=>{
    getList()
})
function getList(){
    myApi('ThemeGet',{type:6}).then(res=>{
        themeId.value=res.data.id
        theme.value=res.data.content[0]
    })
}
function refresh(){
	getList()
}
function save(){
    myApi('ThemeEdit',{id:themeId.value,title:'推广图',content:[theme.value]},'post').then(()=>{       
        ElMessage.success('成功')
    })
}
</script>

<style scoped>
.layout-container{
    display: flex;
    flex-direction: column;
    align-items: center;
    position: relative;
    width: 300px;
    height: 400px;
    background-color: #f0f0f0;
}
.tg-title{
    position: absolute;
    top: 60px; 
    padding-left: 20px;
    padding-right: 10px;  
}
.container {
    display: flex;
}
.component-set{
    margin-left: 20px;
}
</style>