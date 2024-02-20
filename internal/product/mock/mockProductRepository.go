package mocks

import "github.com/sebajax/go-vertical-slice-architecture/internal/user"

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
