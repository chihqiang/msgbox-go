<template>
  <div>
    <div style="margin-bottom: 32px">
      <a-typography-title :level="2">通道管理</a-typography-title>
      <a-typography-paragraph
        >查看和管理您的消息发送通道，支持创建、编辑和删除操作。</a-typography-paragraph
      >
    </div>

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
          placeholder="搜索通道编码、名称或服务商"
          allow-clear
          style="width: 300px"
          @search="handleSearch"
        />
        <a-button type="primary" @click="handleCreate">
          <template #icon>
            <icon-plus />
          </template>
          创建通道
        </a-button>
      </div>
    </a-card>

    <a-card>
      <a-table
        :columns="columns"
        :data="filteredChannels"
        :pagination="pagination"
        row-key="id"
        :loading="loading"
        size="middle"
      >
        <template #operation="{ record }">
          <a-button type="text" @click="handleEdit(record)"> 编辑 </a-button>
          <a-button type="text" status="danger" @click="handleDelete(record)"> 删除 </a-button>
        </template>
      </a-table>
    </a-card>

    <a-modal
      v-model:open="showModal"
      :title="modalTitle"
      @ok="handleSave"
      @cancel="handleCancel"
      width="600px"
    >
      <channel-form v-if="currentChannel" :model="currentChannel" ref="channelForm" />
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, watch } from 'vue'
import { Modal } from '@arco-design/web-vue'
import type { TableColumn } from '@arco-design/web-vue'
import ChannelForm from './Forms/ChannelForm.vue'
import { ChannelItem } from '@/model/channel'
import { createChannel, deleteChannel, listChannels, updateChannel } from '@/api/channel'

const columns: TableColumn<ChannelItem>[] = [
  {
    title: '通道编码',
    dataIndex: 'code',
    key: 'code',
  },
  {
    title: '通道名称',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '服务商名称',
    dataIndex: 'vendor_name_label',
    key: 'vendor_name',
  },
  {
    title: '通道配置',
    dataIndex: 'config',
    key: 'config',
    ellipsis: true,
    render: ({ record }: { record: ChannelItem }) => {
      try {
        const configObj = typeof record.config === 'string' ? JSON.parse(record.config) : record.config
        return JSON.stringify(configObj, null, 2)
      } catch {
        return String(record.config)
      }
    },
  },
  {
    title: '状态',
    dataIndex: 'status',
    key: 'status',
    render: ({ record }: { record: ChannelItem }) => {
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
    key: 'operation',
    fixed: 'right',
  },
]

const channels = ref<ChannelItem[]>([])
const searchKeyword = ref('')
const loading = ref(false)
const showModal = ref(false)
const showCreateModal = ref(false)
const showDeleteConfirm = ref(false)
const currentChannel = ref<ChannelItem | null>(null)
const channelForm = ref<InstanceType<typeof ChannelForm> | null>(null)

const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0,
  onChange: (page: number) => {
    pagination.current = page
    fetchChannels()
  },
})
onMounted(() => {
  fetchChannels()
})

const filteredChannels = computed(() => {
  return channels.value
})
const modalTitle = computed(() => {
  return currentChannel.value?.id ? '编辑通道' : '创建通道'
})
const fetchChannels = async () => {
  loading.value = true
  const res = await listChannels({
    page: pagination.current,
    size: pagination.pageSize,
    keywords: searchKeyword.value,
  })
  channels.value = res.data.data || []
  pagination.total = res.data.total || 0
  loading.value = false
}

const handleSearch = () => {
  pagination.current = 1
  fetchChannels()
}

const handleCreate = () => {
  currentChannel.value = {
    id: 0,
    code: '',
    name: '',
    vendor_name: undefined,
    config: {},
    status: true,
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString(),
  }
  showModal.value = true
  showCreateModal.value = true
}

const handleEdit = (item: ChannelItem) => {
  currentChannel.value = { ...item }
  showModal.value = true
}

const handleDelete = (item: ChannelItem) => {
  currentChannel.value = { ...item }
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除此通道吗？此操作不可恢复。',
    onOk: async () => {
      if (!currentChannel.value) return
      try {
        loading.value = true
        if (currentChannel.value?.id) {
          await deleteChannel(currentChannel.value.id)
        }
        const index = channels.value.findIndex((c) => c.id === currentChannel.value?.id)
        if (index !== -1) {
          channels.value.splice(index, 1)
        }
        showDeleteConfirm.value = false
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

const handleSave = async () => {
  if (!channelForm.value) return

  try {
    const formData = await channelForm.value.validate()
    if (formData) {
      loading.value = true
      if (currentChannel.value?.id) {
        await updateChannel(formData)
      } else {
        await createChannel(formData)
      }
      showModal.value = false
      await fetchChannels()
      if (channelForm.value) {
        channelForm.value.resetFields()
      }
      console.log('保存成功')
    }
  } catch (error) {
    console.error('保存失败:', error)
  } finally {
    loading.value = false
  }
}

const handleCancel = () => {
  showModal.value = false
  if (channelForm.value) {
    channelForm.value.resetFields()
  }
}

watch(searchKeyword, () => {
  pagination.current = 1
})

watch(showCreateModal, (newVal: boolean) => {
  if (newVal) {
    handleCreate()
    showCreateModal.value = false
  }
})
</script>
