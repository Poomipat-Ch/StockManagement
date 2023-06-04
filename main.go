package main

import (
	"log"
	"strconv"

	"github.com/Poomipat-Ch/StockManagement/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// grouping with version api
	v1 := r.Group("/api/v1")
	{
		router := routers.NewRouters(v1) // create new router

		router.AddPingRoutes()
		router.AddUserRoutes()
	}

	var port uint64 = 5000

	log.Printf("Server Started!! on port: %v", port)

	err := r.Run(":" + strconv.FormatUint(port, 10))

	if err != nil {
		panic(err)
	}
}
