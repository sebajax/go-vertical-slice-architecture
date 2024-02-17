package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebajax/go-vertical-slice-architecture/internal/user"
	"github.com/sebajax/go-vertical-slice-architecture/internal/user/api/handler"
)

// UserRouter is the Router for GoFiber App
func UserRouter(app fiber.Router, service user.UserService) {
	app.Post("/", handler.CreateUser(service))
}
