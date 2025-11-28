<template>
  <a-layout>
    <a-layout-content style="padding: 24px; background-color: #f5f5f5;">
      <div style="max-width: 1200px; margin: 0 auto;">
        <!-- Hero Section -->
        <div style="padding: 80px 0;">
          <a-row :gutter="[40, 40]" align="middle">
            <a-col :md="12">
              <div>
                <a-typography-title level="1" style="font-size: 2.5rem; margin-bottom: 16px;">开启你的云消息推送</a-typography-title>
                <a-typography-paragraph style="color: #6b7280; margin-bottom: 24px; font-size: 16px;">
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
              <div style="display: flex; justify-content: center;">
                <img
                  src="https://images.unsplash.com/photo-1665686377065-08ba896d16fd?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=700&h=800&q=80"
                  alt="Hero Image"
                  style="border-radius: 8px; box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1); width: 100%; height: auto;"
                />
              </div>
            </a-col>
          </a-row>
        </div>

        <!-- Features -->
        <div style="padding: 64px 0;">
          <a-row :gutter="[24, 24]">
            <a-col :xs="24" :md="8">
              <a-card hoverable style="padding: 24px; border-radius: 8px;">
                <div style="display: flex; align-items: center; gap: 12px; margin-bottom: 12px;">
                  <div style="width: 40px; height: 40px; background: #e6f4ff; display: flex; align-items: center; justify-content: center; border-radius: 50%;">
                    💬
                  </div>
                  <a-typography-title level="4" style="margin: 0; font-weight: 600; font-size: 16px;">微信模板消息</a-typography-title>
                </div>
                <a-typography-paragraph style="margin: 0; color: #6b7280; font-size: 14px;">微信模板消息即时发送，消息的推送消息将到用户微信上！</a-typography-paragraph>
              </a-card>
            </a-col>

            <a-col :xs="24" :md="8">
              <a-card hoverable style="padding: 24px; border-radius: 8px;">
                <div style="display: flex; align-items: center; gap: 12px; margin-bottom: 12px;">
                  <div style="width: 40px; height: 40px; background: #e6f4ff; display: flex; align-items: center; justify-content: center; border-radius: 50%;">
                    📱
                  </div>
                  <a-typography-title level="4" style="margin: 0; font-weight: 600; font-size: 16px;">短信</a-typography-title>
                </div>
                <a-typography-paragraph style="margin: 0; color: #6b7280; font-size: 14px;">通过短信sms可快速将消息通知到手机！</a-typography-paragraph>
              </a-card>
            </a-col>

            <a-col :xs="24" :md="8">
              <a-card hoverable style="padding: 24px; border-radius: 8px;">
                <div style="display: flex; align-items: center; gap: 12px; margin-bottom: 12px;">
                  <div style="width: 40px; height: 40px; background: #e6f4ff; display: flex; align-items: center; justify-content: center; border-radius: 50%;">
                    ✉️
                  </div>
                  <a-typography-title level="4" style="margin: 0; font-weight: 600; font-size: 16px;">邮箱</a-typography-title>
                </div>
                <a-typography-paragraph style="margin: 0; color: #6b7280; font-size: 14px;">订阅邮箱可将通知消息推送到邮箱客户端！</a-typography-paragraph>
              </a-card>
            </a-col>

            <a-col :xs="24" :md="8">
              <a-card hoverable style="padding: 24px; border-radius: 8px;">
                <div style="display: flex; align-items: center; gap: 12px; margin-bottom: 12px;">
                  <div style="width: 40px; height: 40px; background: #e6f4ff; display: flex; align-items: center; justify-content: center; border-radius: 50%;">
                    🔔
                  </div>
                  <a-typography-title level="4" style="margin: 0; font-weight: 600; font-size: 16px;">钉钉</a-typography-title>
                </div>
                <a-typography-paragraph style="margin: 0; color: #6b7280; font-size: 14px;">对接钉钉机器人发送通知消息。</a-typography-paragraph>
              </a-card>
            </a-col>

            <a-col :xs="24" :md="8">
              <a-card hoverable style="padding: 24px; border-radius: 8px;">
                <div style="display: flex; align-items: center; gap: 12px; margin-bottom: 12px;">
                  <div style="width: 40px; height: 40px; background: #e6f4ff; display: flex; align-items: center; justify-content: center; border-radius: 50%;">
                    📨
                  </div>
                  <a-typography-title level="4" style="margin: 0; font-weight: 600; font-size: 16px;">飞书</a-typography-title>
                </div>
                <a-typography-paragraph style="margin: 0; color: #6b7280; font-size: 14px;">对接飞书机器人快速发送消息。</a-typography-paragraph>
              </a-card>
            </a-col>

            <a-col :xs="24" :md="8">
              <a-card hoverable style="padding: 24px; border-radius: 8px;">
                <div style="display: flex; align-items: center; gap: 12px; margin-bottom: 12px;">
                  <div style="width: 40px; height: 40px; background: #e6f4ff; display: flex; align-items: center; justify-content: center; border-radius: 50%;">
                    ⚙️
                  </div>
                  <a-typography-title level="4" style="margin: 0; font-weight: 600; font-size: 16px;">WebHook</a-typography-title>
                </div>
                <a-typography-paragraph style="margin: 0; color: #6b7280; font-size: 14px;">自定义Webhook地址接收消息。</a-typography-paragraph>
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
