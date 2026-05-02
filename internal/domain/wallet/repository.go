package wallet

import "context"

type Repository interface {
	Create(ctx context.Context, wallet *Wallet) error
	GetByUser(ctx context.Context, uid string) (*Wallet, error)
}
