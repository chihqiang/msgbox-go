<template>
  <a-layout class="register-container">
    <a-layout class="register-wrapper">
      <!-- 左侧品牌区域 -->
      <a-layout-sider width="40%" class="register-brand">
        <div class="brand-content">
          <img src="@/assets/logo.svg" alt="MSGBOX Logo" class="brand-logo" />
          <a-typography-title level="2" class="brand-title">MSGBOX</a-typography-title>
          <a-typography-paragraph class="brand-description">企业级云消息推送平台</a-typography-paragraph>
          <div class="brand-features">
            <a-typography-paragraph class="features-text">安全 · 稳定 · 高效 · 可靠</a-typography-paragraph>
          </div>
        </div>
      </a-layout-sider>

      <!-- 右侧注册表单区域 -->
      <a-layout-content class="register-form-container">
        <div class="form-wrapper">
          <!-- 注册表单标题 -->
          <div class="form-header">
            <a-typography-title level="2" class="form-title">创建账号</a-typography-title>
            <a-typography-paragraph class="form-subtitle">填写以下信息完成注册</a-typography-paragraph>
          </div>

          <!-- 注册表单 -->
          <a-form :model="formState" @finish="handleRegister">
            <!-- 邮箱输入 -->
            <a-form-item
              label="邮箱"
              name="email"
              :rules="[{ required: true, message: '请输入邮箱', type: 'email' }]"
            >
              <a-input
                v-model:value="formState.email"
                placeholder="请输入邮箱"
                prefix-icon="mail"
              />
            </a-form-item>

            <!-- 手机号输入 -->
            <a-form-item
              label="手机号"
              name="phone"
              :rules="[{ required: false, message: '请输入手机号' }]"
            >
              <a-input
                v-model:value="formState.phone"
                placeholder="请输入手机号"
                prefix-icon="mobile"
              />
            </a-form-item>

            <!-- 密码输入 -->
            <a-form-item
              label="密码"
              name="password"
              :rules="[{ required: true, message: '请输入密码', min: 8 }]"
            >
              <a-input-password
                v-model:value="formState.password"
                placeholder="••••••••"
                :visibility-toggle="true"
              />
            </a-form-item>

            <!-- 确认密码输入 -->
            <a-form-item
              label="确认密码"
              name="confirmPassword"
              :rules="[{ validator: validatePassword, trigger: 'change' }]"
            >
              <a-input-password
                v-model:value="formState.confirmPassword"
                placeholder="••••••••"
                :visibility-toggle="true"
              />
            </a-form-item>

            <!-- 注册按钮 -->
            <a-form-item>
              <a-button type="primary" html-type="submit" size="large" class="register-button">
                立即注册
              </a-button>
            </a-form-item>
          </a-form>

          <!-- 登录链接 -->
          <div class="login-link">
            <a-typography-paragraph>
              已有账号?
              <router-link to="/login" class="login-button"> 立即登录 </router-link>
            </a-typography-paragraph>
          </div>
        </div>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import { message } from 'ant-design-vue'
import { register } from '@/api/auth'
import { useRouter } from 'vue-router'
const router = useRouter()
// 表单状态
const formState = reactive({
  email: '',
  phone: '',
  password: '',
  confirmPassword: '',
})

// 密码验证函数
const validatePassword = (
  _rule: { field: string; message?: string; required?: boolean },
  value: string,
) => {
  if (!value) {
    return Promise.reject('请确认密码')
  }
  if (value !== formState.password) {
    return Promise.reject('两次输入的密码不一致')
  }
  return Promise.resolve()
}

const handleRegister = async (values: typeof formState) => {
  // 注册逻辑将在这里实现
  console.log('Register submitted:', values)
  await register(values)
  message.success("注册成功！")
  setTimeout(() => {
      router.push('/login')
  }, 1000)
}
</script>

<style scoped>
/* 注册页面容器 */
.register-container {
  min-height: 100vh;
  background: linear-gradient(135deg, #f0f9ff 0%, #e0f2fe 100%);
  display: flex;
  align-items: center;
  justify-content: center;
}

/* 注册页面包装器 */
.register-wrapper {
  width: 100%;
  max-width: 960px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  overflow: hidden;
}

/* 品牌区域 */
.register-brand {
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
  color: white;
  opacity: 0.9;
  margin-bottom: 24px;
}

/* 品牌特性 */
.brand-features {
  background: rgba(0, 0, 0, 0.2);
  padding: 12px;
  border-radius: 6px;
  display: inline-block;
}

.features-text {
  color: white;
  margin: 0;
}

/* 注册表单容器 */
.register-form-container {
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

/* 注册按钮 */
.register-button {
  width: 100%;
  padding: 12px;
}

/* 登录链接 */
.login-link {
  text-align: center;
  margin-top: 24px;
  padding-top: 24px;
  border-top: 1px solid #f0f0f0;
}

/* 登录按钮 */
.login-button {
  color: #1677ff;
  font-weight: 500;
}
</style>

