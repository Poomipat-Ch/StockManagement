package server

import (
	"net/http"
	"strconv"

	"github.com/Poomipat-Ch/StockManagement/routers"
)

func (s *Server) runHttpServer() error {
	s.mapRoutes()

	srv := &http.Server{
		Addr:    s.config.Host + ":" + strconv.FormatUint(s.config.Port, 10),
		Handler: s.gin,
	}

	s.srv = srv

	return srv.ListenAndServe()
}

func (s *Server) mapRoutes() {
	// grouping with version api
	v1 := s.gin.Group("/api/v1")
	{
		router := routers.NewRouters(v1) // create new router

		router.AddPingRoutes()
		router.AddUserRoutes()
	}
}
