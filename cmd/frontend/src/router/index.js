import { createRouter, createWebHistory } from 'vue-router'

import NotFound from '../views/NotFound.vue'
import FhxHome from '@/components/home/FhxHome.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/', 
      name: 'start',
      component: FhxHome,
    },
    {
      path: '/home',
      name: 'home',
      component: ()=>import('../views/HomeView.vue'),      
    },
    {
      path: '/operation',
      name: 'operation',
      component: ()=>import('../views/OperationView.vue'),
      children: [
        {
          path: ':id?',
          name: 'detail',
          component: ()=>import('../components/operation/FhxOperationDetail.vue'),
        },
      ]
    },
    {
      path: '/unit',
      name: 'unit',
      component: ()=>import('../views/UnitView.vue'),
    },   
    {
      path: '/recipe',
      name: 'recipe',
      component: ()=>import('../views/RecipeView.vue'),
    },
    {path: '/:notfound(.*)', component: NotFound}    
  ]
})

export default router
