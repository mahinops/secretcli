package main

import "fmt"

func main() {
	user := User{}
	userStorage := NewStorage[User](".secrets/user.json")
	err := userStorage.Load(&user)
	if err != nil {
		fmt.Println("No user registered. Please register.")
		err := registerUser(&user, userStorage)
		if err != nil {
			fmt.Println("Registration failed:", err)
			return
		}
	} else {
		err := authenticateUser(&user, userStorage)
		if err != nil {
			fmt.Println("Authentication failed:", err)
			return
		}
	}

	secrets := Secrets{}
	secretStorage := NewStorage[Secrets](".secrets/secrets.json")
	secretStorage.Load(&secrets)
	cmdFlags := NewCommandFlags()
	cmdFlags.Execute(&secrets)
	secretStorage.Save(secrets)
}
