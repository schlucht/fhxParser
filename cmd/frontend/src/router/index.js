import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import FhxOperationDetail from '../components/operation/FhxOperationDetail.vue'
import NotFound from '../views/NotFound.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {path: '/', redirect: '/home'},
    {
      path: '/home',
      name: 'home',
      component: HomeView,
      children: [
        {
          path: ':id?',
          name: 'detail',
          component: FhxOperationDetail
        },
      ]
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
