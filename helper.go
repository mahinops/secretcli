package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func registerUser(user *User, userStorage *Storage[User]) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a password to register: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	err := user.Register(password)
	if err != nil {
		return err
	}

	// Save user data
	err = userStorage.Save(*user)
	if err != nil {
		return err
	}

	fmt.Println("Registration successful! You can now manage your secrets.")
	return nil
}

func authenticateUser(user *User, userStorage *Storage[User]) error {
	// If session is still active, no need to authenticate
	if user.IsSessionActive() {
		return nil
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	err := user.Authenticate(password)
	if err != nil {
		return err
	}

	// Save updated last authentication time
	err = userStorage.Save(*user)
	if err != nil {
		return err
	}

	fmt.Println("Authenticated successfully! You can now manage your secrets.")
	return nil
}
