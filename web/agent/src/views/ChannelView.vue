<template>
  <a-layout>
    <a-layout-content style="padding: 24px; background-color: #f5f5f5; min-height: calc(100vh - 100px);">
      <div style="max-width: 1200px; margin: 0 auto;">
        <!-- 页面标题和说明 -->
        <div style="margin-bottom: 32px;">
          <a-typography-title :level="2">通道管理</a-typography-title>
          <a-typography-paragraph>查看和管理您的消息发送通道，支持创建、编辑和删除操作。</a-typography-paragraph>
        </div>

      <!-- 搜索和创建按钮区域 -->
      <a-card style="margin-bottom: 24px;">
        <div style="display: flex; justify-content: space-between; align-items: center; flex-wrap: wrap; gap: 16px;">
          <a-input-search v-model:value="searchKeyword" placeholder="搜索通道编码、名称或服务商" allow-clear style="width: 300px;"
            @search="handleSearch" />
          <a-button type="primary" @click="handleCreate">
            <template #icon>
              <plus-outlined />
            </template>
            创建通道
          </a-button>
        </div>
      </a-card>

      <!-- 通道列表卡片 -->
      <a-card>
        <a-table :columns="columns" :data-source="filteredChannels" :pagination="pagination" row-key="id"
          :loading="loading" size="middle">
          <template #bodyCell="{ record, column }">
            <template v-if="column.key === 'actions'">
              <a-button-group>
                <a-button type="link" @click="handleEdit(record)">
                  编辑
                </a-button>
                <a-button type="link" danger @click="handleDelete(record)">
                  删除
                </a-button>
              </a-button-group>
            </template>
          </template>
        </a-table>
      </a-card>

      <!-- 创建/编辑通道对话框 -->
      <a-modal v-model:open="showModal" :title="modalTitle" @ok="handleSave" @cancel="handleCancel" width="600px">
        <channel-form v-if="currentChannel" :model="currentChannel" ref="channelForm" />
      </a-modal>

      <!-- 删除确认对话框已通过函数式调用实现 -->
      </div>
    </a-layout-content>
  </a-layout>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, watch } from 'vue';
import { PlusOutlined } from '@ant-design/icons-vue';
import { Modal } from 'ant-design-vue';

// Ant Design Vue组件通过标签形式使用，无需导入组件对象
import type { TableColumnsType } from 'ant-design-vue';
import ChannelForm from './Forms/ChannelForm.vue';
import { ChannelItem } from '@/model/channel';
import { createChannel, deleteChannel, listChannels, updateChannel } from '@/api/channel';

// 表格列配置
const columns: TableColumnsType<ChannelItem> = [
  {
    title: '通道编码',
    dataIndex: 'code',
    key: 'code'
  },
  {
    title: '通道名称',
    dataIndex: 'name',
    key: 'name'
  },
  {
    title: '服务商名称',
    dataIndex: 'vendor_name',
    key: 'vendor_name',
    customRender: ({ text }) => {
      const vendorMap: Record<string, string> = {
        dingtalk: '钉钉机器人',
        wechat: '微信公众号',
        smtp: 'SMTP邮件'
      };
      return vendorMap[text] || text;
    }
  },
  {
    title: '通道配置',
    dataIndex: 'config',
    key: 'config',
    ellipsis: true,
    customRender: ({ text }: { text: string | object }) => {
        try {
          // 尝试解析配置内容
          const configObj = typeof text === 'string' ? JSON.parse(text) : text;
          // 显示格式化后的配置，确保access_token可见
          return JSON.stringify(configObj, null, 2);
        } catch {
          // 如果解析失败，返回原始文本
          return String(text);
        }
      }
  },
  {
    title: '状态',
    dataIndex: 'status',
    key: 'status',
    customRender: ({ text }) => {
      return text ? '启用' : '禁用';
    }
  },
  {
    title: '更新时间',
    dataIndex: 'updated_at',
    key: 'updated_at'
  },
  {
    title: '操作',
    key: 'actions',
    fixed: 'right'
  },
];

