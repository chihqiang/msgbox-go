<template>
  <a-layout>
    <a-layout-content
      style="padding: 24px; background-color: #f5f5f5; min-height: calc(100vh - 100px)"
    >
      <div style="max-width: 1200px; margin: 0 auto">
        <!-- 页面标题和说明 -->
        <div style="margin-bottom: 32px">
          <a-typography-title :level="2">API 密钥管理</a-typography-title>
          <a-typography-paragraph
            >管理您的 API Key 和密钥，用于调用 WPUSH 的推送服务接口。</a-typography-paragraph
          >
        </div>

        <!-- 密钥管理区域 -->
        <a-card style="margin-bottom: 24px">
          <a-row :gutter="[16, 16]">
            <!-- API Key 卡片 -->
            <a-col :span="24" :lg="12">
              <a-card>
                <template #title>
                  <div class="card-header">
                    <span>API Key</span>
                    <a-button-group>
                      <a-button type="text" title="复制" @click="copyToClipboard(apiKey)">
                        <template #icon>
                          <copy-outlined />
                        </template>
                      </a-button>
                    </a-button-group>
                  </div>
                </template>
                <a-typography-paragraph type="secondary">用于标识您的应用</a-typography-paragraph>
                <a-input-group compact>
                  <a-input v-model:value="apiKey" read-only />
                </a-input-group>
              </a-card>
            </a-col>

            <!-- API Secret 卡片 -->
            <a-col :span="24" :lg="12">
              <a-card>
                <template #title>
                  <div class="card-header">
                    <span>API Secret</span>
                    <a-button-group>
                      <a-button type="text" title="复制" @click="copyToClipboard(apiSecret)">
                        <template #icon>
                          <copy-outlined />
                        </template>
                      </a-button>
                      <a-button type="text" title="重新生成" @click="resetApiSecret()">
                        <template #icon>
                          <reload-outlined />
                        </template>
                      </a-button>
                    </a-button-group>
                  </div>
                </template>
                <a-typography-paragraph type="secondary">用于接口签名验证</a-typography-paragraph>
                <a-input-group compact>
                  <a-input v-model:value="apiSecret" read-only />
                </a-input-group>
              </a-card>
            </a-col>
          </a-row>
        </a-card>

        <!-- 安全提示区域 -->
        <a-card style="margin-bottom: 24px">
          <a-alert type="warning" show-icon message="安全提示">
            <template #description>
              <ul>
                <li>请妥善保管您的 API 密钥，不要在公开场合泄露</li>
                <li>如果怀疑密钥已泄露，请立即重新生成</li>
                <li>每个密钥对应一个项目，建议为不同项目创建不同的密钥</li>
                <li>定期更换密钥以提高安全性</li>
              </ul>
            </template>
          </a-alert>
        </a-card>

        <!-- API调用示例 -->
        <a-card>
          <template #title>
            <div class="card-header">
              <span>API调用示例</span>
              <a-button type="text" title="复制" @click="copyCurlExample">
                <template #icon>
                  <copy-outlined />
                </template>
              </a-button>
            </div>
          </template>
          <a-typography-paragraph type="secondary"
            >使用curl调用推送服务接口的示例</a-typography-paragraph
          >
          <a-space direction="vertical" style="width: 100%">
            <div class="curl-example-container">
              <pre class="curl-example">{{ curlExample }}</pre>
            </div>

            <!-- 认证头生成规则说明 -->
            <a-card
              size="small"
              style="margin-top: 16px; border-left: 4px solid #1890ff; padding: 16px"
            >
              <a-typography-title :level="5" style="margin-bottom: 12px"
                >API 认证头生成规则</a-typography-title
              >

              <a-typography-paragraph style="margin-bottom: 12px">
                调用 API 时需要在请求头中添加
                <code>Authorization: Basic</code> 认证信息，生成步骤如下：
              </a-typography-paragraph>

              <div
                style="background: #fafafa; padding: 16px; border-radius: 4px; margin-bottom: 12px"
              >
                <div class="step-item">
                  <div class="step-number">1</div>
                  <div class="step-content">
                    <div class="step-title">准备认证信息</div>
                    <p>组合您的 API Key 和 API Secret，格式为：<code>API Key:API Secret</code></p>
                    <p>
                      示例：<code>{{ apiKey }}:{{ apiSecret }}</code>
                    </p>
                  </div>
                </div>
                <div class="step-item">
                  <div class="step-number">2</div>
                  <div class="step-content">
                    <div class="step-title">Base64 编码</div>
                    <p>将组合后的字符串进行 Base64 编码</p>
                    <p>
                      示例：<code>{{ apiKey }}:{{ apiSecret }}</code> 编码后得到
                      <code>{{ authHeader }}</code>
                    </p>
                  </div>
                </div>
                <div class="step-item">
                  <div class="step-number">3</div>
                  <div class="step-content">
                    <div class="step-title">添加到请求头</div>
                    <p>将编码结果添加到请求头中</p>
                    <p>格式：<code>Authorization: Basic {base64编码结果}</code></p>
                    <p>
                      示例：<code>Authorization: Basic {{ authHeader }}</code>
                    </p>
                  </div>
                </div>
              </div>

              <a-typography-paragraph type="secondary" style="margin-bottom: 0; font-size: 13px">
                💡 <strong>提示：</strong> 当前示例中的认证头是基于您实际的 API Key 和 API Secret
                动态生成的，您可以直接复制使用。
              </a-typography-paragraph>
            </a-card>
          </a-space>
        </a-card>
      </div>
    </a-layout-content>
  </a-layout>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import {
  CopyOutlined as CopyOutlined,
  ReloadOutlined as ReloadOutlined,
} from '@ant-design/icons-vue'
import { message, Modal } from 'ant-design-vue'
import { getAgentInfo, resetSecret } from '@/api/agent'

