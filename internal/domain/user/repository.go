package user

import "context"

type Repository interface {
	Create(ctx context.Context, user *User) error
	FindById(ctx context.Context, id string) (*User, error)
	ExistsUserByIdOrEmail(ctx context.Context, id, email string) (bool, error)
}
