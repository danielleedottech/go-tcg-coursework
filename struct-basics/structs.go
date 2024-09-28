package main

import (
	"fmt"
	"github.com/danielleetech/go-tcg-coursework/struct-basics/user"
	"log"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(firstName, lastName, birthDate)
	if err != nil {
		log.Fatal(err)
	}

	admin := user.NewAdmin("dan@gmail.com", "1234")
	admin.OutputUserDetails()
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
