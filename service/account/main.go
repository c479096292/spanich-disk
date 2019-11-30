package main

import (
	"github.com/c479096292/spinach-disk/common"
	"github.com/c479096292/spinach-disk/service/account/handler"
	proto "github.com/c479096292/spinach-disk/service/account/proto"
	dbproxy "github.com/c479096292/spinach-disk/service/dbproxy/client"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"log"
	"time"
)

func main() {
	reg := consul.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Registry(reg),
		micro.RegisterTTL(time.Second * 10),
		micro.RegisterInterval(time.Second * 5),
		micro.Flags(common.CustomFlags...),
		)

	service.Init()

	dbproxy.Init(service)
	proto.RegisterUserServiceHandler(service.Server(), new(handler.User))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}