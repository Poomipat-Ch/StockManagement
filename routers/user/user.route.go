package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(router *gin.RouterGroup) {
	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This is user route"})
	})
}
