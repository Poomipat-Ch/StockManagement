package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Poomipat-Ch/StockManagement/configs"
	"github.com/Poomipat-Ch/StockManagement/pkg/postgres"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	config *configs.Config
	v      *validator.Validate
	db     *sqlx.DB
	gin    *gin.Engine
	srv    *http.Server
}

func NewServer(config *configs.Config) *Server {
	return &Server{
		config: config,
		v:      validator.New(),
		gin:    gin.Default(),
	}
}

func (s *Server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	db, err := postgres.NewPostgresDatabase(s.config.Postgres)

	if err != nil {
		return err
	}

	s.db = db

	go func() {
		if err := s.runHttpServer(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
			cancel()
		}
	}()

	log.Printf("Server Started!! on port: %v", s.config.Port)

	<-ctx.Done()

	log.Println("Server Shutting Down...")

	if err = s.srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown with error: ", err)
	}

	log.Println("Server Shutdown Properly")

	return nil
}
