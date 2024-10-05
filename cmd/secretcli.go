package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mahinops/secretcli/internal/auth"
	"github.com/mahinops/secretcli/internal/secret"
	"github.com/mahinops/secretcli/internal/storage"
)

const (
	secretCLIPath = "~/secretcli/secrets.json"
	userPath      = "~/secretcli/user.json"
)

func main() {
	user := auth.User{}
	userStorage := storage.NewStorage[auth.User](userPath)
	err := userStorage.Load(&user)
	if err != nil {
		fmt.Println(err)
		fmt.Println("No user registered. Please register.")
		err := auth.RegisterUser(&user, userStorage)
		if err != nil {
			fmt.Println("Registration failed:", err)
			return
		}
	} else {
		err := auth.AuthenticateUser(&user, userStorage)
		if err != nil {
			fmt.Println("Authentication failed:", err)
			return
		}
	}

	// Create separate flag sets for auth and secret
	authFlagSet := flag.NewFlagSet("auth", flag.ExitOnError)
	secretFlagSet := flag.NewFlagSet("secret", flag.ExitOnError)

	// Register auth and secret flags
	userCmdFlag := auth.NewCommandFlags()
	userCmdFlag.RegisterFlags(authFlagSet)

	cmdFlags := secret.NewCommandFlags()
	cmdFlags.RegisterFlags(secretFlagSet)

	// Check if command is provided (auth or secret)
	if len(os.Args) < 2 {
		fmt.Println("Expected 'auth' or 'secret' command.")
		return
	}

	command := os.Args[1]

	// Handle auth command
	if command == "auth" {
		authFlagSet.Parse(os.Args[2:]) // Parse flags for 'auth'
		userCmdFlag.Execute(&user, authFlagSet)
		userStorage.Save(user)
	} else if command == "secret" { // Handle secret command
		secretFlagSet.Parse(os.Args[2:]) // Parse flags for 'secret'
		secrets := secret.Secrets{}
		secretStorage := storage.NewStorage[secret.Secrets](secretCLIPath)
		secretStorage.Load(&secrets)
		cmdFlags.Execute(&secrets, secretFlagSet)
		secretStorage.Save(secrets)
	} else {
		fmt.Println("Invalid command. Use 'auth' or 'secret'.")
	}
}
