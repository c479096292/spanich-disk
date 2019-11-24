package main

import (
	"github.com/c479096292/spinach-disk/common"
	"github.com/c479096292/spinach-disk/service/dbproxy/config"
	"github.com/micro/go-micro"
	"log"
	"time"
	"github.com/micro/cli"
	dbConn "github.com/c479096292/spinach-disk/service/dbproxy/conn"
	dbProxy "github.com/c479096292/spinach-disk/service/dbproxy/proto"
	dbRpc "github.com/c479096292/spinach-disk/service/dbproxy/handler"
)

func startRpcService()  {
	service := micro.NewService(
		micro.Name("go.micro.service.dbproxy"),
		micro.RegisterTTL(time.Second * 10), // 声明超时时间, 避免consul不主动删掉已失去心跳的服务节点
		micro.RegisterInterval(time.Second * 5), // 每5s发生一次心跳
		micro.Flags(common.CustomFlags...),  // 向service.Init传递配置参数.
		)

	service.Init(
		micro.Action(func(c *cli.Context) {
			// 检查是否指定dbhost
			dbhost := c.String("dbhost")
			if len(dbhost) > 0 {
				log.Println("custom db address: " + dbhost)
				config.UpdateDBHost(dbhost)
			}
		}),
	)

	// 初始化db connection
	dbConn.InitDBConn()

	dbProxy.RegisterDBProxyServiceHandler(service.Server(), new(dbRpc.DBProxy))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}

func main() {
	startRpcService()
}