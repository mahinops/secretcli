package auth

import (
	"errors"
	"fmt"
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
	u.Expiry = time.Now().Add(5 * time.Minute)
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
	if u.Expiry.IsZero() {
		fmt.Println("Expiry date is set to 'never'")
		return true
	}
	return time.Since(u.LastAuth) < time.Until(u.Expiry)
}

func (user *User) SetExpiry(expiry string) error {
	fmt.Println("Setting expiry date:", expiry)
	var duration time.Duration
	var err error
	// Handle "never" case
	if expiry == "o" {
		user.Expiry = time.Time{} // Set to zero time to indicate "never"
		fmt.Println("Expiry set to never")
		return nil
	}
	// Parse duration with time.ParseDuration
	duration, err = time.ParseDuration(expiry)
	if err != nil {
		return errors.New("invalid duration format")
	}

	user.Expiry = time.Now().Add(duration)
	fmt.Println("Expiry set to:", user.Expiry)
	return nil
}
