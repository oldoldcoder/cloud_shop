/**
 * 认证相关工具函数
 */

/**
 * 检查用户是否已认证
 * @returns {boolean}
 */
export const isAuthenticated = () => {
  const token = localStorage.getItem('accessToken')
  return !!token
}

/**
 * 获取访问令牌
 * @returns {string|null}
 */
export const getAccessToken = () => {
  return localStorage.getItem('accessToken')
}

/**
 * 获取刷新令牌
 * @returns {string|null}
 */
export const getRefreshToken = () => {
  return localStorage.getItem('refreshToken')
}

/**
 * 获取用户信息
 * @returns {object|null}
 */
export const getUserInfo = () => {
  const userInfo = localStorage.getItem('userInfo')
  return userInfo ? JSON.parse(userInfo) : null
}

/**
 * 设置认证信息
 * @param {string} accessToken 访问令牌
 * @param {string} refreshToken 刷新令牌
 * @param {object} userInfo 用户信息
 */
export const setAuthInfo = (accessToken, refreshToken, userInfo) => {
  localStorage.setItem('accessToken', accessToken || '')
  localStorage.setItem('refreshToken', refreshToken || '')
  localStorage.setItem('userInfo', JSON.stringify(userInfo || {}))
}

/**
 * 清除认证信息
 */
export const clearAuthInfo = () => {
  localStorage.removeItem('accessToken')
  localStorage.removeItem('refreshToken')
  localStorage.removeItem('userInfo')
}

/**
 * 检查令牌是否过期（简单检查，实际应该解析JWT）
 * @returns {boolean}
 */
export const isTokenExpired = () => {
  const token = getAccessToken()
  if (!token) return true
  
  // 这里可以添加JWT解析逻辑来检查过期时间
  // 目前简单返回false，实际项目中应该解析JWT的exp字段
  return false
}
