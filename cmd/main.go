package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/sebajax/go-architecture-angrycoders/api/routes"
	"github.com/sebajax/go-architecture-angrycoders/internal/user"
)

type mockUserRepository struct {}

func NewMockUserRepository() *mockUserRepository {
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
	return nil, false, nil
}

func main() {
	// connect to the database

	// inject dependecies
	mock := NewMockUserRepository() 
	userService := user.NewUserService(mock)

	// create fiber
	app := fiber.New()

	// add fiber middlewares
	app.Use(cors.New())
	app.Use(helmet.New())

	// create health end point
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Status ok - api running")
	})

	// add api group for users
	api := app.Group("/api") // /api
	userApi := api.Group("/users") // /api/user
	routes.UserRouter(userApi, userService)

	// listen in port 8080
	log.Fatal(app.Listen(":3000"))

}