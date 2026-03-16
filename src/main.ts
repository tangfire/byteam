import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from './router'
import App from './App.vue'

// 图片预加载：仅预加载关键图片（如logo），其余懒加载
const preloadImages = () => {
    const images = ['/logo/001.png']
    images.forEach(src => {
        const img = new Image()
        img.src = src
    })
}

preloadImages()

const app = createApp(App)

app.use(ElementPlus)
app.use(router)
app.mount('#app')