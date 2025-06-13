package usecase

func (u *UserUsecase) UpdatePassword(id int, newPass string) error {
	return u.repo.Update(id, newPass)
}