// 密钥数据
const apiKey = ref('')
const apiSecret = ref('')

// API调用示例
const authHeader = computed(() => {
  return btoa(`${apiKey.value}:${apiSecret.value}`)
})

const curlExample = computed(() => {
  return `curl -X POST "${import.meta.env.VITE_GATEWAY_URL}/send"
  -H "Content-Type: application/json"
  -H "Authorization: Basic ${authHeader.value}"
  -d '{
  "template_code": "dingtalk_name_template",
  "receivers": ["18888888888"],
  "variables": {
    "name": "测试"
  }
}'`
})
// 从 API 获取密钥
const fetchAgentInfo = async () => {
  const { data } = await getAgentInfo()
  if (data) {
    apiKey.value = data.agent_no
    apiSecret.value = data.agent_secret || ''
  }
}

onMounted(() => {
  fetchAgentInfo()
})

// 复制到剪贴板
const copyToClipboard = (text: string) => {
  navigator.clipboard
    .writeText(text)
    .then(() => {
      message.success('复制成功')
    })
    .catch(() => {
      message.error('复制失败')
    })
}

// 重新生成密钥
const resetApiSecret = () => {
  Modal.confirm({
    title: '确认重新生成密钥',
    content: '确定要重新生成密钥吗？这将影响您现有的应用程序。',
    okText: '确定',
    cancelText: '取消',
    onOk: async () => {
      const { data } = await resetSecret()
      if (data) {
        apiSecret.value = data.agent_secret || ''
        message.success('密钥已重新生成')
      }
    },
  })
}

// 复制curl示例到剪贴板
const copyCurlExample = () => {
  navigator.clipboard
    .writeText(curlExample.value)
    .then(() => {
      message.success('复制成功')
    })
    .catch(() => {
      message.error('复制失败')
    })
}
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.curl-example-container {
  width: 100%;
}

.curl-example {
  background: #f6f8fa;
  padding: 16px;
  border-radius: 6px;
  overflow-x: auto;
  white-space: pre-wrap;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  font-size: 14px;
  line-height: 1.6;
  border: 1px solid #e1e4e8;
  color: #24292e;
  margin: 0;
}

/* 自定义步骤条样式 */
.step-item {
  display: flex;
  align-items: flex-start;
  margin-bottom: 16px;
  position: relative;
}

.step-item:last-child {
  margin-bottom: 0;
}

.step-number {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 24px;
  height: 24px;
  background-color: #1890ff;
  color: white;
  border-radius: 50%;
  font-size: 12px;
  font-weight: bold;
  margin-right: 12px;
  flex-shrink: 0;
}

.step-content {
  flex: 1;
}

.step-title {
  font-weight: bold;
  margin-bottom: 4px;
  font-size: 14px;
  color: #333;
}

.step-content p {
  margin: 4px 0;
  font-size: 13px;
  color: #666;
}
</style>


