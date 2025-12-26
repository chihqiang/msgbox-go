<template>
  <div class="login-wrapper">
    <div class="login-brand">
      <div class="brand-content">
        <img src="@/assets/logo.svg" alt="MSGBOX Logo" class="brand-logo" />
        <a-typography-title :level="2" class="brand-title">MSGBOX</a-typography-title>
        <a-typography-paragraph class="brand-description">企业级云消息推送平台</a-typography-paragraph>
        <div class="brand-features">
          <p>安全 · 稳定 · 高效 · 可靠</p>
        </div>
      </div>
    </div>

    <div class="login-form-container">
      <div class="form-wrapper">
        <div class="form-header">
          <a-typography-title :level="2" class="form-title">欢迎回来</a-typography-title>
          <a-typography-paragraph class="form-subtitle">请输入您的账号信息登录</a-typography-paragraph>
        </div>

        <a-form ref="formRef" :model="formState" @finish="handleLogin">
          <a-form-item label="邮箱" name="email" :rules="[{ required: true, message: '请输入邮箱', type: 'email' }]">
            <a-input v-model:value="formState.email" placeholder="请输入邮箱" />
          </a-form-item>

          <a-form-item label="密码" name="password" :rules="[{ required: true, message: '请输入密码', min: 8 }]">
            <a-input-password v-model:value="formState.password" placeholder="••••••••" />
          </a-form-item>

          <div class="form-options">
            <a-checkbox v-model:checked="formState.rememberMe">记住我</a-checkbox>
            <a href="#" class="link">忘记密码?</a>
          </div>

          <a-form-item>
            <a-button type="primary" html-type="submit" size="large" :loading="loading" :disabled="loading" class="btn-block">
              登录
            </a-button>
          </a-form-item>
        </a-form>

        <div class="register-link">
          <p>
            还没有账号?
            <router-link to="/register" class="link">
              立即注册
            </router-link>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { Message } from '@arco-design/web-vue'
import { login } from '@/api/auth'
import { setToken } from '@/utils/cookie'
import { useRouter } from 'vue-router'

const router = useRouter()
const formState = reactive({
  email: '',
  password: '',
  rememberMe: false
})
const loading = ref(false)
const handleLogin = async (loginData: typeof formState) => {
  const { data } = await login(loginData)
  setToken(data.token, data.expires_in)
  Message.success("登录成功")
  setTimeout(() => {
    router.push('/keys')
  }, 1000)
}
</script>

<style scoped>
.login-wrapper {
  display: flex;
  min-height: 100vh;
}

.login-brand {
  width: 40%;
  background: linear-gradient(135deg, #1e40af 0%, #1e3a8a 100%);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
}

.brand-content {
  text-align: center;
  padding: 24px;
}

.brand-logo {
  width: 80px;
  height: 80px;
  margin: 0 auto 24px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  padding: 8px;
}

.brand-title {
  color: white;
  margin-bottom: 16px;
}

.brand-description {
  color: rgba(255, 255, 255, 0.9);
  margin-bottom: 24px;
}

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

.login-form-container {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 40px;
}

.form-wrapper {
  width: 100%;
  max-width: 400px;
}

.form-header {
  text-align: center;
  margin-bottom: 40px;
}

.form-title {
  color: #1f2937;
  margin-bottom: 8px;
}

.form-subtitle {
  color: #6b7280;
}

.form-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.link {
  color: #165DFF;
}

.btn-block {
  width: 100%;
}

.register-link {
  text-align: center;
  margin-top: 24px;
  padding-top: 24px;
  border-top: 1px solid #f0f0f0;
}

@media (max-width: 768px) {
  .login-wrapper {
    flex-direction: column;
  }

  .login-brand {
    width: 100%;
    padding: 32px 16px;
  }

  .brand-logo {
    width: 60px;
    height: 60px;
  }

  .brand-title {
    font-size: 24px;
  }

  .login-form-container {
    padding: 24px 16px;
  }

  .form-wrapper {
    max-width: 100%;
  }

  .form-header {
    margin-bottom: 24px;
  }
}
</style>
