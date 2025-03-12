import {createRouter, createWebHashHistory} from 'vue-router'


import HomeView from '../components/HomeView.vue'

import ResearchDirectionView from "../components/ResearchDirectionView.vue";
import DrBaoyaoYangView from "../components/DrBaoyaoYangView.vue";
import ContactView from "../components/ContactView.vue";
import InternationalJournalsConferencesView from "../components/InternationalJournalsConferencesView.vue";
import GroupView from "../components/GroupView.vue";
import PatentsView from '../components/PatentViews.vue'
import ResearchProjectsView from "../components/ResearchProjectsView.vue";
import NewsView from "../components/NewsView.vue";


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
