package wallet

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	Id        uuid.UUID
	UserId    string
	CreatedAt time.Time
}
