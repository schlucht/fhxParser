import { createRouter, createWebHistory } from 'vue-router'

import App from '@/App.vue';
import NotFound from '../views/NotFound.vue'
import OperationView from '../views/OperationView.vue'
import UnitView from '../views/UnitView.vue'

import FhxOperationDetail from '../components/operation/FhxOperationDetail.vue'
import FhxHome from '@/components/home/FhxHome.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/', 
      name: 'start',
      component: App,
    },
    {
      path: '/home',
      name: 'home',
      component: FhxHome,      
    },
    {
      path: '/operation',
      name: 'operation',
      component: OperationView,
      children: [
        {
          path: ':id?',
          name: 'detail',
          component: FhxOperationDetail
        },
      ]
    },
    {
      path: '/unit',
      name: 'unit',
      component: UnitView,
    },
    {path: '/:notfound(.*)', component: NotFound}
    // {
    //   path: '/about',
    //   name: 'about',
    //   // route level code-splitting
    //   // this generates a separate chunk (About.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () => import('../views/AboutView.vue')
    // }
  ]
})

export default router
