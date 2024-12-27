package secret

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/mahinops/secretcli/internal/utils"
)

type CmdFlags struct {
	Add    bool
	List   bool
	Delete int
	Edit   int
	Export bool
}

func NewCommandFlags() *CmdFlags {
	cf := CmdFlags{}
	return &cf
}

func (cf *CmdFlags) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&cf.Add, "add", false, "Add a new secret")
	fs.BoolVar(&cf.List, "list", false, "List all secrets")
	fs.IntVar(&cf.Delete, "del", -1, "Delete a secret by index")
	fs.IntVar(&cf.Edit, "edit", -1, "Edit a secret by index")
	fs.BoolVar(&cf.Export, "export", false, "Export secret in JSON")
}

func (cf *CmdFlags) Execute(secrets *Secrets, fs *flag.FlagSet) {
	switch {
	case cf.List:
		cf.listSecrets(secrets)
	case cf.Delete != -1:
		cf.deleteSecret(secrets)
	case cf.Add:
		cf.addSecret(secrets)
	case cf.Edit != -1:
		cf.editSecret(secrets, cf.Edit)
	case cf.Export:
		cf.exportSecret(secrets)
	default:
		fmt.Println("Invalid Command. Use --help to see available commands.")
	}
}

func (cf *CmdFlags) exportSecret(secrets *Secrets) {
	if err := secrets.Export(); err != nil {
		fmt.Println("Error exporting secret:", err)
		return
	}
	fmt.Println("Secret exported successfully!")
}

func (cf *CmdFlags) listSecrets(secrets *Secrets) {
	if err := secrets.ListSecrets(); err != nil {
		fmt.Println("Error listing secrets:", err)
		return
	}
}

func (cf *CmdFlags) deleteSecret(secrets *Secrets) {
	if err := secrets.Delete(cf.Delete); err != nil {
		fmt.Println("Error deleting secret:", err)
		return
	}
	fmt.Println("Secret deleted successfully!")
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
	password, err := getPassword(secrets)

	if err != nil {
		return
	}

	scanner.Scan()

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
	if err := secrets.Add(title, username, password, note, email, website); err != nil {
		fmt.Println("Error adding secret:", err)
		return
	}
	fmt.Println("Secret added successfully!")
}

func getPassword(secrets *Secrets) (string, error) {
	fmt.Print("Enter Password (press Tab for a suggested password): ")
	password := ""

	if err := keyboard.Open(); err != nil {
		return "", fmt.Errorf("error opening keyboard: %w", err)
	}
	defer keyboard.Close()

	for {
		char, key, err := keyboard.GetKey()

		if err != nil {
			return "", fmt.Errorf("error reading input: %w", err)
		}

		if key == keyboard.KeyEnter {
			break
		}

		if key == keyboard.KeyBackspace {
			if len(password) > 0 {
				password = password[:len(password)-1]
				printPassword(password)
			}
			continue
		}

		if key == keyboard.KeyTab {
			password = secrets.Suggest()
			printPassword(password)
			continue
		}

		if char != 0 {
			password += string(char)
			fmt.Print(string(char))
		}
	}

	return password, nil
}

func printPassword(password string) {
	fmt.Printf("\rEnter Password (press Tab for a suggested password): %s", password)
	fmt.Print("\033[K")
}

// Edit an existing secret
func (cf *CmdFlags) editSecret(secrets *Secrets, index int) {
	if err := secrets.Validate(index); err != nil {
		fmt.Println("Error validating secret index:", err)
		return
	}

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
				secret.Password, _ = utils.Encrypt(scanner.Text())
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

				// Call the Edit method in the secrets package
				if err := secrets.Edit(index, secret); err != nil {
					fmt.Println("Error updating secret:", err)
					return
				}
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
