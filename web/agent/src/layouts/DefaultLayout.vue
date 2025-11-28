<template>
  <!-- 使用Ant Design的Layout组件 -->
  <a-layout>
    <!-- 使用独立的SiteHeader组件 -->
    <a-layout-header>
      <SiteHeader />
    </a-layout-header>

    <!-- 主内容区域 -->
    <a-layout-content>
      <div class="p-4">
        <slot></slot>
      </div>
    </a-layout-content>

    <!-- 使用独立的SiteFooter组件 -->
    <a-layout-footer>
      <SiteFooter />
    </a-layout-footer>
  </a-layout>
</template>

<script setup lang="ts">
// 导入组件
import SiteHeader from '../components/SiteHeader.vue'
import SiteFooter from '../components/SiteFooter.vue'
</script>

<style scoped>
.p-4 {
  padding: 1rem;
}

/* 为导航栏添加白色背景 */
:deep(.ant-layout-header) {
  background-color: white !important;
  border-bottom: 1px solid #f0f0f0;
}
</style>
