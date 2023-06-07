package models

import "github.com/Poomipat-Ch/StockManagement/pkg/utils"

type User struct {
	ID             int    `json:"id" db:"id"`
	FirstName      string `json:"first_name" db:"first_name" validate:"required,min=2,max=100"`
	LastName       string `json:"last_name" db:"last_name" validate:"required,min=2,max=100"`
	Email          string `json:"email" db:"email" validate:"required,email"`
	PasswordHashed string `json:"-" db:"password_hashed" validate:"required,min=8,max=100"`
}

func CreateNewUser(firstName string, lastName string, email string, password string) (*User, error) {

	passwordHashed, err := utils.HashPassword(password)

	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:      firstName,
		LastName:       lastName,
		Email:          email,
		PasswordHashed: passwordHashed,
	}, nil
}
