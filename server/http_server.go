package server

import (
	"net/http"
	"strconv"

	"github.com/Poomipat-Ch/StockManagement/routers"
	"github.com/Poomipat-Ch/StockManagement/routers/user"
	"github.com/Poomipat-Ch/StockManagement/services"
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

	userService := services.NewUserService(s.db, s.v)

	v1 := s.gin.Group("/api/v1")
	{
		routes := []routers.Routers{
			user.NewUserRouter(v1, userService),
		}

		routers.MapRoutes(routes...)
	}
}
