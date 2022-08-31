import { createRouter, createWebHistory } from 'vue-router'

import Home from '@/containers/Home.vue'
import Keys from '@/containers/Keys.vue'
import Login from '@/containers/Login.vue'
import Recover from '@/containers/Recover.vue'
import Register from '@/containers/Register.vue'
import Send from '@/containers/Send.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/home', name: 'home', component: Home },
    { path: '/keys', name: 'keys', component: Keys },
    { path: '/', name: 'login', component: Login },
    { path: '/recover', name: 'recover', component: Recover },
    { path: '/register', name: 'register', component: Register },
    { path: '/send', name: 'send', component: Send },
  ],
})

export default router
