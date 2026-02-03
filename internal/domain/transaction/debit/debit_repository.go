package debit

type Repository interface {
	Create(debit Debit) error
}
