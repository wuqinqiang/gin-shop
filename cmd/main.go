package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wuqinqiang/gin-shop/routes"
)

func main() {
	r := gin.Default()
	routes.RegisterRoute(r)
	_ = r.Run(":9091")
}
