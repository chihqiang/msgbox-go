<template>
  <a-layout style="min-height: 100vh; background: linear-gradient(135deg, #f0f9ff 0%, #e0f2fe 100%); display: flex; align-items: center; justify-content: center;">
    <a-layout style="width: 100%; max-width: 960px; box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1); border-radius: 8px; overflow: hidden;">
      <!-- 左侧品牌区域 -->
      <a-layout-sider width="40%" style="background: linear-gradient(135deg, #1e40af 0%, #1e3a8a 100%); color: white; display: flex; align-items: center; justify-content: center;">
        <div style="text-align: center; padding: 24px;">
          <img src="@/assets/logo.svg" alt="MSGBOX Logo" style="width: 80px; height: 80px; margin: 0 auto 24px; background: rgba(255, 255, 255, 0.2); border-radius: 50%; padding: 8px;" />
          <a-typography-title level="2" style="color: white; margin-bottom: 16px;">MSGBOX</a-typography-title>
          <a-typography-paragraph style="color: rgba(255, 255, 255, 0.9); margin-bottom: 24px;">企业级云消息推送平台</a-typography-paragraph>
          <div style="background: rgba(0, 0, 0, 0.2); padding: 12px; border-radius: 6px; display: inline-block;">
            <p style="margin: 0; color: white;">安全 · 稳定 · 高效 · 可靠</p>
          </div>
        </div>
      </a-layout-sider>

      <!-- 右侧登录表单区域 -->
      <a-layout-content style="padding: 40px; background: white;">
        <div style="max-width: 400px; margin: 0 auto;">
          <!-- 登录表单标题 -->
          <div style="text-align: center; margin-bottom: 40px;">
            <a-typography-title level="2" style="color: #1f2937; margin-bottom: 8px;">欢迎回来</a-typography-title>
            <a-typography-paragraph style="color: #6b7280;">请输入您的账号信息登录</a-typography-paragraph>
          </div>

          <!-- 登录表单 -->
          <a-form ref="formRef" :model="formState" @finish="handleLogin">
            <!-- 邮箱输入 -->
            <a-form-item label="邮箱" name="email" :rules="[{ required: true, message: '请输入邮箱', type: 'email' }]">
              <a-input v-model:value="formState.email" placeholder="请输入邮箱" prefix-icon="mail" />
            </a-form-item>

            <!-- 密码输入 -->
            <a-form-item label="密码" name="password" :rules="[{ required: true, message: '请输入密码' }]">
              <a-input-password v-model:value="formState.password" placeholder="••••••••" :visibility-toggle="true" />
            </a-form-item>

            <!-- 记住我和忘记密码 -->
            <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 24px;">
              <a-checkbox v-model:checked="formState.rememberMe">记住我</a-checkbox>
              <a href="#">忘记密码?</a>
            </div>

            <!-- 登录按钮 -->
            <a-form-item>
              <a-button type="primary" html-type="submit" size="large" :loading="loading" :disabled="loading" style="width: 100%; padding: 12px;">
                登录
              </a-button>
            </a-form-item>
          </a-form>

          <!-- 注册链接 -->
          <div style="text-align: center; margin-top: 24px; padding-top: 24px; border-top: 1px solid #f0f0f0;">
            <p>
              还没有账号?
              <router-link to="/register" style="color: #1677ff; font-weight: 500;">
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


