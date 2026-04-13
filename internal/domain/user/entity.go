package user

import "time"

type User struct {
	Id        string
	Email     string
	CreatedAt time.Time
}
