package auth

import (
	"errors"
	"time"

	"github.com/mahinops/secretcli/internal/utils"
)

// Register implements UserService
func (u *User) Register(password string) error {
	if password == "" {
		return errors.New("password cannot be empty")
	}
	u.PasswordHash = utils.HashPassword(password)
	u.LastAuth = time.Now()
	return nil
}

// Authenticate implements UserService
func (u *User) Authenticate(password string) error {
	if utils.HashPassword(password) != u.PasswordHash {
		return errors.New("invalid password")
	}
	u.LastAuth = time.Now()
	return nil
}

// IsSessionActive implements UserService
func (u *User) IsSessionActive() bool {
	return time.Since(u.LastAuth) < 5*time.Minute
}
