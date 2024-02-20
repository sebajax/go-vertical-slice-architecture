package user

import "time"

// Const for error messages
const (
	ErrorEmailExists string = "ERROR_EMAIL_EXISTS"
)

// User Domain
type User struct {
	Id             int
	IdentityNumber string
	FirstName      string
	LastName       string
	Email          string
	DateOfBirth    time.Time
	CreatedAt      time.Time
}

// Create a new user instance
func New(in string, n string, fn string, ln string, e string, dob time.Time) (*User, error) {
	return &User{
		IdentityNumber: in,
		FirstName:      fn,
		LastName:       ln,
		Email:          e,
		DateOfBirth:    dob,
	}, nil
}
