package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebajax/go-architecture-angrycoders/internal/user"
	"github.com/sebajax/go-architecture-angrycoders/pkg/apperror"
	"github.com/sebajax/go-architecture-angrycoders/pkg/messages"
)

// Creates a new user into the database
func CreateUser(service user.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get body request
		var requestBody user.User
		// Validate the body
		err := c.BodyParser(&requestBody)
		if err != nil {
			// Map the error and response via the middleware
			return apperror.BadRequest("BAD_REQUEST")
		}

		// Execute the service
		result, err := service.CreateUser(&requestBody)
		if err != nil {
			// if service response an error return via the middleware
			return err
		}

		// Success execution
		return c.Status(fiber.StatusCreated).JSON(messages.SuccessResponse(&result))
	}
}
