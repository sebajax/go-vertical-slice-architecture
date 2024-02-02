package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/sebajax/go-architecture-angrycoders/internal/user"
	"github.com/sebajax/go-architecture-angrycoders/pkg/apperror"
	"github.com/sebajax/go-architecture-angrycoders/pkg/messages"
)

func CreateUser(service user.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody user.User
		err := c.BodyParser(&requestBody)
		if err != nil {
			err := apperror.BadRequest("BAD_REQUEST")
			c.Status(err.Code)
			return c.JSON(messages.ErrorResponse(err))
		}

		result, err := service.CreateUser(&requestBody)
		if err != nil {
			errors.Unwrap(err)
			return c.JSON(messages.ErrorResponse(err))
		}

		c.Status(fiber.StatusCreated)
		return c.JSON(messages.SuccessResponse(&result))
	}
}
