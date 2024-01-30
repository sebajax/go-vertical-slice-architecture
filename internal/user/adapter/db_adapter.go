package adapter

import (
	"gorm.io/gorm"
)

type DBUserAdapter struct {
	db *gorm.DB
}

func (dbUserAdapter *DBUserAdapter) AddUser(userModel *DBUserModel) (uint, error) {
	result := dbUserAdapter.db.Create(&userModel)
	return userModel.ID, result.Error
}

func (dbUserAdapter *DBUserAdapter) GetUserByEmail(email string) (DBUserModel, error) {
	var userModel DBUserModel
	result := dbUserAdapter.db.Where("email = ?", email).First(&userModel)
	return userModel, result.Error
}

/*
func (dbUserAdapter *DBUserAdapter) GetUsers() ([]*user.User, error) {

}

func  (dbUserAdapter *DBUserAdapter) GetUser(id int) (user.User, error) {

}
*/