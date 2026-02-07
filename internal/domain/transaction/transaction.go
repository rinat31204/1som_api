package transaction

import (
	"1som/internal/domain/currency"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID            uuid.UUID
	AccountNumber string
	Amount        int64
	Currency      currency.Currency
	Description   string
	Type          Type
	CreatedAt     time.Time
}
