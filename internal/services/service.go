package services

import (
	"github.com/DrusGalkin/auth-protos/gen/go/auth"
	"github.com/DrusGalkin/user-info/internal/repository"
	"github.com/DrusGalkin/user-info/internal/storage/mysql"
	"github.com/DrusGalkin/user-info/internal/transport/http/handlers"
	"github.com/DrusGalkin/user-info/internal/usecase"
	"go.uber.org/zap"
	"time"
)

type App struct {
	log     *zap.Logger
	db      *mysql.Storage
	repo    repository.Repository
	usecase usecase.Usecase
	handler handlers.Handler
	Timeout time.Duration
	Client  auth.AuthClient
}

func New(db *mysql.Storage, log *zap.Logger, timeout time.Duration, client auth.AuthClient) *App {
	app := &App{
		log:     log,
		db:      db,
		repo:    repository.New(db, log, timeout),
		Timeout: timeout,
		Client:  client,
	}

	app.usecase = usecase.New(app.repo)
	app.handler = handlers.New(app.usecase, timeout, client)

	return app
}

func (app *App) Handler() handlers.Handler {
	return app.handler
}
