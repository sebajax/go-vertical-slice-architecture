package user

import "time"

/*
* User package api errors
 */

const (
	// User-related errors
	ErrorUserNotFound           string = "ERROR_USER_NOT_FOUND"
	ErrorUserInvalidCredentials string = "ERROR_INVALID_CREDENTIALS"
	ErrorUserInactive           string = "ERROR_USER_INACTIVE"
	ErrorUserUnauthorized       string = "ERROR_USER_UNAUTHORIZED"
	ErrorUserPasswordMismatch   string = "ERROR_PASSWORD_MISMATCH"

	// Authentication & Authorization errors
	ErrorAuthTokenInvalid            string = "ERROR_AUTH_TOKEN_INVALID"
	ErrorAuthTokenExpired            string = "ERROR_AUTH_TOKEN_EXPIRED"
	ErrorAuthInsufficientPermissions string = "ERROR_INSUFFICIENT_PERMISSIONS"

	// Resource-related errors
	ErrorResourceNotFound    string = "ERROR_RESOURCE_NOT_FOUND"
	ErrorResourceConflict    string = "ERROR_RESOURCE_CONFLICT"
	ErrorResourceUnavailable string = "ERROR_RESOURCE_UNAVAILABLE"

	// Validation errors
	ErrorValidationFailed     string = "ERROR_VALIDATION_FAILED"
	ErrorInvalidInput         string = "ERROR_INVALID_INPUT"
	ErrorMissingRequiredField string = "ERROR_MISSING_REQUIRED_FIELD"

	// System errors
	ErrorSystemFailure      string = "ERROR_SYSTEM_FAILURE"
	ErrorServiceUnavailable string = "ERROR_SERVICE_UNAVAILABLE"
	ErrorOperationTimeout   string = "ERROR_OPERATION_TIMEOUT"
)

type User struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

func New(name string, email string, dateOfBirth time.Time) (User, error) {
	return User{
		Name:        name,
		Email:       email,
		DateOfBirth: dateOfBirth,
	}, nil
}

