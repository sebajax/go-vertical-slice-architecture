package user

import "time"

/*
* User package api errors
 */
const (
	ErrorUserEmailExists string = "ERROR_USER_EMAIL_EXISTS"
)

type User struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Email       string `json:"email"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

func New(name string, email string, dateOfBirth time.Time) (User, error) {
	return User{
		Name: name,
		Email: email, 
		DateOfBirth: dateOfBirth,
	}, nil
}