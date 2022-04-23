import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const routes = [
    {
        path: '/',
        name: 'home',
        component: HomeView
    },
    {
        path: '/about',
        name: 'about',
        component: () => import('../views/AboutView.vue')
    },
    {
        path: '/hoge',
        name: 'hoge',
        component: () => import('../views/HogeView.vue')
    },
    {
        path: '/sync',
        name: 'sync',
        component: () => import('../views/SyncView')
    },
    {
        path: '/sync2',
        name: 'sync2',
        component: () => import('../views/SyncView2')
    },
    {
        path: '/watch',
        name: 'watch',
        component: () => import('../views/WatchView')
    }
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router
