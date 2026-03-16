import { createRouter, createWebHashHistory } from 'vue-router'

const router = createRouter({
    history: createWebHashHistory(import.meta.env.BASE_URL),
    routes: [
        { path: '/', name: 'home', component: () => import('../pages/HomeView.vue') },
        { path: '/about', name: 'about', component: () => import('../pages/AboutView.vue') },
        { path: '/news', name: 'news', component: () => import('../pages/NewsView.vue') },
        { path: '/research-direction', name: 'research-direction', component: () => import('../pages/ResearchDirectionView.vue') },
        { path: '/research-projects', name: 'research-projects', component: () => import('../pages/ResearchProjectsView.vue') },
        { path: '/dr-Baoyao-Yang', name: 'dr-Baoyao-Yang', component: () => import('../pages/DrBaoyaoYangView.vue') },
        { path: '/our-group', name: 'our-group', component: () => import('../pages/GroupView.vue') },
        { path: '/Alumni', name: 'alumni', component: () => import('../pages/AlumniView.vue') },
        { path: '/international-journals-conferences', name: 'publications', component: () => import('../pages/InternationalJournalsConferencesView.vue') },
        { path: '/patents', name: 'patents', component: () => import('../pages/PatentViews.vue') },
        { path: '/contact', name: 'contact', component: () => import('../pages/ContactView.vue') },
        { path: '/vKnow', name: 'vKnow', component: () => import('../pages/VknowView.vue') },
        { path: '/VideoMind', name: 'VideoMind', component: () => import('../pages/VideoMindView.vue') },
        { path: '/video-player-XiaoqiZheng01', name: 'video-player-XiaoqiZheng01', component: () => import('../pages/VideoPlayerXiaoqiZheng01View.vue') },
        { path: '/video-player-XianrunXu01', name: 'video-player-XianrunXu01', component: () => import('../pages/VideoPlayerXianrunXu01View.vue') },
        { path: '/video-player-YaliMa01', name: 'video-player-YaliMa01', component: () => import('../pages/VideoPlayerYaliMa01View.vue') }
    ]
})

export default router