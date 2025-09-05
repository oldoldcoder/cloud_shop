import { createRouter, createWebHistory } from 'vue-router'
import HelloWorld from '@/views/HelloWorld.vue'
import userServiceRoutes from './modules/user'
import { isAuthenticated } from '@/utils/auth'

const routes = [
  {
    path: '/',
    name: 'Index',
    component: HelloWorld,
    meta: { requiresAuth: true }  // 需要认证
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
  const isAuth = isAuthenticated()
  
  // 定义不需要认证的公开路由
  const publicRoutes = ['/login', '/register', '/forgot-password', '/reset-password']
  
  if (!isAuth && !publicRoutes.includes(to.path)) {
    // 未登录且访问需要认证的页面，重定向到登录页
    return next('/login')
  }
  
  if (isAuth && to.path === '/login') {
    // 已登录但访问登录页，重定向到首页
    return next('/')
  }
  
  next()
})

export default router
