<template>
  <a-layout style="min-height: 100vh; background: linear-gradient(135deg, #f0f9ff 0%, #e0f2fe 100%); display: flex; align-items: center; justify-content: center; padding: 20px;">
    <a-layout style="width: 100%; max-width: 960px; box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1); border-radius: 8px; overflow: hidden;">
      <!-- 左侧品牌区域 -->
      <a-layout-sider width="40%" style="background: linear-gradient(135deg, #1e40af 0%, #1e3a8a 100%); color: white; display: flex; align-items: center; justify-content: center;">
        <div style="text-align: center; padding: 24px;">
          <img src="@/assets/logo.svg" alt="MSGBOX Logo" style="width: 80px; height: 80px; margin: 0 auto 24px; background: rgba(255, 255, 255, 0.2); border-radius: 50%; padding: 8px;" />
          <a-typography-title level="2" style="color: white; margin-bottom: 16px;">MSGBOX</a-typography-title>
          <a-typography-paragraph style="color: white; opacity: 0.9; margin-bottom: 24px;">企业级云消息推送平台</a-typography-paragraph>
          <div style="background: rgba(0, 0, 0, 0.2); padding: 12px; border-radius: 6px; display: inline-block;">
            <a-typography-paragraph style="color: white; margin: 0;">安全 · 稳定 · 高效 · 可靠</a-typography-paragraph>
          </div>
        </div>
      </a-layout-sider>

      <!-- 右侧注册表单区域 -->
      <a-layout-content style="padding: 40px; background: white;">
        <div style="max-width: 400px; margin: 0 auto;">
          <!-- 注册表单标题 -->
          <div style="text-align: center; margin-bottom: 40px;">
            <a-typography-title level="2" style="color: #1f2937; margin-bottom: 8px;">创建账号</a-typography-title>
            <a-typography-paragraph style="color: #6b7280;">填写以下信息完成注册</a-typography-paragraph>
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
              <a-button type="primary" html-type="submit" size="large" style="width: 100%; padding: 12px;">
                立即注册
              </a-button>
            </a-form-item>
          </a-form>

          <!-- 登录链接 -->
          <div style="text-align: center; margin-top: 24px; padding-top: 24px; border-top: 1px solid #f0f0f0;">
            <a-typography-paragraph>
              已有账号?
              <router-link to="/login" style="color: #1677ff; font-weight: 500;"> 立即登录 </router-link>
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


