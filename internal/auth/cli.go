package auth

import (
	"flag"
	"fmt"
)

type CmdFlags struct {
	Expiry string
}

func NewCommandFlags() *CmdFlags {
	cf := CmdFlags{}
	return &cf
}

func (cf *CmdFlags) RegisterFlags(fs *flag.FlagSet) {
	fs.StringVar(&cf.Expiry, "set-expiry", "", "Set expiry date for a secret")
}

func (cf *CmdFlags) Execute(user *User, fs *flag.FlagSet) {
	fs.Parse(flag.Args()) // Ensure flag parsing occurs here
	if fs.Parsed() {      // Check if the flag set is parsed correctly
		switch {
		case cf.Expiry != "":
			cf.setExpiry(user)
		default:
			fmt.Println("Invalid Command. Use --help to see available commands.")
		}
	} else {
		fmt.Println("Error parsing flags.")
	}
}

func (cf *CmdFlags) setExpiry(user *User) {
	if err := user.SetExpiry(cf.Expiry); err != nil {
		fmt.Println("Error setting expiry date:", err)
		return
	}
	fmt.Println("Expiry date set successfully!")
}
