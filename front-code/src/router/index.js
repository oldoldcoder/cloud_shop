import { createRouter, createWebHistory } from 'vue-router'
import HelloWorld from '../views/HelloWorld.vue'

const routes = [
  {
    path: '/',
    name: 'Hello',
    component: HelloWorld
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
