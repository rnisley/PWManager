package new

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rnisley/PWManager/db"
	"github.com/rnisley/PWManager/logger"
	"github.com/rnisley/PWManager/session"
)

// Create a NewUser as authorized by the user 'user'
func NewUser(user string) {
	if !db.NoUsers(){
		logger.Log(//) add appropriate log
		log.Fatalf("This application is already initialized in this location.")
	}

	newUser := getNewUsername()
	newPassHash, err := session.GetPassword()
	if err != nil {
		log.Fatalf("Unable to get password hash")
	}

	err = db.SetUserPassHash("PWManager", newUser, newPassHash)
	if err != nil {
		log.Fatalf("Unable to create new user")
	}
	logger.Log(2, user) // update log call
}

// getUserMessage prompts the user for the message to send
// and returns it
func getNewUsername() string {
	fmt.Println("Enter the username for the new user: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	return strings.Trim(text, "\n\t ")
}
