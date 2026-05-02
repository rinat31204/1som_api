package repository

import (
	"1som/internal/domain/wallet"
	"context"
	"database/sql"
	"errors"
)

type WalletRepository struct {
	conn DbTx
}

func NewWalletRepository(conn DbTx) wallet.Repository {
	return &WalletRepository{conn: conn}
}

func (r *WalletRepository) Create(ctx context.Context, u *wallet.Wallet) error {
	_, err := r.conn.ExecContext(ctx,
		`INSERT INTO wallets (id, user_id, created_at) VALUES ($1, $2, $3)`,
		u.Id, u.UserId, u.CreatedAt)
	return err
}

func (r *WalletRepository) GetByUser(ctx context.Context, uid string) (*wallet.Wallet, error) {
	var w wallet.Wallet
	err := r.conn.QueryRowContext(ctx,
		`SELECT id, user_id, created_at FROM wallets WHERE user_id = $1`, uid).Scan(
		&w.Id, &w.UserId, &w.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // или ошибка domain.ErrUserNotFound
		}
		return nil, err
	}
	return &w, nil
}
