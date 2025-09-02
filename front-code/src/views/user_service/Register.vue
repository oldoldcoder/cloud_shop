<template>
  <div class="auth-page">
    <el-card class="auth-card">
      <h2 class="title">注册</h2>
      <el-form :model="form" :rules="rules" ref="formRef" label-width="80px">
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
          <el-button type="primary" :loading="loading" @click="onSubmit" style="width:100%">注册</el-button>
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

    const rules = {
      username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
      email: [{ required: true, message: '请输入邮箱', trigger: 'blur' }],
      password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
      confirmPassword: [{ required: true, message: '请确认密码', trigger: 'blur' }]
    }

    const onSubmit = () => {
      formRef.value.validate(async (valid) => {
        if (!valid) return
        if (form.value.password !== form.value.confirmPassword) {
          ElMessage.error('两次输入的密码不一致')
          return
        }
        loading.value = true
        try {
          const payload = { ...form.value }
          const { data } = await registerApi(payload)
          if (data.code === 200) {
            ElMessage.success('注册成功，请登录')
            router.push('/login')
          } else {
            ElMessage.error(data.message || '注册失败')
          }
        } catch (e) {
          ElMessage.error('注册失败')
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
