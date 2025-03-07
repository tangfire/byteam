import { createRouter, createWebHistory } from 'vue-router'


import HomeView from '../components/HomeView.vue'

import CurrentProjectsView from "../components/CurrentProjectsView.vue";
import DrBaoyaoYangView from "../components/DrBaoyaoYangView.vue";
import ContactView from "../components/ContactView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
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
      path:'/current-projects',
      name:'current-projects',
      component: CurrentProjectsView
    },
    {
      path:'/dr-Baoyao-Yang',
      name:'dr-Baoyao-Yang',
      component:DrBaoyaoYangView
    },
    {
      path:'/contact',
      name:'contact',
      component:ContactView
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
