package repository

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (u *UserRepository) Update(id int, newPassword string) error {
	const op = "Repository.Update"
	log := u.log.With(zap.String("op", op))

	query := `UPDATE oldmine.users
			  SET password_hash = ?
			  WHERE id = ?`

	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	stmt, err := u.db.PrepareContext(ctx, query)
	defer stmt.Close()

	if err != nil {
		log.Error(op, zap.Error(err))

		return fmt.Errorf("%s ошибка запроса: %w", op, err)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	result, err := stmt.ExecContext(ctx, password, id)
	if err != nil {
		log.Error(op, zap.Error(err))

		return fmt.Errorf("%s Ошибка изменения пароля: %w", op, err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		log.Error(op, zap.Error(err))

		return fmt.Errorf("%s Ошибка изменения пароля: %w", op, err)
	}

	if affected == 0 {
		log.Warn(op, zap.String("rows", fmt.Sprintf("%d", affected)))

		return fmt.Errorf("%s Ошибка изменения пароля: %d", op, affected)
	}

	return nil
}
