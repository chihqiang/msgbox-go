<template>
  <div class="template-form-container">
    <a-form
      v-if="formModel"
      :model="formModel"
      :rules="rules"
      layout="vertical"
      ref="formRef"
      class="modern-form"
    >
      <a-form-item label="通道" name="channel_id" class="form-item">
        <a-select
          v-model:value="formModel.channel_id"
          show-search
          placeholder="请选择通道"
          style="width: 200px"
          :options="channelOptions"
          :filter-option="filterOption"
          @change="handleChange"
        ></a-select>
      </a-form-item>
      <a-form-item label="模板名称" name="name" class="form-item">
        <a-input v-model:value="formModel.name" placeholder="请输入模板名称" class="modern-input" />
      </a-form-item>
      <a-form-item label="模板编码" name="code" class="form-item">
        <a-input v-model:value="formModel.code" placeholder="请输入模板编码" class="modern-input" />
      </a-form-item>

      <a-form-item label="服务商编码" name="vendor_code" class="form-item">
        <a-input
          v-model:value="formModel.vendor_code"
          placeholder="请输入服务商编码"
          class="modern-input"
        />
      </a-form-item>

      <a-form-item label="服务商签名" name="signature" class="form-item">
        <a-input
          v-model:value="formModel.signature"
          placeholder="请输入签名"
          class="modern-input"
        />
      </a-form-item>

      <a-form-item label="模板内容" name="content" class="form-item">
        <a-textarea
          v-model:value="formModel.content"
          placeholder="请输入模板内容"
          :auto-size="{ minRows: 2, maxRows: 5 }"
        />
      </a-form-item>

      <!-- 使用水平布局的子表单来确保状态字段的标签和开关在同一行 -->
      <a-form-item label="状态" name="status">
        <a-switch v-model:checked="formModel.status" />
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, watch, onMounted } from 'vue'
import type { FormInstance } from 'ant-design-vue'
import { TemplateItem } from '@/model/template'
import { SelectOption } from '@/model/base'
import { listChannels } from '@/api/channel'
import { ChannelItem } from '@/model/channel'
// Props定义
interface Props {
  model: TemplateItem | null
}
// 定义props，默认值为null
const props = withDefaults(defineProps<Props>(), { model: null })
// 编辑模式判断已移除，因为模板表单不再需要根据编辑状态显示不同字段
// 表单引用
const formRef = ref<FormInstance | null>(null)
// 本地响应式数据，避免直接修改props
const formModel = ref<TemplateItem | null>(null)
// 表单验证规则
const rules = reactive({
  name: [{ required: true, message: '请输入模板名称', trigger: 'blur' }],
  code: [{ required: true, message: '请输入模板编码', trigger: 'blur' }],
  content: [{ required: true, message: '请输入模板内容', trigger: 'blur' }],
  channel_id: [{ required: true, message: '请选择通道', trigger: ['blur', 'change'] }],
})

const channelOptions = reactive<SelectOption[]>([])
const filterOption = (input: string, option: SelectOption) => {
  const optionValue = String(option.value)
  return optionValue.toLowerCase().indexOf(input.toLowerCase()) >= 0
}
const handleChange = (value: string | number) => {
  if (formModel.value) {
    formModel.value.channel_id = Number(value)
  }
}
onMounted(() => {
  fetchchannelOptions()
})
// 从后端获取通道列表
const fetchchannelOptions = async () => {
  try {
    const response = await listChannels({ page: 1, size: 100 })
    // 清空数组并添加新数据，确保id是有效的number类型
    const validChannels =
      response.data?.data
        ?.filter((item: ChannelItem) => typeof item.id === 'number' && item.id > 0)
        ?.map((item: ChannelItem) => ({
          label: item.name,
          value: item.id as number,
        })) || []

    channelOptions.splice(0, channelOptions.length, ...validChannels)
  } catch (error) {
    console.error('获取通道列表失败:', error)
  }
}

// 监听props变化，更新本地数据
watch(
  () => props.model,
  (newVal) => {
    if (newVal) {
      formModel.value = {
        ...newVal,
      }
    }
  },
  { immediate: true, deep: true },
)

// 暴露方法给父组件
defineExpose({
  validate: async (): Promise<TemplateItem | null> => {
    if (formRef.value && formModel.value) {
      try {
        // 验证基本字段
        await formRef.value.validate()
        // 返回验证通过后的表单数据
        return { ...formModel.value }
      } catch (error) {
        console.error('表单验证失败:', error)
        // 滚动到第一个错误字段
        const firstErrorElement = document.querySelector('.ant-form-item-has-error')
        if (firstErrorElement) {
          firstErrorElement.scrollIntoView({ behavior: 'smooth', block: 'center' })
        }
        return null
      }
    }
    return null
  },
  resetFields: () => {
    if (formRef.value) {
      formRef.value.resetFields()
    }
  },
})
</script>

<style scoped></style>
