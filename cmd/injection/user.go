package injection

import (
	"github.com/sebajax/go-vertical-slice-architecture/internal/user"
	"go.uber.org/dig"
)

// INI: DELETE THIS ONCE DATABASE IS WORKING
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

// END: DELETE THIS ONCE DATABASE IS WORKING

// user instance
var UserServiceProvider user.UserService

// user init container
func initUserComponents() func(u user.UserService) {
	// user instance
	return func(u user.UserService) {
		UserServiceProvider = u
	}
}

// user provider
func provideUserComponents(container *dig.Container) {
	container.Provide(NewMockUserRepository)
	container.Provide(user.NewUserService)
}
