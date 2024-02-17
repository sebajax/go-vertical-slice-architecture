package injection

import (
	"github.com/sebajax/go-vertical-slice-architecture/internal/user"
	"github.com/sebajax/go-vertical-slice-architecture/internal/user/mocks"
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
	// user provider injection
	container.Provide(mocks.NewMockUserRepository)
	container.Provide(user.NewUserService)
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
