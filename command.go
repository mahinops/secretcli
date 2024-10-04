package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

type CmdFlags struct {
	Add    bool
	List   bool
	Delete int
	Edit   int
}

// NewCommandFlags initializes command flags
func NewCommandFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.BoolVar(&cf.Add, "add", false, "Add a new secret")
	flag.BoolVar(&cf.List, "list", false, "List all secrets")
	flag.IntVar(&cf.Delete, "del", -1, "Delete a secret by index")
	flag.IntVar(&cf.Edit, "edit", -1, "Edit a secret by index")
	flag.Parse()
	return &cf
}

// Execute processes the command flags and prompts for secret details
func (cf *CmdFlags) Execute(secrets *Secrets) {
	switch {
	case cf.Add:
		cf.addSecret(secrets)
	case cf.List:
		secrets.list()
	case cf.Delete != -1:
		secrets.delete(cf.Delete)
	case cf.Edit != -1:
		cf.editSecret(secrets, cf.Edit)
	default:
		fmt.Println("Invalid Command. Use --help to see available commands.")
	}
}

// Add a new secret
func (cf *CmdFlags) addSecret(secrets *Secrets) {
	var title, username, password, note, email, website string
	fmt.Println("Adding a new secret...")

	scanner := bufio.NewScanner(os.Stdin)

	// Input Title
	fmt.Print("Enter Title: ")
	scanner.Scan()
	title = scanner.Text()

	// Input Username
	fmt.Print("Enter Username: ")
	scanner.Scan()
	username = scanner.Text()

	// Input Password
	fmt.Print("Enter Password: ")
	scanner.Scan()
	password = scanner.Text()

	// Input Note (optional)
	fmt.Print("Enter Note (optional): ")
	scanner.Scan()
	note = scanner.Text()

	// Input Email (optional)
	fmt.Print("Enter Email (optional): ")
	scanner.Scan()
	email = scanner.Text()

	// Input Website (optional)
	fmt.Print("Enter Website (optional): ")
	scanner.Scan()
	website = scanner.Text()

	// Check for errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Call the add method to add the new secret
	err := secrets.add(title, username, password, note, email, website)
	if err != nil {
		fmt.Println("Error adding secret:", err)
		return
	}
	fmt.Println("Secret added successfully!")
}

// Edit an existing secret
func (cf *CmdFlags) editSecret(secrets *Secrets, index int) {
	// Fetch the existing secret
	secret := (*secrets)[index]

	fmt.Println("Editing secret:", secret.Title)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Select field to edit:")
		fmt.Println("1. Title")
		fmt.Println("2. Username")
		fmt.Println("3. Password")
		fmt.Println("4. Note")
		fmt.Println("5. Email")
		fmt.Println("6. Website")
		fmt.Println("7. Save and Exit")
		fmt.Print("Enter your choice (1-7): ")

		if scanner.Scan() {
			choice := scanner.Text()
			switch choice {
			case "1":
				fmt.Print("Enter new Title: ")
				scanner.Scan()
				secret.Title = scanner.Text()
			case "2":
				fmt.Print("Enter new Username: ")
				scanner.Scan()
				secret.Username = scanner.Text()
			case "3":
				fmt.Print("Enter new Password: ")
				scanner.Scan()
				secret.Password, _ = encrypt(scanner.Text())
			case "4":
				fmt.Print("Enter new Note: ")
				scanner.Scan()
				secret.Note = scanner.Text()
			case "5":
				fmt.Print("Enter new Email: ")
				scanner.Scan()
				secret.Email = scanner.Text()
			case "6":
				fmt.Print("Enter new Website: ")
				scanner.Scan()
				secret.Website = scanner.Text()
			case "7":
				secret.UpdatedAt = new(time.Time)
				*secret.UpdatedAt = time.Now()
				(*secrets)[index] = secret // Update the secret in the slice
				fmt.Println("Secret updated successfully!")
				return
			default:
				fmt.Println("Invalid choice. Please try again.")
			}
		}

		// Check for errors
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
	}
}
