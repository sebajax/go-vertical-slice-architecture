package injection

import (
	"os"

	productservice "github.com/sebajax/go-vertical-slice-architecture/internal/product/service"
	userservice "github.com/sebajax/go-vertical-slice-architecture/internal/user/service"
	"github.com/sebajax/go-vertical-slice-architecture/pkg/database"
	"go.uber.org/dig"
)

// container instance
var container *dig.Container

// new user service instace
var UserServiceProvider *userservice.UserService

// new product service instace
var ProductServiceProvider *productservice.ProductService

// provide components for injection
func ProvideComponents() {
	// create a new container
	container = dig.New()

	// generate db config instance
	err := container.Provide(func() *database.DbConfig {
		config := database.NewDbConfig(
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
		)
		config.WithMigration(os.Getenv("MIGRATIONS_PATH"))

		return config
	})
	if err != nil {
		panic(err)
	}

	// provide the database connection injection
	err = container.Provide(database.InitPool)
	if err != nil {
		panic(err)
	}

	// user provider injection
	userservice.ProvideUserComponents(container)

	// product provider injection
	productservice.ProvideProductComponents(container)
}

// init service container
func InitComponents() error {
	// user service init componets injection
	UserServiceProvider = userservice.NewUserService()
	err := UserServiceProvider.InitUserComponents(container)
	if err != nil {
		panic(err)
	}

	// product service init componets injection
	ProductServiceProvider = productservice.NewProductService()
	err = ProductServiceProvider.InitProductComponents(container)
	if err != nil {
		panic(err)
	}

	return err
}
