<template>
  <div>
    <!-- 页面标题和说明 -->
    <a-typography-title :level="2" style="margin-bottom: 8px">模版管理</a-typography-title>
    <a-typography-paragraph style="margin-bottom: 32px"
      >查看和管理您的消息模版，支持创建、编辑和删除操作。</a-typography-paragraph
    >

    <!-- 搜索和创建按钮区域 -->
    <a-card style="margin-bottom: 24px">
      <div
        style="
          display: flex;
          justify-content: space-between;
          align-items: center;
          flex-wrap: wrap;
          gap: 16px;
        "
      >
        <a-input-search
          v-model:value="searchKeyword"
          placeholder="搜索模板编码、名称或服务商"
          allow-clear
          style="width: 300px"
          @search="handleSearch"
        />
        <a-button type="primary" @click="handleCreate">
          <template #icon>
            <plus-outlined />
          </template>
          创建模版
        </a-button>
      </div>
    </a-card>

    <!-- 模版列表卡片 -->
    <a-card>
      <a-table
        :columns="columns"
        :data-source="filteredTemplates"
        :pagination="pagination"
        row-key="id"
        :loading="loading"
        size="middle"
      >
        <template #bodyCell="{ record, column }">
          <template v-if="column.key === 'actions'">
            <a-button-group>
              <a-button type="text" @click="handleEdit(record)"> 编辑 </a-button>
              <a-button type="text" status="danger" @click="handleDelete(record)"> 删除 </a-button>
            </a-button-group>
          </template>
        </template>
      </a-table>
    </a-card>

    <!-- 创建/编辑通道对话框 -->
    <a-modal
      v-model:open="showModal"
      :title="modalTitle"
      @ok="handleSave"
      @cancel="handleCancel"
      width="600px"
    >
      <template-form v-if="currentTemplate" :model="currentTemplate" ref="templateForm" />
    </a-modal>

    <!-- 删除确认对话框已通过函数式调用实现 -->
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, watch } from 'vue'
import { Modal } from '@arco-design/web-vue'
import type { TableColumn } from '@arco-design/web-vue'
import TemplateForm from '@/views/Forms/TemplateForm.vue'
import { TemplateItem } from '@/model/template'
import { createTemplate, deleteTemplate, listTemplates, updateTemplate } from '@/api/template'

// Ant Design Vue组件通过标签形式使用，不需要导入组件对象

// 表格列配置
const columns: TableColumn<TemplateItem>[] = [
  {
    title: '模版编码',
    dataIndex: 'code',
    key: 'code',
  },
  {
    title: '模版名称',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '服务商编码',
    dataIndex: 'vendor_code',
    key: 'vendor_code',
  },
  {
    title: '服务商签名',
    dataIndex: 'signature',
    key: 'signature',
  },
  {
    title: '模版内容',
    dataIndex: 'content',
    key: 'content',
    ellipsis: true,
  },
  {
    title: '状态',
    dataIndex: 'status',
    key: 'status',
    customRender: ({ record }: { record: TemplateItem }) => {
      return record.status ? '启用' : '禁用'
    },
  },
  {
    title: '更新时间',
    dataIndex: 'updated_at',
    key: 'updated_at',
  },
  {
    title: '操作',
    key: 'actions',
    fixed: 'right',
  },
]

// 响应式数据
const templates = ref<TemplateItem[]>([])
const searchKeyword = ref('')
const loading = ref(false)
const showModal = ref(false)
const showDeleteConfirm = ref(false)
const currentTemplate = ref<TemplateItem | null>(null)
const templateForm = ref<InstanceType<typeof TemplateForm> | null>(null)
// 分页配置
const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0,
  onChange: (page: number) => {
    pagination.current = page
    fetchTemplates() // 页码变更时重新获取数据
  },
})
// 生命周期钩子：组件挂载时获取模版列表
onMounted(() => {
  fetchTemplates()
})

// 搜索功能 - 直接返回templates数组，搜索在后端API中进行
const filteredTemplates = computed(() => {
  return templates.value
})
// 计算属性：模态框标题
const modalTitle = computed(() => {
  return currentTemplate.value?.id ? '编辑模版' : '创建模版'
})
// 获取模版列表（调用后端API，支持搜索）
const fetchTemplates = async () => {
  loading.value = true
  const res = await listTemplates({
    page: pagination.current,
    size: pagination.pageSize,
    keywords: searchKeyword.value, // 传入搜索关键词到后端API
  })
  templates.value = res.data.data || []
  pagination.total = res.data.total || 0
  loading.value = false
}

// 搜索处理
const handleSearch = () => {
  pagination.current = 1
  fetchTemplates() // 搜索时重新获取数据
}

// 打开创建模版模态框
const handleCreate = () => {
  // 初始化新的模板对象
  currentTemplate.value = {
    id: 0,
    channel_id: null,
    name: '',
    code: '',
    vendor_code: '',
    signature: '',
    content: '',
    status: true,
    used_count: 0,
    created_at: new Date().toISOString(),
    updated_at: new Date().toISOString(),
  }
  showModal.value = true
}

// 打开编辑模版模态框
const handleEdit = (item: TemplateItem) => {
  currentTemplate.value = { ...item }
  showModal.value = true
}

// 打开删除确认对话框
const handleDelete = (item: TemplateItem) => {
  currentTemplate.value = { ...item }
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除此模版吗？此操作不可恢复。',
    onOk: async () => {
      if (!currentTemplate.value) return
      try {
        // 模拟API调用
        loading.value = true
        if (currentTemplate.value?.id) {
          await deleteTemplate(currentTemplate.value.id)
        }
        // 从本地数据中移除
        const index = templates.value.findIndex((t) => t.id === currentTemplate.value?.id)
        if (index !== -1) {
          templates.value.splice(index, 1)
        }
        // 关闭确认对话框
        showDeleteConfirm.value = false
        // 显示成功消息
        console.log('删除成功')
      } catch (error) {
        console.error('删除失败:', error)
      } finally {
        loading.value = false
      }
    },
    onCancel: () => {
      showDeleteConfirm.value = false
    },
  })
}

// 保存模版（创建或编辑）
const handleSave = async () => {
  if (!templateForm.value) return

  try {
    const formData = await templateForm.value.validate()
    if (formData) {
      // 模拟API调用
      loading.value = true
      if (currentTemplate.value?.id) {
        await updateTemplate(formData)
      } else {
        await createTemplate(formData)
      }
      // 保存成功后关闭模态框
      showModal.value = false
      // 重新获取模版列表以显示最新数据
      await fetchTemplates()
      // 重置表单
      if (templateForm.value) {
        templateForm.value.resetFields()
      }
      console.log('保存成功')
    }
  } catch (error) {
    console.error('保存失败:', error)
  } finally {
    loading.value = false
  }
}

// 取消操作
const handleCancel = () => {
  showModal.value = false
  if (templateForm.value) {
    templateForm.value.resetFields()
  }
}

// 删除功能的处理逻辑已通过闭包方式直接实现在Modal.confirm配置中

// 监听搜索关键词变化
watch(searchKeyword, () => {
  pagination.current = 1
  // 注意：不再手动设置total，而是通过fetchChannels从后端获取
})
</script>
