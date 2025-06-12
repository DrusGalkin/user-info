package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DrusGalkin/user-info/internal/domain/models"
	"go.uber.org/zap"
)

func (u *UserRepository) ByID(uid int) (*models.User, error) {
	const op = "Repository.ByID"
	log := u.log.With(zap.String("op", op), zap.Int("uid", uid))

	query := `SELECT id, email, username FROM users WHERE id = ?`

	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	stmt, err := u.db.PrepareContext(ctx, query)

	if err != nil {
		log.Warn(op, zap.Error(err), zap.String("query", query))

		return nil, fmt.Errorf("%s: ошибка запроса : %w", op, err)
	}

	defer stmt.Close()

	var user models.User
	if err := stmt.QueryRowContext(ctx, uid).Scan(&user.ID, &user.Username, &user.Email); err != nil {
		log.Error(op, zap.Error(err))

		return nil, ErrUserNotFound
	}

	return &user, nil
}

func (u *UserRepository) ByEmail(email string) (*models.User, error) {
	const op = "Repository.ByEmail"
	log := u.log.With(zap.String("op", op), zap.String("email", email))

	query := `SELECT id, email, username FROM users WHERE email = ?`

	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	stmt, err := u.db.PrepareContext(ctx, query)

	if err != nil {
		log.Warn(op, zap.Error(err), zap.String("query", query))

		return nil, fmt.Errorf("%s: ошибка запроса : %w", op, err)
	}

	defer stmt.Close()

	var user models.User
	if err := stmt.QueryRowContext(ctx, email).Scan(&user.ID, &user.Email, &user.Username); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Debug("Нет пользователя с данным email", zap.String("email", email))
			return nil, ErrUserNotFound
		}

		log.Error("Ошибка выполнения запроса", zap.Error(err))
		return nil, fmt.Errorf("%s: ошибка выполнения запроса: %w", op, err)
	}

	return &user, nil
}

func (u *UserRepository) ByUsername(username string) (*models.User, error) {
	const op = "Repository.ByUsername"
	log := u.log.With(zap.String("op", op), zap.String("email", username))

	query := `SELECT id, email, username FROM users WHERE username = ?`

	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	stmt, err := u.db.PrepareContext(ctx, query)

	if err != nil {
		log.Warn(op, zap.Error(err), zap.String("query", query))

		return nil, fmt.Errorf("%s: ошибка запроса : %w", op, err)
	}

	defer stmt.Close()

	var user models.User
	if err := stmt.QueryRowContext(ctx, username).Scan(&user.ID, &user.Username, &user.Email); err != nil {
		log.Error(op, zap.Error(err))

		return nil, ErrUserNotFound
	}

	return &user, nil
}

func (u *UserRepository) All() ([]*models.User, error) {
	const op = "Repository.All"
	log := u.log.With(zap.String("op", op))

	query := `SELECT id, email, username FROM users`

	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	stmt, err := u.db.PrepareContext(ctx, query)

	if err != nil {
		log.Warn(op, zap.Error(err), zap.String("query", query))

		return nil, fmt.Errorf("%s: ошибка запроса : %w", op, err)
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Warn(op, zap.Error(err), zap.String("query", query))

		return nil, fmt.Errorf("%s: ошибка запроса : %w", op, err)
	}

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Username); err != nil {
			log.Warn(op, zap.Error(err))
			continue
		}
		users = append(users, &user)
	}

	return users, nil
}
