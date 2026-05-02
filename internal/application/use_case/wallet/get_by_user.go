package wallet

import (
	"1som/internal/domain/account"
	"1som/internal/domain/wallet"
	"context"
)

type GetWalletUseCase interface {
	Execute(ctx context.Context, uid string) (*Dto, error)
}

type GetWalletUseCaseImpl struct {
	walletRepo  wallet.Repository
	accountRepo account.Repository
}

func NewWalletUseCaseImpl(
	walletRepo wallet.Repository,
	accountRepo account.Repository) *GetWalletUseCaseImpl {
	return &GetWalletUseCaseImpl{
		walletRepo:  walletRepo,
		accountRepo: accountRepo,
	}
}

func (u *GetWalletUseCaseImpl) Execute(ctx context.Context, uid string) (*Dto, error) {
	walletDb, err := u.walletRepo.GetByUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	accountDb, err := u.accountRepo.GetByWallet(ctx, walletDb.Id)
	if err != nil {
		return nil, err
	}
	return &Dto{
		uid,
		accountDb.AccountNumber,
		accountDb.Balance,
	}, nil
}
