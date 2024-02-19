package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebajax/go-vertical-slice-architecture/internal/user/api/handler"
	"github.com/sebajax/go-vertical-slice-architecture/internal/user/service"
)

// UserRouter is the Router for GoFiber App
func UserRouter(app fiber.Router, s *service.UserService) {
	app.Post("/", handler.CreateUser(s.CreateUserServiceProvider))
}
