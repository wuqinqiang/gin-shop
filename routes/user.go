package routes

import "github.com/gin-gonic/gin"

func RegisterUser(e *gin.Engine) {
	g := e.Group("user")
	g.GET("/index", func(c *gin.Context) {
		c.JSON(200,gin.H{"message":"nihao"})
	})
}
