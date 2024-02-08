package user

import "time"

// Const for error messages
const (
	ErrorEmailExists string = "ERROR_EMAIL_EXISTS"
)

// Domain types
// type name string
// type email string
// type dateOfBirth time.Time

// User Domain
type User struct {
	Id          int
	Name        string
	Email       string
	DateOfBirth time.Time
	CreatedAt   time.Time
}

// Create a new user instance
func New(n string, e string, dob time.Time) (*User, error) {
	return &User{
		Name:        n,
		Email:       e,
		DateOfBirth: dob,
	}, nil
}
