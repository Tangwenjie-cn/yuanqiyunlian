
<template>
	<!--搜索栏-->
	<div class="page-header">
		<el-form></el-form>
		<el-button type="primary" icon="refresh" @click="refresh">刷新</el-button>
	</div>
	
	<!--按钮栏-->
	<div class="page-buttons">
		<el-button type="primary" @click="addDialog=true">添加参数</el-button>
	</div>
	<!--列表栏-->
	<div class="page-table">
		<el-tabs v-model="groupId" @tab-change="tabChange">
			<el-tab-pane v-for="item in groupList" :label="item.name" :name="item.value">
				<el-table :data="list" stripe border table-layout="fixed">
				    <el-table-column align="center" prop="id" label="ID" width="80" />
					<el-table-column label="名称" prop="title" width="150" align="center" />
					<el-table-column label="参数名" prop="key" width="150" align="center" />
					<el-table-column label="参数值" prop="value" width="150" align="center" show-overflow-tooltip />
					<el-table-column label="排序" prop="sort" width="80" align="center" />
					<el-table-column label="说明" prop="remark" width="180" align="center" />
					<el-table-column label="创建时间" width="180" align="center">
						<template #default="scope">
							{{new Date(scope.row.ctime).toLocaleString()}}
						</template>
					</el-table-column>
					<el-table-column label="更新时间" width="180" align="center">
						<template #default="scope">
							{{new Date(scope.row.utime).toLocaleString()}}
						</template>
					</el-table-column>
					<el-table-column label="操作" width="150">
						<template #default="scope">
							<el-button type="primary" size="small" @click="edit(scope)">编辑</el-button>						
							<el-popconfirm confirm-button-text="删除" cancel-button-text="取消" title="确定删除?" @confirm="del(scope)">
								<template #reference>
									<el-button type="danger" size="small">删除</el-button>
								</template>
							</el-popconfirm>
					    </template>
					</el-table-column>
				</el-table>
				<el-pagination class="page-pagination" v-model:page-size="limit" v-model:current-page="page" :page-sizes="[10,20, 50, 100, 200]" layout="total,sizes,prev,pager,next" :total="total"/>
			</el-tab-pane>
		</el-tabs>		
	</div>
	<!--弹出层-->
	<el-dialog v-model="addDialog" title="添加">
		<el-form :model="addForm" ref="refAddForm" label-width="80px">
			<el-form-item label="参数组" prop="group" required>
				<el-select v-model="addForm.group" placeholder="请选择">
					<el-option v-for="item in groupList" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item label="参数标题" prop="title" required>
				<el-input v-model="addForm.title"></el-input>
			</el-form-item>
			<el-form-item label="参数名" prop="key" required>
				<el-input v-model="addForm.key" placeholder="请输入字母，数字，下划线组合"></el-input>
			</el-form-item>
			<el-form-item label="类型" prop="type" required>
				<el-select v-model="addForm.type" placeholder="请选择" @change="typeChange">
					<el-option :value="1" label="文本框" />
					<el-option :value="2" label="文本域" />
					<el-option :value="3" label="列表" />
					<el-option :value="4" label="开关" />
					<el-option :value="5" label="多项选择" />
					<el-option :value="6" label="图片上传" />
					<el-option :value="7" label="文件上传" />
					<el-option :value="8" label="单选" />
				</el-select>
			</el-form-item>
			<el-form-item label="参数值" v-if="param" prop="param" required>
				<el-input v-model="addForm.param" type="textarea" placeholder="名称=值&#10;名称=值" />
			</el-form-item>
			<el-form-item label="值">
				<el-input v-model="addForm.value"/>			
			</el-form-item>
			<el-form-item label="排序" >
				<el-input v-model="addForm.sort" type="number" />			
			</el-form-item>
			<el-form-item label="说明">
				<el-input v-model="addForm.remark" type="textarea" />			
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="add">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
	<el-dialog v-model="editDialog" title="编辑">		
		<el-form :model="editForm" ref="refEditForm" label-width="80px">
			<el-form-item label="参数组" prop="group" required>
				<el-select v-model="editForm.group" placeholder="请选择">
					<el-option v-for="item in groupList" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item label="参数标题" prop="title" required>
				<el-input v-model="editForm.title"></el-input>
			</el-form-item>
			<el-form-item label="参数名" prop="key" required>
				<el-input v-model="editForm.key" placeholder="请输入字母，数字，下划线组合"></el-input>
			</el-form-item>
			<el-form-item label="类型" prop="type" required>
				<el-select v-model="editForm.type" placeholder="请选择" @change="typeChange">
					<el-option :value="1" label="文本框" />
					<el-option :value="2" label="文本域" />
					<el-option :value="3" label="列表" />
					<el-option :value="4" label="开关" />
					<el-option :value="5" label="多项选择" />
					<el-option :value="6" label="图片上传" />
					<el-option :value="7" label="文件上传" />
					<el-option :value="8" label="单选" />
				</el-select>
			</el-form-item>
			<el-form-item label="参数值" v-if="param" prop="param" required>
				<el-input v-model="editForm.param" type="textarea" placeholder="名称=值&#10;名称=值" />
			</el-form-item>
			<el-form-item label="值">
				<el-input v-model="editForm.value" />			
			</el-form-item>
			<el-form-item label="排序" >
				<el-input v-model="editForm.sort" type="number" />			
			</el-form-item>
			<el-form-item label="说明" >
				<el-input v-model="editForm.remark" type="textarea" />			
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="editSub">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
</template>

