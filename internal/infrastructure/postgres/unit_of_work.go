package postgres

import (
	"1som/internal/domain/unitOfWork"
	"1som/internal/infrastructure/postgres/repository"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"1som/internal/domain/account"
	"1som/internal/domain/transaction"
	"1som/internal/domain/user"
	"1som/internal/domain/wallet"
)

type UnitOfWork struct {
	db *sql.DB
	tx *sql.Tx

	userRepo        user.Repository
	walletRepo      wallet.Repository
	accountRepo     account.Repository
	transactionRepo transaction.Repository
}

func NewUnitOfWork(db *sql.DB) unitOfWork.UnitOfWork {
	return &UnitOfWork{db: db}
}

func NewUnitOfWorkFactory(db *sql.DB) func() unitOfWork.UnitOfWork {
	return func() unitOfWork.UnitOfWork {
		return NewUnitOfWork(db)
	}
}

func (u *UnitOfWork) Begin(ctx context.Context) error {
	if u.tx != nil {
		return fmt.Errorf("transaction already started")
	}
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	u.tx = tx

	u.userRepo = repository.NewUserRepository(u.tx)
	u.walletRepo = repository.NewWalletRepository(u.tx)
	u.accountRepo = repository.NewAccountRepository(u.tx)
	u.transactionRepo = repository.NewTransactionRepository(u.tx)

	slog.Debug("transaction started")
	return nil
}

func (u *UnitOfWork) Commit(ctx context.Context) error {
	if u.tx == nil {
		return fmt.Errorf("no active transaction")
	}
	err := u.tx.Commit()
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}
	u.tx = nil
	slog.Debug("transaction committed")
	return nil
}

func (u *UnitOfWork) Rollback(ctx context.Context) error {
	if u.tx == nil {
		return fmt.Errorf("no active transaction")
	}
	err := u.tx.Rollback()
	if err != nil && !errors.Is(err, sql.ErrTxDone) {
		return fmt.Errorf("rollback tx: %w", err)
	}
	u.tx = nil
	slog.Debug("transaction rolled back")
	return nil
}

func (u *UnitOfWork) Complete() error {
	if u.tx != nil {
		slog.Warn("unit of work completed with an active transaction, rolling back")
		return u.Rollback(context.Background())
	}
	return nil
}

func (u *UnitOfWork) UserRepo() user.Repository {
	return u.userRepo
}

func (u *UnitOfWork) WalletRepo() wallet.Repository {
	return u.walletRepo
}

func (u *UnitOfWork) AccountRepo() account.Repository {
	return u.accountRepo
}

func (u *UnitOfWork) TransactionRepo() transaction.Repository {
	return u.transactionRepo
}
