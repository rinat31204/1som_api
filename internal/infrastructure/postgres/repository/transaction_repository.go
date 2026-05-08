package repository

import (
	"1som/internal/domain/transaction"
	"context"
)

type TransactionRepository struct {
	conn DbTx
}

func NewTransactionRepository(conn DbTx) transaction.Repository {
	return &TransactionRepository{conn: conn}
}

func (r *TransactionRepository) Create(ctx context.Context, u *transaction.Transaction) error {
	_, err := r.conn.ExecContext(ctx,
		`INSERT INTO transactions (id,
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
