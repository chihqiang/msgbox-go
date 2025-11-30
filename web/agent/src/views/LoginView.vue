<template>
  <a-layout class="login-container">
    <a-layout class="login-wrapper">
      <!-- 左侧品牌区域 -->
      <a-layout-sider width="40%" class="login-brand">
        <div class="brand-content">
          <img src="@/assets/logo.svg" alt="MSGBOX Logo" class="brand-logo" />
          <a-typography-title :level="2" class="brand-title">MSGBOX</a-typography-title>
          <a-typography-paragraph class="brand-description">企业级云消息推送平台</a-typography-paragraph>
          <div class="brand-features">
            <p>安全 · 稳定 · 高效 · 可靠</p>
          </div>
        </div>
      </a-layout-sider>

      <!-- 右侧登录表单区域 -->
      <a-layout-content class="login-form-container">
        <div class="form-wrapper">
          <!-- 登录表单标题 -->
          <div class="form-header">
            <a-typography-title :level="2" class="form-title">欢迎回来</a-typography-title>
            <a-typography-paragraph class="form-subtitle">请输入您的账号信息登录</a-typography-paragraph>
          </div>

          <!-- 登录表单 -->
          <a-form ref="formRef" :model="formState" @finish="handleLogin">
            <!-- 邮箱输入 -->
            <a-form-item label="邮箱" name="email" :rules="[{ required: true, message: '请输入邮箱', type: 'email' }]">
              <a-input v-model:value="formState.email" placeholder="请输入邮箱" prefix-icon="mail" />
            </a-form-item>

            <!-- 密码输入 -->
            <a-form-item label="密码" name="password" :rules="[{ required: true, message: '请输入密码', min: 8 }]">
              <a-input-password v-model:value="formState.password" placeholder="••••••••" :visibility-toggle="true" />
            </a-form-item>

            <!-- 记住我和忘记密码 -->
            <div class="form-options">
              <a-checkbox v-model:checked="formState.rememberMe">记住我</a-checkbox>
              <a href="#" class="forgot-password">忘记密码?</a>
            </div>

            <!-- 登录按钮 -->
            <a-form-item>
              <a-button type="primary" html-type="submit" size="large" :loading="loading" :disabled="loading" class="login-button">
                登录
              </a-button>
            </a-form-item>
          </a-form>

          <!-- 注册链接 -->
          <div class="register-link">
            <p>
              还没有账号?
              <router-link to="/register" class="register-button">
                立即注册
              </router-link>
            </p>
          </div>
        </div>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { message } from 'ant-design-vue'

// Ant Design Vue组件通过标签形式使用，无需导入组件对象
import { login } from '@/api/auth'
import { setToken } from '@/utils/cookie'
import { useRouter } from 'vue-router'
// import type { LoginRequest } from '@/api/auth'
// import { setToken, setJson, set } from '@/utils/cookie'
// 路由实例
const router = useRouter()
// 表单状态
const formState = reactive({
  email: '',
  password: '',
  rememberMe: false
})
// 加载状态
const loading = ref(false)
const handleLogin = async (loginData: typeof formState) => {
  const {  data } = await login(loginData)
  setToken(data.token, data.expires_in)
  message.success("登录成功")
  setTimeout(() => {
      router.push('/keys')
  }, 1000)
}
</script>

<style scoped>
/* 登录页面容器 */
.login-container {
  min-height: 100vh;
  background: linear-gradient(135deg, #f0f9ff 0%, #e0f2fe 100%);
  display: flex;
  align-items: center;
  justify-content: center;
}

/* 登录页面包装器 */
.login-wrapper {
  width: 100%;
  max-width: 960px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  overflow: hidden;
}

/* 品牌区域 */
.login-brand {
  background: linear-gradient(135deg, #1e40af 0%, #1e3a8a 100%);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* 品牌内容 */
.brand-content {
  text-align: center;
  padding: 24px;
}

/* 品牌Logo */
.brand-logo {
  width: 80px;
  height: 80px;
  margin: 0 auto 24px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  padding: 8px;
}

/* 品牌标题 */
.brand-title {
  color: white;
  margin-bottom: 16px;
}

/* 品牌描述 */
.brand-description {
  color: rgba(255, 255, 255, 0.9);
  margin-bottom: 24px;
}

/* 品牌特性 */
.brand-features {
  background: rgba(0, 0, 0, 0.2);
  padding: 12px;
  border-radius: 6px;
  display: inline-block;
}

.brand-features p {
  margin: 0;
  color: white;
}

/* 登录表单容器 */
.login-form-container {
  padding: 40px;
  background: white;
}

/* 表单包装器 */
.form-wrapper {
  max-width: 400px;
  margin: 0 auto;
}

/* 表单头部 */
.form-header {
  text-align: center;
  margin-bottom: 40px;
}

/* 表单标题 */
.form-title {
  color: #1f2937;
  margin-bottom: 8px;
}

/* 表单副标题 */
.form-subtitle {
  color: #6b7280;
}

/* 表单选项 */
.form-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

/* 忘记密码链接 */
.forgot-password {
  color: #1677ff;
}

/* 登录按钮 */
.login-button {
  width: 100%;
  padding: 12px;
}

/* 注册链接 */
.register-link {
  text-align: center;
  margin-top: 24px;
  padding-top: 24px;
  border-top: 1px solid #f0f0f0;
}

/* 注册按钮 */
.register-button {
  color: #1677ff;
  font-weight: 500;
}
</style>