<script setup>
import {ref,onMounted,inject} from "vue"
	const myApi=inject('myApi');
	const ElMessage=inject('ElMessage');
	const list=ref([])
	const groupList=ref([])
	const groupId=ref('1');
	const addDialog=ref(false)
	const editDialog=ref(false)
	const addForm=ref({})
	const editForm=ref({})
	const refAddForm=ref()
	const refEditForm=ref()
	const searchForm=ref({});
	const page=ref(1);
	const limit=ref(10);	
	const total=ref(0);
	const param=ref(false)
	onMounted(()=>{
		getList()
		getGroup()
	})
	function refresh(){
		getList()
		getGroup()
	}
	function tabChange(e){
		page.value=1
		getList()
	}
	function getGroup(){
		myApi('setGroupList').then((res)=>{
			groupList.value=res.data
		})
	}
	function getList(){
		searchForm.value.page=page.value
		searchForm.value.limit=limit.value
		searchForm.value.group=Number(groupId.value) 
		myApi('setList',searchForm.value,'post').then((res)=>{
			list.value=res.data.data
			total.value=res.data.total
		})
	}
	function add(){
		refAddForm.value.validate(valid=>{
			if(valid){
				let data={...addForm.value}
				data.group=Number(data.group)
				data.sort=Number(data.sort)
				myApi('addSet',data,'post').then((res)=>{
					addDialog.value=false
					ElMessage({message:res.msg,type:'success'})
					getList()
					refAddForm.value.resetFields()
				})
			}
		})
		
	}
	function edit(scope){
		myApi('getSet',{id:scope.row.id}).then((res)=>{
			editDialog.value=true
			typeChange(scope.row.type)
			editForm.value=res.data
			editForm.value.group=editForm.value.group.toString()
		})
	}
	function editSub(){
		refEditForm.value.validate(valid=>{
			if(valid){
				let data={...editForm.value}
				data.group=Number(data.group)
				data.sort=Number(data.sort)
				myApi('editSet',data,'post').then((res)=>{
					editDialog.value=false
					ElMessage({message:res.msg,type:'success'})
					getList()
				})
			}
		})
		
	}
	function del(scope){
		myApi('delSet',{id:scope.row.id},'post').then((res)=>{
			ElMessage({message:res.msg,type:'success'})
			getList()
		})
	}
	function typeChange(val){
		if(val==3||val==5||val==8){
			param.value=true		
		}else if(val==4){
			param.value=false
			addForm.value.value='1'
		}else{
			param.value=false
		}
	}
</script>

<style>
	
</style>

