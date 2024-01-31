package user

import "time"

// Const for error messages
const (
	ErrorEmailExists string = "ERROR_EMAIL_EXISTS"
)

// User Domain
type User struct {
	Id          int       `json:"id"`
	Name        string    `json:"name" validate:"required, gte=5"`
	Email       string    `json:"email" validate:"required, email"`
	DateOfBirth time.Time `json:"date_of_birth" validate:"required datetime"`
	CreatedAt 	time.Time `json:"created_at"`
}

// Create a new user instance
func New(name string, email string, dateOfBirth time.Time, createdAt time.Time) (*User, error) {
	return &User{
		Name:        name,
		Email:       email,
		DateOfBirth: dateOfBirth,
		CreatedAt: 	 createdAt,
	}, nil
}
