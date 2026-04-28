package wallet

import "1som/internal/domain/account"

type GetWalletUseCase interface {
	Execute(accountNumber string) (WalletDto, error)
}

type GetWalletUseCaseImpl struct {
	repo account.Repository
}

func NewGetWallet(repo account.Repository) *GetWalletUseCaseImpl {
	return &GetWalletUseCaseImpl{repo: repo}
}

func (w *GetWalletUseCaseImpl) Execute(accountNumber string) (WalletDto, error) {

}
