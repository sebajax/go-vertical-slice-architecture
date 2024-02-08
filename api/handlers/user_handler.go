package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sebajax/go-architecture-angrycoders/internal/user"
	"github.com/sebajax/go-architecture-angrycoders/pkg/apperror"
	"github.com/sebajax/go-architecture-angrycoders/pkg/messages"
)

// Body request schema for CreateUser
type UserSchema struct {
	IdentityNumber string    `json:"identity_number" validate:"required,min=6"`
	FirstName      string    `json:"first_name" validate:"required,min=2"`
	LastName       string    `json:"last_name" validate:"required,min=2"`
	Email          string    `json:"email" validate:"required,email"`
	DateOfBirth    time.Time `json:"date_of_birth" validate:"required"`
}

// Creates a new user into the database
func CreateUser(service user.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get body request
		var body UserSchema
		// Validate the body
		err := c.BodyParser(&body)
		if err != nil {
			// Map the error and response via the middleware
			log.Error(err)
			return err
		}

		// Validate schema
		serr, err := Validate(body)
		if err != nil {
			log.Error(serr)
			return apperror.BadRequest(serr)
		}

		// No schema errores then map body to domain
		user := &user.User{
			IdentityNumber: body.IdentityNumber,
			FirstName:      body.FirstName,
			LastName:       body.LastName,
			Email:          body.Email,
			DateOfBirth:    body.DateOfBirth,
		}

		// Execute the service
		result, err := service.CreateUser(user)
		if err != nil {
			// if service response an error return via the middleware
			log.Error(err)
			return err
		}

		// Success execution
		return c.Status(fiber.StatusCreated).JSON(messages.SuccessResponse(&result))
	}
}
