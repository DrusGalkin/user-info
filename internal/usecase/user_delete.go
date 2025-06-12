package usecase

func (u *UserUsecase) DeleteUserByID(uid int) error {
	return u.repo.DeleteByID(uid)
}

func (u *UserUsecase) DeleteUserByUsername(username string) error {
	return u.repo.DeleteByUsername(username)
}

func (u *UserUsecase) DeleteUserByEmail(email string) error {
	return u.repo.DeleteByEmail(email)
}
