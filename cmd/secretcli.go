package main

import (
	"fmt"

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

	secrets := secret.Secrets{}
	secretStorage := storage.NewStorage[secret.Secrets](secretCLIPath)
	secretStorage.Load(&secrets)
	cmdFlags := secret.NewCommandFlags()
	cmdFlags.Execute(&secrets)
	secretStorage.Save(secrets)
}
