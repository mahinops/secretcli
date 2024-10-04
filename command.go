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

func (cf *CmdFlags) Execute(secrets *Secrets) {
	switch {
	case cf.Add:
		cf.addSecret(secrets)
	case cf.List:
		secrets.list()
	case cf.Delete != -1:
		secrets.delete(cf.Delete)
	// case cf.Edit != -1:
	// 	cf.editSecret(secrets, cf.Edit)
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
