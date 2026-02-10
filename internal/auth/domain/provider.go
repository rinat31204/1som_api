package domain

import (
	"context"
)

type Provider interface {
	VerifyToken(ctx context.Context, token string) (*User, error)
}
