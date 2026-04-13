package unitOfWork

import (
	"1som/internal/domain/account"
	"1som/internal/domain/transaction"
	"1som/internal/domain/user"
	"1som/internal/domain/wallet"
	"context"
)

type UnitOfWork interface {
	Begin(ctx context.Context) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	Complete() error
	UserRepo() user.Repository
	WalletRepo() wallet.Repository
	AccountRepo() account.Repository
	TransactionRepo() transaction.Repository
}
