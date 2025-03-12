import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from './router'
import App from './App.vue'

const app = createApp(App)
const preloadImage = new Image()
preloadImage.src = '/background/background.png'
app.use(ElementPlus)
app.use(router)
app.mount('#app')