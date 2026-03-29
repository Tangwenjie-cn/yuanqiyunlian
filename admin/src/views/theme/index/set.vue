
<template>
    <div>
        <el-divider>{{ componentSetTitle }}</el-divider>
        <el-form>
            <el-form-item label="背景色" v-if="backgroundColor(model.type)">
                <el-color-picker v-model="model.data.backgroundColor" />
                <el-input style="width: auto;margin-left: 5px;" v-model="model.data.backgroundColor" />
            </el-form-item>
            <el-form-item label="标题" v-if="title(model.type)">
                <el-input style="width: auto;margin-left: 5px;" v-model="model.data.title" />
            </el-form-item>
            <el-form-item label="图标颜色" v-if="model.type=='user'">
                <el-color-picker v-model="model.data.iconColor" />
                <el-input style="width: auto;margin-left: 5px;" v-model="model.data.iconColor" />
            </el-form-item>
            <el-form-item label="文字颜色" v-if="fontSet(model.type)">
                <el-color-picker v-model="model.data.color" />
                <el-input style="width: auto;margin-left: 5px;" v-model="model.data.color" />
            </el-form-item>
            <el-form-item label="文字大小" v-if="fontSet(model.type)">
                <el-input style="width: auto;margin-left: 5px;" v-model="model.data.fontSize" />
            </el-form-item>
            <el-form-item label="文字颜色" v-if="model.type === 'status'">
                <el-radio-group v-model="model.data.color">
                    <el-radio value="#ffffff">白色</el-radio>
                    <el-radio value="#000000">黑色</el-radio>
                </el-radio-group>
            </el-form-item>
            <el-form-item label="高度" v-if="heightSet(model.type)">
                <el-input style="width: auto;margin-left: 5px;" v-model="model.data.height" />
            </el-form-item>
            <el-form-item label="渐变开始色" v-if="gradientColor(model.type)">
                <el-color-picker v-model="model.data.gradientStartColor" />
                <el-input style="width: auto;margin-left: 5px;" v-model="model.data.gradientStartColor" />
            </el-form-item>
            <el-form-item label="渐变结束色" v-if="gradientColor(model.type)">
                <el-color-picker v-model="model.data.gradientEndColor" />
                <el-input style="width: auto;margin-left: 5px;" v-model="model.data.gradientEndColor" />
            </el-form-item>
            <el-form-item label="排列方式" v-if="arrange(model.type)">
                <el-radio-group v-model="model.data.arrange">
                    <template v-if="model.type === 'goods' || model.type === 'article'">
                        <el-radio :value="1">一列</el-radio>
                        <el-radio :value="2">二列</el-radio>
                    </template>
                    <template v-else-if="model.type === 'classify'">
                        <el-radio :value="3">三列</el-radio>
                        <el-radio :value="4">四列</el-radio>
                        <el-radio :value="5">五列</el-radio>
                    </template>
                    <template v-else-if="model.type === 'operate'">
                        <el-radio :value="1">一列</el-radio>
                        <el-radio :value="2">二列</el-radio>
                        <el-radio :value="3">三列</el-radio>
                        <el-radio :value="4">四列</el-radio>
                        <el-radio :value="5">五列</el-radio>
                    </template>                    
                </el-radio-group>
            </el-form-item>
            <el-form-item label="数据内容" v-if="dataContent(model.type)">
                <el-button @click="addData()"><el-icon><Plus /></el-icon>增加</el-button>
            </el-form-item>
            <div style="display: flex;flex-direction: column;" v-if="dataContent(model.type)">                    
                <div style="display: flex;flex-direction: column;border: 1px solid #0000006b;margin-bottom: 5px;padding: 5px;" 
                v-for="(item,index) in model.data.list">
                    <uploadImg v-model="item.imgUrl"></uploadImg>
                    <div style="display: flex;align-items: center;">
                        <div style="display: flex;flex-direction: column;margin-top: 5px;">
                            <div>名称 <el-input style="width: auto;margin-bottom: 5px;" v-model="item.name"></el-input></div>
                            <div>链接 <el-input style="width: auto;" v-model="item.link"></el-input></div>
                        </div>
                        <el-icon style="margin-left: 5px;" @click="model.data.list.splice(index,1)"><Close /></el-icon>
                    </div>                   
                </div>
            </div>
            <el-form-item label="图片" v-if="image(model.type)">
                <uploadImg v-model="model.data.imgUrl"></uploadImg>
            </el-form-item>
            <el-form-item label="链接" v-if="link(model.type)">
                <el-input style="width: auto;margin-left: 5px;" v-model="model.data.link" />
            </el-form-item>
            <el-form-item label="数量" v-if="limit(model.type)">
                <el-input style="width: auto;margin-left: 5px;" v-model="model.data.limit" />
            </el-form-item>
            <el-form-item label="类型" v-if="goods(model.type)">
                <el-radio-group v-model="model.data.type">
                    <el-radio :value="0">全部</el-radio>
                    <el-radio :value="1">热门</el-radio>
                    <el-radio :value="2">置顶</el-radio>
                </el-radio-group>
            </el-form-item>
            <el-form-item label="分类" v-if="sortId(model.type)">
                <!-- <el-input style="width: auto;margin-left: 5px;"/> -->
                <el-cascader v-model="model.data.sortId" :options="sortData" :props="props" clearable>
				</el-cascader>
            </el-form-item>
        </el-form>
    </div>
