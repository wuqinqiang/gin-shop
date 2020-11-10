package routes

import "github.com/gin-gonic/gin"

func RegisterProduct(e *gin.Engine) {
	g := e.Group("product")
	g.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "nihao"})
	})
}
