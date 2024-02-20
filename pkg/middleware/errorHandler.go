package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sebajax/go-architecture-angrycoders/pkg/messages"
	"github.com/sebajax/go-vertical-slice-architecture/pkg/apperror"
	"github.com/sebajax/go-vertical-slice-architecture/pkg/message"
)

// ErrorHandler is a middleware that converts AppError to fiber.Error.
func ErrorHandler(c *fiber.Ctx) error {
	// Try to execute the next middleware/handler
	err := c.Next()

	// Check if there was an error
	if err != nil {
		// Log the error, handle it, or send a custom response
		if e, ok := err.(*apperror.AppError); ok {
			log.Error(message.ErrorResponse(e))
			return c.Status(e.Code).JSON(messages.ErrorResponse(e))
		}

		// An internal server error ocurred trying to cast error to apperror
		log.Error(message.ErrorResponse(err))
		return c.Status(fiber.StatusInternalServerError).JSON(message.ErrorResponse(err))
	}

	// If no error, continue the chain
	return nil
}