</template>
<script setup>
    import { ref,watch, inject} from 'vue'
    import uploadImg from "../../../components/upload/uploadImg.vue"
    const model = defineModel()
    const componentSetTitle = ref('组件设置')
    const sortData=ref([])
    const myApi=inject('myApi');
    const props=ref({
		checkStrictly: true,
		label: 'name',
		value: 'id',
		emitPath: false,
	});
    watch(model, (newVal) => {
        let type = newVal.type
        switch(type){
            case 'status':
                componentSetTitle.value = '状态栏设置'
                break
            case 'search-carousel':
                componentSetTitle.value = '搜索轮播设置'
                break
            case 'carousel':
                componentSetTitle.value = '轮播图设置'
                break
            case 'classify':
                componentSetTitle.value = '分类设置'
                break
            case 'search':
                componentSetTitle.value = '搜索设置'
                break
            case 'notice':
                componentSetTitle.value = '公告设置'
                break
            case 'article':
                componentSetTitle.value = '文章设置'
                myApi('articleSort',{page:1,limit:1000},'post').then(res=>{
                    sortData.value=res.data.data
                })
                break
            case 'user':
                componentSetTitle.value = '用户设置'
                break
            case 'operate':
                componentSetTitle.value = '操作设置'
                break 
            case 'goods':
                componentSetTitle.value = '商品设置'
                myApi('goodsSort',{},'post').then((res)=>{
                    sortData.value=res.data
                })
                break 
            case 'image':
                componentSetTitle.value = '图片设置'
                break
            case 'video':
                componentSetTitle.value = '视频设置'
                break
            default:
                componentSetTitle.value = '组件设置'
        }
    })
    function addData(){
        let data = {
            imgUrl:'',
            name:'',
            link:''
        }
        model.value.data.list.push(data)
    }
    function backgroundColor(type){
        if(type==='search-carousel' || type==='carousel' || type==='image' || type==='video'){
            return false
        }else{
            return true
        }
    }
    function gradientColor(type){
        if(type==='search-carousel'||type==='carousel' || type==='goods' || type==='user'){
            return true
        }else{
            return false
        }
    }
    function fontSet(type){
        if(type==='classify' || type==='notice' || type==='goods' || type==='user' || type === 'operate'){
            return true
        }else{
            return false
        }
    }
    function arrange(type){
        if(type==='classify' || type==='goods'　|| type==='article' || type === 'operate'){
            return true
        }else{
            return false
        }
    }
    function dataContent(type){
        if(type==='classify' || type==='search-carousel' || type==='carousel' || type === 'operate'){
            return true
        }else{
            return false
        }
    }
    function image(type){
        if(type==='image'){
            return true
        }else{
            return false
        }
    }
    function link(type){
        if(type==='image' || type==='video'){
           return true
        }else{
            return false
        }
    }
    function title(type){
        if(type==='goods' || type==='status'){
            return true
        }else{
            return false
        }
    }
    function goods(type){
        if(type==='goods'){
            return true
        }else{
            return false
        }
    }
    function limit(type){
        if(type==='goods' || type==='article'){
            return true
        }else{
            return false
        }
    }
    function sortId(type){
        if(type==='goods' || type==='article'){
            return true
        }else{
            return false
        }
        
    }
    function heightSet(type){
        if(type==='image' || type==='video'){
            return true
        }else{
            return false
        }
    }
</script>