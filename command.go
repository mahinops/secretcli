// command.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type CmdFlags struct {
	Add    bool
	List   bool
	Delete int
}

// NewCommandFlags initializes command flags
func NewCommandFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.BoolVar(&cf.Add, "add", false, "Add a new secret")
	flag.BoolVar(&cf.List, "list", false, "List all secrets")
	flag.IntVar(&cf.Delete, "del", -1, "Delete a secrets by index")
	flag.Parse()
	return &cf
}

// Execute processes the command flags and prompts for secret details
func (cf *CmdFlags) Execute(secrets *Secrets) {
	switch {
	case cf.Add:
		var title, username, password, note, email, website string
		fmt.Println("Adding a new secret with title:", cf.Add)

		// Create a new scanner
		scanner := bufio.NewScanner(os.Stdin)

		// Input Title
		fmt.Print("Enter Title: ")
		if scanner.Scan() {
			title = scanner.Text()
		}

		// Input Username
		fmt.Print("Enter Username: ")
		if scanner.Scan() {
			username = scanner.Text()
		}

		// Input Password
		fmt.Print("Enter Password: ")
		if scanner.Scan() {
			password = scanner.Text()
		}

		// Input Note (optional)
		fmt.Print("Enter Note (optional): ")
		if scanner.Scan() {
			note = scanner.Text()
		}

		// Input Email (optional)
		fmt.Print("Enter Email (optional): ")
		if scanner.Scan() {
			email = scanner.Text()
		}

		// Input Website (optional)
		fmt.Print("Enter Website (optional): ")
		if scanner.Scan() {
			website = scanner.Text()
		}

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

	case cf.List:
		// Call the list method to list all secrets
		secrets.list()

	case cf.Delete != -1:
		secrets.delete(cf.Delete)

	default:
		fmt.Println("Invalid Command. Use --help to see available commands.")
	}
}
