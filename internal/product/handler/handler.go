package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebajax/go-vertical-slice-architecture/internal/product/service"
)

// ProductRouter is the Router for GoFiber App
func ProductRouter(app fiber.Router, s *service.ProductService) {
	app.Post("/", CreateProduct(s.CreateProductServiceProvider))
}
