import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from './router'
import App from './App.vue'

// 导入自动生成的图片列表
import preloadImageList from './preload-images.json'

const app = createApp(App)

// 预加载所有图片的函数
const preloadImages = () => {
    preloadImageList.forEach(src => {
        const img = new Image()
        img.src = src
    })
}

// 调用预加载函数
preloadImages()

app.use(ElementPlus)
app.use(router)
app.mount('#app')