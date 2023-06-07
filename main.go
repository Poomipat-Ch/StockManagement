package main

import (
	"log"

	"github.com/Poomipat-Ch/StockManagement/configs"
	"github.com/Poomipat-Ch/StockManagement/server"
	_ "github.com/lib/pq"
)

func main() {

	// Config
	config, err := configs.InitConfig()

	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(config)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
