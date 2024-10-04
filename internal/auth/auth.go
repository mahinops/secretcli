package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"
)

type User struct {
	PasswordHash string    `json:"password_hash"`
	LastAuth     time.Time `json:"last_auth"`
}

func (u *User) Register(password string) error {
	if password == "" {
		return errors.New("password cannot be empty")
	}
	u.PasswordHash = hashPassword(password)
	u.LastAuth = time.Now()
	return nil
}

func (u *User) Authenticate(password string) error {
	if hashPassword(password) != u.PasswordHash {
		return errors.New("invalid password")
	}
	u.LastAuth = time.Now()
	return nil
}

func (u *User) IsSessionActive() bool {
	return time.Since(u.LastAuth) < 5*time.Minute
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
