package route

import (
	"github.com/c479096292/spinach-disk/handler"
	"github.com/gin-gonic/gin"
)

func LoadRouterHandler()  {
	router := gin.Default()
	router.Static("/static", "./static")
	//router.LoadHTMLGlob("static/view/*")
	//router.Use(middleware.LimitIP())
	//router.GET("/login", func(c *gin.Context) {
	//	c.JSON(200,"login page")
	//})
	//router.Use(middleware.JwtAuth())

	router.GET("/index", func(c *gin.Context) {
		c.JSON(200,"index page")
	})

	fileGroup := router.Group("/file")
	{
		fileGroup.GET("/upload", handler.UploadHandler)
		fileGroup.POST("/upload", handler.UploadHandler)

		// 秒传接口
		//fileGroup.POST("/fastupload", api.TryFastUploadHandler)
		//
		//// 分块上传接口
		//fileGroup.POST("/mpupload/init", api.InitialMultipartUploadHandler)
		//fileGroup.POST("/mpupload/uppart", api.UploadPartHandler)
		//fileGroup.POST("/mpupload/complete", api.CompleteUploadHandler)
		//
	}
	//userGroup := router.Group("/user")
	//{
	//	userGroup.GET("/total",controller.GetPersonTotal)
	//	userGroup.POST("/paged", controller.GetPersonPage)
	//	userGroup.POST("/new", controller.AddNewPerson)
	//	userGroup.POST("/find", controller.FindPersonByID)
	//	userGroup.POST("/del", controller.DelPerson)
	//	userGroup.POST("/pwd", controller.ModifyPassword)
	//}


	router.Run(":9000")
}
