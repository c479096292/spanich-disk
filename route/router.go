package route

import (
	"github.com/c479096292/spinach-disk/handler"
	"github.com/gin-gonic/gin"
)

func LoadRouterHandler()  {
	router := gin.Default()

	//router.Use(middleware.LimitIP())
	//router.GET("/login", func(c *gin.Context) {
	//	c.JSON(200,"login page")
	//})
	//router.Use(middleware.JwtAuth())

	router.GET("/index", func(c *gin.Context) {
		c.JSON(200,"index page")
	})

	fileGroup := router.Group("/article")
	{
		fileGroup.GET("/total", handler.UploadHandler)
		fileGroup.POST("/paged", controller.GetArticlePaged())
		fileGroup.POST("/articles", controller.GetArticlesByPersonID())
		fileGroup.POST("/find", controller.FindArticleByTitle())
		fileGroup.POST("/new", controller.CreateNewArticle())

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
