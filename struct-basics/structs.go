package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *user) outputUserDetails() {
	fmt.Println((*u).firstName, u.lastName, u.birthDate, u.createdAt)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName string, lastName string, birthDate string) (*user, error) {

	if firstName == "" {
		return nil, errors.New("first name is blank")
	}

	if lastName == "" {
		return nil, errors.New("last name is blank")
	}

	if birthDate == "" {
		return nil, errors.New("birthDate is blank")
	}

	return &user{
		firstName,
		lastName,
		birthDate,
		time.Now(),
	}, nil
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := newUser(firstName, lastName, birthDate)
	if err != nil {
		log.Fatal(err)
	}
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
