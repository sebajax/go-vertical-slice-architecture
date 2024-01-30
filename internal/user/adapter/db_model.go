package adapter

import (
	"time"

	"github.com/sebajax/go-architecture-angrycoders/internal/user"
	"gorm.io/gorm"
)

type DBUserModel struct {
	gorm.Model
	Id          int
	Name        string
	Email       string
	DateOfBirth time.Time
}

func (userModel *DBUserModel) transformUserModel(user *user.User) {
	userModel.Name = user.Name
	userModel.Email = user.Email       
	userModel.DateOfBirth = user.DateOfBirth
}
