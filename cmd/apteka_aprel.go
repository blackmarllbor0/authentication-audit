package main

import (
	"aptekaaprel/config"
	"aptekaaprel/internal/app/server"
	"aptekaaprel/internal/app/server/handlers"
	userHandlers "aptekaaprel/internal/app/server/handlers/users"
	userServices "aptekaaprel/internal/app/server/services/users"
	"aptekaaprel/internal/pkg/repository/postgres"
	userRepositories "aptekaaprel/internal/pkg/repository/postgres/models/users"
	"github.com/joho/godotenv"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	envService := config.NewEnv()
	repository := postgres.NewRepository(envService)

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

	s := server.NewServer(envService)
	if err := s.Run(hand.Router()); err != nil {
		log.Fatal(err)
	}
}
