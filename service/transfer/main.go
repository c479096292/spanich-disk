package main

import (
	"fmt"
	"github.com/c479096292/spinach-disk/common"
	"github.com/c479096292/spinach-disk/config"
	"github.com/c479096292/spinach-disk/mq"
	dbproxy "github.com/c479096292/spinach-disk/service/dbproxy/client"
	"github.com/c479096292/spinach-disk/service/transfer/process"
	_ "github.com/micro/go-plugins/service/kubernetes"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"log"
	"time"
)

func startRPCService() {
	reg := consul.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	service := micro.NewService(
		micro.Name("go.micro.service.transfer"),
		micro.Registry(reg),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
		micro.Flags(common.CustomFlags...),
	)
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 检查是否指定mqhost
			mqhost := c.String("mqhost")
			if len(mqhost) > 0 {
				log.Println("custom mq address: " + mqhost)
				mq.UpdateRabbitHost(mqhost)
			}
		}),
	)

	// 初始化dbproxy client
	dbproxy.Init(service)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startTranserService() {
	if !config.AsyncTransferEnable {
		log.Println("异步转移文件功能目前被禁用，请检查相关配置")
		return
	}
	log.Println("文件转移服务启动中，开始监听转移任务队列...")

	// 初始化mq client
	mq.Init()

	mq.StartConsume(
		config.TransOSSQueueName,
		"transfer_oss",
		process.Transfer)
}

func main() {
	// 文件转移服务
	go startTranserService()

	// rpc 服务
	startRPCService()
}


