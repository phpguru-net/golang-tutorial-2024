// we'll create a struct in go, it is custom type for complex data structure
package user

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type ID string

func (id ID) GenerateUUID() string {
	return uuid.New().String()
}

type User struct {
	id        ID
	firstName string
	lastName  string
	prize     string
}

type Winner struct {
	email string
	User
}

// receiver argument
func (u *User) DisplayUser() string {
	output := ""
	output += fmt.Sprintf("ID: %v\nName: %v %v\nPrize: %v", u.id, u.firstName, u.lastName, u.prize)
	return output
}

func (u *Winner) DisplayUser() string {
	output := ""
	output += fmt.Sprintf("%v\nEmail:%v", u.User.DisplayUser(), u.email)
	return output
}

func (u *User) SetPrize(prize string) {
	u.prize = prize
}

// A pattern to create new struct
func New(id ID, firstName string, lastName string, prize string) (*User, error) {
	if id == "" || firstName == "" || lastName == "" {
		return nil, errors.New("ID, First name and last name are required!")
	}
	return &User{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		prize:     prize,
	}, nil
}

func NewWinner(email string, u *User) *Winner {
	return &Winner{
		email: email,
		User: User{
			id:        u.id,
			firstName: u.firstName,
			lastName:  u.lastName,
			prize:     u.prize,
		},
	}
}
