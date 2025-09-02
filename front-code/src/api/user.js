import axios from 'axios'

const prefix = '/api/users'

// 创建 axios 实例 - 使用 Vite 代理，不设置 baseURL
const api = axios.create({
  timeout: 10000
})

// 请求拦截器：添加 JWT 令牌
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('accessToken')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器：处理令牌过期
api.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response?.status === 401) {
      // 令牌过期，尝试刷新
      try {
        const refreshToken = localStorage.getItem('refreshToken')
        if (refreshToken) {
          const response = await api.post(`${prefix}/token/refresh`, { refreshToken })
          if (response.data.code === 200) {
            localStorage.setItem('accessToken', response.data.data.accessToken)
            localStorage.setItem('refreshToken', response.data.data.refreshToken)
            
            // 重试原请求
            error.config.headers.Authorization = `Bearer ${response.data.data.accessToken}`
            return api.request(error.config)
          }
        }
      } catch (refreshError) {
        // 刷新失败，清除令牌
        localStorage.removeItem('accessToken')
        localStorage.removeItem('refreshToken')
        localStorage.removeItem('userInfo')
        // 可以在这里添加路由跳转到登录页
      }
    }
    return Promise.reject(error)
  }
)

// 用户认证相关 API
export const register = (payload) => api.post(`${prefix}/register`, payload)
export const login = (payload) => api.post(`${prefix}/login`, payload)
export const logout = () => api.post(`${prefix}/logout`)

// 密码管理
export const forgotPassword = (payload) => api.post(`${prefix}/password/forgot`, payload)
export const resetPassword = (payload) => api.post(`${prefix}/password/reset`, payload)
export const changePassword = (payload) => api.post(`${prefix}/password/change`, payload)

// 用户信息管理
export const getUserInfo = () => api.get(`${prefix}/me`)
export const updateUserInfo = (payload) => api.put(`${prefix}/me`, payload)
export const refreshToken = (payload) => api.post(`${prefix}/token/refresh`, payload)

// 用户管理（管理员）
export const getUserList = (params) => api.get(`${prefix}/list`, { params })
export const getUserById = (id) => api.get(`${prefix}/list`, { params })
export const updateUser = (id, payload) => api.put(`${prefix}/${id}`, payload)
export const deleteUser = (id) => api.delete(`${prefix}/${id}`)
