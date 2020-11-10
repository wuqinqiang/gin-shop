package routes

import "github.com/gin-gonic/gin"

func RegisterRoute(e *gin.Engine) {
	RegisterCategory(e)
	RegisterUser(e)
	RegisterProduct(e)
}
