
<template>
    <div style="border: 1px solid #ccc;width: 100%;">
        <Toolbar
        style="border-bottom: 1px solid #ccc"
        :editor="editorRef"
        :defaultConfig="toolbarConfig"
        />
        <Editor
        style="height: 500px; overflow-y: hidden;"
        v-model="model"
        :defaultConfig="editorConfig"
        @onCreated="handleCreated"
        @onChange="handleChange"
        />
    </div>
</template>
<script setup>
import '@wangeditor/editor/dist/css/style.css';	
import {onBeforeUnmount,shallowRef,inject} from "vue";
import { Editor, Toolbar } from '@wangeditor/editor-for-vue';

const myApi=inject('myApi');
const model = defineModel()
const editorRef = shallowRef();
const toolbarConfig = {};
const emit = defineEmits(['delfile'])
let saveImg=[];
let initImg=[]
let saveVideo=[];
let initVideo=[]
const editorConfig = {
	placeholder: '请输入内容...',
	MENU_CONF:{
		uploadImage:{
			// 小于该值就插入 base64 格式（而不上传）
			base64LimitSize: 10 * 1024, // 10kb
			maxNumberOfFiles:1,
			async customUpload(file, insertFn){
                myApi('uploadFile',{file:file},'post',{'Content-Type':'multipart/form-data'}).then(res=>{
                    insertFn(res.data)
                })
			}
		},
        insertImage:{
            onInsertedImage(imageNode){ //插入图片回调
                if (imageNode == null) return
                if(imageNode.src.indexOf('base64')<0){
                    saveImg.push(imageNode)
                }                
            }
        },
        uploadVideo:{
            async customUpload(file, insertFn) {
                myApi('uploadFile',{file:file},'post',{'Content-Type':'multipart/form-data'}).then(res=>{
                    insertFn(res.data)
                })
            }
        },
        insertVideo:{
            onInsertedVideo(videoNode){
                if (videoNode == null) return
                saveVideo.push(videoNode)
            }
        }
	}
};	
onBeforeUnmount(() => {
    if(editorRef.value){
        editorRef.value.destroy()
    }
});
function handleCreated(editor){
	editorRef.value = editor // 记录 editor 实例，重要！
    initImg=editor.getElemsByType('image')   
    initVideo=editor.getElemsByType('video')    
}
function handleChange(editor){
    let imgs=editor.getElemsByType('image')
    let videos=editor.getElemsByType('video')

    let delfile=[]
    if(saveImg.length>0){ //处理新增后又删除的图片列表
        for (let index = 0,len=saveImg.length; index < len; index++) {
            let src = saveImg[index].src;
            let i=imgs.findIndex(val => val.src==src)           
            if(i<0){
                delfile.push(src)
            }
        }
    }
    if(initImg.length>0){ //处理编辑后删除的图片列表
        for (let index = 0,len=initImg.length; index < len; index++) {
            let src = initImg[index].src;
            let i=imgs.findIndex(val => val.src==src)
            if(i<0){
                delfile.push(src)
            }
        }
    }
    if(saveVideo.length>0){ //处理新增后又删除的视频列表
        for (let index = 0,len=saveVideo.length; index < len; index++) {
            let src = saveVideo[index].src;
            let i=videos.findIndex(val => val.src==src)
            if(i<0){
                delfile.push(src)
            }
        }
    }
    if(initVideo.length>0){ //处理编辑后又删除的视频列表
        for (let index = 0,len=initVideo.length; index < len; index++) {
            let src = initVideo[index].src;
            let i=videos.findIndex(val => val.src==src)
            if(i<0){
                delfile.push(src)
            }
        }
        
    }
    emit('delfile',delfile)

}    
</script>