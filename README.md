# PWManager

A Password Manager utility written in Go. Encrypt usernames and passwords in a database behind a master password to be retrieved at a later date.

This is a part of the University of Wyoming's Secure Software Design Course (Spring 2023). 

## Versioning

`PWManager` is built with:
- go version go1.20.2 linux/amd64

## Usage

`go run main.go` Is used to launch the program, but must be followed by exactly one of the following tags.

`-setup`   Specifies you want to setup the PWManager. This can only be run once in any given installation location. This will prompt you to set the master password.

`-add`  Specifies you want to add a new application with a username and password pair. This will prompt you for the master password, then ask you for an app name or web address, username, and password to insert.

`-get`  Specifies you want to retreive a username and password combination. This will prompt you for the master password, then ask you for an app name or web address and return the appropriate username password combination.

`-update`  Specifies you want to update a username and password for an app that already exists in the database. This will prompt you for the master password, then ask you for an app name or web address, username, and password to update.

## Database

Data gets stored into the local database file dd.db. Delete this file if you don't set up a user properly on the first go or if you wish to purge all saved logins.

## Documentation

See `Project Design Documentation.pdf` for full design specifications.
