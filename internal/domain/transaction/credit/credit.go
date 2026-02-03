package credit

import (
	"1som/internal/domain/currency"
	"time"

	"github.com/google/uuid"
)

type Credit struct {
	Id            uuid.UUID
	AccountNumber string
	Amount        int64
	Currency      currency.Currency
	Description   string
	CreatedAt     time.Time
}
