<template>
  <div class="channel-form-container">
    <a-form v-if="formModel" :model="formModel" :rules="rules" layout="vertical" ref="formRef" class="modern-form">
      <!-- 创建时显示code字段，编辑时隐藏 -->
      <div class="form-row" v-if="!isEdit">
        <a-form-item label="通道编码" name="code" class="form-item">
          <a-input v-model:value="formModel.code" placeholder="请输入通道编码" class="modern-input" />
        </a-form-item>

        <a-form-item label="通道名称" name="name" class="form-item">
          <a-input v-model:value="formModel.name" placeholder="请输入通道名称" class="modern-input" />
        </a-form-item>
      </div>

      <div class="form-row" v-else>
        <a-form-item label="通道名称" name="name" class="form-item">
          <a-input v-model:value="formModel.name" placeholder="请输入通道名称" class="modern-input" />
        </a-form-item>
      </div>

      <div class="form-row">
        <a-form-item label="服务商（不同的服务商不同的配置）" name="vendor_name" class="form-item full-width">
          <a-select v-model:value="formModel.vendor_name" placeholder="请选择服务商名称" class="modern-select"
            @change="handleChangeConfig">
            <a-select-option v-for="vendor in configs" :key="vendor.name" :value="vendor.name">
              {{ vendor.label }}
            </a-select-option>
          </a-select>
        </a-form-item>
      </div>

      <!-- 动态配置表单 -->
      <div v-if="currentVendor" class="config-form-container form-item full-width">
        <div class="config-section-title">通道配置</div>
        <div v-for="(configItem, index) in currentVendor.configs" :key="index" class="config-section">
          <a-form-item :label="configItem.label" :required="configItem.required" class="nested-form-item">
            <a-input v-if="configItem.type === 'text'" v-model:value="configForm[configItem.name]"
              :placeholder="configItem.placeholder" :default-value="configItem.default" :required="configItem.required"
              class="modern-input" @blur="() => validateDynamicFields()" />
            <div v-if="dynamicFieldErrors[configItem.name]" class="ant-form-item-explain">
              {{ dynamicFieldErrors[configItem.name] }}
            </div>
          </a-form-item>
        </div>
      </div>


      <a-form-item label="状态" name="status" class="form-item status-item">
        <div class="status-container">
          <span class="status-label">{{ formModel.status ? '启用' : '禁用' }}</span>
          <a-switch v-model:checked="formModel.status" size="medium" />
        </div>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, watch, computed, onMounted } from 'vue'
import type { FormInstance } from '@arco-design/web-vue'
import { ChannelItem, ConfigItem, Configs } from "@/model/channel"
import { getChannelConfigs } from '@/api/channel';
// Props定义
interface Props {
  model: ChannelItem | null
}
// 定义props，默认值为null
const props = withDefaults(defineProps<Props>(), { model: null })
// 判断是否为编辑模式
const isEdit = computed(() => {
  return props.model?.id !== undefined && props.model.id !== 0
})
// 表单引用
const formRef = ref<FormInstance | null>(null)
// 本地响应式数据，避免直接修改props
const formModel = ref<ChannelItem | null>(null)
// 配置表单数据
const configForm = ref<Record<string, string>>({})
// 服务商列表
const configs = ref<Configs[]>()
// 表单验证规则
const rules = reactive({
  code: [{ required: !isEdit.value, message: '请输入通道编码', trigger: 'blur' }],
  name: [{ required: true, message: '请输入通道名称', trigger: 'blur' }],
  vendor_name: [{ required: true, message: '请选择服务商', trigger: 'change' }],
})
// 重置配置表单
const resetConfigForm = () => {
  configForm.value = {}
}
// 服务商列表
const fetchChannelConfigs = async () => {
  const { data } = await getChannelConfigs()
  configs.value = data
}
// 初始化服务商列表
onMounted(() => {
  fetchChannelConfigs()
})

