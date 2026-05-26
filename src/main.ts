import { createApp } from 'vue'
import router from './router'
import App from './App.vue'
import { installElementPlus } from './plugins/element-plus'

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

installElementPlus(app)
app.use(router)
app.mount('#app')
