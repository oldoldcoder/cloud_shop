<template>
  <div class="auth-page">
    <el-card class="auth-card">
      <h2 class="title">找回密码</h2>
      <el-form :model="form" :rules="rules" ref="formRef" label-width="80px" @submit.prevent>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="form.email" placeholder="请输入注册邮箱" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" native-type="button" :loading="loading" @click="onSubmit" style="width:100%">发送重置邮件</el-button>
        </el-form-item>
        <div class="actions">
          <el-button link type="primary" @click="goLogin">返回登录</el-button>
        </div>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { forgotPassword as forgotPasswordApi } from '@/api/user'

export default {
  name: 'ForgotPassword',
  setup() {
    const router = useRouter()
    const formRef = ref(null)
    const loading = ref(false)
    const form = ref({ email: '' })

    const rules = {
      email: [
        { required: true, message: '请输入邮箱', trigger: 'blur' },
        { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
      ]
    }

    const onSubmit = () => {
      formRef.value.validate(async (valid, fields) => {
        if (!valid) {
          console.log('忘记密码表单验证失败:', fields)
          // 显示第一个验证错误
          const firstError = Object.values(fields)[0]
          if (firstError && firstError.length > 0) {
            ElMessage.error(firstError[0].message)
          }
          return
        }
        loading.value = true
        try {
          const response = await forgotPasswordApi(form.value)
          console.log('忘记密码响应:', response)
          
          if (response.data && response.data.code === 200) {
            ElMessage.success('验证码已发送，请前往重置密码')
            router.push({ path: '/reset-password', query: { email: form.value.email } })
          } else {
            ElMessage.error(response.data?.message || '发送失败')
          }
        } catch (e) {
          console.error('忘记密码错误:', e)
          ElMessage.error(e.response?.data?.message || '发送失败，请稍后重试')
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
