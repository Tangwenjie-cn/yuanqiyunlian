
<template>
    <el-form :model="formModel" label-width="80px" ref="refForm">
        <el-tabs>
            <el-tab-pane label="基础"> 
                <el-form-item label="类型" prop="type" required>
                    <el-radio-group v-model="formModel.type">
                        <el-radio-button :value="1" label="文档"/>
                        <el-radio-button :value="2" label="文章"/>
                        <el-radio-button :value="3" label="视频"/>
                        <el-radio-button :value="4" label="音频"/>
                        <el-radio-button :value="5" label="网盘"/>
                    </el-radio-group>
                </el-form-item>                
                <el-form-item label="封面" prop="thumb" required>
                    <uploadImg v-model="formModel.thumb"></uploadImg>
                </el-form-item>
                <el-form-item 
                v-if="formModel.type == 1 || formModel.type == 3 || formModel.type == 4 || formModel.type == 5" 
                label="链接" 
                prop="link" required>
                    <uploadFile v-model="formModel.link"></uploadFile>
                    <el-alert title="视情况选择上传或直接填写链接地址" type="info" style="margin-top: 5px;" show-icon :closable="false" />
                </el-form-item>
                <el-form-item label="提取码" v-if="formModel.type == 5">
                    <el-input v-model="formModel.code" style="width: 48%;"></el-input>
                </el-form-item> 
                <el-form-item label="标题" prop="title" required>
                    <el-input v-model="formModel.title"></el-input>
                </el-form-item>
                <el-form-item label="分类" prop="sort_ids" required>
                    <el-cascader v-model="formModel.sort_ids" :options="sort"  :props="cascaderProps">
                    </el-cascader>
                </el-form-item>
                <el-form-item label="栏目" prop="cid">
                    <el-select v-model="formModel.cid" placeholder="请选择" clearable>
                        <el-option label="未选择" :value="0"></el-option>
                        <el-option v-for="item in columns" :label="item.name" :value="item.id">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="简介">
                    <el-input v-model="formModel.introduction" type="textarea"></el-input>
                </el-form-item>
                <el-form-item label="成本价">
                    <el-input-number v-model="formModel.cost_price" :precision="2">
                        <template #prefix>
                            <el-text>￥</el-text>
                        </template>
                    </el-input-number>
                </el-form-item>
                <el-form-item label="原价">
                    <el-input-number v-model="formModel.original_price" :precision="2">
                        <template #prefix>
                            <el-text>￥</el-text>
                        </template>
                    </el-input-number>
                    <el-text type="info" style="margin-left: 10px;">划线价展示用</el-text>  
                </el-form-item>
                <el-form-item label="价格" prop="price" required>
                    <el-input-number v-model="formModel.price" :precision="2">
                        <template #prefix>
                            <el-text>￥</el-text>
                        </template>
                    </el-input-number>
                </el-form-item>
                <el-form-item label="积分" prop="points">
                    <el-input-number v-model="formModel.points">
                    </el-input-number>
                    <el-text type="info" style="margin-left: 10px;">0不允许积分兑换</el-text>
                </el-form-item>
                <el-form-item v-if="formModel.type == 1" label="总页数" prop="pages" required>
                    <el-input-number v-model="formModel.pages"></el-input-number>
                </el-form-item>
                <el-form-item label="浏览量" prop="views" required>
                    <el-input-number v-model="formModel.views"></el-input-number>
                </el-form-item>
                <el-form-item label="状态" prop="status" required>
                    <el-radio v-model="formModel.status" label="上架" :value="1"></el-radio>
                    <el-radio v-model="formModel.status" label="下架" :value="-1"></el-radio>				
                </el-form-item>
            </el-tab-pane>
            <el-tab-pane label="内容详情">
                <el-form-item label="商品详情" prop="content" required>
                    <editor v-model="formModel.content" @delfile="getDelFile"></editor>			
                </el-form-item>
                <el-form-item v-if="formModel.type == 2" label="收费内容" prop="fee_content" required>
                    <editor v-model="formModel.fee_content" @delfile="getDelFile"></editor>			
                </el-form-item>
            </el-tab-pane>
            <el-tab-pane label="其他配置">
                <el-form-item label="分销">
                    <el-switch
                        v-model="formModel.retail_on"
                        inline-prompt
                        active-text="开启"
                        inactive-text="关闭"
                        :active-value="1"
                        :inactive-value="-1"
                    />			
                </el-form-item>
                
                <el-form-item label="佣金设置" v-if="formModel.retail_on==1">
                    <el-radio-group v-model="formModel.retail_set">
                        <el-radio :value="1">公用设置</el-radio>
                        <el-radio :value="2">单独设置</el-radio>
                    </el-radio-group>			
                </el-form-item>
                <el-form-item label="分配方式" v-if="formModel.retail_on==1">
                    <el-radio-group v-model="formModel.retail_type">
                        <el-radio :value="1">固定金额</el-radio>
                        <el-radio :value="2">百分比</el-radio>
                    </el-radio-group>			
                </el-form-item>
                <el-form-item label="佣金" v-if="formModel.retail_on==1">
                    <el-input-number v-model="formModel.retail_yj"></el-input-number>
                    <el-text type="info" style="margin-left: 10px;">百分比填写百分比值，如5%填5</el-text>
                </el-form-item>
                <el-form-item label="置顶">
                    <el-switch
                        v-model="formModel.is_top"
                        inline-prompt
                        active-text="是"
                        inactive-text="否"
                        :active-value="1"
                        :inactive-value="-1"
                    />			
                </el-form-item>
                <!-- <el-form-item label="激励广告">
                    <el-switch
                        v-model="formModel.is_adv"
                        inline-prompt
                        active-text="开启"
                        inactive-text="关闭"
                        :active-value="1"
                        :inactive-value="-1"
                    />			
                </el-form-item> -->
                <el-form-item label="分享获取">
                    <el-switch
                        v-model="formModel.is_share"
                        inline-prompt
                        active-text="开启"
                        inactive-text="关闭"
                        :active-value="1"
                        :inactive-value="-1"
                    />			
                </el-form-item>
                <el-form-item label="分享数量">
                    <el-input-number v-model="formModel.share_nums"></el-input-number>
                    <el-text type="info" style="margin-left: 10px;">分享多少次后可免费获取</el-text>
                </el-form-item>
            </el-tab-pane>
        </el-tabs>
        <el-form-item >
			<el-button type="primary" @click="save">提交</el-button>
		</el-form-item>
    </el-form>   
