<template>
	<el-form :model="formModel" label-width="80px" ref="refForm">
		<el-form-item label="分类" prop="sid" required>
			<el-select v-model="formModel.sid" placeholder="请选择">
				<el-option v-for="item in sort" :value="item.id" :label="item.name"></el-option>
			</el-select>
		</el-form-item>
		<el-form-item label="标题" prop="title" required>
			<el-input v-model="formModel.title"></el-input>
		</el-form-item>
		<el-form-item label="封面">
			<uploadImg v-model="formModel.thumb"></uploadImg>
		</el-form-item>
		<el-form-item label="状态" prop="status" required>
			<el-radio v-model="formModel.status" label="启用" :value="1" />
			<el-radio v-model="formModel.status" label="禁用" :value="-1" />			
		</el-form-item>
		<el-form-item label="富文本" prop="content" required>
			<editor v-model="formModel.content" @delfile="getDelFile"></editor>			
		</el-form-item>
		<el-form-item >
			<el-button type="primary" @click="save">提交</el-button>
		</el-form-item>	
	</el-form>
</template>

<script setup>
import {ref,onMounted,inject} from "vue";
import uploadImg from "../../components/upload/uploadImg.vue"
import editor from '../../components/editor/editor.vue';
const myApi=inject('myApi');
const ElMessage=inject('ElMessage');
const emit = defineEmits(['submit']);
const sort=ref()
const refForm=ref()	
const formModel=ref({status:1,content:''});
onMounted(()=>{
	myApi('articleSort',{page:1,limit:1000},'post').then(res=>{
		sort.value=res.data.data
	})
})	
function save(){	 
	refForm.value.validate(valid=>{
		if(valid){
			myApi('addArticle',formModel.value,'post').then((res)=>{
				ElMessage({message:res.msg,type:'success'})
				emit('submit');
				refForm.value.resetFields()
				delfile(editorDelFile)
			})
		}
	})	
}
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
</script>

<style>
</style>