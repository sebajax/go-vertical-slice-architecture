package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/sebajax/go-architecture-angrycoders/api/middlewares"
	"github.com/sebajax/go-architecture-angrycoders/api/routes"
	"github.com/sebajax/go-architecture-angrycoders/internal/user"
	"go.uber.org/dig"
)

type mockUserRepository struct{}

func NewMockUserRepository() user.UserRepository {
	return &mockUserRepository{}
}

func (mock *mockUserRepository) Save(u *user.User) (int64, error) {
	return 1, nil
}

func (mock *mockUserRepository) GetByEmail(email string) (*user.User, bool, error) {
	/*return &user.User{
		Id: 1,
		Email: "juan@example.com",
		Name: "Juan",
		DateOfBirth: time.Now(),
		CreatedAt: time.Now(),
	}, true, nil*/
	return nil, true, nil
}

var container *dig.Container

var us user.UserService

func main() {
	// connect to the database

	// prepare all components for dependency injection
	ProvideUserComponents()

	// instantiate all components with dependency injection
	err := initComponents()
	if err != nil {
		panic(err)
	}

	// userService := user.NewUserService(repo)

	// create fiber
	app := fiber.New()

	// add fiber middlewares
	app.Use(cors.New())
	app.Use(helmet.New())

	// custom middlewares
	app.Use(middlewares.ErrorHandler)

	// create health end point
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Status ok - api running")
	})

	// add api group for users
	api := app.Group("/api")       // /api
	userApi := api.Group("/users") // /api/user
	routes.UserRouter(userApi, us)

	// listen in port 8080
	log.Fatal(app.Listen(":3000"))

}

func initComponents() error {
	return container.Invoke(
		func(u user.UserService) {
			us = u
		},
	)
}

// inject provider for user components
func ProvideUserComponents() {
	container = dig.New()
	container.Provide(NewMockUserRepository)
	container.Provide(user.NewUserService)
}
