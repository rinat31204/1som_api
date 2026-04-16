package repository

import (
	"context"
	"database/sql"
	"errors"

	"1som/internal/domain/user"
)

type userRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func NewUserRepository(conn interface{}) user.Repository {
	switch v := conn.(type) {
	case *sql.DB:
		return &userRepository{db: v}
	case *sql.Tx:
		return &userRepository{tx: v}
	default:
		panic("unsupported connection type")
	}
}

func (r *userRepository) getExecer() exec {
	if r.tx != nil {
		return r.tx
	}
	return r.db
}

type exec interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

func (r *userRepository) Create(u *user.User) error {
	_, err := r.getExecer().ExecContext(context.Background(),
		`INSERT INTO users (id, email, created_at) VALUES ($1, $2, $3)`,
		u.Id, u.Email, u.CreatedAt)
	return err
}

func (r *userRepository) ExistsUserByIdOrEmail(id, email string) (bool, error) {
	var exists bool
	err := r.getExecer().QueryRowContext(context.Background(),
		`SELECT EXISTS(SELECT 1 FROM users WHERE id = $1 OR email = $2)`, id, email).Scan(&exists)
	return exists, err
}

func (r *userRepository) FindById(id string) (*user.User, error) {
	var usr user.User
	err := r.getExecer().QueryRowContext(context.Background(),
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
