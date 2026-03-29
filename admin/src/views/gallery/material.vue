
<template>
    <el-container>
        <el-aside width="200px">
            <el-menu>
                <el-menu-item :index="0">全部</el-menu-item>
                <el-menu-item v-for="item in sort" :index="item.id">{{item.name}}</el-menu-item>
            </el-menu>
        </el-aside>
        <el-container>
            <el-header>
                <el-button>上传</el-button>
                <el-button type="primary">使用</el-button>
                <el-button type="danger">删除</el-button>
            </el-header>
            <el-main>Main</el-main>
            <el-footer>Footer</el-footer>
        </el-container>
    </el-container>
</template>
<script setup>
    import {ref,onMounted,watch,inject} from "vue";
	import uploadImg from "../../components/upload/uploadImg.vue"
    const myApi=inject('myApi');
	const ElMessage=inject('ElMessage');
	const list=ref([]);
	const sort=ref()
    const page=ref(1);
	const limit=ref(10);
	const total=ref(0);
    onMounted(()=>{
		getList()
		getSortList()
	});
    function getList(){
		searchForm.value.page=page.value
		searchForm.value.limit=limit.value
		myApi('imgList',searchForm.value,'post').then((res)=>{
			list.value=res.data.data;
			total.value=res.data.total;
		})
	}
	function getSortList(){
		myApi('imgSort',{page:1,limit:100},'post').then((res)=>{
			sort.value=res.data.data;
		})
	}
</script>