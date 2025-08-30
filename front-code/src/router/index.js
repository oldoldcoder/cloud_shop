import { createRouter, createWebHistory } from 'vue-router'
import HelloWorld from '@/views/HelloWorld.vue'
import userServiceRoutes from './modules/user_service'

const routes = [
  {
    path: '/',
    name: 'Index',
    component: HelloWorld
  },
  ...userServiceRoutes,
  {
    path: '/:pathMatch(.*)*',
    redirect: '/login'
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  if (to.path === '/') {
    if (!from.name) {
      return next('/login')
    }
  }
  next()
})

export default router
