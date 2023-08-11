package main

import (
	"auth_audit/config"
	"auth_audit/internal/app/repository/postgres"
	"auth_audit/internal/app/server"
	"auth_audit/internal/app/server/handlers"
	"auth_audit/internal/app/server/services"
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

	userRepo := postgres.NewUserRepository(DB)
	sessionRepo := postgres.NewSessionRepository(DB)
	authAuditRepo := postgres.NewAuthenticationAudit(DB)

	userService := services.NewUserService(userRepo)
	sessionService := services.NewSessionService(sessionRepo)
	authAuditService := services.NewAuthAuditService(authAuditRepo)
	authService := services.NewAuthService(userService, sessionService, authAuditService)

	handler := handlers.NewHandler(authService)

	srv := server.NewServer(configService)
	if err := srv.Run(handler.Router()); err != nil {
		log.Fatal(err)
	}
}
