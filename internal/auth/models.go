package auth

import "time"

// User struct represents the user model.
type User struct {
	PasswordHash string    `json:"password_hash"`
	LastAuth     time.Time `json:"last_auth"`
	Expiry       time.Time `json:"expiry"`
}

// UserService defines methods for user operations
type UserService interface {
	Register(password string) error
	Authenticate(password string) error
	IsSessionActive() bool
	SetExpiry(expiry string) error
}
