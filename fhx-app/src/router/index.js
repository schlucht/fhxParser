import { createRouter, createWebHistory } from 'vue-router'
import OtsBody from '../components/ots-body.vue'

const routes = [
    {
        path: "/",
        name: "Home",
        component: OtsBody
    }
]

const router = createRouter({history: createWebHistory(), routes})

export default router