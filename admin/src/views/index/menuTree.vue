
<template>
    <template v-if="props.item!=undefined">
        <template v-if="props.item.children">
            <el-sub-menu :index="props.item.id">
                <template #title>
                    <el-icon>
                        <component :is="props.item.icon" />
                    </el-icon>
                    <span>{{props.item.name}}</span>
                </template>
                <menuTree v-for="itemChildren in props.item.children" :item="itemChildren"></menuTree>
            </el-sub-menu>
        </template>
                    
        <el-menu-item v-else @click="toPath(item)" :index="props.item.id">
            <el-icon>
                <component :is="props.item.icon" />
            </el-icon>
            <span>{{props.item.name}}</span>
        </el-menu-item>
    </template>
      
</template>
												
<script setup>
import {inject} from "vue";
const props= defineProps({item: Object})
const store = inject('store');
function toPath(item){
    if(item!=undefined){
        store.treeMenuClick=item 
    }      
}
</script>  