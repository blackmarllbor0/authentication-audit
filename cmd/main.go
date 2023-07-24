package main

import (
	"auth_audit/config"
	"auth_audit/internal/app/server"
	"auth_audit/internal/app/server/handlers"
	userHandlers "auth_audit/internal/app/server/handlers/users"
	userServices "auth_audit/internal/app/server/services/users"
	"auth_audit/internal/pkg/repository/postgres"
	userRepositories "auth_audit/internal/pkg/repository/postgres/models/users"
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

	DB, err := repository.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := repository.Disconnect(); err != nil {
			log.Fatal(err)
		}
	}()

	userRepository := userRepositories.NewUser(DB)
	userService := userServices.NewService(userRepository)
	userHandler := userHandlers.NewUsers(userService)
	hand := handlers.NewHandler(userHandler)

	s := server.NewServer(configService)
	if err := s.Run(hand.Router()); err != nil {
		log.Fatal(err)
	}
}
