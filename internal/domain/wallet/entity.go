package wallet

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	CreatedAt time.Time
}
