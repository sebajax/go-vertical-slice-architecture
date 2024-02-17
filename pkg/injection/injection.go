package injection

import (
	"os"

	"github.com/sebajax/go-vertical-slice-architecture/internal/user"
	"github.com/sebajax/go-vertical-slice-architecture/internal/user/infrastructure"
	"github.com/sebajax/go-vertical-slice-architecture/pkg/database"
	"go.uber.org/dig"
)

// container instance
var container *dig.Container

// user service instance
var UserServiceProvider user.UserService

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
	err = container.Provide(infrastructure.NewUserRepository)
	if err != nil {
		panic(err)
	}

	err = container.Provide(user.NewUserService)
	if err != nil {
		panic(err)
	}
}

// init service container
func InitComponents() error {
	// user service
	err := container.Invoke(
		func(us user.UserService) {
			UserServiceProvider = us
		},
	)

	// product instance

	return err
}
