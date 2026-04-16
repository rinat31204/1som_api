package repository

import (
	"1som/internal/domain/transaction"
	"context"
	"database/sql"
)

type transactionRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func NewTransactionRepository(conn interface{}) transaction.Repository {
	switch v := conn.(type) {
	case *sql.DB:
		return &transactionRepository{db: v}
	case *sql.Tx:
		return &transactionRepository{tx: v}
	default:
		panic("unsupported connection type")
	}
}

func (r *transactionRepository) execute() transactionExec {
	if r.tx != nil {
		return r.tx
	}
	return r.db
}

type transactionExec interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

func (r *transactionRepository) Create(u *transaction.Transaction) error {
	_, err := r.execute().ExecContext(context.Background(),
		`INSERT INTO wallets (id,
                     account_number,
                     amount,
                     currency,
                     description,
                     type,
                     created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		u.ID,
		u.AccountNumber,
		u.Amount,
		u.Currency,
		u.Description,
		u.Type,
		u.CreatedAt,
	)
	return err
}
