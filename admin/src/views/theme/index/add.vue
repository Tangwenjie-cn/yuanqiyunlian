<template> 
<div class="container">
    <div class="component-library">
        <h3>组件库</h3>
        <div>
            <el-button class="lib-btn" type="primary" plain @click="addItem('status')">状态栏</el-button>
        </div>
        <div>
            <el-button class="lib-btn" type="primary" plain @click="addItem('search-carousel')">搜索轮播图</el-button>
        </div>
        <div>
            <el-button class="lib-btn" type="primary" plain @click="addItem('carousel')">轮播图</el-button>
        </div>
        <div>
            <el-button class="lib-btn" type="primary" plain @click="addItem('classify')">分类</el-button>
        </div>
        <div>
            <el-button class="lib-btn" type="primary" plain @click="addItem('notice')">公告</el-button>
        </div>
        <div>
            <el-button class="lib-btn" type="primary" plain @click="addItem('search')">搜索框</el-button>
        </div>
        <div>
            <el-button class="lib-btn" type="primary" plain @click="addItem('image')">图片</el-button>
        </div>
        <div>
            <el-button class="lib-btn" type="primary" plain @click="addItem('video')">视频</el-button>
        </div>
        <div>
            <el-button class="lib-btn" type="primary" plain @click="addItem('goods')">商品</el-button>
        </div>
        <div>
            <el-button class="lib-btn" type="primary" plain @click="addItem('article')">文章</el-button>
        </div>
    </div>
    <div class="diy-block">
        <div style="display: flex;align-items: center;height: 35px;">
            <el-button size="small" type="danger" plain @click="removeItem">移除选中组件</el-button>
            <el-input v-model="title" placeholder="请输入主题名称" style="width: auto;margin-left: 5px;"></el-input>
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
                    <!-- 轮播图 -->
                    <div v-if="item.type === 'carousel'" class="block" :class="{selected: item.i == selectedId,unselected: item.i != selectedId}">
                        <div :style="{'background': `linear-gradient(${item.data.gradientStartColor},${item.data.gradientEndColor})`}">
                            <el-carousel height="128px">
                                <el-carousel-item v-for="item1 in item.data.list">
                                    <img :src="item1.imgUrl" height="128px" width="100%">
                                </el-carousel-item>
                            </el-carousel>
                        </div>
                    </div>
                    <!-- 搜索轮播图 -->
                    <div v-if="item.type === 'search-carousel'" class="block" :class="{selected: item.i == selectedId,unselected: item.i != selectedId}">
                        <div :style="{'background': `linear-gradient(${item.data.gradientStartColor},${item.data.gradientEndColor})`}">
                            <div class="search-block">
                                <div style="background-color: #ffff;" class="search-input">请输入搜索内容

                                </div>
                            </div>
                            <el-carousel height="128px">
                                <el-carousel-item v-for="item1 in item.data.list">
                                    <img :src="item1.imgUrl" height="128px" width="100%">
                                </el-carousel-item>
                            </el-carousel>
                        </div>
                    </div>
                    <!-- 分类 -->
                    <div v-if="item.type === 'classify'" class="block" :class="{selected: item.i == selectedId,unselected: item.i != selectedId}">
                        <div class="classify-block" :style="{'background-color': item.data.backgroundColor,'grid-template-columns': `repeat(${item.data.arrange}, 1fr)`}">
                            <div v-for="item1 in item.data.list" class="classify-item">
                                <img :src="item1.imgUrl">
                                <span class="classify-name" :style="{'color': item.data.color, 'font-size': item.data.fontSize}">{{ item1.name }}</span>
                            </div>
                        </div>
                    </div>
                    <!-- 公告 -->
                    <div v-if="item.type === 'notice'" class="block" :class="{selected: item.i == selectedId,unselected: item.i != selectedId}">
                        <div class="notice-block" :style="{'background-color': item.data.backgroundColor}">
                            <span class="iconfont icon-laba" :style="{'color': item.data.color, 'font-size': item.data.fontSize}"></span>
                            <span class="notice-content" :style="{'color': item.data.color, 'font-size': item.data.fontSize}">公告内容</span>
                        </div>
                    </div>
                    <!-- 搜索 -->
                    <div v-if="item.type === 'search'" class="block" :class="{selected: item.i == selectedId,unselected: item.i != selectedId}">
                        <div class="search-block">
                            <div style="background-color: #ffff;" class="search-input">请输入搜索内容

                            </div>
                        </div>
                    </div>
                    <!-- 图片 -->
                    <div v-if="item.type === 'image'" class="block" :class="{selected: item.i == selectedId,unselected: item.i != selectedId}">
                        <img :src="item.data.imgUrl" class="block-img">
                    </div>
                    <!-- 视频 -->
                    <div v-if="item.type === 'video'" class="block" :class="{selected: item.i == selectedId,unselected: item.i != selectedId}">
                        <video :src="item.data.link" class="block-video" controls></video>
                    </div>
                    <!-- 商品 -->
                    <div v-if="item.type === 'goods'" class="block" :class="{selected: item.i == selectedId,unselected: item.i != selectedId}">
                        <div class="top row-center" :style="{'background': `linear-gradient(to right,${item.data.gradientStartColor},${item.data.gradientEndColor})`}">
                            <span :style="{'color': item.data.color, 'font-size': item.data.fontSize}">{{ item.data.title }}</span>
                            <span :style="{'color': item.data.color, 'font-size': item.data.fontSize}">更多》</span>
                        </div>
                        <div class="goods-block" :style="{'background-color': item.data.backgroundColor,'grid-template-columns': `repeat(${item.data.arrange}, 1fr)`}">
                            <div class="two-item" v-if="item.data.arrange === 2">
                                <img class="two-img" src="/src/assets/images/avatar.jpeg" mode=""></img>
                                <div class="two-block column"> 
                                    <span class="two-title">商品标题商品标题</span>
                                    <div class="two-bottom row-center">
                                        <div class="column">
                                            <span class="price">￥250.00</span>
                                            <span class="price-underline">￥500.00</span>
                                        </div>
                                        <div class="column">
                                            <span class="text">5分钟前</span>
                                            <span class="text">
                                                <el-icon style="font-size: 14px;color: #999;"><View /></el-icon>
                                                610
                                            </span>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div class="two-item" v-if="item.data.arrange === 2">
                                <img class="two-img" src="/src/assets/images/avatar.jpeg"></img>
                                <div class="two-block column"> 
                                    <span class="two-title">商品标题商品标题</span>
                                    <div class="two-bottom row-center">
                                        <div class="column">
                                            <span class="price">￥250.00</span>
                                            <span class="price-underline">￥500.00</span>
                                        </div>
                                        <div class="column">
                                            <span class="text">5分钟前</span>
                                            <span class="text">
                                                <el-icon style="font-size: 14px;color: #999;"><View /></el-icon>
                                                610
                                            </span>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div class="one-item" v-if="item.data.arrange === 1">
                                <img class="one-img" src="/src/assets/images/avatar.jpeg" mode=""></img>
                                <div class="one-right-block column"> 
                                    <span class="one-title">商品标题商品标题商品标题商品标题商品标题商品标题</span>
                                    <div class="one-right-bottom row-center">
                                        <div class="column">
                                            <span class="price">￥250.00</span>
                                            <span class="price-underline">￥500.00</span>
                                        </div>
                                        <div class="column">
                                            <span class="text">5分钟前</span>
                                            <span class="text">
                                                <el-icon style="font-size: 14px;color: #999;"><View /></el-icon>
                                                6.6千
                                            </span>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div v-if="item.type === 'article'" class="block" :class="{selected: item.i == selectedId,unselected: item.i != selectedId}">
                        <div class="goods-block" :style="{'background-color': item.data.backgroundColor,'grid-template-columns': `repeat(${item.data.arrange}, 1fr)`}">
                            <div class="article-one-item" v-if="item.data.arrange === 1">
                                <img class="article-one-img" src="/src/assets/images/avatar.jpeg" mode=""></img>
                                <div class="article-one-text column"> 
                                    <span class="article-one-title">文章标题文章标题文章标题文章标题</span>
                                    <span class="article-time">2小时前</span>
                                </div>
                            </div>
                            <div class="article-two-item" v-if="item.data.arrange === 2">
                                <img class="article-two-img" src="/src/assets/images/avatar.jpeg" mode=""></img>
                                <div class="article-two-text column"> 
                                    <span class="article-two-title">文章标题文章标题文章标题文章标题</span>
                                    <span class="article-time">2小时前</span>
                                </div>
                            </div>
                            <div class="article-two-item" v-if="item.data.arrange === 2">
                                <img class="article-two-img" src="/src/assets/images/avatar.jpeg" mode=""></img>
                                <div class="article-two-text column"> 
                                    <span class="article-two-title">文章标题文章标题文章标题文章标题</span>
                                    <span class="article-time">2小时前</span>
                                </div>
                            </div>
                        </div>
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
import { ref,inject } from 'vue'
import componentSet from './set.vue'
import {randomStr} from '../../../config/functions.js'
const myApi=inject('myApi');
const ElMessage=inject('ElMessage');
const layout = ref([])
const selectedId = ref(null)
const selectedIndex = ref(null)
const isComponentSet= ref(false)
const title = ref('')
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
        case 'search-carousel':
            item.h = 15
            data={
                gradientStartColor:'#f56b59',
                gradientEndColor:'#ffffff',
                list:[]
            }
            item.data=data
            break
        case 'carousel':
            item.h = 11
            data={
                gradientStartColor:'#f56b59',
                gradientEndColor:'#ffffff',
                list:[]
            }
            item.data=data
            break
        case 'classify':
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
        case 'notice':
            item.h = 3
            data={
                backgroundColor:'#ffffff',
                color:'#e21515',
                fontSize:'14px',
            }
            item.data=data
            break
        case 'search':
            item.h = 4
            data={
                backgroundColor:'#ffffff',
            }
            item.data=data
            break
        case 'image':
            item.h = 6
            data={
                imgUrl:'',
                link:''
            }
            item.data=data
            break
        case 'video':
            item.h = 12
            data={
                link:'',
                imgUrl:'',
            }
            item.data=data
            break
        case 'goods':
            item.h = 21
            data={
                backgroundColor:'#ffffff',
                arrange:2,
                gradientStartColor:'#f56b59',
                gradientEndColor:'#c68686',
                title:'热门商品',
                color:'#ffffff',
                fontSize:'14px',
                limit:2,
                type:0,
                sortId:0,
            }
            item.data=data
            break
        case 'article':
            item.h = 19
            data={
                backgroundColor:'#ffffff',
                arrange:2,
                limit:2,
                sortId:1,
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
    if (title.value == '') {
        return ElMessage.warning('请输入主题名称')
    }

    myApi('ThemeAddIndex',{title:title.value,content:[...layout.value].sort((a, b) => a.y - b.y)},'post').then(res=>{       
        ElMessage.success('添加成功')
        setTimeout(()=>{
            window.location.href='/theme/index'
        },1000)
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
.notice-block{
    display: flex;
    align-items: center;
    padding: 6px;
}
.notice-block .icon-laba{
    color: #ea1a1a;
    font-size: 20px;
    margin-right: 5px;
}
.notice-block .notice-content{
    font-size: 12px;
    color: #211f1f;
}
.search-block{
    display: flex;
    justify-content: center;
    align-items: center;
    height: 40px;
}
.search-input{
    width: 200px;
    height: 30px;
    border-radius: 20px;
    display: flex;
    align-items: center;
    padding-left: 10px;
}
.selected{
    border: 3px solid #3858d6;
}
.unselected{
    border: 1px dashed rgb(119, 119, 119);
}
.block{
    position: relative;
    margin-bottom: 5px;
}
.component-set{
    margin-left: 10px;
}
.block-img{
    width: 100%;
    height: 60px;
}
.block-video{
    width: 100%;
    height: 120px;
}
.top{
    height: 48px;
    border-top-right-radius: 10rpx;
    border-top-left-radius: 10rpx;
    justify-content: space-between;
    padding: 0 10px;
}
.goods-block{
    display: grid;
}
.two-item{
    border-radius: 10px;
    background-color: #ffffff;
    width: 160px;
    height: 180px;
    display: flex;
    flex-direction: column;
}
.column{
    display: flex;
    flex-direction: column;
}
.row-center{
    display: flex;
    align-items: center;
}
.two-img{
    width: 100%;
    height: 120px;
}
.two-block{
    padding: 5px;
}
.two-title{
    font-size: 16px;
    color: #000000cc;  
    overflow: hidden;       
    text-overflow: ellipsis;
    white-space: nowrap;
    width: 160px;
    min-width: 0;
    font-weight: 300;
}
.two-bottom{
    justify-content: space-between;
}
.text{
    color: #999;
    font-size: 12px;
}
.one-item{
    display: flex;
    border-radius: 10rpx;
    background-color: #ffffff;
    padding: 10px;
    margin-top: 10px;
}
.one-img{
    width: 160px;
    height: 120px;
}
.one-right-block{
    margin-left: 10px;
    width: 60%;
    justify-content: space-between;
}
.one-title{
    font-size: 16px;
    font-weight: 300;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
    text-overflow: ellipsis;
}
.one-right-bottom{
    justify-content: space-between;
}
.price{
    color: #e05211;
    font-size: 14px;
    font-weight: 600;
}
.price-underline{
    color: #999;
    font-size: 12px;
    text-decoration: line-through;
}
.article-one-item{
		border-radius: 10px;
		background-color: #ffffff;
		padding: 10px;
		display: flex;
		margin-bottom: 10px;
	}
	.article-one-img{
		width: 160px;
		height: 120px;
	}
	.article-one-text{
		margin-left: 10px;
		justify-content: space-between;
		width: 60%;
	}
	.article-one-title{
		font-size: 16px;
		color: #000000cc;  
		overflow: hidden;       
		text-overflow: ellipsis; 
	}
	.article-time{
		font-size: 12px;
		color: #999999;
	}
    .article-two-item{
		border-radius: 10px;
		background-color: #ffffff;
		width: 150px;
	}
	.article-two-img{
        display: flex;
		width: 100%;
		height: 155px;
	}
	.article-two-text{
		width: 100%;
		padding: 10px;
	}
	.article-two-title{
		font-size: 16px;
		color: #000000cc;  
		overflow: hidden;       
		text-overflow: ellipsis;
		white-space: nowrap;
		width: 150px;
		min-width: 0; 
	}
</style>