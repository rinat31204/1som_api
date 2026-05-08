package repository

import (
	"1som/internal/domain/account"
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type AccountRepository struct {
	conn DbTx
}

func NewAccountRepository(conn DbTx) account.Repository {
	return &AccountRepository{conn: conn}
}

func (r *AccountRepository) Create(ctx context.Context, u *account.Account) error {
	_, err := r.conn.ExecContext(ctx,
		`INSERT INTO accounts (id,
                     				wallet_id,
                     				account_number,
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

func (r *AccountRepository) GetByWallet(ctx context.Context, walletId uuid.UUID) (*account.Account, error) {
	var acc account.Account
	err := r.conn.QueryRowContext(ctx,
		`SELECT id,
					  wallet_id,
					  account_number,
					  created_at,
					  updated_at,
					  balance,
					  currency FROM accounts WHERE wallet_id = $1`, walletId).Scan(
		&acc.Id,
		&acc.WalletId,
		&acc.AccountNumber,
		&acc.CreatedAt,
		&acc.UpdatedAt,
		&acc.Balance,
		&acc.Currency)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // или ошибка domain.ErrUserNotFound
		}
		return nil, err
	}
	return &acc, nil
}
