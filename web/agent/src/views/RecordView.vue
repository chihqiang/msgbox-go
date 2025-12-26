<template>
  <div>
    <!-- 页面标题和说明 -->
    <div style="margin-bottom: 32px">
      <a-typography-title :level="2">发送记录</a-typography-title>
      <a-typography-paragraph>查看消息发送记录，支持搜索和分页浏览。</a-typography-paragraph>
    </div>

    <!-- 搜索和筛选区域 -->
    <a-card style="margin-bottom: 24px">
      <a-space size="middle" wrap>
        <a-input-search
          v-model:value="searchKeyword"
          placeholder="搜索关键词（接收人、通道名称等）"
          allow-clear
          style="width: 300px"
          @search="handleSearch"
        />
        <a-button type="primary" @click="handleSearch">
          搜索
        </a-button>
        <a-button @click="handleReset"> 重置 </a-button>
      </a-space>
    </a-card>

    <!-- 记录列表卡片 -->
    <a-card>
      <a-table
        :columns="columns"
        :data-source="records"
        :pagination="pagination"
        row-key="id"
        :loading="loading"
        size="middle"
        :scroll="{ x: 1200 }"
      >
      </a-table>
    </a-card>

    <!-- 详情查看对话框 -->
    <a-modal
      v-model:open="showDetailModal"
      :title="detailModalTitle"
      @cancel="handleDetailCancel"
      width="800px"
    >
      <pre>{{ detailContent }}</pre>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, h } from 'vue'
import type { TableColumn } from '@arco-design/web-vue'
import { RecordItem } from '@/model/record'
import { listRecords } from '@/api/record'

// 表格列配置
const columns: TableColumn<RecordItem>[] = [
  {
    title: '接收人',
    dataIndex: 'receiver',
    key: 'receiver',
  },
  {
    title: '通道名称',
    dataIndex: 'channel_name',
    key: 'channel_name',
  },
  {
    title: '服务商',
    dataIndex: 'vendor_name',
    key: 'vendor_name',
  },
  {
    title: '模版标题',
    dataIndex: 'title',
    key: 'title',
    ellipsis: true,
  },
  {
    title: '发送内容',
    dataIndex: 'content',
    key: 'content',
    ellipsis: true,
    customRender: ({ record }: { record: RecordItem }) => {
      return h(
        'div',
        {
          class: 'content-cell',
          onClick: () => showContent(record),
          title: '点击查看完整内容',
        },
        record.content,
      )
    },
  },

  {
    title: '状态信息',
    dataIndex: 'status_msg',
    key: 'status_msg',
    ellipsis: true,
  },
  {
    title: '发送时间',
    dataIndex: 'send_time',
    key: 'send_time',
    ellipsis: true,
  },
  {
    title: '送达时间',
    dataIndex: 'delivery_time',
    key: 'delivery_time',
    ellipsis: true,
  },
  {
    title: '错误信息',
    dataIndex: 'error',
    key: 'error',
    ellipsis: true,
    customRender: ({ record }: { record: RecordItem }) => {
      if (!record.error) return null
      return h(
        'div',
        {
          class: 'error-cell',
          onClick: () => showError(record),
          title: '点击查看完整错误信息',
        },
        record.error,
      )
    },
  },
  {
    title: '响应信息',
    dataIndex: 'response',
    key: 'response',
    ellipsis: true,
    customRender: ({ record }: { record: RecordItem }) => {
      if (!record.response) return null
      return h(
        'div',
        {
          class: 'response-cell',
          onClick: () => showResponse(record),
          title: '点击查看完整响应',
        },
        record.response,
      )
    },
  },
  {
    title: '通道配置',
    dataIndex: 'channel_config',
    key: 'channel_config',
    customRender: ({ record }: { record: RecordItem }) => {
      return h(
        'div',
        {
          class: 'config-cell',
          onClick: () => showConfig(record.channel_config),
          title: '点击查看通道配置',
        },
        '查看配置',
      )
    },
  },
  {
    title: '变量',
    dataIndex: 'variables',
    key: 'variables',
    customRender: ({ record }: { record: RecordItem }) => {
      if (!record.variables) return null
      return h(
        'div',
        {
          class: 'config-cell',
          onClick: () => showConfig(record.variables),
          title: '点击查看变量',
        },
        '查看变量',
      )
    },
  },
]

// 响应式数据
const records = ref<RecordItem[]>([])
const searchKeyword = ref('')
const loading = ref(false)

// 分页配置
const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0,
  onChange: (page: number) => {
    pagination.current = page
    fetchRecords() // 页码变更时重新获取数据
  },
})

// 详情对话框
const showDetailModal = ref(false)
const detailModalTitle = ref('详情查看')
const detailContent = ref('')

// 生命周期钩子：组件挂载时获取记录列表
onMounted(() => {
  fetchRecords()
})

// 获取记录列表（调用后端API，支持搜索）
const fetchRecords = async () => {
  loading.value = true
  const res = await listRecords({
    page: pagination.current,
    size: pagination.pageSize,
    keywords: searchKeyword.value,
  })
  records.value = res.data.data || []
  pagination.total = res.data.total || 0
  loading.value = false
}

// 搜索处理
const handleSearch = () => {
  pagination.current = 1
  fetchRecords() // 搜索时重新获取数据
}

// 重置搜索
const handleReset = () => {
  searchKeyword.value = ''
  pagination.current = 1
  fetchRecords()
}

// 显示内容详情
const showContent = (record: RecordItem) => {
  detailModalTitle.value = '发送内容详情'
  detailContent.value = record.content
  showDetailModal.value = true
}

// 显示错误信息
const showError = (record: RecordItem) => {
  detailModalTitle.value = '错误信息详情'
  detailContent.value = record.error || '无错误信息'
  showDetailModal.value = true
}

// 显示响应信息
const showResponse = (record: RecordItem) => {
  detailModalTitle.value = '响应信息详情'
  detailContent.value = record.response || '无响应信息'
  showDetailModal.value = true
}

// 显示配置信息
const showConfig = (config?: Record<string, unknown>) => {
  detailModalTitle.value = '配置信息详情'
  if (!config) {
    detailContent.value = '暂无配置信息'
  } else {
    try {
      detailContent.value = JSON.stringify(config, null, 2)
    } catch {
      detailContent.value = String(config)
    }
  }
  showDetailModal.value = true
}

// 关闭详情对话框
const handleDetailCancel = () => {
  showDetailModal.value = false
}
</script>

<style scoped>
.content-cell,
.error-cell,
.response-cell {
  cursor: pointer;
  color: #1890ff;
  text-decoration: underline;
}

.content-cell:hover,
.error-cell:hover,
.response-cell:hover,
.config-cell:hover {
  color: #40a9ff;
}

.config-cell {
  cursor: pointer;
  color: #1890ff;
}

:deep(.ant-table-cell) {
  word-break: break-word;
}
</style>
