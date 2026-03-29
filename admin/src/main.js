import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import {request} from './config/api.js';
import { ElMessage } from 'element-plus';
import store from './store/index.js';
import './assets/iconfont/iconfont.css'
const app=createApp(App)
app.use(router)
app.use(ElementPlus,{locale: zhCn})
app.mount('#app')

app.provide('myApi',request)
app.provide('router',router)
app.provide('store',store)
app.provide('ElMessage',ElMessage)
// 注册所有图标组件
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}