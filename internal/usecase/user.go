package usecase

import (
	"github.com/DrusGalkin/user-info/internal/domain/models"
	"github.com/DrusGalkin/user-info/internal/repository"
)

type Usecase interface {
	UserByID(uid int) (*models.User, error)
	UserByEmail(email string) (*models.User, error)
	UserByUsername(username string) (*models.User, error)
	AllUsers() ([]*models.User, error)
	DeleteUserByID(uid int) error
	DeleteUserByUsername(username string) error
	DeleteUserByEmail(email string) error
	UpdatePassword(id int, newPass string) error
}

type UserUsecase struct {
	repo repository.Repository
}

func New(repo repository.Repository) Usecase {
	return &UserUsecase{repo: repo}
}
