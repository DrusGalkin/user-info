package repository

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

func (u *UserRepository) DeleteByID(uid int) error {
	const op = "Repository.DeleteByID"
	log := u.log.With(zap.String("op", op), zap.Int("uid", uid))

	query := `DELETE FROM users WHERE id = ?`

	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		log.Error(op, zap.Error(err))

		return fmt.Errorf("%s ошибка запроса: %w", op, err)
	}

	result, err := stmt.ExecContext(ctx, uid)
	if err != nil {
		log.Error("Ошибка удаления пользователя", zap.Error(err))

		return fmt.Errorf("%s: %w", op, err)
	}

	updateRows, err := result.RowsAffected()
	if err != nil {
		log.Error("Пользователь не удален", zap.Int("id", uid))

		return fmt.Errorf("%s, измененные строки %d : %w", op, updateRows, err)
	}

	if updateRows == 0 {
		log.Warn("Пользователь не удален, у вас нет прав или его не существует", zap.Int("id", uid))

		return ErrUserNotFound
	}

	return nil
}

func (u *UserRepository) DeleteByUsername(username string) error {
	const op = "Repository.DeleteByUsername"
	log := u.log.With(zap.String("op", op), zap.String("username", username))

	query := `DELETE FROM users WHERE username = ?`

	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		log.Error(op, zap.Error(err))

		return fmt.Errorf("%s ошибка запроса: %w", op, err)
	}

	result, err := stmt.ExecContext(ctx, username)
	if err != nil {
		log.Error("Ошибка удаления пользователя", zap.Error(err))

		return fmt.Errorf("%s: %w", op, err)
	}

	updateRows, err := result.RowsAffected()
	if err != nil || updateRows == 0 {
		log.Error("Пользователь не удален", zap.Error(err))

		return fmt.Errorf("%s, измененные строки %d : %w", op, updateRows, err)
	}

	return nil
}

func (u *UserRepository) DeleteByEmail(email string) error {
	const op = "Repository.DeleteByEmail"
	log := u.log.With(zap.String("op", op), zap.String("email", email))

	query := `DELETE FROM users WHERE email = $1`
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		log.Error(op, zap.Error(err))

		return fmt.Errorf("%s ошибка запроса: %w", op, err)
	}

	result, err := stmt.ExecContext(ctx, email)
	if err != nil {
		log.Error("Ошибка удаления пользователя", zap.Error(err))

		return fmt.Errorf("%s: %w", op, err)
	}

	updateRows, err := result.RowsAffected()
	if err != nil || updateRows == 0 {
		log.Error("Пользователь не удален", zap.Error(err))

		return fmt.Errorf("%s, измененные строки %d : %w", op, updateRows, err)
	}

	return nil
}
