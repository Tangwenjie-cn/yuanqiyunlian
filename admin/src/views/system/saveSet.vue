<template>
	<!--搜索栏-->
	<div class="page-header">
		<el-form></el-form>
		<el-button type="primary" icon="refresh" @click="refresh">刷新</el-button>
	</div>
	<!--列表栏-->
	<div class="page-table">
		<el-form>
			<el-tabs v-model="tableTabsValue">
				<el-tab-pane v-for="item in groupList" :label="item.name" :name="item.value">
					<el-form-item v-for="(item1,index) in list[item.value]" :label="item1.title" label-position="top">
						<template v-if="item1.type==1">
							<el-input v-model="item1.value" :placeholder="item1.remark"></el-input>
						</template>
						<template v-else-if="item1.type==2">
							<el-input v-model="item1.value" type="textarea" :placeholder="item1.remark"></el-input>
						</template>
						<template v-else-if="item1.type==3">
							<el-select v-model="item1.value"  placeholder="请选择">								
								<el-option v-for="item2 in item1.list" :value="item2.value" :label="item2.name" />
							</el-select>
						</template>
						<template v-else-if="item1.type==4">
							<el-switch v-model="item1.value" 
							active-value="1"
							inactive-value="0"
							active-text="开启"
    						inactive-text="关闭"
							inline-prompt />
						</template>
						<template v-else-if="item1.type==5">
							<el-checkbox-group v-model="item1.value">
								<el-checkbox v-for="item2 in item1.list" :label="item2.value" :value="item2.value">{{item2.name}}</el-checkbox>
							</el-checkbox-group>
						</template>
						<template v-else-if="item1.type==6">
							<uploadImg v-model="item1.value"></uploadImg>
						</template>
						<template v-else-if="item1.type==7">
							<uploadFile v-model="item1.value"></uploadFile>
						</template>
						<template v-else-if="item1.type==8">
							<el-radio-group v-model="item1.value">
								<el-radio v-for="item2 in item1.list" :value="item2.value">{{item2.name}}</el-radio>
							</el-radio-group>
						</template>
						<el-alert v-if="item1.remark" :title="item1.remark" style="margin-top: 5px;" type="info" show-icon :closable="false" />
					</el-form-item>
				</el-tab-pane>
			</el-tabs>
			<el-form-item>
				<el-button type="primary" @click="save">保存</el-button>
			</el-form-item>
		</el-form>	
	</div>
	
</template>

<script setup>
	import {ref,onMounted,inject} from "vue"
	import uploadImg from "../../components/upload/uploadImg.vue"
	import uploadFile from "../../components/upload/uploadFile.vue"
	const myApi=inject('myApi');
	const ElMessage=inject('ElMessage');
	const list=ref([])
	const groupList=ref([])
	const tableTabsValue = ref('1')
	onMounted(()=>{
		getGroup()
		getList()
	})
	function refresh(){
		getGroup()
		getList()
	}
	function getGroup(){
		myApi('setGroupList').then((res)=>{
			groupList.value=res.data
		})
	}
	function getList(){
		myApi('getSetAllInfo').then((res)=>{
			list.value=res.data
		})
	}
	function save(){
		myApi('saveAllSet',list.value,'post').then((res)=>{
			ElMessage({message:res.msg,type:'success'})
		})
	}
</script>

<style>
	
</style>