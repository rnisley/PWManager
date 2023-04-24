package new

import (
	"fmt"
	"log"
	"os"

	"github.com/rnisley/PWManager/db"
	"github.com/rnisley/PWManager/logger"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

// Perform first time setup
func Initialize() {
	if !NoUsers() {
		logger.Log(1, "User attempted to initilize app again")
		log.Fatalf("This application is already initialized in this location.")
	}

	newPassHash, err := getPassword()
	if err != nil {
		log.Fatalf("Unable to get password hash")
	}

	err = setUserPassHash(newPassHash)
	if err != nil {
		log.Println(err)
		log.Fatalf("Unable to create new user")
	}
	logger.Log(0, "App initialized correctly")
}

// GetPassword will read in a password from stdin using the terminal
// no-echo utility ReadPassword. it will then salt and hash it with
// bcrypt
func getPassword() (string, error) {
	pass, err := ReadPass()
	if err != nil {
		return "", err
	}

	return saltAndHash(pass)
}

// SetUserPassHash will insert the master password hash
// into the logins table
func setUserPassHash(hash string) error {
	db := db.Connect().Db

	_, err := db.Exec(`
		INSERT INTO Logins (app, username, passhash)
		VALUES (
			"PWManager",
			"User",
			?
		);
	`, hash)
	return err
}

// a nice wrapper to encapsulate the bcrypt generateFromPassword func
// to make salting and hashing easier
func saltAndHash(pass []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// using the built in password read utility, get a password from stdin
func ReadPass() ([]byte, error) {
	fmt.Println("Password: ")
	pass, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}
	return pass, nil
}

// Returns true if no master password has been set
// otherwise returns false
func NoUsers() bool {
	db := db.Connect().Db

	var usersExist bool
	err := db.QueryRow("SELECT IIF(COUNT(*),'true', 'false') FROM Logins;").Scan(&usersExist)
	if err != nil {
		fmt.Println(err)
		return true
	}

	return !usersExist
}
