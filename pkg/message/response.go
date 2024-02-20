package message

import "github.com/gofiber/fiber/v2"

// SuccessResponseSlice is the list SuccessResponse that will be passed in the response by Handler
func SuccessResponseSlice(data *[]any) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// SuccessResponse is the primitive type SuccessResponse that will be passed in the response by Handler
func SuccessResponse(data any) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// ErrorResponse is the ErrorResponse that will be passed in the response by Handler
func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
