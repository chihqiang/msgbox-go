GO_ZERO_VERSION=v1.9.2

init-env:
	@command -v goctl >/dev/null 2>&1 || { go install github.com/zeromicro/go-zero/tools/goctl@$(GO_ZERO_VERSION); }
	@goctl env check --install --verbose --force
	@go get github.com/zeromicro/go-zero@$(GO_ZERO_VERSION)
	cp -R devops/goctl/template/ ~/.goctl/1.9.2/

docker-start-env:
	docker-compose -f devops/docker/docker-compose-env.yml up -d

docker-stop-env:
	docker-compose -f devops/docker/docker-compose-env.yml down

generate-api:
	goctl api format --dir .
	goctl api go -api services/agent/api/agent.api -dir services/agent/api -style gozero
	goctl api go -api services/gateway/api/gateway.api -dir services/gateway/api -style gozero

generate-swagger:
	goctl api format --dir .
	goctl api swagger -api services/agent/api/agent.api -dir  services/agent/api
	goctl api swagger -api services/gateway/api/gateway.api -dir  services/gateway/api

generate-rpc:
	goctl rpc protoc services/gateway/rpc/gateway.proto --go_out=services/gateway/rpc/pb --go-grpc_out=services/gateway/rpc/pb --zrpc_out=services/gateway/rpc --client=true
