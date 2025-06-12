package main

import (
	log "github.com/DrusGalkin/auth-grpc-service/pkg/lib/logger"
	"github.com/DrusGalkin/user-info/internal/app"
	"github.com/DrusGalkin/user-info/internal/config"
)

func main() {
	// Конфиг
	cfg := config.MustLoadConfig()

	// Логгер
	logger := log.SetupLogger(cfg.Env)
	defer logger.Sync()

	logger.Info("Логгер и конфиг найден")

	// Сервер
	app.New(logger).Run(cfg)
}
