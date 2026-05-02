package repository

import (
	"context"
	"database/sql"
	"errors"

	"1som/internal/domain/user"
)

type UserRepository struct {
	conn DbTx
}

func NewUserRepository(conn DbTx) user.Repository {
	return &UserRepository{conn: conn}
}

func (r *UserRepository) Create(ctx context.Context, u *user.User) error {
	_, err := r.conn.ExecContext(ctx,
		`INSERT INTO users (id, email, created_at) VALUES ($1, $2, $3)`,
		u.Id, u.Email, u.CreatedAt)
	return err
}

func (r *UserRepository) ExistsUserByIdOrEmail(ctx context.Context, id, email string) (bool, error) {
	var exists bool
	err := r.conn.QueryRowContext(ctx,
		`SELECT EXISTS(SELECT 1 FROM users WHERE id = $1 OR email = $2)`, id, email).Scan(&exists)
	return exists, err
}

func (r *UserRepository) FindById(ctx context.Context, id string) (*user.User, error) {
	var usr user.User
	err := r.conn.QueryRowContext(ctx,
		`SELECT id, email, created_at FROM users WHERE id = $1`, id).
		Scan(&usr.Id, &usr.Email, &usr.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // или ошибка domain.ErrUserNotFound
		}
		return nil, err
	}
	return &usr, nil
}
