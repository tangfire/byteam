import {createRouter, createWebHashHistory} from 'vue-router'


import HomeView from '../components/HomeView.vue'

import CurrentProjectsView from "../components/CurrentProjectsView.vue";
import DrBaoyaoYangView from "../components/DrBaoyaoYangView.vue";
import ContactView from "../components/ContactView.vue";
import InternationalJournalsConferencesView from "../components/InternationalJournalsConferencesView.vue";
import GroupView from "../components/GroupView.vue";


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
    },
    {
      path:'/international-journals-conferences',
      name:'international-journals-conferences',
      component:InternationalJournalsConferencesView
    },
    {
      path:'/our-group',
      name:"our-group",
      component:GroupView
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
