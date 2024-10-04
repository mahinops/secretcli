package secret

import "time"

type Secret struct {
	Title     string
	Username  string
	Password  string
	Note      string
	Email     string
	Website   string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type Secrets []Secret

// SecretService defines methods for secret operations.
type SecretService interface {
	Add(title, username, password, note, email, website string) error
	ListSecrets() error
	Delete(index int) error
	Validate(index int) error
}
