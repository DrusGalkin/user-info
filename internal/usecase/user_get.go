package usecase

import "github.com/DrusGalkin/user-info/internal/domain/models"

func (u *UserUsecase) UserByID(uid int) (*models.User, error) {
	return u.repo.ByID(uid)
}

func (u *UserUsecase) UserByEmail(email string) (*models.User, error) {
	return u.repo.ByEmail(email)
}

func (u *UserUsecase) UserByUsername(username string) (*models.User, error) {
	return u.repo.ByUsername(username)
}

func (u *UserUsecase) AllUsers() ([]*models.User, error) {
	return u.repo.All()
}
