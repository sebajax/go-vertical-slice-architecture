package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	producthandler "github.com/sebajax/go-vertical-slice-architecture/internal/product/handler"
	userhandler "github.com/sebajax/go-vertical-slice-architecture/internal/user/handler"
	"github.com/sebajax/go-vertical-slice-architecture/pkg/injection"
	"github.com/sebajax/go-vertical-slice-architecture/pkg/middleware"
)

func main() {
	// prepare all components for dependency injection
	injection.ProvideComponents()

	// initiate service components with dependency injection
	if err := injection.InitComponents(); err != nil {
		panic(err)
	}

	// create fiber
	app := fiber.New()

	// add fiber middlewares
	app.Use(cors.New())
	app.Use(helmet.New())

	// custom middlewares
	app.Use(middleware.ErrorHandler)

	// create health end point
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Status ok - api running")
	})

	// create api group
	api := app.Group("/api") // /api

	// add api group for user
	userApi := api.Group("/user") // /api/user
	userhandler.UserRouter(userApi, injection.UserServiceProvider)

	// add api group for product
	productApi := api.Group("/product") // /api/product
	producthandler.ProductRouter(productApi, injection.ProductServiceProvider)

	// listen in port 8080
	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("API_PORT"))))
}
