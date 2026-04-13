package transaction

type Repository interface {
	Create(transaction *Transaction) error
}
