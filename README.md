# PWManager

A Password Manager utility written in Go. Encrypt usernames and passwords in a database behind a master password to be retrieved at a later date.

This is a part of the University of Wyoming's Secure Software Design Course (Spring 2023). 

## Versioning

`PWManager` is built with:
- go version go1.20.2 linux/amd64

## Usage

`go run main.go --help` for instructions

Then run `go run main.go -new` and you will be prompted to create the master password.

## Database

Data gets stored into the local database file dd.db. Delete this file if you don't set up a user properly on the first go or if you wish to purge all saved logins.
