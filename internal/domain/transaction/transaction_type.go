package transaction

type Type string

const (
	Debit  Type = "debit"
	Credit Type = "credit"
	Refund Type = "refund"
)
