<template>
  <div class="auth-page">
    <el-card class="auth-card">
      <h2 class="title">登录</h2>
      <el-form :model="form" :rules="rules" ref="formRef" label-width="80px">
        <el-form-item label="账号" prop="principal">
          <el-input v-model="form.principal" placeholder="用户名/邮箱/手机号" clearable />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="form.password" type="password" show-password placeholder="请输入密码" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="loading" @click="onSubmit" style="width: 100%">登录</el-button>
        </el-form-item>
        <div class="actions">
          <el-button link type="primary" @click="goRegister">没有账号？去注册</el-button>
          <el-button link type="primary" @click="goForgot">忘记密码？</el-button>
        </div>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { login as loginApi } from '@/api/user_service'

export default {
  name: 'Login',
  setup() {
    const router = useRouter()
    const formRef = ref(null)
    const loading = ref(false)
    const form = ref({ principal: '', password: '' })

    const rules = {
      principal: [{ required: true, message: '请输入账号', trigger: 'blur' }],
      password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
    }

    const onSubmit = () => {
      formRef.value.validate(async (valid) => {
        if (!valid) return
        loading.value = true
        try {
          const { data } = await loginApi(form.value)
          if (data.code === 200 && data.data) {
            const { accessToken, refreshToken, userInfo } = data.data
            localStorage.setItem('accessToken', accessToken || '')
            localStorage.setItem('refreshToken', refreshToken || '')
            localStorage.setItem('userInfo', JSON.stringify(userInfo || {}))
            ElMessage.success('登录成功')
            router.push('/')
          } else {
            ElMessage.error(data.message || '登录失败')
          }
        } catch (e) {
          ElMessage.error('登录失败')
        } finally {
          loading.value = false
        }
      })
    }

    const goRegister = () => router.push('/register')
    const goForgot = () => router.push('/forgot-password')

    return { formRef, form, rules, loading, onSubmit, goRegister, goForgot }
  }
}
</script>

<style scoped>
.auth-page {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #f5f7fa;
}

.auth-card {
  width: 420px;
  padding: 10px 10px 0 10px;
}

.title {
  text-align: center;
  margin-bottom: 20px;
}

.actions {
  display: flex;
  justify-content: space-between;
}
</style>
