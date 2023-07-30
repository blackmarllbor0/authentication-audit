package main

import (
	"auth_audit/config"
	"auth_audit/internal/app/repository/postgres"
	"github.com/spf13/viper"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	configService := config.NewConfig(viper.New())
	if err := configService.LoadConfig("config", "yaml", "config"); err != nil {
		log.Fatal(err)
	}

	repository := postgres.NewRepository(configService)

	_, err := repository.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := repository.Disconnect(); err != nil {
			log.Fatal(err)
		}
	}()
}
