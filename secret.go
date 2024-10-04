package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

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

func (secret *Secrets) add(title, username, password, note, email, website string) error {
	if title == "" {
		return errors.New("title is required")
	}

	encryptPassword, err := encrypt(password)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(encryptPassword)
	// Create a new Secret instance
	newSecret := Secret{
		Title:     title,
		Username:  username,
		Password:  encryptPassword,
		Note:      note,
		Email:     email,
		Website:   website,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}
	// Append the new secret to the slice
	*secret = append(*secret, newSecret)
	return nil
}

func (secrets *Secrets) list() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Username", "Password", "Note", "Email", "Website", "Created At", "Updated At")

	for index, secret := range *secrets {
		updatedAt := ""
		if secret.UpdatedAt != nil {
			updatedAt = secret.UpdatedAt.Format(time.RFC3339)
		}
		decryptPassword := ""
		if len(secret.Password) != 0 {
			var err error
			decryptPassword, err = decrypt(secret.Password)
			if err != nil {
				fmt.Println(err)
			}
		}
		table.AddRow(strconv.Itoa(index), secret.Title, secret.Username, decryptPassword, secret.Note, secret.Email, secret.Website, secret.CreatedAt.Format(time.RFC3339), updatedAt)
	}
	table.Render()
}
