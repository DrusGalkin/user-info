package usecase

import (
	"github.com/DrusGalkin/user-info/internal/lib/bcrypt"
)

func (u *UserUsecase) UpdatePassword(id int, newPass string) error {
	hashPassword, err := bcrypt.HashPassword(newPass)
	if err != nil {
		return err
	}

	return u.repo.Update(id, hashPassword)
}
