package user

import "time"

// Const for error messages
const (
	ErrorEmailExists string = "ERROR_EMAIL_EXISTS"
)

// Domain types
type name string
type email string
type dateOfBirth time.Time

// User Domain
type User struct {
	Id          int         `json:"id"`
	Name        name        `json:"name" validate:"required, gte=5"`
	Email       email       `json:"email" validate:"required, email"`
	DateOfBirth dateOfBirth `json:"date_of_birth" validate:"required datetime"`
	CreatedAt   time.Time   `json:"created_at"`
}

// Create a new user instance
func New(n string, e string, dob time.Time) (*User, error) {
	return &User{
		Name:        name(n),
		Email:       email(e),
		DateOfBirth: dateOfBirth(dob),
	}, nil
}