// 将表单数据转换为JSON
const syncConfigToJson = () => {
  if (!formModel.value) return

  // 验证必填配置项
  const missingRequiredFields: string[] = []
  if (currentVendor.value) {
    currentVendor.value.configs.forEach((configItem: ConfigItem) => {
      if (configItem.required) {
        const value = configForm.value[configItem.name]
        if (!value || value.trim() === '') {
          missingRequiredFields.push(configItem.label)
        }
      }
    })
  }

  // 如果有必填项未填写，抛出错误
  if (missingRequiredFields.length > 0) {
    const errorMsg = `请填写以下必填配置项: ${missingRequiredFields.join(', ')}`
    console.error(errorMsg)
    // 不抛出错误，而是标记验证错误，让validate方法处理
    return false
  }

  // 构建配置对象
  const config: Record<string, string> = {}
  if (currentVendor.value) {
    currentVendor.value.configs.forEach((configItem: ConfigItem) => {
      const value = configForm.value[configItem.name]
      // 只保存非空值
      if (value !== undefined && value !== null && value.trim() !== '') {
        config[configItem.name] = value.trim() // 去除首尾空格
      }
    })
  }

  // 确保formModel.value存在config属性
  if (!formModel.value.config) {
    formModel.value.config = {}
  }

  // 更新配置对象
  formModel.value.config = config
  return true
}


// 当前选中的服务商
const currentVendor = computed(() => {
  if (!formModel.value?.vendor_name) return null
  return configs?.value?.find(vendor => vendor.name === formModel.value?.vendor_name) || null
})

// 存储当前的配置JSON，以便在服务商数据加载后重新解析
const currentConfigJson = ref('{}');

// 解析配置JSON到表单
const parseConfig = (configJson: string) => {
  // 保存当前配置JSON
  currentConfigJson.value = configJson;

  try {
    // 确保configJson是字符串类型
    const configStr = typeof configJson === 'string' ? configJson : JSON.stringify(configJson || {})
    const config = JSON.parse(configStr || '{}') as Record<string, unknown>
    resetConfigForm()

    // 检查currentVendor是否已准备好
    if (currentVendor.value) {
      // 设置表单字段值
      currentVendor.value.configs.forEach((configItem: ConfigItem) => {
        const configValue = config[configItem.name]

        // 类型安全的赋值，确保值是字符串类型
        if (configValue !== undefined && configValue !== null) {
          configForm.value[configItem.name] = String(configValue)
        } else if (configItem.default) {
          configForm.value[configItem.name] = configItem.default
        }
        // 即使没有默认值和配置值，也为必填项创建空字段以支持验证
        else if (configItem.required) {
          configForm.value[configItem.name] = ''
        }
      })
    }
  } catch (error) {
    console.error('配置解析错误:', error)
    // 即使解析失败，也为必填配置项创建空字段以支持验证
    if (currentVendor.value) {
      resetConfigForm()
      currentVendor.value.configs.forEach((configItem: ConfigItem) => {
        if (configItem.required) {
          configForm.value[configItem.name] = ''
        }
      })
    }
  }
}

// 处理服务商变化
const handleChangeConfig = () => {
  // 清空之前的验证错误
  dynamicFieldErrors.value = {}

  // 重置配置表单
  resetConfigForm()

  // 设置默认配置值和初始化必填字段
  if (currentVendor.value) {
    currentVendor.value.configs.forEach(config => {
      // 设置默认值（如果有）
      if (config.default) {
        configForm.value[config.name] = config.default
      }
      // 对于必填项，即使没有默认值也要创建空字段
      else if (config.required) {
        configForm.value[config.name] = ''
      }
    })

    // 只有在表单存在且有配置的情况下才同步到JSON
    if (formModel.value) {
      syncConfigToJson()
    }
  }

  // 触发表单重新验证
  if (formRef.value) {
    formRef.value.validateFields(['vendor_name'])
  }
}
// 监听props变化，更新本地数据
watch(
  () => props.model,
  (newVal) => {
    if (newVal) {
      // 处理字段映射，确保兼容性
      formModel.value = {
        ...newVal,
      }
      // 解析配置并更新表单
      // 如果config是字符串则直接使用，如果是对象则转换为字符串
      const configStr = typeof newVal.config === 'string' ? newVal.config : JSON.stringify(newVal.config || {})
      parseConfig(configStr)
    } else {
      formModel.value = null
      resetConfigForm()
    }
  },
  { immediate: true, deep: true },
)

// 监听currentVendor变化，当服务商数据加载完成后重新解析配置
watch(
  () => currentVendor.value,
  (newVendor) => {
    // 只有当服务商数据已加载且有配置数据时，才重新解析
    if (newVendor && currentConfigJson.value) {
      // 使用保存的配置JSON重新解析，确保access_token等字段正确初始化
      parseConfig(currentConfigJson.value);
    }
  },
  { immediate: true }
)

