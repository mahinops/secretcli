package main

import (
	"errors"
	"fmt"
	"time"
)

type Secret struct {
	Title     string
	Username  string
	Password  string
	Note      string
	Email     string
	Website   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Secrets []Secret

func (secret *Secrets) add(title, username, password, note, email, website string) error {
	if title == "" {
		return errors.New("title is required")
	}
	// Create a new Secret instance
	newSecret := Secret{
		Title:     title,
		Username:  username,
		Password:  password,
		Note:      note,
		Email:     email,
		Website:   website,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Append the new secret to the slice
	*secret = append(*secret, newSecret)

	fmt.Println(newSecret)
	return nil
}
