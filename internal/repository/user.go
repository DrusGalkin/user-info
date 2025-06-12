package repository

import (
	"errors"
	"github.com/DrusGalkin/user-info/internal/domain/models"
	"github.com/DrusGalkin/user-info/internal/storage/mysql"
	"go.uber.org/zap"
	"time"
)

type Repository interface {
	ByID(uid int) (*models.User, error)
	ByEmail(email string) (*models.User, error)
	ByUsername(username string) (*models.User, error)
	All() ([]*models.User, error)
	DeleteByID(uid int) error
	DeleteByUsername(username string) error
	DeleteByEmail(email string) error
	Update(id int, newPassword []byte) error
}

var (
	ErrUserNotFound = errors.New("Пользователь не найден")
)

type UserRepository struct {
	db      *mysql.Storage
	log     *zap.Logger
	timeout time.Duration
}

func New(db *mysql.Storage, log *zap.Logger, timeout time.Duration) Repository {
	return &UserRepository{
		db:      db,
		log:     log,
		timeout: timeout,
	}
}
