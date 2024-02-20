package service

import (
	"github.com/sebajax/go-vertical-slice-architecture/internal/user/infrastructure"
	"go.uber.org/dig"
)

// user service instance
type UserService struct {
	CreateUserServiceProvider CreateUserService
}

func NewUserService() *UserService {
	return &UserService{}
}

// provide components for injection
func ProvideUserComponents(c *dig.Container) {
	// repositorory provider injection
	err := c.Provide(infrastructure.NewUserRepository)
	if err != nil {
		panic(err)
	}

	//service provider injection
	err = c.Provide(NewCreateUserService)
	if err != nil {
		panic(err)
	}
}

// init service container
func (us *UserService) InitUserComponents(c *dig.Container) error {
	// create user service
	err := c.Invoke(
		func(s CreateUserService) {
			us.CreateUserServiceProvider = s
		},
	)

	return err
}
