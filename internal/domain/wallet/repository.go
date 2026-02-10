package wallet

type Repository interface {
	Create(wallet Wallet) error
}
