package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterCategory(e *gin.Engine) {
	g := e.Group("category")
	g.GET("/index", func(c *gin.Context) {
		c.JSON(200,gin.H{"message":"nihao"})
	})
}
