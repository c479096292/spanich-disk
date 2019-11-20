package handler

import (
	"fmt"
	"github.com/c479096292/spinach-disk/db"
	"github.com/c479096292/spinach-disk/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

func UploadHandler(c *gin.Context)  {
	if c.Request.Method == "GET" {
		loc, _  := os.Getwd()
		fmt.Println("location:",loc)
		c.Redirect(http.StatusMovedPermanently,"./static/view/index.html")
		return
	} else if c.Request.Method == "POST" {
		file, head, err := c.Request.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data, err:%s\n", err.Error())
			return
		}

		if err != nil {
			fmt.Printf("Failed to acquire file handle, err:%s\n", err.Error())
			return
		}
		defer file.Close()
		fileMeta := utils.FileMeta{
			FileName:head.Filename,
			FileSize:head.Size,
			Location:"/tmp/"+ head.Filename,
			UploadAt:time.Now().Format("2006-01-02 15:04:05"),
		}

		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("Failed to create file, err:%s\n", err.Error())
			return
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save data into file, err:%s\n", err.Error())
			return
		}

		newFile.Seek(0, 0)
		fileMeta.FileSha1 = utils.FileSha1(newFile)
		newFile.Seek(0, 0)
		ok := db.SaveMeta(fileMeta.FileSha1,fileMeta.FileName,fileMeta.FileSize,fileMeta.Location)
		if ok {
			c.Redirect(http.StatusMovedPermanently,"/static/view/home.html")
		} else {
			c.JSON(500,"Upload Failed")
			return
		}
	}
}


