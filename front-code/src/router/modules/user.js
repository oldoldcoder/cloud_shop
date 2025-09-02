import Login from '@/views/user_service/Login.vue'
import Register from '@/views/user_service/Register.vue'
import ForgotPassword from '@/views/user_service/ForgotPassword.vue'

export default [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  },
  {
    path: '/forgot-password',
    name: 'ForgotPassword',
    component: ForgotPassword
  }
]