</template>
<script setup>
import {ref,onMounted,inject} from "vue";
import uploadImg from "../../../components/upload/uploadImg.vue"
import uploadFile from "../../../components/upload/uploadFile.vue"
import editor from '../../../components/editor/editor.vue';
const myApi=inject('myApi');
const ElMessage=inject('ElMessage');
const emit = defineEmits(['submit']);
const sort=ref()
const columns=ref([])
const refForm=ref()	
const formModel=ref({})
const props = defineProps(['id']);
const cascaderProps=ref({
    multiple: true,
    checkStrictly: true,
    emitPath: false,
    value: 'id',
    label: 'name',
})
let editorDelFile
function getDelFile(files){
	editorDelFile=files
}
function delfile(files){
	if(!files){
		return
	}
	if(files.length>0){
		for (let index = 0,len=files.length; index < len; index++) {
			let src = files[index];
			myApi('delFile',{"path":src},'post')
		}
	}	
}
onMounted(()=>{
    myApi('getGoods',{id:props.id}).then((res)=>{
		formModel.value={...res.data}
        formModel.value.sort_ids=[]
        for(let i=0;i<res.data.gs_correlation.length;i++){
			formModel.value.sort_ids.push(res.data.gs_correlation[i].sid)
		}
        formModel.value.content=res.data.goods_content.content
        formModel.value.fee_content=res.data.goods_content.fee_content
    })
    myApi('goodsSort',{},'post').then((res)=>{
		sort.value=res.data
	})
    myApi('GoodsColumn').then((res)=>{
		columns.value=res.data
	})
})
function save(){	 
	refForm.value.validate(valid=>{
		if(valid){            
			myApi('editGoods',formModel.value,'post').then((res)=>{
				ElMessage({message:res.msg,type:'success'})
				delfile(editorDelFile)
				refForm.value.resetFields()
                emit('submit');
			})
		}
	})	
}
</script>
