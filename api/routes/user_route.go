package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebajax/go-architecture-angrycoders/api/handlers"
	"github.com/sebajax/go-architecture-angrycoders/internal/user"
)

// UserRouter is the Router for GoFiber App
func UserRouter(app fiber.Router, service user.UserService) {
	app.Post("/", handlers.CreateUser(service))
}