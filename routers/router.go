package routers

import (
	"net/http"

	"github.com/Poomipat-Ch/StockManagement/routers/user"
	"github.com/gin-gonic/gin"
)

type Routers struct {
	router *gin.RouterGroup
}

func NewRouters(router *gin.RouterGroup) *Routers {
	return &Routers{router: router}
}

func (r *Routers) AddPingRoutes() {
	r.router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})
}

func (r *Routers) AddUserRoutes() {
	userGroup := r.router.Group("/user")
	user.AddUserRoutes(userGroup)
}
