package main

import (
	"fmt"

	"github.com/mahinops/secretcli/internal/auth"
	"github.com/mahinops/secretcli/internal/command"
	"github.com/mahinops/secretcli/internal/helper"
	"github.com/mahinops/secretcli/internal/secret"
	"github.com/mahinops/secretcli/internal/storage"
)

func main() {
	user := auth.User{}
	userStorage := storage.NewStorage[auth.User](".secrets/user.json")
	err := userStorage.Load(&user)
	if err != nil {
		fmt.Println("No user registered. Please register.")
		err := helper.RegisterUser(&user, userStorage)
		if err != nil {
			fmt.Println("Registration failed:", err)
			return
		}
	} else {
		err := helper.AuthenticateUser(&user, userStorage)
		if err != nil {
			fmt.Println("Authentication failed:", err)
			return
		}
	}

	secrets := secret.Secrets{}
	secretStorage := storage.NewStorage[secret.Secrets](".secrets/secrets.json")
	secretStorage.Load(&secrets)
	cmdFlags := command.NewCommandFlags()
	cmdFlags.Execute(&secrets)
	secretStorage.Save(secrets)
}
