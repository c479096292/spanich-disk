package upload

import (
	"fmt"
	"github.com/c479096292/spinach-disk/common"
	"github.com/c479096292/spinach-disk/config"
	"github.com/c479096292/spinach-disk/mq"
	"github.com/micro/go-micro"
	dbproxy "github.com/c479096292/spinach-disk/service/dbproxy/client"
	upRpc "github.com/c479096292/spinach-disk/service/upload/handler"
	cfg "github.com/c479096292/spinach-disk/service/upload/config"
	upProto "github.com/c479096292/spinach-disk/service/upload/proto"
	"github.com/c479096292/spinach-disk/service/upload/route"
	"github.com/micro/cli"
	"log"
	"os"
	"time"
)

func startRPCService() {
	service := micro.NewService(
		micro.Name("go.micro.service.upload"),
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
	// 初始化mq client
	mq.Init()

	upProto.RegisterUploadServiceHandler(service.Server(), new(upRpc.Upload))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startAPIService() {
	router := route.Router()
	router.Run(cfg.UploadServiceHost)
	// service := web.NewService(
	// 	web.Name("go.micro.web.upload"),
	// 	web.Handler(router),
	// 	web.RegisterTTL(10*time.Second),
	// 	web.RegisterInterval(5*time.Second),
	// )
	// if err := service.Init(); err != nil {
	// 	log.Fatal(err)
	// }

	// if err := service.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}

func main() {
	os.MkdirAll(config.TempLocalRootDir, 0777)
	os.MkdirAll(config.TempPartRootDir, 0777)

	// api 服务
	go startAPIService()

	// rpc 服务
	startRPCService()
}