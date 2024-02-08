package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sebajax/go-architecture-angrycoders/api/handlers"
	"github.com/sebajax/go-architecture-angrycoders/internal/user"
)

type UserSchema struct {
	Name        string    `json:"name" validate:"required, gte=5"`
	Email       string    `json:"email" validate:"required, email"`
	DateOfBirth time.Time `json:"date_of_birth" validate:"required datetime"`
}

// UserRouter is the Router for GoFiber App
func UserRouter(app fiber.Router, service user.UserService) {
	app.Post("/", handlers.CreateUser(service))
}
