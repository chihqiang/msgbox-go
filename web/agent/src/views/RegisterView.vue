<template>
  <div class="register-wrapper">
    <div class="register-brand">
      <div class="brand-content">
        <img src="@/assets/logo.svg" alt="MSGBOX Logo" class="brand-logo" />
        <a-typography-title :level="2" class="brand-title">MSGBOX</a-typography-title>
        <a-typography-paragraph class="brand-description"
          >企业级云消息推送平台</a-typography-paragraph
        >
        <div class="brand-features">
          <a-typography-paragraph class="features-text"
            >安全 · 稳定 · 高效 · 可靠</a-typography-paragraph
          >
        </div>
      </div>
    </div>

    <div class="register-form-container">
      <div class="form-wrapper">
        <div class="form-header">
          <a-typography-title :level="2" class="form-title">创建账号</a-typography-title>
          <a-typography-paragraph class="form-subtitle"
            >填写以下信息完成注册</a-typography-paragraph
          >
        </div>

        <a-form :model="formState" @finish="handleRegister">
          <a-form-item
            label="邮箱"
            name="email"
            :rules="[{ required: true, message: '请输入邮箱', type: 'email' }]"
          >
            <a-input v-model:value="formState.email" placeholder="请输入邮箱" />
          </a-form-item>

          <a-form-item
            label="手机号"
            name="phone"
            :rules="[{ required: false, message: '请输入手机号' }]"
          >
            <a-input
              v-model:value="formState.phone"
              placeholder="请输入手机号"
            />
          </a-form-item>

          <a-form-item
            label="密码"
            name="password"
            :rules="[{ required: true, message: '请输入密码', min: 8 }]"
          >
            <a-input-password
              v-model:value="formState.password"
              placeholder="••••••••"
            />
          </a-form-item>

          <a-form-item
            label="确认密码"
            name="confirmPassword"
            :rules="[{ validator: validatePassword, trigger: 'change' }]"
          >
            <a-input-password
              v-model:value="formState.confirmPassword"
              placeholder="••••••••"
            />
          </a-form-item>
          <a-form-item
            label="邀请码"
            name="code"
            :rules="[{ required: true, message: '请输入邀请码' }]"
          >
            <a-input
              v-model:value="formState.code"
              placeholder="请输入邀请码"
            />
          </a-form-item>
          <a-form-item>
            <a-button type="primary" html-type="submit" size="large" class="btn-block">
              立即注册
            </a-button>
          </a-form-item>
        </a-form>

        <div class="login-link">
          <a-typography-paragraph>
            已有账号?
            <router-link to="/login" class="link"> 立即登录 </router-link>
          </a-typography-paragraph>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import { Message } from '@arco-design/web-vue'
import { register } from '@/api/auth'
import { useRouter } from 'vue-router'
const router = useRouter()
const formState = reactive({
  email: '',
  phone: '',
  password: '',
  code: '',
  confirmPassword: '',
})

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
  console.log('Register submitted:', values)
  await register(values)
  Message.success('注册成功！')
  setTimeout(() => {
    router.push('/login')
  }, 1000)
}
</script>

<style scoped>
.register-wrapper {
  display: flex;
  min-height: 100vh;
}

.register-brand {
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

.features-text {
  color: white;
  margin: 0;
}

.register-form-container {
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

.btn-block {
  width: 100%;
}

.login-link {
  text-align: center;
  margin-top: 24px;
  padding-top: 24px;
  border-top: 1px solid #f0f0f0;
}

.link {
  color: #165DFF;
}

@media (max-width: 768px) {
  .register-wrapper {
    flex-direction: column;
  }

  .register-brand {
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

  .register-form-container {
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
