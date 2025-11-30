<template>
  <a-layout>
    <a-layout-content class="home-content">
      <div class="home-container">
        <!-- Hero Section -->
        <div>
          <a-row align="middle">
            <a-col :md="12">
              <div>
                <a-typography-title class="hero-title" :level="1">开启你的云消息推送</a-typography-title>
                <a-typography-paragraph class="hero-description">
                  一行代码即可实现微信模板消息、短信、邮件、钉钉、飞书等平台消息一键推送！
                </a-typography-paragraph>
                <a-space size="middle">
                  <a-button type="primary" size="large" href="#">
                    立即开始
                  </a-button>
                  <a-button size="large" href="#">
                    查看文档
                  </a-button>
                </a-space>
              </div>
            </a-col>
            <a-col :md="12">
              <div class="image-container">
                <img
                  src="@/assets/hometalk.jpg"
                  alt="home talk"
                />
              </div>
            </a-col>
          </a-row>
        </div>

        <!-- Features -->
        <div class="features-section">
          <a-row :gutter="[24, 24]">
            <a-col :xs="24" :md="8">
              <a-card hoverable class="feature-card">
                <div class="feature-header">
                  <div class="feature-icon">
                    💬
                  </div>
                  <a-typography-title :level="4" class="feature-title">微信模板消息</a-typography-title>
                </div>
                <a-typography-paragraph class="feature-description">微信模板消息即时发送，消息的推送消息将到用户微信上！</a-typography-paragraph>
              </a-card>
            </a-col>

            <a-col :xs="24" :md="8">
              <a-card hoverable class="feature-card">
                <div class="feature-header">
                  <div class="feature-icon">
                    📱
                  </div>
                  <a-typography-title :level="4" class="feature-title">短信</a-typography-title>
                </div>
                <a-typography-paragraph class="feature-description">通过短信sms可快速将消息通知到手机！</a-typography-paragraph>
              </a-card>
            </a-col>

            <a-col :xs="24" :md="8">
              <a-card hoverable class="feature-card">
                <div class="feature-header">
                  <div class="feature-icon">
                    ✉️
                  </div>
                  <a-typography-title :level="4" class="feature-title">邮箱</a-typography-title>
                </div>
                <a-typography-paragraph class="feature-description">订阅邮箱可将通知消息推送到邮箱客户端！</a-typography-paragraph>
              </a-card>
            </a-col>

            <a-col :xs="24" :md="8">
              <a-card hoverable class="feature-card">
                <div class="feature-header">
                  <div class="feature-icon">
                    🔔
                  </div>
                  <a-typography-title  :level="4" class="feature-title">钉钉</a-typography-title>
                </div>
                <a-typography-paragraph class="feature-description">对接钉钉机器人发送通知消息。</a-typography-paragraph>
              </a-card>
            </a-col>

            <a-col :xs="24" :md="8">
              <a-card hoverable class="feature-card">
                <div class="feature-header">
                  <div class="feature-icon">
                    📨
                  </div>
                  <a-typography-title :level="4" class="feature-title">飞书</a-typography-title>
                </div>
                <a-typography-paragraph class="feature-description">对接飞书机器人快速发送消息。</a-typography-paragraph>
              </a-card>
            </a-col>

            <a-col :xs="24" :md="8">
              <a-card hoverable class="feature-card">
                <div class="feature-header">
                  <div class="feature-icon">
                    ⚙️
                  </div>
                  <a-typography-title :level="4" class="feature-title">WebHook</a-typography-title>
                </div>
                <a-typography-paragraph class="feature-description">自定义Webhook地址接收消息。</a-typography-paragraph>
              </a-card>
            </a-col>
          </a-row>
        </div>
      </div>
    </a-layout-content>
  </a-layout>
</template>

<script setup lang="ts">

</script>

<style scoped>
.home-content {
  padding: 24px;
  background-color: #f5f5f5;
}

.home-container {
  max-width: 1200px;
  margin: 0 auto;
}

.hero-title {
  font-size: 3.5rem;
}

.hero-description {
  color: #6b7280;
}

.image-container {
  display: flex;
  justify-content: center;
}

.features-section {
  padding: 64px 0;
}

.feature-card {
  padding: 24px;
  border-radius: 8px;
}

.feature-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 12px;
}

.feature-icon {
  width: 40px;
  height: 40px;
  background: #e6f4ff;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
}

.feature-title {
  margin: 0;
  font-weight: 600;
  font-size: 16px;
}

.feature-description {
  margin: 0;
  color: #6b7280;
  font-size: 14px;
}
</style>
