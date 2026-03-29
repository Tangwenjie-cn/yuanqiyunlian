import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
export default defineConfig({
  	plugins: [vue()],
  	server: {
  	    host: '127.0.0.1',//自定义主机名
  		port: 5200,//自定义端口
	},
})
