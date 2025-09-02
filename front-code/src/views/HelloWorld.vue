<template>
  <div class="hello-container">
    <el-card class="hello-card">
      <div class="hello-content">
        <h1 class="hello-title">Hello World!</h1>
        <p class="hello-subtitle">欢迎来到 Cloud Shop 云商城</p>
        
        <!-- 用户信息显示 -->
        <div v-if="userInfo" class="user-info">
          <el-avatar :size="64" :src="userInfo.avatar || ''">
            {{ userInfo.username?.charAt(0)?.toUpperCase() || 'U' }}
          </el-avatar>
          <h3>欢迎回来，{{ userInfo.username || userInfo.email || '用户' }}！</h3>
          <p>用户ID: {{ userInfo.id }}</p>
        </div>
        
        <el-button type="primary" size="large" @click="showMessage">
          点击我
        </el-button>
        
        <el-button type="danger" size="large" @click="handleLogout" style="margin-left: 10px;">
          退出登录
        </el-button>
        
        <div class="hello-info">
          <p>这是一个基于 Vue 3 + Element Plus 的前端项目模板</p>
          <p>当前时间: {{ currentTime }}</p>
          <p>认证状态: {{ isAuth ? '已登录' : '未登录' }}</p>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script>
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { isAuthenticated, getUserInfo, clearAuthInfo } from '@/utils/auth'

export default {
  name: 'HelloWorld',
  setup() {
    const router = useRouter()
    const currentTime = ref('')
    const userInfo = ref(null)
    const isAuth = ref(false)
    let timer = null

    const updateTime = () => {
      currentTime.value = new Date().toLocaleString('zh-CN')
    }

    const updateAuthStatus = () => {
      isAuth.value = isAuthenticated()
      userInfo.value = getUserInfo()
    }

    const showMessage = () => {
      ElMessage.success('Hello! 你点击了按钮！')
    }

    const handleLogout = async () => {
      try {
        await ElMessageBox.confirm('确定要退出登录吗？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })
        
        clearAuthInfo()
        ElMessage.success('已退出登录')
        router.push('/login')
      } catch {
        // 用户取消操作
      }
    }

    onMounted(() => {
      updateTime()
      updateAuthStatus()
      timer = setInterval(updateTime, 1000)
    })

    onUnmounted(() => {
      if (timer) {
        clearInterval(timer)
      }
    })

    return {
      currentTime,
      userInfo,
      isAuth,
      showMessage,
      handleLogout
    }
  }
}
</script>

<style scoped>
.hello-container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 20px;
}

.hello-card {
  width: 100%;
  max-width: 600px;
  text-align: center;
  border-radius: 15px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
}

.hello-content {
  padding: 40px 20px;
}

.hello-title {
  font-size: 3rem;
  color: #409eff;
  margin-bottom: 20px;
  font-weight: bold;
}

.hello-subtitle {
  font-size: 1.5rem;
  color: #606266;
  margin-bottom: 30px;
}

.hello-info {
  margin-top: 30px;
  padding-top: 20px;
  border-top: 1px solid #ebeef5;
}

.hello-info p {
  margin: 10px 0;
  color: #909399;
  font-size: 0.9rem;
}

.user-info {
  margin: 20px 0;
  padding: 20px;
  background: #f8f9fa;
  border-radius: 10px;
  border: 1px solid #e9ecef;
}

.user-info h3 {
  margin: 10px 0;
  color: #409eff;
}

.user-info p {
  margin: 5px 0;
  color: #606266;
}

.el-button {
  margin: 20px 0;
}
</style>
