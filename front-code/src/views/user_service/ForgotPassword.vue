<template>
  <div class="auth-page">
    <el-card class="auth-card">
      <h2 class="title">找回密码</h2>
      <el-form :model="form" label-width="80px">
        <el-form-item label="邮箱">
          <el-input v-model="form.email" placeholder="请输入注册邮箱" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit" style="width:100%">发送重置邮件</el-button>
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

export default {
  name: 'ForgotPassword',
  setup() {
    const router = useRouter()
    const form = ref({ email: '' })

    const onSubmit = async () => {
      // TODO: 调后端忘记密码接口
      await new Promise(r => setTimeout(r, 500))
      ElMessage.success('重置邮件已发送，请查收')
      router.push('/login')
    }

    const goLogin = () => router.push('/login')

    return { form, onSubmit, goLogin }
  }
}
</script>

<style scoped>
.auth-page { min-height: 100vh; display:flex; justify-content:center; align-items:center; background:#f5f7fa; }
.auth-card { width: 480px; padding: 10px 10px 0 10px; }
.title { text-align:center; margin-bottom: 20px; }
.actions { display:flex; justify-content: space-between; }
</style>
