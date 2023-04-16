import { createRouter, createWebHistory } from 'vue-router'
import OtsBody from '../components/ots-body.vue'
import Download from '../components/download/download.vue'

const routes = [
    {
        path: "/",
        name: "Home",
        component: OtsBody,
    },
    {
        path: "/download",
        name: "Download",
        component: Download,
    },
]

const router = createRouter({history: createWebHistory(), routes})

export default router