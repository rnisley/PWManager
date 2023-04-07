package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Run with -help for help information")
	}

	var new, add, get, update, help bool
	flag.BoolVar(&help, "help", false, "Get help")
	flag.BoolVar(&update, "update", false, "run the utility in update mode")
	flag.BoolVar(&new, "new", false, "first time setup")
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
		fmt.Println("  -new   The verb flag to specify you want to create a new user")
		fmt.Println("  -add  The verb flag to specify you want to add a new password")
		fmt.Println("  -get  The verb flag to specify you want to retreive a password")
		fmt.Println("  -update  The verb flag to specify you want to update a username and password")
		pad()
		return
	}

	if !add && !update && !new && !get {
		fmt.Println("Please specify a verb for the utility.")
		fmt.Println("Valid verbs: get, add, update, help")
		return
	}

	if (new && (help || add || get || update)) || (add && (new || help || get || update)) || (get && (new || add || help || update)) || (update && (new || add || get || help)) {
		fmt.Println("Please specify only one verb")
		return
	}

	if get {
		read.getPW()
	} else if add {
		send.addPW()
	} else if new {
		new.initialize()
	} else if update {
		update.updatePW()
	}
}
