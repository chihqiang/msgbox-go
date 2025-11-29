<template>
  <a-layout-header>
    <a-row align="middle" justify="space-between" style="width: 100%; height: 100%;">
      <!-- Left: Logo -->
      <a-col>
        <a-space>
          <img src="@/assets/logo.svg" alt="MSGBOX Logo" style="width: 48px; height: 48px;" />
          <span style="font-size: 18px; font-weight: 700;">MSGBOX</span>
        </a-space>
      </a-col>

      <!-- Center: Desktop Navigation -->
      <a-col>
        <a-menu
          mode="horizontal"
          :selected-keys="[currentPath]"
          class="site-menu-horizontal"
          @click="(e: any) => navigateToRoute(e.key)"
        >
          <a-menu-item
            v-for="route in navRoutes"
            :key="route.path"
          >
            {{ route.meta.title }}
          </a-menu-item>
        </a-menu>
      </a-col>

      <!-- Right: Desktop Login/Logout Button -->
      <a-col>
        <!-- 登录按钮 - 未登录时显示 -->
        <a-button v-if="!isLoggedIn" type="primary" @click="navigateToLogin">
          登录
        </a-button>
        <!-- 退出按钮 - 已登录时显示 -->
        <a-button v-else type="default" @click="handleLogout">
          退出
        </a-button>
      </a-col>
    </a-row>
  </a-layout-header>
</template>

<script setup lang="ts">
// Header组件 - 网站顶部导航栏
import { useRouter, useRoute } from 'vue-router'
import { Button as AButton } from 'ant-design-vue'
import { computed, createVNode } from 'vue'
import { getToken, removeToken } from '@/utils/cookie'
import { Modal } from 'ant-design-vue';
import { ExclamationCircleOutlined } from '@ant-design/icons-vue';

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
    icon: createVNode(ExclamationCircleOutlined),
    okText: '确认',
    okType: 'danger',
    onOk: () => {
      // 执行退出登录操作
      removeToken()
      // 刷新当前页面
      window.location.reload()
    },
    onCancel: () => {
    }
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
