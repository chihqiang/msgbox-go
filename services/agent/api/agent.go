// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"chihqiang/msgbox-go/services/agent/api/internal/config"
	"chihqiang/msgbox-go/services/agent/api/internal/handler"
	"chihqiang/msgbox-go/services/agent/api/internal/svc"
	"chihqiang/msgbox-go/services/common/validators"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/agent-api.yaml", "the config file")

func init() {
	httpx.SetValidator(validators.New())
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.PrintRoutes()
	server.Start()
}
