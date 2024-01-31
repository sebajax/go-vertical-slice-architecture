package infrastructure

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
)

// User handlers for routing
type UserHandler struct {
	service *Service
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// Create a new handler instance
func NewHandler(s *Service) *UserHandler {
	return &UserHandler{service: s}
}

// CreateUser handles the POST request to create a new user.
func (h *Handler) CreateUser(c *fiber.Ctx) error {
	
}