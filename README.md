# msgbox-go

MsgBox 是一个轻量级、可扩展的消息中心，用于统一处理系统中的通知、消息分发与事件推送。它提供灵活的消息路由机制、可插拔的发送渠道、以及清晰的接口设计，帮助开发者快速构建稳定、高可维护的消息体系。

## 项目介绍

MsgBox 旨在解决分布式系统中的消息通知管理难题，通过统一的 API 接口和灵活的配置，支持多种消息发送渠道（如钉钉、微信、邮件等），实现消息的高效分发、追踪和管理。

## 主要特性

- **统一消息接口**：提供标准化的 API，屏蔽不同渠道的实现差异
- **可插拔通道架构**：支持钉钉、微信、邮件等多种消息发送渠道，易于扩展
- **消息模板管理**：支持创建、编辑和管理消息模板，提高消息发送效率
- **工作流处理**：支持串行和并行的消息处理工作流，灵活控制消息发送逻辑
- **消息追踪**：完整记录消息发送状态、响应和投递信息，便于问题排查
- **多租户支持**：基于 Agent 的多租户架构，支持不同应用独立管理消息
- **高可用设计**：支持消息重试、错误处理和故障恢复机制
- **Web 管理界面**：提供直观的管理界面，方便配置和监控

## 系统架构

### 核心组件

1. **Gateway 服务**：统一入口，处理认证和请求路由
2. **Agent 服务**：多租户隔离，管理消息模板和发送记录
3. **Common 组件**：公共库，包含通道实现、模型定义和工具函数
4. **Web 管理界面**：基于 Vue.js 的管理控制台

### 技术栈

- **后端**：Go 1.23.12 + go-zero 框架
- **数据库**：支持 MySQL/PostgreSQL
- **ORM**：GORM
- **前端**：Vue.js + TypeScript
- **部署**：Docker + Docker Compose

## 快速开始

### 环境要求

- Go 1.23.12+
- MySQL/PostgreSQL
- Node.js 22+
- Docker (可选，用于容器化部署)

### 安装部署

#### 1. 克隆项目

```bash
git clone https://github.com/chihqiang/msgbox-go.git
cd msgbox-go
```

#### 2. 配置数据库

修改配置文件`services/agent/api/etc/agent-api.yaml`和`services/gateway/api/etc/gateway-api.yaml`，设置数据库连接信息：

```bash
DB:
  DBType: mysql
  Username: root
  Password: "123456"
  Host: 127.0.0.1
  Port: 3306
  Database: msgbox
```

#### 3. 启动服务

#### 3.1 启动基础环境（使用 Docker）

如果需要快速搭建开发环境，可以使用 Docker Compose 启动基础服务（如数据库等）：

```bash
cd deploy/docker
docker-compose -f docker-compose-env.yml up -d
```

#### 3.2 运行应用程序

在启动基础环境后，需要单独启动应用程序组件。以下提供两种运行方式（选择其一即可）：

**方式一：直接运行（开发环境推荐）**
```bash
# 启动 Agent 服务
cd services/agent/api
go run agent.go -f etc/agent-api.yaml

# 启动 Gateway 服务（在另一个终端）
cd services/gateway/api
go run gateway.go -f etc/gateway-api.yaml
```

**方式二：构建后运行（生产环境推荐）**
```bash
# 构建并启动 Agent 服务
cd services/agent/api
go build -ldflags="-s -w" -tags no_k8s  -o agent-api agent.go
./agent-api -f etc/agent-api.yaml

# 构建并启动 Gateway 服务（在另一个终端）
cd services/gateway/api
go build -ldflags="-s -w" -tags no_k8s -o gateway-api gateway.go
./gateway-api -f etc/gateway-api.yaml
```

**启动前端（在另一个终端）**
```bash
cd web/agent
pnpm install
pnpm run dev
```

## 项目结构

```
├── deploy/           # 部署相关文件
│   ├── docker/       # Docker 配置
│   └── goctl/        # 代码生成模板
├── pkg/              # 通用工具包
│   ├── clientx/      # HTTP 客户端
│   ├── cryptox/      # 加密工具
│   ├── htmlx/        # HTML 表单处理
│   ├── stringx/      # 字符串工具
│   ├── timex/        # 时间工具
│   └── workflow/     # 工作流处理
├── services/         # 服务层
│   ├── agent/        # 代理服务
│   ├── common/       # 公共组件
│   └── gateway/      # 网关服务
└── web/              # 前端代码
    └── agent/        # 管理界面
```

## 开发指南

### 后端开发

1. **创建新的消息通道**：
   - 在 `services/common/channels/senders` 目录下创建新的发送器实现
   - 实现 `senders.ISender` 接口

2. **添加新的 API**：
   - 使用 goctl 工具生成 API 代码
   - 编辑 API 描述文件

### 前端开发

1. **安装依赖**：
   ```bash
   cd web/agent
   pnpm install
   ```

2. **开发模式**：
   ```bash
   pnpm run dev
   ```

3. **构建生产版本**：
   ```bash
   pnpm run build
   ```

## 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 联系方式

如有问题或建议，欢迎提交 Issue 或联系项目维护者。

