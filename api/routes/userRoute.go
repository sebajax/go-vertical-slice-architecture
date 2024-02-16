package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebajax/go-vertical-slice-architecture/api/handlers"
	"github.com/sebajax/go-vertical-slice-architecture/internal/user"
)

// UserRouter is the Router for GoFiber App
func UserRouter(app fiber.Router, service user.UserService) {
	app.Post("/", handlers.CreateUser(service))
}
