package account

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, account *Account) error
	GetByWallet(ctx context.Context, walletId uuid.UUID) (*Account, error)
}
