package user

import (
	"github.com/sebajax/go-architecture-angrycoders/internal/user/adapter"
)

type IUserRepository interface {
	AddUser(userModel *adapter.UserModel) (int, error)
	GetUserByEmail(email string) (adapter.UserModel, error)
	// GetUsers() ([]*User, error)
	// GetUser(id int) (User, error)
}

type IUserModel interface {
	transformUserModel(user *User) adapter.DBUserModel
}