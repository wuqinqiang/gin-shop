package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wuqinqiang/gin-shop/routes"
	"go.uber.org/dig"
)

func initContainer() *dig.Container {
	d := dig.New()
	return d
}

func main() {
	r := gin.Default()
	routes.RegisterRoute(r)
	_ = r.Run(":9091")
}
