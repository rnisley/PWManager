package main

import (
	"flag"
	"fmt"

	"github.com/rnisley/PWManager/new"
	"github.com/rnisley/PWManager/actions"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Run with -help for help information")
	}

	var setup, add, get, update, help bool
	flag.BoolVar(&help, "help", false, "Get help")
	flag.BoolVar(&update, "update", false, "run the utility in update mode")
	flag.BoolVar(&setup, "setup", false, "first time setup")
	flag.BoolVar(&add, "add", false, "run the utility in add mode")
	flag.BoolVar(&get, "get", false, "run the utility in get mode")
	flag.Parse()

	if help {
		pad := func() {
			fmt.Printf("\n\n")
		}

		pad()
		fmt.Println(" Welcome to PasswordManager.")
		pad()
		fmt.Println("Args:")
		fmt.Println("  -setup   The verb flag to specify you want to setup the PWManager")
		fmt.Println("  -add  The verb flag to specify you want to add a new password")
		fmt.Println("  -get  The verb flag to specify you want to retreive a password")
		fmt.Println("  -update  The verb flag to specify you want to update a username and password")
		pad()
		return
	}

	if !add && !update && !setup && !get {
		fmt.Println("Please specify a verb for the utility.")
		fmt.Println("Valid verbs: get, add, update, help")
		return
	}

	if (setup && (help || add || get || update)) || (add && (setup || help || get || update)) || (get && (setup || add || help || update)) || (update && (setup || add || get || help)) {
		fmt.Println("Please specify only one verb")
		return
	}

	if get {
		actions.GetPW()
	} else if add {
		actions.AddPW()
	} else if setup {
		new.Initialize()
	} else if update {
		actions.UpdatePW()
	}
}
