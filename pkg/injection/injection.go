package injection

import (
	"os"

	userservice "github.com/sebajax/go-vertical-slice-architecture/internal/user/service"
	"github.com/sebajax/go-vertical-slice-architecture/pkg/database"
	"go.uber.org/dig"
)

// container instance
var container *dig.Container

// new service instace
var UserServiceProvider *userservice.UserService

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

}

// init service container
func InitComponents() error {
	// user service init componets injection
	UserServiceProvider = userservice.NewUserService()
	err := UserServiceProvider.InitUserComponents(container)

	return err
}
