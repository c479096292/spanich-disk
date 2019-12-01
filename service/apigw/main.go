package main

import (
	"github.com/c479096292/spinach-disk/service/apigw/router"
)


func main() {
	r := router.Router()
	r.Run(":8080")
}