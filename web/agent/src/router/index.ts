import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
import RegisterView from '@/views/RegisterView.vue'
import KeysView from '@/views/KeysView.vue'
import ChannelView from '@/views/ChannelView.vue'
import TemplateView from '@/views/TemplateView.vue'
import RecordView from '@/views/RecordView.vue'
import DefaultLayout from '@/layouts/DefaultLayout.vue';


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: {
        layout: DefaultLayout,
        title: '首页',
        showInNav: true
      },
    },
    {
      path: '/keys',
      name: 'keys',
      component: KeysView,
      meta: {
        layout: DefaultLayout,
        title: 'API密钥管理',
        showInNav: true
      },
    },
       {
      path: '/record',
      name: 'record',
      component: RecordView,
      meta: {
        layout: DefaultLayout,
        title: '记录管理',
        showInNav: true
      },
    },
    {
      path: '/channel',
      name: 'channel',
      component: ChannelView,
      meta: {
        layout: DefaultLayout,
        title: '通道管理',
        showInNav: true
      },
    },
        {
      path: '/template',
      name: 'template',
      component: TemplateView,
      meta: {
        layout: DefaultLayout,
        title: '模板管理',
        showInNav: true
      },
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView,
      meta: {
        layout: DefaultLayout,
        title: '登录',
        showInNav: false
      },
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView,
      meta: {
        layout: DefaultLayout,
        title: '注册',
        showInNav: false
      },
    },
  ],
})

// 全局前置守卫 - 更新页面标题
router.beforeEach((to, _from, next) => {
  // 获取页面标题，如果没有设置则使用默认标题
  const pageTitle = to.meta?.title || 'MSGBOX'
  // 更新文档标题
  document.title = `${pageTitle} - MSGBOX`
  // 继续导航
  next()
})

export default router
