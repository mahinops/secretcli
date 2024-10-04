package auth

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mahinops/secretcli/internal/storage"
)

// RegisterUser prompts the user for a password and registers the user
func RegisterUser(user UserService, userStorage *storage.Storage[User]) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a password to register: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read password: %w", err)
	}
	password = strings.TrimSpace(password)

	if err := user.Register(password); err != nil {
		return fmt.Errorf("registration failed: %w", err)
	}

	// Save user data
	if err := userStorage.Save(*user.(*User)); err != nil {
		return fmt.Errorf("failed to save user data: %w", err)
	}

	fmt.Println("Registration successful! You can now manage your secrets.")
	return nil
}

// AuthenticateUser prompts the user for their password and authenticates the user
func AuthenticateUser(user UserService, userStorage *storage.Storage[User]) error {
	// If session is still active, no need to authenticate
	if user.IsSessionActive() {
		fmt.Println("Session is still active. No need to authenticate.")
		return nil
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read password: %w", err)
	}
	password = strings.TrimSpace(password)

	if err := user.Authenticate(password); err != nil {
		return fmt.Errorf("authentication failed: %w", err)
	}

	// Save updated last authentication time
	if err := userStorage.Save(*user.(*User)); err != nil {
		return fmt.Errorf("failed to save user data: %w", err)
	}

	fmt.Println("Authenticated successfully! You can now manage your secrets.")
	return nil
}
