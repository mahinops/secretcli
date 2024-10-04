// command.go
package main

import (
	"flag"
	"fmt"
)

type CmdFlags struct {
	Add    string
	List   bool
	Delete int
}

// NewCommandFlags initializes command flags
func NewCommandFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "title", "", "Add a new secret")
	flag.BoolVar(&cf.List, "list", false, "List all secrets")
	flag.IntVar(&cf.Delete, "del", -1, "Delete a secrets by index")
	flag.Parse()
	return &cf
}

// Execute processes the command flags and prompts for secret details
func (cf *CmdFlags) Execute(secrets *Secrets) {
	switch {
	case cf.Add != "":
		var username, password, note, email, website string
		fmt.Println("Adding a new secret with title:", cf.Add)

		fmt.Print("Enter Username: ")
		fmt.Scanln(&username)

		fmt.Print("Enter Password: ")
		fmt.Scanln(&password)

		fmt.Print("Enter Note (optional): ")
		fmt.Scanln(&note)

		fmt.Print("Enter Email (optional): ")
		fmt.Scanln(&email)

		fmt.Print("Enter Website (optional): ")
		fmt.Scanln(&website)

		// Call the add method to add the new secret
		err := secrets.add(cf.Add, username, password, note, email, website)
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
