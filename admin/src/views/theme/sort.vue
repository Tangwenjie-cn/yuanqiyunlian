
<template>
    <!--搜索栏-->
	<div class="page-header">
		<el-form :inline="true" :model="searchForm">
		</el-form>
		<el-button type="primary" icon="refresh" @click="refresh">刷新</el-button>
	</div> 
<div class="container">
    <el-radio-group v-model="theme.type">        
        <el-card>
            <img src="/src/assets/images/QQ20260108-155805.png" style="width: 300px; height: 450px" />
            <template #footer><el-radio :value="1">样式1</el-radio></template>
        </el-card>
        <el-card style="margin-left: 20px;">
            <img src="/src/assets/images/QQ20260108-171725.png" style="width: 300px; height: 450px" />
            <template #footer><el-radio :value="2">样式2</el-radio></template>
        </el-card>
    </el-radio-group>
    <div style="margin-top: 20px;width: 100%;">
        <el-button type="primary" @click="save">保存</el-button>
    </div>
    
</div>  
</template>

<script setup>
import { ref,inject, onMounted } from 'vue'
const myApi=inject('myApi');
const ElMessage=inject('ElMessage');
const theme=ref({
    type: 1,
})
const themeId=ref(0)
onMounted(()=>{
    getList()
})
function getList(){
    myApi('ThemeGet',{type:2}).then(res=>{
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

.container {
    display: flex;
    width: 100%;
    flex-wrap: wrap;
}

</style>