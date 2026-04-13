package repository

import (
	"1som/internal/domain/wallet"
	"context"
	"database/sql"
)

type walletRepository struct {
	db *sql.DB // for non‑transactional queries
	tx *sql.Tx // optional – used when inside a UoW
}

// NewUserRepository returns a repository that works with either *sql.DB or *sql.Tx.
func NewWalletRepository(conn interface{}) wallet.Repository {
	switch v := conn.(type) {
	case *sql.DB:
		return &walletRepository{db: v}
	case *sql.Tx:
		return &walletRepository{tx: v}
	default:
		panic("unsupported connection type")
	}
}

func (r *walletRepository) execute() walletExec {
	if r.tx != nil {
		return r.tx
	}
	return r.db
}

type walletExec interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

func (r *walletRepository) Create(u *wallet.Wallet) error {
	_, err := r.execute().ExecContext(context.Background(),
		`INSERT INTO wallets (id, user_id, created_at) VALUES ($1, $2, $3)`,
		u.Id, u.UserId, u.CreatedAt)
	return err
}
