package service

const EventUserVerified = "user_verified"

type UserVerifiedEvent struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}
