package logger

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
)

func Log(event int) {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime)

	switch event {
	case 0:
		InfoLogger.Println("App initialized correctly")
	case 1:
		WarningLogger.Println("User attempted to initilize app again")
	case 2:
		WarningLogger.Println("Incorrect masterpass given to add new login")
	case 3:
		InfoLogger.Println("Attempted to make new login for app already in DB")
	case 4:
		InfoLogger.Println("Added new login to db")
	case 5:
		WarningLogger.Println("Incorrect masterpass given to password lookup")
	case 6:
		InfoLogger.Println("User successfully looked up login credentials")
	case 7:
		WarningLogger.Println("Incorrect masterpass given to update login credentials")
	case 8:
		InfoLogger.Println("User successfully updated login credentials")
	}
}
