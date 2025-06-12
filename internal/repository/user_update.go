package repository

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

func (u *UserRepository) Update(id int, newPassword []byte) error {
	const op = "Repository.Update"
	log := u.log.With(zap.String("op", op))

	query := `UPDATE users
			  SET password = ?
			  WHERE id = ?`

	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	stmt, err := u.db.PrepareContext(ctx, query)

	if err != nil {
		log.Error(op, zap.Error(err))

		return fmt.Errorf("%s ошибка запроса: %w", op, err)
	}

	defer stmt.Close()

	if err := stmt.QueryRowContext(ctx, newPassword, id).Scan(&id); err != nil {
		log.Error(op, zap.Error(err))

		return fmt.Errorf("%s Ошибка изменения паролья: %w", op, err)
	}

	return nil
}
