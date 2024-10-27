package main

import (
	"log"

	"github.com/Nokeni/GODS/config"
	"github.com/Nokeni/GODS/internal/db"
	"github.com/Nokeni/GODS/internal/web"
	"github.com/spf13/viper"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	database, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("failed to init database: %v", err)
	}

	server, err := web.NewHTTPServer(database)
	if err != nil {
		log.Fatalf("failed to init web server: %v", err)
	}

	if err := server.Run(":" + viper.GetString("WEB_PORT")); err != nil {
		log.Fatalf("failed to run web server: %v", err)
	}
}
