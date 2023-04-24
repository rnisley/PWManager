package actions

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/rnisley/PWManager/db"
	"github.com/rnisley/PWManager/logger"
	"github.com/rnisley/PWManager/new"
	"golang.org/x/crypto/bcrypt"
)

// AddPW takes an app name and then
// prompt the user for a new username and password combo.
func AddPW() {
	var masterPass []byte
	err := authenticate(&masterPass)
	if err != nil {
		logger.Log(1, "Incorrect masterpass given to add new login")
		log.Fatalf("Unable to authenticate user")
	}

	app := getAppName()
	if LoginExists(app) {
		logger.Log(0, "Attempted to make new login for app already in DB")
		log.Fatalf("App login already exists. Try -update")
	}

	MPString := string(masterPass)
	user := encrypt(getUserName(app), MPString)
	password := encrypt(getAppPW(app), MPString)
	saveLogin(app, user, password)
	logger.Log(0,"Added new login to db")
}

// GetPW retrieves and decrypts the users password credentials and prints to the console.
func GetPW() {
	var masterPass []byte
	err := authenticate(&masterPass)
	if err != nil {
		logger.Log(1, "Incorrect masterpass given to password lookup")
		log.Fatalf("Unable to authenticate user")
	}

	app := getAppName()
	if LoginExists(app) {
		db := db.Connect().Db

		var username string
		var passhash string
		if err := db.QueryRow("SELECT username, passhash FROM Logins WHERE app=?;", app).Scan(&username, &passhash); err != nil {
			log.Println(err)
			log.Fatalf("An unexpected error occured checking the database")
		}
		var str string
		str = string(masterPass)
		var userBytes []byte
		userBytes = []byte(username)
		username = decrypt(userBytes, str)

		var passBytes []byte
		passBytes = []byte(passhash)
		passhash = decrypt(passBytes, str)

		fmt.Println("The username for " + app + " is:")
		fmt.Println(username)
		fmt.Println("The password for " + app + " is:")
		fmt.Println(passhash)
	} else {
		logger.Log(0, "User attempted to lookup an app that is not in the database")
		log.Fatalf("Application name is not found in database")
	}
	logger.Log(0, "User successfully looked up login credentials")
}

// UpdatePW takes an app name and then
// prompt the user for a new username and password combo.
func UpdatePW() {
	var masterPass []byte
	err := authenticate(&masterPass)
	if err != nil {
		logger.Log(1, "Incorrect masterpass given to update login credentials")
		log.Fatalf("Unable to authenticate user")
	}

	app := getAppName()
	if LoginExists(app) {
		MPString := string(masterPass)
		user := encrypt(getUserName(app), MPString)
		password := encrypt(getAppPW(app), MPString)
		updateLogin(app, user, password)

	} else {
		logger.Log(0, "User attempted to update an app that is not in the database")
		log.Fatalf("Application name is not found in database")
	}
	logger.Log(0, "User successfully updated login credentials")
}

// getAppName prompts the user for the app name or web address to save
// and returns it
func getAppName() string {
	fmt.Println("Enter app name or web address: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	text = strings.ReplaceAll(text, "\n", "")
	return strings.ToLower(text)
}

// getUserName prompts the user for the username to save with the app
// and returns it
func getUserName(app string) []byte {
	fmt.Println("Enter user name for " + app + ": ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	text = strings.ReplaceAll(text, "\n", "")
	return []byte(text)
}

// getUserName prompts the user for the password to save with the app
// and returns it
func getAppPW(app string) []byte {
	fmt.Println("Enter password for " + app + ": ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	text = strings.ReplaceAll(text, "\n", "")
	return []byte(text)
}

// prompts user for app Master Password and compares it to the one hash in the db
// also stores password in "pass" which is passed by reference so the password can
// be used as the passphrase for AES encryption.
func authenticate(pass *[]byte) error {

	if new.NoUsers() {
		log.Fatalf("Application has not been initialized. Try -setup")
	}

	var err error
	*pass, err = new.ReadPass()
	if err != nil {
		log.Fatalf("Error reading in password")
	}

	hash, err := getUserPassHash()
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword([]byte(hash), *pass)
}

// pulls master password hash from db
func getUserPassHash() (string, error) {
	db := db.Connect().Db

	var hash string
	if err := db.QueryRow("SELECT passhash FROM Logins WHERE id = 1;").Scan(&hash); err != nil {
		if err == sql.ErrNoRows {
			return "", &ErrNoUser{}
		} else {
			log.Println(err)
			log.Fatalf("An unexpected error occured checking the database")
		}
	}
	return hash, nil
}

// ErrNoUser is a generic error for no user existing
type ErrNoUser struct{}

func (e *ErrNoUser) Error() string {
	return "user does not exist"
}

// LoginExists takes an app name or url 'app' and returns true if that
// app exists in the database and false otherwise
func LoginExists(app string) bool {
	db := db.Connect().Db

	var id int
	if err := db.QueryRow("SELECT id FROM Logins WHERE app=?;", app).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Println(err)
			log.Fatalf("An unexpected error occured checking the database")
		}
	}
	return true
}

// saveLogin will process the transaction to place a login
// into the database
func saveLogin(app string, user []byte, password []byte) {
	database := db.Connect().Db

	database.Exec(`
		INSERT INTO Logins (app, username, passhash)
		VALUES (
			?, ?, ?
		);
	`, app, user, password)
}

// updateLogin will process the transaction to update an existing login
// in the database
func updateLogin(app string, user []byte, password []byte) {
	database := db.Connect().Db

	database.Exec(`
		UPDATE Logins
		SET username = ?, 
		passhash = ?
		WHERE 
			app = ?
		;
	`, user, password, app)
}

//AES Encryption Code Below

//Encrypts data using AES
func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

//Decrypts data using AES
func decrypt(data []byte, passphrase string) string {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return string(plaintext)
}

//Makes hash of string to use for AES encryption
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
