import axios from 'axios'

const prefix = '/api/users'

export const register = (payload) => axios.post(`${prefix}/register`, payload)
export const login = (payload) => axios.post(`${prefix}/login`, payload)
export const forgotPassword = (payload) => axios.post(`${prefix}/password/forgot`, payload)
