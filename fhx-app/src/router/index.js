import { createRouter, createWebHistory } from 'vue-router'
import OtsBody from '../components/ots-body.vue'
import Dropzone from '../components/download/dropzone.vue'

const routes = [
    {
        path: "/",
        name: "Home",
        component: OtsBody,
    },
    {
        path: "/download",
        name: "Download",
        component: Dropzone,
    },
]

const router = createRouter({history: createWebHistory(), routes})

export default router