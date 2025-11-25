GO_ZERO_VERSION=v1.9.2

init-env:
	@echo "检查是否安装 goctl 工具，如果未安装则进行安装"
	@command -v goctl >/dev/null 2>&1 || { go install github.com/zeromicro/go-zero/tools/goctl@$(GO_ZERO_VERSION); }
	@echo "检查 goctl 环境并安装必要组件"
	@goctl env check --install --verbose --force
	@echo "获取指定版本的 go-zero"
	@go get github.com/zeromicro/go-zero@$(GO_ZERO_VERSION)
	@echo "拷贝自定义模板到 goctl 默认模板目录"
	cp -R devops/goctl/template/ ~/.goctl/1.9.2/

docker-start-env:
	@echo "使用 docker-compose 启动开发环境容器"
	docker-compose -f deploy/docker/docker-compose-env.yml up -d

docker-stop-env:
	@echo "使用 docker-compose 停止开发环境容器"
	docker-compose -f deploy/docker/docker-compose-env.yml down
