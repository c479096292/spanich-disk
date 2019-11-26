package route

import (
	"github.com/c479096292/spinach-disk/service/download/api"
	"github.com/gin-gonic/gin"
	"github.com/c479096292/spinach-disk/utils"
)

// Router : 路由表配置
func Router() *gin.Engine {
	// gin framework, 包括Logger, Recovery
	router := gin.Default()

	// 处理静态资源
	router.Static("/static/", "./static")

	// // 加入中间件，用于校验token的拦截器(将会从account微服务中验证)
	// router.Use(handler.HTTPInterceptor())

	// 使用gin插件支持跨域请求
	router.Use(utils.Cors())

	// Use之后的所有handler都会经过拦截器进行token校验

	// 文件下载相关接口
	router.GET("/file/download", api.DownloadHandler)
	router.POST("/file/downloadurl", api.DownloadURLHandler)

	return router
}
