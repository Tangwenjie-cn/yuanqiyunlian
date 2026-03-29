<template>
	<!--搜索栏-->
	<div class="page-header">
		<el-form :inline="true" :model="searchForm">
			<el-form-item class="page-search-form-item page-search-select">
				<el-select v-model="searchForm.status" clearable placeholder="状态">
					<el-option value="1" label="正常"></el-option>
					<el-option value="-1" label="禁用"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item class="page-search-form-item page-search-select">
				<el-select v-model="searchForm.sid" clearable placeholder="分类">
					<el-option v-for="item in sort" :value="item.id" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item class="page-search-form-item">
				<el-button type="primary" icon="search" @click="getList()" />
			</el-form-item>
		</el-form>
		<el-button type="primary" icon="refresh" @click="refresh">刷新</el-button>
	</div>
	
	<!--按钮栏-->
	<div class="page-buttons">
		<el-button type="primary" @click="addDialog=true">添加图片</el-button>
	</div>
	<!--列表栏-->
	<div class="page-table">
		<el-table :data="list" stripe border table-layout="fixed" style="width: 100%">
		    <el-table-column align="center" prop="id" label="ID" width="80" />
			<el-table-column label="缩略图" width="120" align="center">
				<template #default="scope">
					<el-image :src="scope.row.thumb" fit="contain">
						<template #error>
							<el-icon><Picture /></el-icon>
						</template>
					</el-image>
				</template>
			</el-table-column>
			<el-table-column label="分类" width="120" align="center">
				<template #default="scope">
					{{scope.row.ImgSort.name}}
				</template>
			</el-table-column>			
			<el-table-column label="状态" width="120" align="center">
				<template #default="scope">
					<span v-if="scope.row.status==1" style="color: #67C23A;">正常</span>
					<span v-if="scope.row.status==-1" style="color: #F56C6C;">禁用</span>
				</template>
			</el-table-column>
			<el-table-column label="创建时间" width="180" align="center">
				<template #default="scope">
					{{new Date(scope.row.ctime).toLocaleString()}}
				</template>
			</el-table-column>
			<el-table-column fixed="right" label="操作">
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
		<el-pagination class="page-pagination" 
		v-model:page-size="limit" 
		v-model:current-page="page" 
		:page-sizes="[10,20, 50, 100, 200]" 
		layout="total,sizes,prev,pager,next" :total="total"/>
	</div>
	<!--弹出层-->
	<el-dialog v-model="addDialog" title="添加">
		<el-form :model="addForm" label-width="80px" ref="refAddForm">
			<el-form-item label="图片" prop="thumb" required>
				<uploadImg v-model="addForm.thumb"></uploadImg>
			</el-form-item>
			<el-form-item label="分类" prop="sid" required>
				<el-select v-model="addForm.sid" placeholder="请选择">
					<el-option v-for="item in sort" :value="item.id" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item label="状态" prop="status" required>
				<el-radio v-model="addForm.status" label="启用" :value="1" />
				<el-radio v-model="addForm.status" label="禁用" :value="-1" />				
			</el-form-item>
			<el-form-item label="排序">
				<el-input type="number" v-model="addForm.sort"></el-input>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="add">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
	<el-dialog v-model="editDialog" title="编辑">		
		<el-form :model="editForm" ref="refEditForm" label-width="80px">
			<el-form-item label="封面" prop="thumb" required>
				<uploadImg v-model="editForm.thumb"></uploadImg>
			</el-form-item>
			<el-form-item label="分类" prop="sid" required>
				<el-select v-model="editForm.sid" placeholder="请选择">
					<el-option v-for="item in sort" :value="item.id" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item label="状态" prop="status" required>
				<el-radio v-model="editForm.status" label="启用" :value="1" />
				<el-radio v-model="editForm.status" label="禁用" :value="-1" />				
			</el-form-item>
			<el-form-item label="排序">
				<el-input type="number" v-model="editForm.sort"></el-input>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="editSub">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
</template>

<script setup>
	import {ref,onMounted,watch,inject} from "vue";
	import uploadImg from "../../components/upload/uploadImg.vue"
	import {rfcToTime} from "../../config/functions"
	const myApi=inject('myApi');
	const ElMessage=inject('ElMessage');
	const list=ref([]);
	const sort=ref()
	const addDialog=ref(false);
	const editDialog=ref(false);
	const addForm=ref({status:1,sort:0});
	const editForm=ref({});
	const refAddForm=ref()
	const refEditForm=ref()
	const searchForm=ref({});
	const page=ref(1);
	const limit=ref(10);
	const total=ref(0);
	onMounted(()=>{
		getList()
		getSortList()
	});
	function refresh(){
		searchForm.value={},
		getList()
		getSortList()
	}
	function getList(){
		searchForm.value.page=page.value
		searchForm.value.limit=limit.value
		myApi('imgList',searchForm.value,'post').then((res)=>{
			list.value=res.data.data;
			total.value=res.data.total;
		})
	}
	function getSortList(){
		myApi('imgSort',{page:1,limit:1000}).then((res)=>{
			sort.value=res.data.data;
		})
	}
	function add(){
		refAddForm.value.validate(valid=>{
			if(valid){
				let data={...addForm.value}
				data.sort=Number(data.sort)
				myApi('addImg',data,'post').then((res)=>{
					addDialog.value=false
					ElMessage({message:res.msg,type:'success'})
					getList()
					refAddForm.value.resetFields()
				})
			}
		})
		
	}
	function edit(scope){
		editDialog.value=true
		editForm.value={...scope.row}
	}
	function editSub(){
		refEditForm.value.validate(valid=>{
			if(valid){
				let data={...editForm.value}
				data.sort=Number(data.sort)
				myApi('editImg',data,'post').then((res)=>{
					editDialog.value=false
					ElMessage({message:res.msg,type:'success'})
					getList()
				})
			}
		})
		
	}
	function del(scope){
		myApi('delImg',{id:scope.row.id},'post').then((res)=>{
			ElMessage({message:res.msg,type:'success'})
			getList()
		})
	}
	watch([page,limit],()=>{
		getList()
	});
</script>

<style>
	
</style>