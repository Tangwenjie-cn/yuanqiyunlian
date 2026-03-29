
<template>
	<!--搜索栏-->
	<div class="page-header">
		<el-form :inline="true" :model="searchForm">
			
		</el-form>
		<el-button type="primary" icon="refresh" @click="refresh">刷新</el-button>
	</div>
	
	<!--按钮栏-->
	<div class="page-buttons">
		<el-button type="primary" @click="addDialog=true">添加</el-button>
	</div>
	<!--列表栏-->
	<div class="page-table">
		<el-table :data="list" stripe border table-layout="fixed" row-key="id" style="width: 100%">
			<el-table-column align="center" label="树形" width="80" />
		    <el-table-column align="center" prop="id" label="ID" width="80" />
			<el-table-column label="缩略图" width="120" align="center">
				<template #default="scope">
					<el-image :src="scope.row.thumb" fit="contain" style="width: 39px;height: 39px;">
						<template #error>
							<el-icon size="39"><Picture /></el-icon>
						</template>
					</el-image>
				</template>
			</el-table-column>
		    <el-table-column label="名称" prop="name" align="center" />
            <el-table-column label="创建时间" width="180" align="center">
				<template #default="scope">
					{{new Date(scope.row.ctime).toLocaleString()}}
				</template>
			</el-table-column>
            <el-table-column label="状态" align="center" width="80">
                <template #default="scope">
                    <el-text v-if="scope.row.status==1" type="success">正常</el-text>
                    <el-text v-if="scope.row.status==-1" type="danger">禁用</el-text>
                </template>                
            </el-table-column>
			<el-table-column fixed="right" width="150" label="操作">
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
	</div>
	<!--弹出层-->
	<el-dialog v-model="addDialog" title="添加">
		<el-form :model="addForm" ref="refAddForm" label-width="80px">
			<el-form-item label="上级分类" prop="pid" required>
				<el-cascader v-model="addForm.pid" :options="list" :props="props">
				</el-cascader>
			</el-form-item>
			<el-form-item label="名称" prop="name" required>
				<el-input v-model="addForm.name"></el-input>
			</el-form-item>
			<el-form-item label="图标" prop="thumb">
                <uploadImg v-model="addForm.thumb"></uploadImg>
            </el-form-item>
			<el-form-item label="排序" prop="sort" required>
				<el-input-number v-model="addForm.sort"></el-input-number>
			</el-form-item>
			<el-form-item label="状态" prop="status" required>
                <el-radio v-model="addForm.status" label="启用" :value="1"></el-radio>
                <el-radio v-model="addForm.status" label="禁用" :value="-1"></el-radio>				
		    </el-form-item>
			<el-form-item label-width="80px">
				<el-button type="primary" @click="add">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
	<el-dialog v-model="editDialog" title="编辑">
		<el-form :model="editForm" ref="refEditForm" label-width="80px">
			<el-form-item label="上级分类" prop="pid" required>
				<el-cascader v-model="editForm.pid" :options="list"  :props="props">
				</el-cascader>
			</el-form-item>
			<el-form-item label="名称" prop="name" required>
				<el-input v-model="editForm.name"></el-input>
			</el-form-item>
			<el-form-item label="图标" prop="thumb">
                <uploadImg v-model="editForm.thumb"></uploadImg>
            </el-form-item>
			<el-form-item label="排序" prop="sort" required>
				<el-input-number v-model="editForm.sort"></el-input-number>
			</el-form-item>
            <el-form-item label="状态" prop="status" required>
                <el-radio v-model="addForm.status" label="启用" :value="1"></el-radio>
                <el-radio v-model="addForm.status" label="禁用" :value="-1"></el-radio>				
		    </el-form-item>
			<el-form-item>
				<el-button type="primary" @click="editSub">提交</el-button>
			</el-form-item>	
		</el-form>
	</el-dialog>
</template>

<script setup>
	import {ref,onMounted,inject} from "vue";
	import uploadImg from "../../components/upload/uploadImg.vue"
	const myApi=inject('myApi');
	const ElMessage=inject('ElMessage');
	const list=ref([]);
	const addDialog=ref(false);
	const editDialog=ref(false);
	const addForm=ref({pid:0,status:1,sort:0});
	const editForm=ref({});
	const refAddForm=ref()
	const refEditForm=ref()
	const searchForm=ref({});
	const props=ref({
		checkStrictly: true,
		value: 'id',
    	label: 'name',
		emitPath: false,
	});
	onMounted(()=>{
		getList()
	});
	function refresh(){
		searchForm.value={},
		getList()
	}
	function getList(){
		myApi('goodsSort',searchForm.value,'post').then((res)=>{
			list.value=res.data
		})
	}
	function add(){
		refAddForm.value.validate(valid=>{
			if(valid){
				myApi('addGoodsSort',addForm.value,'post').then((res)=>{
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
		editForm.value=scope.row		
	}
	function editSub(){
		refEditForm.value.validate(valid=>{
			if(valid){
				myApi('upGoodsSort',editForm.value,'post').then((res)=>{
					editDialog.value=false
					ElMessage({message:res.msg,type:'success'})
					getList()
                    refEditForm.value.resetFields()
				})
			}
		})		
	}
	function del(scope){
		myApi('delGoodsSort',{id:scope.row.id},'post').then((res)=>{
			ElMessage({message:res.msg,type:'success'})
			getList()
		})
	}
	
</script>

<style>
	
</style>
