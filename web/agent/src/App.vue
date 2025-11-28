<template>
  <component :is="currentLayout">
    <router-view />
  </component>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import DefaultLayout from './layouts/DefaultLayout.vue'

const route = useRoute()

// 动态计算当前页面应该使用的布局
const currentLayout = computed(() => {
  // 如果路由配置了layout，则使用配置的layout，否则使用默认布局
  return route.meta.layout || DefaultLayout
})
</script>
