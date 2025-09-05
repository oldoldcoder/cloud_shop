<template>
  <div class="auth-page">
    <el-card class="auth-card">
      <h2 class="title">重置密码</h2>
      <el-form :model="form" :rules="rules" ref="formRef" label-width="100px" @submit.prevent>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="form.email" placeholder="请输入邮箱" clearable />
        </el-form-item>
        <el-form-item label="验证码" prop="verificationCode">
          <el-input v-model="form.verificationCode" placeholder="请输入邮箱验证码" clearable />
        </el-form-item>
        <el-form-item label="新密码" prop="newPassword">
          <el-input v-model="form.newPassword" type="password" show-password placeholder="请输入新密码" />
        </el-form-item>
        <el-form-item label="确认密码" prop="confirmPassword">
          <el-input v-model="form.confirmPassword" type="password" show-password placeholder="请再次输入新密码" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" native-type="button" :loading="loading" @click="onSubmit" style="width:100%">提交</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
  </template>

<script>
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { ElMessage } from 'element-plus'
import { resetPassword as resetPasswordApi, verifyVerificationCode as verifyCodeApi } from '@/api/user'

export default {
  name: 'ResetPassword',
  setup() {
    const router = useRouter()
    const route = useRoute()
    const formRef = ref(null)
    const loading = ref(false)
    const form = ref({
      email: route.query.email || '',
      verificationCode: '',
      newPassword: '',
      confirmPassword: ''
    })

    const validateConfirm = (rule, value, callback) => {
      if (value !== form.value.newPassword) callback(new Error('两次输入的密码不一致'))
      else callback()
    }

    const rules = {
      email: [{ required: true, message: '请输入邮箱', trigger: 'blur' }, { type: 'email', message: '邮箱格式不正确', trigger: 'blur' }],
      verificationCode: [{ required: true, message: '请输入验证码', trigger: 'blur' }],
      newPassword: [{ required: true, message: '请输入新密码', trigger: 'blur' }, { min: 6, max: 20, message: '密码长度在6到20个字符', trigger: 'blur' }],
      confirmPassword: [{ required: true, message: '请确认密码', trigger: 'blur' }, { validator: validateConfirm, trigger: 'blur' }]
    }

    const onSubmit = () => {
      formRef.value.validate(async (valid, fields) => {
        if (!valid) {
          const first = Object.values(fields)[0]
          if (first && first.length) ElMessage.error(first[0].message)
          return
        }
        loading.value = true
        try {
          // 先校验验证码
          const verifyRes = await verifyCodeApi({ email: form.value.email, verificationCode: form.value.verificationCode })
          if (!verifyRes.data || verifyRes.data.code !== 200 || verifyRes.data.data !== true) {
            ElMessage.error(verifyRes.data?.message || '验证码校验失败')
            return
          }
          // 再重置密码（后端当前的 reset 接口示例为 token+newPassword，这里简化用 email 作为 token）
          const payload = { resetToken: form.value.email, newPassword: form.value.newPassword, confirmPassword: form.value.confirmPassword }
          const res = await resetPasswordApi(payload)
          if (res.data && res.data.code === 200) {
            ElMessage.success('密码重置成功，请登录')
            router.push('/login')
          } else {
            ElMessage.error(res.data?.message || '重置失败')
          }
        } catch (e) {
          ElMessage.error(e.response?.data?.message || '重置失败，请稍后再试')
        } finally {
          loading.value = false
        }
      })
    }

    return { formRef, form, rules, loading, onSubmit }
  }
}
</script>

<style scoped>
.auth-page { min-height: 100vh; display:flex; justify-content:center; align-items:center; background:#f5f7fa; }
.auth-card { width: 480px; padding: 10px 10px 0 10px; }
.title { text-align:center; margin-bottom: 20px; }
</style>


