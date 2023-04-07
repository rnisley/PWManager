package logger

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
)

func Log(event int, user string) {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime)

	switch event {
	case 0:
		InfoLogger.Println(user + " read messages")
	case 1:
		InfoLogger.Println("Message sent to " + user)
	case 2:
		InfoLogger.Println(user + " created a new user")
	case 3:
		WarningLogger.Println("Wrong password used to read user " + user)
	case 4:
		WarningLogger.Println("Attempted to read messages for user that does not exist :" + user)
	case 5:
		WarningLogger.Println("Attempted to send message to user that does not exist :" + user)
	case 6:
		WarningLogger.Println("Attempted to create new user using invalid user :" + user)
	case 7:
		WarningLogger.Println(user + " attempted to create new user using invalid password")
	case 8:
		WarningLogger.Println(user + " attempted to send message, but " + user + " is not a valid user")
	case 9:
		WarningLogger.Println(user + " attempted to send message using invalid password")
	}

}

func NoAuth(user string) {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime)
	WarningLogger.Println("Message for " + user + " could not be authenticated.")
}
