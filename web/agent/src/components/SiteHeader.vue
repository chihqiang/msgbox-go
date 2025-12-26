<template>
  <a-layout-header class="site-header">
    <div class="site-header-inner">
      <!-- Left: Logo -->
      <div class="header-logo" @click="navigateToRoute('/')">
        <img src="@/assets/logo.svg" alt="MSGBOX Logo" class="logo-img"/>
        <span class="logo-text">MSGBOX</span>
      </div>

      <!-- Center: Desktop Navigation -->
      <div class="header-nav">
        <a-menu
          mode="horizontal"
          :selected-keys="[currentPath]"
          :popup-max-height="360"
        >
          <a-menu-item
            v-for="route in navRoutes"
            :key="route.path"
            @click="navigateToRoute(route.path)"
          >
            {{ route.meta.title }}
          </a-menu-item>
        </a-menu>
      </div>

      <!-- Right: Desktop Login/Logout Button -->
      <div class="header-actions">
        <!-- 登录按钮 - 未登录时显示 -->
        <a-button v-if="!isLoggedIn" type="primary" size="small" @click="navigateToLogin">
          登录
        </a-button>
        <!-- 退出按钮 - 已登录时显示 -->
        <a-button v-else size="small" @click="handleLogout">
          退出
        </a-button>
      </div>
    </div>
  </a-layout-header>
</template>

<script setup lang="ts">
// Header组件 - 网站顶部导航栏
import { useRouter, useRoute } from 'vue-router'
import { computed } from 'vue'
import { getToken, removeToken } from '@/utils/cookie'
import { Modal } from '@arco-design/web-vue'

// 获取路由实例
const router = useRouter()
const route = useRoute()

// 获取当前路径，用于激活状态判断
const currentPath = computed(() => route.path)

// 获取所有应该在导航中显示的路由
const navRoutes = computed(() => {
  return router.getRoutes().filter(route => route.meta.showInNav)
})

// 检查用户是否已登录
const isLoggedIn = computed(() => !!getToken())

// 退出登录函数
const handleLogout = () => {
  Modal.confirm({
    title: '确认退出登录吗？',
    content: '确定要退出当前账号吗？',
    okText: '确认',
    onOk: () => {
      removeToken()
      window.location.reload()
    },
  })
}

// 导航到登录页面
const navigateToLogin = () => {
  router.push('/login')
}

// 导航到指定路由
const navigateToRoute = (path: string) => {
  router.push(path)
}
</script>

<style scoped>
.site-header {
  background: #fff;
  padding: 0 16px;
  height: 56px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  position: sticky;
  top: 0;
  z-index: 100;
}

.site-header-inner {
  display: flex;
  align-items: center;
  width: 100%;
  height: 100%;
  gap: 8px;
}

.header-logo {
  display: flex;
  align-items: center;
  cursor: pointer;
  flex-shrink: 0;
}

.logo-img {
  width: 24px;
  height: 24px;
  margin-right: 8px;
}

.logo-text {
  font-size: 14px;
  font-weight: 700;
  color: #1d2129;
  white-space: nowrap;
}

.header-nav {
  flex: 5;
  overflow: hidden;
}

.header-actions {
  flex: 1;
}
</style>
