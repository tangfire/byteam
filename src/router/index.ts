import {createRouter, createWebHashHistory} from 'vue-router'


import HomeView from '../pages/HomeView.vue'

import ResearchDirectionView from "../pages/ResearchDirectionView.vue";
import DrBaoyaoYangView from "../pages/DrBaoyaoYangView.vue";
import ContactView from "../pages/ContactView.vue";
import InternationalJournalsConferencesView from "../pages/InternationalJournalsConferencesView.vue";
import GroupView from "../pages/GroupView.vue";
import PatentsView from '../pages/PatentViews.vue'
import ResearchProjectsView from "../pages/ResearchProjectsView.vue";
import NewsView from "../pages/NewsView.vue";


const router = createRouter({
    history: createWebHashHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView,
        },
        // {
        //   path:'/research',
        //   name:'research',
        //   component: ResearchView,
        // },
        {
            path: '/research-direction',
            name: 'research-direction',
            component: ResearchDirectionView
        },
        {
            path: '/dr-Baoyao-Yang',
            name: 'dr-Baoyao-Yang',
            component: DrBaoyaoYangView
        },
        {
            path: '/contact',
            name: 'contact',
            component: ContactView
        },
        {
            path: '/international-journals-conferences',
            name: 'international-journals-conferences',
            component: InternationalJournalsConferencesView
        },
        {
            path: '/our-group',
            name: "our-group",
            component: GroupView
        },
        {
            path: '/patents',
            name: 'patents',
            component: PatentsView
        },
        {
            path: '/research-projects',
            name: 'research-projects',
            component: ResearchProjectsView

        },
        {
            path:'/news',
            name:'news',
            component: NewsView
        }



        // {
        //   path: '/about',
        //   name: 'about',
        //   // route level code-splitting
        //   // this generates a separate chunk (About.[hash].js) for this route
        //   // which is lazy-loaded when the route is visited.
        //   component: () => import('../views/AboutView.vue'),
        // },
    ],
})

export default router
