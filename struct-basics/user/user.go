package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u *User) OutputUserDetails() {
	fmt.Println((*u).firstName, u.lastName, u.birthDate, u.createdAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{email: email, password: password, User: User{
		firstName: "ADMIN",
		lastName:  "ADMIN",
		birthDate: "------",
		createdAt: time.Now(),
	}}
}

func New(firstName string, lastName string, birthDate string) (*User, error) {

	if firstName == "" {
		return nil, errors.New("first name is blank")
	}

	if lastName == "" {
		return nil, errors.New("last name is blank")
	}

	if birthDate == "" {
		return nil, errors.New("birthDate is blank")
	}

	return &User{
		firstName,
		lastName,
		birthDate,
		time.Now(),
	}, nil
}
