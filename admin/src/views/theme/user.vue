
<template> 
<div class="container">
    <div class="component-library">
        <h3>组件库</h3>
        <div>
            <el-button class="lib-btn" type="primary" plain @click="addItem('status')">状态栏</el-button>
        </div>
        <div>
            <el-button class="lib-btn" type="primary" plain @click="addItem('user')">用户信息</el-button>
        </div>
        <div>
            <el-button class="lib-btn" type="primary" plain @click="addItem('operate')">操作栏</el-button>
        </div>
        <div>
            <el-button class="lib-btn" type="primary" plain @click="addItem('image')">图片</el-button>
        </div>
    </div>
    <div class="diy-block">
        <div style="display: flex;align-items: center;height: 35px;">
            <el-button size="small" type="danger" plain @click="removeItem">移除选中组件</el-button>
            <el-button size="small" type="primary" style="margin-left: 5px;" @click="save">保存</el-button>
        </div>
        <div class="layout-container">
            <GridLayout 
            v-model:layout="layout"
            :col-num="1"  
            :margin="[2,2]"  
            :row-height="10">
                <GridItem
                v-for="(item, index) in layout"
                :x="item.x"
                :y="item.y"
                :w="item.w"
                :h="item.h"
                :i="item.i"
                @click="handleItemClick(item, index)">
                <!-- 状态栏 -->
                <div v-if="item.type === 'status'" class="block" :class="{selected: item.i == selectedId,unselected: item.i != selectedId}">
                    <div class="diy-block-status" :style="{'background-color': item.data.backgroundColor}">
                        <div class="diy-block-status-left">
                            <span>8:30</span>
                        </div>
                        <div class="diy-block-status-right">
                            <span class="iconfont icon-xinhao2 icon-status"></span>
                            <span class="iconfont icon-dianchi-man icon-status"></span>
                        </div>
                    </div>
                    <div class="diy-block-title" :style="{'background-color': item.data.backgroundColor}">
                        <span :style="{'color': item.data.color}">{{ item.data.title }}</span>
                    </div>
                </div> 
                    <!-- 操作 -->
                    <div v-if="item.type === 'operate'" class="block" :class="{selected: item.i == selectedId,unselected: item.i != selectedId}">
                        <template v-if="item.data.arrange==1">
                            <div class="one-operate" :style="{'background-color': item.data.backgroundColor}" >
                                <div class="one-operate-item" v-for="item1 in item.data.list">
                                    <div style="display: flex;align-items: center;">
                                        <img style="width: 32px;height: 32px;margin:0 10px;" :src="item1.imgUrl">
                                        <span :style="{'color': item.data.color, 'font-size': item.data.fontSize}">{{ item1.name }}</span>
                                    </div>
                                    <el-icon size="22"><ArrowRightBold /></el-icon>
                                </div>
                            </div>
                        </template>
                        <template v-else>
                            <div class="classify-block"
                            :style="{'background-color': item.data.backgroundColor,'grid-template-columns': `repeat(${item.data.arrange}, 1fr)`}">
                                <div v-for="item1 in item.data.list" class="classify-item">
                                    <img :src="item1.imgUrl">
                                    <span class="classify-name" :style="{'color': item.data.color, 'font-size': item.data.fontSize}">{{ item1.name }}</span>
                                </div>
                            </div>
                        </template>
                        
                    </div>
                    <!-- 用户信息 -->
                    <div v-if="item.type === 'user'" class="block" :class="{selected: item.i == selectedId,unselected: item.i != selectedId}">
                        <div class="user-block" :style="{'background-color': item.data.backgroundColor}">
                            <div class="user-center">
                                <div style="display: flex;align-items: center;">	
                                    <img src="/src/assets/images/avatar.jpeg" class="user-avatar" />
                                    <div style="display: flex;flex-direction: column;">
                                        <span class="user-name" :style="{'color': item.data.color}">用户昵称</span>				
                                        <span class="user-text" :style="{'color': item.data.color}">UID:1000</span>
                                    </div>
                                </div>
                                <div class="user-right">
                                    <el-icon size="26" :color="item.data.color"><ArrowRightBold /></el-icon>
                                </div>
                            </div>
                        </div>
                        <div class="vip-block" :style="{'background': `linear-gradient(to right,${item.data.gradientStartColor},${item.data.gradientEndColor})`}">
                            <div class="vip-info">
                                <span class="iconfont icon-f-vip" :style="{'color': item.data.iconColor}"></span>
                                <span class="vip-text" :style="{'color': item.data.color}">开通会员特权</span>
                                <span :style="{'color': item.data.color}">立即开通</span>
                            </div>
                        </div>
                    </div>
                    <!-- 图片 -->
                    <div v-if="item.type === 'image'" class="block" :class="{selected: item.i == selectedId,unselected: item.i != selectedId}">
                        <img :src="item.data.imgUrl" class="block-img">
                    </div>
                </GridItem>
                
            </GridLayout>
        </div>
        
    </div>
    <div class="component-set">
        <componentSet v-if="isComponentSet" v-model="layout[selectedIndex]"></componentSet>       
    </div>
