package repository

import (
	"1som/internal/domain/account"
	"context"
	"database/sql"
)

type accountRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func NewAccountRepository(conn interface{}) account.Repository {
	switch v := conn.(type) {
	case *sql.DB:
		return &accountRepository{db: v}
	case *sql.Tx:
		return &accountRepository{tx: v}
	default:
		panic("unsupported connection type")
	}
}

func (r *accountRepository) execute() accountExec {
	if r.tx != nil {
		return r.tx
	}
	return r.db
}

type accountExec interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

func (r *accountRepository) Create(u *account.Account) error {
	_, err := r.execute().ExecContext(context.Background(),
		`INSERT INTO wallets (id,
                     				wallet_id,
                     				account_numberl,
                     				created_at,
                     				updated_at,
                     				balance,
                     				currency) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		u.Id,
		u.WalletId,
		u.AccountNumber,
		u.CreatedAt,
		u.UpdatedAt,
		u.Balance,
		u.Currency)
	return err
}
