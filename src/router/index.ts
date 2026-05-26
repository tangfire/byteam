import { createRouter, createWebHashHistory } from 'vue-router'
import { authStore } from '../api/client'

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
        { path: '/undergraduate', name: 'undergraduate', component: () => import('../pages/UndergraduateView.vue') },
        { path: '/international-journals-conferences', name: 'publications', component: () => import('../pages/InternationalJournalsConferencesView.vue') },
        { path: '/patents', name: 'patents', component: () => import('../pages/PatentViews.vue') },
        { path: '/contact', name: 'contact', component: () => import('../pages/ContactView.vue') },
        { path: '/vKnow', name: 'vKnow', component: () => import('../pages/VknowView.vue') },
        { path: '/VideoMind', name: 'VideoMind', component: () => import('../pages/VideoMindView.vue') },
        { path: '/video/:slug', name: 'video-page', component: () => import('../pages/SiteVideoPageView.vue') },
        { path: '/video-player-XiaoqiZheng01', name: 'video-player-XiaoqiZheng01', redirect: '/video/video-xiaoqi-zheng-01' },
        { path: '/video-player-XianrunXu01', name: 'video-player-XianrunXu01', redirect: '/video/video-xianrun-xu-01' },
        { path: '/video-player-YaliMa01', name: 'video-player-YaliMa01', redirect: '/video/video-yali-ma-01' },
        { path: '/admin/login', name: 'admin-login', component: () => import('../pages/admin/AdminLoginView.vue'), meta: { adminPublic: true } },
        {
            path: '/admin',
            component: () => import('../pages/admin/AdminLayoutView.vue'),
            meta: { requiresAdmin: true },
            children: [
                { path: '', redirect: '/admin/dashboard' },
                { path: 'dashboard', name: 'admin-dashboard', component: () => import('../pages/admin/AdminDashboardView.vue') },
                { path: 'pages', name: 'admin-pages', component: () => import('../pages/admin/AdminPagesView.vue') },
                { path: 'news', name: 'admin-news', component: () => import('../pages/admin/AdminNewsView.vue') },
                { path: 'people', name: 'admin-people', component: () => import('../pages/admin/AdminPeopleView.vue') },
                { path: 'undergraduates', name: 'admin-undergraduates', component: () => import('../pages/admin/AdminUndergraduatesView.vue') },
                { path: 'publications', name: 'admin-publications', component: () => import('../pages/admin/AdminPublicationsView.vue') },
                { path: 'patents', name: 'admin-patents', component: () => import('../pages/admin/AdminPatentsView.vue') },
                { path: 'research-projects', name: 'admin-research-projects', component: () => import('../pages/admin/AdminResearchProjectsView.vue') },
                { path: 'media', name: 'admin-media', component: () => import('../pages/admin/AdminMediaView.vue') },
                { path: 'trash', name: 'admin-trash', component: () => import('../pages/admin/AdminTrashView.vue') },
                { path: 'guide', name: 'admin-guide', component: () => import('../pages/admin/AdminGuideView.vue') }
            ]
        }
    ]
})

router.beforeEach((to) => {
    if (to.meta.requiresAdmin && !authStore.token) {
        return { name: 'admin-login', query: { redirect: to.fullPath } }
    }
    if (to.name === 'admin-login' && authStore.token) {
        return { name: 'admin-dashboard' }
    }
})

export default router
