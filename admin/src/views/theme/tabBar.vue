
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
            <div class="block" :style="{backgroundColor: theme.backgroundColor,color: theme.color}">
                <div class="tabbar-item" v-for="item in theme.list">
                    <img :src="item.imgUrl" class="tabbar-img">
                    <span class="tabbar-text">{{ item.name }}</span>
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
            <el-form-item label="文字颜色">
                <el-color-picker v-model="theme.color" />
                <el-input style="width: auto;margin-left: 5px;" v-model="theme.color" />
            </el-form-item>
            <el-form-item label="选中颜色">
                <el-color-picker v-model="theme.selectColor" />
                <el-input style="width: auto;margin-left: 5px;" v-model="theme.selectColor" />
            </el-form-item>
            <el-form-item label="导航内容">
                <el-button @click="addData()"><el-icon><Plus /></el-icon>增加</el-button>
            </el-form-item>
            <div style="display: flex;flex-direction: column;margin-bottom: 10px;">                    
                <div style="display: flex;flex-direction: column;margin-top: 10px;" v-for="(item,index) in theme.list">
                    <div style="display: flex;justify-content: space-around;">
                        <div style="display: flex;flex-direction: column;align-items: center;">
                            <uploadImg v-model="item.imgUrl"></uploadImg>
                            <span>图片</span>
                        </div>
                        <div style="display: flex;flex-direction: column;align-items: center;">
                            <uploadImg v-model="item.selectImgUrl"></uploadImg>
                            <span>选中图片</span>
                        </div>
                    </div>                                     
                    <div style="display: flex;flex-direction: column;margin-left: 5px;justify-content: space-around;">
                        <div style="margin-bottom: 5px;">名称 <el-input style="width: auto;" v-model="item.name"></el-input></div>
                        <div style="margin-bottom: 5px;">链接 <el-input style="width: auto;" v-model="item.link"></el-input></div>
                    </div>
                    <el-button type="danger" size="small" @click="theme.list.splice(index,1)">删除</el-button>
                </div>
            </div>
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
    backgroundColor: '#ffffff',
    color: '#3d3d3d',
    selectColor: '#f56b59',
    list:[]
})
const themeId=ref(0)
onMounted(()=>{
    getList()
})
function addData(){
    let data = {
        imgUrl:'',
        name:'',
        link:'',
        selectImgUrl:''
    }
    theme.value.list.push(data)
}
function getList(){
    myApi('ThemeGet',{type:5}).then(res=>{
        themeId.value=res.data.id
        theme.value=res.data.content[0]
    })
}
function refresh(){
	getList()
}
function save(){
    myApi('ThemeEdit',{id:themeId.value,title:'导航',content:[theme.value]},'post').then(()=>{       
        ElMessage.success('成功')
    })
}
</script>

<style scoped>
.layout-container{
    height: 480px;
    overflow: auto;
    background-color: #eaeaea;
    position: relative;
    border: 1px solid #000000;
}
.container {
    display: flex;
}
.block{
    position: absolute;
    display: flex;
    align-items: center;
    justify-content: space-around;
    bottom: 0;
    left: 0;
    right: 0;
    padding: 5px 0;
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
.tabbar-item{
    display: flex;
    flex-direction: column;
}
.tabbar-img{
    width: 32px;
    height: 32px;
}
.tabbar-text{
    text-align: center;
    font-size: 12px;
}
.component-set{
    margin-left: 20px;
}
</style>