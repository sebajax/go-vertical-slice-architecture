package mocks

import (
	"time"

	"github.com/sebajax/go-vertical-slice-architecture/internal/user"
)

type mockUserRepository struct{}

func NewMockUserRepository() user.UserRepository {
	return &mockUserRepository{}
}

func (mock *mockUserRepository) Save(u *user.User) (int64, error) {
	return 1, nil
}

func (mock *mockUserRepository) GetByEmail(email string) (*user.User, bool, error) {
	return &user.User{
		Id:             1,
		IdentityNumber: "9999999",
		FirstName:      "Joe",
		LastName:       "Doe",
		DateOfBirth:    time.Now(),
		Email:          "joe@doe.com",
		CreatedAt:      time.Now(),
	}, true, nil
}
