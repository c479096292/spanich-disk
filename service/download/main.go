package main

import (
	"fmt"
	"github.com/c479096292/spinach-disk/common"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"time"
	dbproxy "github.com/c479096292/spinach-disk/service/dbproxy/client"
	dlProto "github.com/c479096292/spinach-disk/service/download/proto"
	dlRpc "github.com/c479096292/spinach-disk/service/download/handler"
	"github.com/c479096292/spinach-disk/service/download/route"
	cfg "github.com/c479096292/spinach-disk/service/download/config"
)

func startRPCService() {
	reg := consul.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	service := micro.NewService(
		micro.Name("go.micro.service.download"),
		micro.Registry(reg),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
		micro.Flags(common.CustomFlags...),
	)
	service.Init()

	// 初始化dbproxy client
	dbproxy.Init(service)

	dlProto.RegisterDownloadServiceHandler(service.Server(), new(dlRpc.Download))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startAPIService() {
	router := route.Router()
	router.Run(cfg.DownloadServiceHost)
}

func main() {
	// api 服务
	go startAPIService()

	// rpc 服务
	startRPCService()
}