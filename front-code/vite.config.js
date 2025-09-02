import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src')
    }
  },
  server: {
    host: '0.0.0.0',  // 这里加上
    port: 12403,
    open: true,
    proxy: {
      '/api': {
        target: 'http://localhost:9080',
        changeOrigin: true,
        configure: (proxy, options) => {
          // 代理请求时的日志
          proxy.on('proxyReq', (proxyReq, req, res) => {
            console.log('🚀 [Vite代理] 转发请求:', {
              method: req.method,
              url: req.url,
              target: options.target + req.url,
              timestamp: new Date().toISOString()
            });
          });
          
          // 代理响应时的日志
          proxy.on('proxyRes', (proxyRes, req, res) => {
            console.log('📥 [Vite代理] 收到响应:', {
              method: req.method,
              url: req.url,
              statusCode: proxyRes.statusCode,
              statusMessage: proxyRes.statusMessage,
              timestamp: new Date().toISOString()
            });
          });
          
          // 代理错误时的日志
          proxy.on('error', (err, req, res) => {
            console.error('❌ [Vite代理] 代理错误:', {
              method: req.method,
              url: req.url,
              error: err.message,
              timestamp: new Date().toISOString()
            });
          });
          
          // 代理连接时的日志
          proxy.on('proxyReqWs', (proxyReq, req, socket) => {
            console.log('🔌 [Vite代理] WebSocket 连接:', {
              url: req.url,
              timestamp: new Date().toISOString()
            });
          });
        }
      }
    }
  }
})