// 响应式数据
const channels = ref<ChannelItem[]>([]);
const searchKeyword = ref('');
const loading = ref(false);
const showModal = ref(false);
const showCreateModal = ref(false);
const showDeleteConfirm = ref(false);
const currentChannel = ref<ChannelItem | null>(null);
const channelForm = ref<InstanceType<typeof ChannelForm> | null>(null);

// 分页配置
const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0,
  onChange: (page: number) => {
    pagination.current = page;
    fetchChannels(); // 页码变更时重新获取数据
  }
});
// 生命周期钩子：组件挂载时获取通道列表
onMounted(() => {
  fetchChannels();
});

// 搜索功能 - 直接返回channels数组，搜索在后端API中进行
const filteredChannels = computed(() => {
  return channels.value;
});
// 计算属性：模态框标题
const modalTitle = computed(() => {
  return currentChannel.value?.id ? '编辑通道' : '创建通道';
});
// 获取通道列表（调用后端API，支持搜索）
const fetchChannels = async () => {
  loading.value = true;
  const res = await listChannels({
    page: pagination.current,
    size: pagination.pageSize,
    keywords: searchKeyword.value // 传入搜索关键词到后端API
  })
  channels.value = res.data.data || [];
  pagination.total = res.data.total || 0;
  loading.value = false;
};

// 搜索处理
const handleSearch = () => {
  pagination.current = 1;
  fetchChannels(); // 搜索时重新获取数据
};

// 打开创建通道模态框
const handleCreate = () => {
  currentChannel.value = {
    id: 0,
    code: '',
    name: '',
    vendor_name: undefined,
    config: {},
    status: true,
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString()
  };
  showModal.value = true;
  showCreateModal.value = true;
};

// 打开编辑通道模态框
const handleEdit = (item: ChannelItem) => {
  currentChannel.value = { ...item };
  showModal.value = true;
};

// 打开删除确认对话框
const handleDelete = (item: ChannelItem) => {
  currentChannel.value = { ...item };
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除此通道吗？此操作不可恢复。',
    onOk: async () => {
      if (!currentChannel.value) return;
      try {
        // 模拟API调用
        loading.value = true;
        if (currentChannel.value?.id) {
          await deleteChannel(currentChannel.value.id);
        }
        // 从本地数据中移除
        const index = channels.value.findIndex(c => c.id === currentChannel.value?.id);
        if (index !== -1) {
          channels.value.splice(index, 1);
        }
        // 关闭确认对话框
        showDeleteConfirm.value = false;
        // 显示成功消息
        console.log('删除成功');
      } catch (error) {
        console.error('删除失败:', error);
      } finally {
        loading.value = false;
      }
    },
    onCancel: () => {
      showDeleteConfirm.value = false;
    }
  });
};

// 保存通道（创建或编辑）
const handleSave = async () => {
  if (!channelForm.value) return;

  try {
    const formData = await channelForm.value.validate();
    if (formData) {
      // 模拟API调用
      loading.value = true;
      if (currentChannel.value?.id) {
        await updateChannel(formData);
      } else {
        await createChannel(formData);
      }
      // 保存成功后关闭模态框
      showModal.value = false;
      // 重新获取通道列表以显示最新数据
      await fetchChannels();
      // 重置表单
      if (channelForm.value) {
        channelForm.value.resetFields();
      }
      console.log('保存成功');
    }
  } catch (error) {
    console.error('保存失败:', error);
  } finally {
    loading.value = false;
  }
};

// 取消操作
const handleCancel = () => {
  showModal.value = false;
  if (channelForm.value) {
    channelForm.value.resetFields();
  }
};

// 删除功能的处理逻辑已通过闭包方式直接实现在Modal.confirm配置中

// 监听搜索关键词变化
watch(searchKeyword, () => {
  pagination.current = 1;
  // 注意：不再手动设置total，而是通过fetchChannels从后端获取
});

// 监听创建模态框显示状态
watch(showCreateModal, (newVal: boolean) => {
  if (newVal) {
    handleCreate();
    showCreateModal.value = false;
  }
});
</script>


