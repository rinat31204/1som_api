package user

import "time"

type User struct {
	id        string
	Email     string
	CreatedAt time.Time
}
