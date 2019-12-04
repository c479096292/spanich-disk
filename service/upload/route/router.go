package route

import (
	"github.com/c479096292/spinach-disk/service/upload/api"
	"github.com/c479096292/spinach-disk/utils"
	"github.com/gin-gonic/gin"
)

// Router : 路由表配置
func Router() *gin.Engine {
	router := gin.Default()
	router.Use(utils.Cors())
	// 处理静态资源
	router.Static("/static/", "./static")

	// // 加入中间件，用于校验token的拦截器(将会从account微服务中验证)
	// router.Use(handler.HTTPInterceptor())


	// Use之后的所有handler都会经过拦截器进行token校验

	// 文件上传相关接口
	router.POST("/file/upload", api.DoUploadHandler)
	// 秒传接口
	router.POST("/file/fastupload", api.TryFastUploadHandler)

	// 分块上传接口
	router.POST("/file/mpupload/init", api.InitialMultipartUploadHandler)
	router.POST("/file/mpupload/uppart", api.UploadPartHandler)
	router.POST("/file/mpupload/complete", api.CompleteUploadHandler)

	return router
}
