import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from './router'
import App from './App.vue'

const app = createApp(App)


// 预加载背景图片
const preloadBackgroundImage = new Image()
preloadBackgroundImage.src = '/background/ContactBackground.png'

// 预加载 currentprojects 目录下的图片
const preloadCurrentProjectsImages = [
    '/currentprojects/001.ppg',
    '/currentprojects/002.png',
    '/currentprojects/003.png',
    '/currentprojects/004.jpg',
    '/currentprojects/005.jpg',
    '/currentprojects/006.png',
    '/currentprojects/007.png',
    '/currentprojects/yangbaoyao.jpg',
];

preloadCurrentProjectsImages.forEach(src => {
    const img = new Image();
    img.src = src;
});

// 预加载 publications 目录下的图片
const preloadPublicationsImages = [
    '/publications/001.png',
    '/publications/002.png',
    '/publications/003.png',
    '/publications/004.png',
    '/publications/005.png',
    '/publications/006.png',
    '/publications/007.png',
    '/publications/008.png',
    '/publications/009.png',
];

preloadPublicationsImages.forEach(src => {
    const img = new Image();
    img.src = src;
});


app.use(ElementPlus)
app.use(router)
app.mount('#app')