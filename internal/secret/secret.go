package secret

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
	"github.com/mahinops/secretcli/internal/utils"
)

// Function to add a new secret
func (secrets *Secrets) Add(title, username, password, note, email, website string) error {
	if title == "" {
		return errors.New("title is required")
	}

	encryptPassword, err := utils.Encrypt(password)
	if err != nil {
		return err
	}

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
	*secrets = append(*secrets, newSecret)
	return nil
}

// Function to list all secrets
func (secrets *Secrets) ListSecrets() error {
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
			decryptPassword, err = utils.Decrypt(secret.Password)
			if err != nil {
				return err
			}
		}
		table.AddRow(strconv.Itoa(index), secret.Title, secret.Username, decryptPassword, secret.Note, secret.Email, secret.Website, secret.CreatedAt.Format(time.RFC3339), updatedAt)
	}
	table.Render()
	return nil
}

// Validate secret index
func (secrets *Secrets) Validate(index int) error {
	if index < 0 || index >= len(*secrets) {
		err := errors.New("invalid index")
		return err
	}
	return nil
}

// Delete a secret
func (secrets *Secrets) Delete(index int) error {
	if err := secrets.Validate(index); err != nil {
		return err
	}
	*secrets = append((*secrets)[:index], (*secrets)[index+1:]...)
	return nil
}

// Edit an existing secret
func (secrets *Secrets) Edit(index int, updatedSecret Secret) error {
	if err := secrets.Validate(index); err != nil {
		return err
	}

	// Update the secret
	(*secrets)[index] = updatedSecret
	return nil
}
