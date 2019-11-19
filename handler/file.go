package handler

import (
	"fmt"
	"github.com/c479096292/spinach-disk/utils"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func UploadHandler(c *gin.Context)  {
	if c.Request.Method == "GET" {
		c.Redirect("200",)
		return
	} else if c.Request.Method == "POST" {
		file , err := c.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data, err:%s\n", err.Error())
			return
		}
		fileObj, err := file.Open()
		if err != nil {
			fmt.Printf("Failed to acquire file handle, err:%s\n", err.Error())
			return
		}
		defer fileObj.Close()
		fileMeta := utils.FileMeta{
			FileName:file.Filename,
			FileSize:file.Size,
			Location:"/tmp/"+ file.Filename,
			UploadAt:time.Now().Format("2006-01-02 15:04:05"),
		}
		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("Failed to create file, err:%s\n", err.Error())
			return
		}
		defer newFile.Close()

	}
}
