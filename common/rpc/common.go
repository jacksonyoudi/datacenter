package main

import (
	"flag"
	"fmt"

	"github.com/jacksonyoudi/datacenter/common/rpc/common"
	"github.com/jacksonyoudi/datacenter/common/rpc/internal/config"
	"github.com/jacksonyoudi/datacenter/common/rpc/internal/server"
	"github.com/jacksonyoudi/datacenter/common/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/common.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewCommonServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		common.RegisterCommonServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
