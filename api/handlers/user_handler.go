package handlers

import (
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
			return c.Status(err.Code).JSON(messages.ErrorResponse(err))
		}

		result, err := service.CreateUser(&requestBody)
		if err != nil {
			// Log the error, handle it, or send a custom response
			if e, ok := err.(*apperror.AppError); ok {
				return c.Status(e.Code).JSON(messages.ErrorResponse(e))		
			}
			// An internal server error ocurred trying to cast error to apperror 
			return c.Status(fiber.StatusInternalServerError).JSON(messages.ErrorResponse(err))
		}

		// No errors
		return c.Status(fiber.StatusCreated).JSON(messages.SuccessResponse(&result))
	}
}
