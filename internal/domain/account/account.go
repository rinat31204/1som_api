package account

import (
	"1som/internal/domain/currency"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id            uuid.UUID
	WalletId      uuid.UUID
	AccountNumber string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Balance       int64
	Currency      currency.Currency
}