// 监听配置表单变化
watch(
  () => configForm.value,
  () => {
    syncConfigToJson()
  },
  { deep: true }
)
// 验证动态配置字段
const validateDynamicFields = (): boolean => {
  let isValid = true
  const errors: Record<string, string> = {}

  if (currentVendor.value) {
    currentVendor.value.configs.forEach((configItem: ConfigItem) => {
      if (configItem.required) {
        const value = configForm.value[configItem.name]
        if (!value || value.trim() === '') {
          errors[configItem.name] = `请输入${configItem.label}`
          isValid = false
        }
      }
    })
  }

  // 存储验证错误信息，供UI显示
  dynamicFieldErrors.value = errors
  return isValid
}

// 动态字段验证错误
const dynamicFieldErrors = ref<Record<string, string>>({})

// 暴露方法给父组件
defineExpose({
  validate: async (): Promise<ChannelItem | null> => {
    if (formRef.value && formModel.value) {
      try {
        // 验证基本字段
        await formRef.value.validate()

        // 验证并同步配置到JSON
        const configValid = syncConfigToJson()
        if (!configValid) {
          // 如果配置验证失败，运行动态字段验证以显示具体错误
          validateDynamicFields()
          throw new Error('动态配置字段验证失败')
        }

        // 再次验证动态配置字段以确保完整性
        const dynamicValid = validateDynamicFields()
        if (!dynamicValid) {
          throw new Error('动态配置字段验证失败')
        }

        // 确保配置对象不为空
        if (!formModel.value.config || Object.keys(formModel.value.config).length === 0) {
          console.warn('配置数据为空')
        }

        // 返回验证通过后的表单数据
        return { ...formModel.value } // 返回副本避免直接引用
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
      resetConfigForm()
      // 清空动态字段验证错误
      dynamicFieldErrors.value = {}
    }
  },
})
</script>

<style scoped>
/* 表单容器 */
.channel-form-container {
  padding: 8px;
}

/* 现代化表单样式 */
.modern-form {
  padding: 0;
  background-color: transparent;
  border-radius: 0;
  box-shadow: none;
  border: none;
}

/* 表单项样式 */
.form-item {
  margin-bottom: 24px;
  transition: all 0.3s ease;
  flex: 1;
  min-width: 0;
  margin-right: 20px;
  opacity: 1;
  visibility: visible;
}

.form-item:last-child {
  margin-right: 0;
}

.form-row {
  display: flex;
  gap: 20px;
  margin-bottom: 4px;
  align-items: flex-start;
}

.full-width {
  flex: 1 1 100%;
  margin-right: 0;
}

/* 配置表单样式 */
.config-form-container {
  background-color: #fafafa;
  border-radius: 8px;
  padding: 16px;
  border: 1px solid #e8e8e8;
}

.config-section-title {
  font-weight: 600;
  font-size: 14px;
  color: #1a1a1a;
  margin-bottom: 16px;
}

.config-section {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

/* 嵌套表单项样式 */
.nested-form-item {
  margin-bottom: 16px;
  margin-right: 0;
}

.nested-form-item:last-child {
  margin-bottom: 0;
}

/* 选择器样式 */
.modern-select {
  width: 100%;
  transition: all 0.3s ease;
  border-radius: 6px !important;
}

.modern-select:hover {
  border-color: #4096ff !important;
}

.modern-select:focus-within {
  border-color: #4096ff !important;
  box-shadow: 0 0 0 2px rgba(64, 158, 255, 0.2) !important;
}

/* 标签样式 */
.ant-form-item-label>label {
  font-weight: 600;
  font-size: 14px;
  color: #1a1a1a;
  padding-bottom: 8px;
  transition: color 0.3s ease;
}

.ant-form-item-label>label:hover {
  color: #4096ff;
}

/* 输入框样式 */
.modern-input,
.modern-select .ant-select-selector,
.modern-textarea {
  border-radius: 10px;
  border: 1.5px solid #e8e8e8;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  font-size: 14px;
  height: 42px;
  background-color: #ffffff;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.03);
  opacity: 1 !important;
  visibility: visible !important;
}

/* 输入框悬停效果 */
.modern-input:hover,
.modern-select .ant-select-selector:hover,
.modern-textarea:hover {
  border-color: #69b1ff;
  box-shadow: 0 2px 8px rgba(64, 158, 255, 0.15);
  transform: translateY(-1px);
}

/* 输入框聚焦效果 */
.modern-input:focus,
.modern-select.ant-select-focused .ant-select-selector,
.modern-textarea:focus {
  border-color: #4096ff;
  box-shadow: 0 0 0 3px rgba(64, 158, 255, 0.2);
  outline: none;
  transform: translateY(-1px);
}

/* 文本域特殊处理 */
.modern-textarea {
  height: auto;
  min-height: 120px;
  resize: vertical;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  font-size: 13px;
  line-height: 1.5;
}



/* 状态项特殊样式 */
.status-item {
  margin-top: 8px;
}

.status-container {
  display: flex;
  align-items: center;
  gap: 12px;
}

.status-label {
  font-size: 14px;
  font-weight: 500;
  color: #595959;
  transition: all 0.3s ease;
}

/* 状态标签样式 */
.status-label {
  font-size: 14px;
  font-weight: 500;
  color: #595959;
  margin-right: 12px;
  line-height: 32px;
}

/* 修复配置文本域的显示 */
.modern-textarea {
  height: auto !important;
  min-height: 120px !important;
  resize: vertical !important;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  font-size: 13px;
  line-height: 1.5;
  opacity: 1 !important;
  visibility: visible !important;
  display: block !important;
}

/* 错误状态样式 */
.ant-form-item-has-error .modern-input,
.ant-form-item-has-error .modern-select .ant-select-selector,
.ant-form-item-has-error .modern-textarea {
  border-color: #ff4d4f;
  box-shadow: 0 0 0 2px rgba(255, 77, 79, 0.1);
}

.ant-form-item-has-error .modern-input:hover,
.ant-form-item-has-error .modern-select .ant-select-selector:hover,
.ant-form-item-has-error .modern-textarea:hover {
  border-color: #ff7875;
  box-shadow: 0 0 0 3px rgba(255, 77, 79, 0.15);
}

.ant-form-item-has-error .modern-input:focus,
.ant-form-item-has-error .modern-select.ant-select-focused .ant-select-selector,
.ant-form-item-has-error .modern-textarea:focus {
  border-color: #ff4d4f;
  box-shadow: 0 0 0 3px rgba(255, 77, 79, 0.25);
}

/* 错误提示样式 */
.ant-form-item-explain {
  color: #ff4d4f;
  font-size: 12px;
  margin-top: 6px;
  padding-left: 4px;
}

/* 必填项标识 */
.ant-form-item-required::before {
  color: #ff4d4f;
  margin-right: 4px;
  font-weight: bold;
}

/* 加载状态 */
.ant-loading {
  color: #4096ff;
}

/* 自定义滚动条 */
.modern-textarea::-webkit-scrollbar,
.ant-select-dropdown::-webkit-scrollbar {
  width: 6px;
}

.modern-textarea::-webkit-scrollbar-track,
.ant-select-dropdown::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.modern-textarea::-webkit-scrollbar-thumb,
.ant-select-dropdown::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
  transition: background 0.3s ease;
}

.modern-textarea::-webkit-scrollbar-thumb:hover,
.ant-select-dropdown::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .form-row {
    flex-direction: column;
    gap: 0;
  }

  .form-item {
    margin-right: 0;
    margin-bottom: 20px;
  }

  .channel-form-container {
    padding: 4px;
  }

  .ant-form-item-label>label {
    font-size: 13px;
  }

  .modern-input,
  .modern-select .ant-select-selector,
  .modern-textarea {
    font-size: 13px;
    height: 38px;
  }

  .modern-textarea {
    min-height: 100px;
  }

  .status-container {
    gap: 8px;
  }

  .status-label {
    font-size: 13px;
  }
}

/* 动画效果 */
@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.form-item {
  animation: slideIn 0.4s ease forwards;
}

.form-item:nth-child(1) {
  animation-delay: 0.1s;
}

.form-item:nth-child(2) {
  animation-delay: 0.2s;
}

.form-item:nth-child(3) {
  animation-delay: 0.3s;
}

.form-item:nth-child(4) {
  animation-delay: 0.4s;
}

.form-item:nth-child(5) {
  animation-delay: 0.5s;
}

.form-item:nth-child(6) {
  animation-delay: 0.6s;
}
</style>
