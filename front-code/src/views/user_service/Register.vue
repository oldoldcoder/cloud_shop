<template>
  <div class="auth-page">
    <el-card class="auth-card">
      <h2 class="title">注册</h2>
      <el-form :model="form" :rules="rules" ref="formRef" label-width="80px" @submit.prevent>
        <el-form-item label="用户名" prop="username">
          <el-input v-model="form.username" placeholder="请输入用户名" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="form.email" placeholder="请输入邮箱" />
        </el-form-item>
        <el-form-item label="手机号" prop="phone">
          <el-input v-model="form.phone" placeholder="请输入手机号（可选）" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="form.password" type="password" show-password placeholder="请输入密码" />
        </el-form-item>
        <el-form-item label="确认密码" prop="confirmPassword">
          <el-input v-model="form.confirmPassword" type="password" show-password placeholder="请再次输入密码" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" native-type="button" :loading="loading" @click="onSubmit" style="width:100%">注册</el-button>
        </el-form-item>
        <div class="actions">
          <el-button link type="primary" @click="goLogin">已有账号？去登录</el-button>
        </div>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { register as registerApi } from '@/api/user'

export default {
  name: 'Register',
  setup() {
    const router = useRouter()
    const formRef = ref(null)
    const loading = ref(false)
    const form = ref({ username: '', email: '', phone: '', password: '', confirmPassword: '' })

    const validateConfirmPassword = (rule, value, callback) => {
      if (value !== form.value.password) {
        callback(new Error('两次输入的密码不一致'))
      } else {
        callback()
      }
    }

    const rules = {
      username: [
        { required: true, message: '请输入用户名', trigger: 'blur' },
        { min: 3, max: 20, message: '用户名长度在3到20个字符', trigger: 'blur' }
      ],
      email: [
        { required: true, message: '请输入邮箱', trigger: 'blur' },
        { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
      ],
      phone: [
        { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号', trigger: 'blur' }
      ],
      password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
        { min: 6, max: 20, message: '密码长度在6到20个字符', trigger: 'blur' }
      ],
      confirmPassword: [
        { required: true, message: '请确认密码', trigger: 'blur' },
        { validator: validateConfirmPassword, trigger: 'blur' }
      ]
    }

    const onSubmit = () => {
      formRef.value.validate(async (valid, fields) => {
        if (!valid) {
          console.log('表单验证失败:', fields)
          // 显示第一个验证错误
          const firstError = Object.values(fields)[0]
          if (firstError && firstError.length > 0) {
            ElMessage.error(firstError[0].message)
          }
          return
        }
        if (form.value.password !== form.value.confirmPassword) {
          ElMessage.error('两次输入的密码不一致')
          return
        }
        loading.value = true
        try {
          const payload = { ...form.value }
          console.log('发送注册请求:', payload)
          const response = await registerApi(payload)
          console.log('注册响应:', response)
          
          if (response.data && Number(response.data.code) === 200) {
            ElMessage.success('注册成功，请登录')
            router.push('/login')
          } else {
            ElMessage.error(response.data?.message || '注册失败')
          }
        } catch (e) {
          console.error('注册错误:', e)
          if (e.code === 'ECONNABORTED' || e.message?.includes('timeout')) {
            ElMessage.error('请求超时，请稍后再试')
          } else if (e.message === 'canceled') {
            ElMessage.error('请求被取消，请重试')
          } else {
            ElMessage.error(e.response?.data?.message || '注册失败，请稍后重试')
          }
        } finally {
          loading.value = false
        }
      })
    }

    const goLogin = () => router.push('/login')

    return { formRef, form, rules, loading, onSubmit, goLogin }
  }
}
</script>

<style scoped>
.auth-page { min-height: 100vh; display:flex; justify-content:center; align-items:center; background:#f5f7fa; }
.auth-card { width: 480px; padding: 10px 10px 0 10px; }
.title { text-align:center; margin-bottom: 20px; }
.actions { display:flex; justify-content: space-between; }
</style>
