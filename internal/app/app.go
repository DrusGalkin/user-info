package app

import (
	"fmt"
	"github.com/DrusGalkin/auth-grpc-service/pkg/lib/grpc_client"
	"github.com/DrusGalkin/auth-protos/gen/go/auth"
	"github.com/DrusGalkin/user-info/internal/config"
	"github.com/DrusGalkin/user-info/internal/services"
	"github.com/DrusGalkin/user-info/internal/storage/mysql"
	"github.com/DrusGalkin/user-info/internal/storage/redis"
	"github.com/DrusGalkin/user-info/internal/transport/http"
	"github.com/DrusGalkin/user-info/internal/transport/http/handlers"
	"go.uber.org/zap"
	"time"
)

type App struct {
	Storage *mysql.Storage
	Logger  *zap.Logger
}

func New(log *zap.Logger) *App {
	return &App{
		Storage: mysql.New(),
		Logger:  log,
	}
}

func (app *App) Run(cfg *config.Config) {

	handler, client := app.mustLoadService(
		cfg.Sever.Timeout,
		fmt.Sprintf(":%d", cfg.GRPC.Port),
	)

	rdb := redis.New(
		cfg.Redis.Address,
		cfg.Redis.Port,
		cfg.Sever.Timeout,
		cfg.Redis.TTL,
	)

	engine := http.SetupRouters(
		handler,
		cfg.Sever.Timeout,
		client,
		rdb,
	)

	if err := engine.Listen(fmt.Sprintf(":%d", cfg.Sever.Port)); err != nil {
		panic(err)
	}
}

func (app *App) mustLoadService(timeout time.Duration, grpcAddress string) (handlers.Handler, auth.AuthClient) {
	const op = "app.MustLoadService"

	client, err := grpc_client.NewClient(grpcAddress)
	if err != nil {
		panic(fmt.Errorf("%s: %w", op, err))

	}

	serviceApp := services.New(app.Storage, app.Logger, timeout, client)
	return serviceApp.Handler(), client
}
