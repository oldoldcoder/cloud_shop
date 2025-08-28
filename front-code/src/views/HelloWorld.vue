<template>
  <div class="hello-container">
    <el-card class="hello-card">
      <div class="hello-content">
        <h1 class="hello-title">Hello World!</h1>
        <p class="hello-subtitle">欢迎来到 Cloud Shop 云商城</p>
        <el-button type="primary" size="large" @click="showMessage">
          点击我
        </el-button>
        <div class="hello-info">
          <p>这是一个基于 Vue 3 + Element Plus 的前端项目模板</p>
          <p>当前时间: {{ currentTime }}</p>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script>
import { ref, onMounted, onUnmounted } from 'vue'
import { ElMessage } from 'element-plus'

export default {
  name: 'HelloWorld',
  setup() {
    const currentTime = ref('')
    let timer = null

    const updateTime = () => {
      currentTime.value = new Date().toLocaleString('zh-CN')
    }

    const showMessage = () => {
      ElMessage.success('Hello! 你点击了按钮！')
    }

    onMounted(() => {
      updateTime()
      timer = setInterval(updateTime, 1000)
    })

    onUnmounted(() => {
      if (timer) {
        clearInterval(timer)
      }
    })

    return {
      currentTime,
      showMessage
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

.el-button {
  margin: 20px 0;
}
</style>