</div>  
</template>

<script setup>
import { GridLayout, GridItem} from 'grid-layout-plus'
import { ref,inject, onMounted } from 'vue'
import {randomStr} from '../../config/functions.js'
import componentSet from './index/set.vue'
const myApi=inject('myApi');
const ElMessage=inject('ElMessage');
const layout = ref([])
const selectedId = ref(null)
const selectedIndex = ref(null)
const isComponentSet= ref(false)
const themeId=ref(0)
onMounted(()=>{
    myApi('ThemeGet',{type:4}).then(res=>{
        themeId.value=res.data.id
        layout.value=res.data.content
    })
    
})
//点击组件
function handleItemClick(item, index){
    isComponentSet.value = true
    selectedId.value = item.i
    selectedIndex.value = index    
}
//添加组件
function addItem(type) {
    let item={x: 0, y: layout.value.length, w: 1, i: randomStr(6), type: type,data:null}
    let data={}
    switch(type){
        case 'status':
            item.h = 5
            data={
                backgroundColor:'#f56b59',
                color:'#ffffff',
                title:'首页'
            }
            item.data=data
            break
        case 'operate':
            item.h = 17
            data={
                backgroundColor:'#ffffff',
                color:'#000000',
                fontSize:'14px',
                arrange:5,
                list:[]
            }
            item.data=data
            break
        case 'user':
            item.h = 12
            data={
                backgroundColor:'#f56b59',
                color:'#ffffff',
                gradientStartColor:'#eaeaaf',
                gradientEndColor:'#5e5e40',
            }
            item.data=data
            break
        case 'image':
        item.h = 6
        data={
            imgUrl:'/src/assets/images/avatar.jpeg',
            link:''
        }
        item.data=data
        break
        default:
            break
    }
    layout.value.push(item)
}
//删除组件
function removeItem(){
    if(layout.value.length == 1){
        return ElMessage.warning('至少保留一个组件')
    }
    let index = layout.value.findIndex(item => item.i == selectedId.value)  
    selectedId.value = layout.value[0].i
    selectedIndex.value = 0
    if (index > -1) {
        layout.value.splice(index, 1)
    }
}
function save(){
    myApi('ThemeEdit',{id:themeId.value,title:'个人中心',content:[...layout.value].sort((a, b) => a.y - b.y)},'post').then(res=>{       
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
.component-library {
    display: flex;
    flex-direction: column;
    width: 160px;
    background-color: #f5f5f5;
    padding: 10px;
    margin-right: 20px;
    align-items: center;
}
.lib-btn{
    margin-bottom: 10px;
    width: 99px;
    margin-left: 0px;
}
.selected{
    border: 3px solid #3858d6;
}
.unselected{
    border: 1px dashed rgb(119, 119, 119);
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
.classify-block{
    display: grid;
    grid-gap: 2px;
    padding: 6px;
}
.classify-block img{
    height:42px; 
    width:42px;
    border-radius: 10px;
}
.classify-item{
    display: flex;
    flex-direction: column;
    align-items: center;
}
.classify-name{
    font-size: 12px;
    color: #211f1f;
}
.one-operate{
    display: flex;
    flex-direction: column;
}
.one-operate-item{
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid #00000021;
    height: 46px;
}
.block{
    position: relative;
    margin-bottom: 5px;
}
.user-block{
    display: flex;
    width: 100%;
    height: 120px;
    border-bottom-left-radius: 75% 20%;
    border-bottom-right-radius: 75% 20%;
    justify-content: space-between;
}
.user-center{
    display: flex;
    align-items: center;
    width: 100%;
    height: 80px;
    justify-content: space-between;
}
.user-avatar{
    width: 60px;
    height: 60px;
    border-radius: 50%;
    margin-left: 20px;
    background-color: #ffffff;
}
.user-name{
    font-size: 16px;
    margin-left: 10px;
}
.user-text{
    font-size: 12px;
    margin-left: 10px;
}
.user-right{
    margin-right: 20px;
}
.vip-block{
    width: auto;
    height: 50px;
    border-radius: 10px;
    margin-top: -35px;
    margin-left: 45px;
    margin-right: 45px;
}
.vip-info{
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
}
.vip-text{
    font-size: 12px;
    margin: 0 5px;
}
.component-set{
    margin-left: 10px;
}
.block-img{
    width: 100%;
    height: 60px;
}
</style>