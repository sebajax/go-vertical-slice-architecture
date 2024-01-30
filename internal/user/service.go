package user

import (
	"log"

	"github.com/sebajax/go-architecture-angrycoders/pkg/apperror"
)

type UserService struct {
	userRepository IUserRepository
	userModel      IUserModel
}

func NewUserService(userRepository IUserRepository, userModel IUserModel) UserService {
	return UserService{
		userRepository,
		userModel,
	}
}

func (userService *UserService) CreateUser(user *User) (int, error) {
	_, err := userService.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		// log the error & return the error
		log.Println(user, ErrorUserEmailExists)
		return 0, apperror.BadRequest(ErrorUserEmailExists)
	}

	userModel := userService.userModel.transformUserModel(user)
	userId, err := userService.userRepository.AddUser(&userModel)
	if err != nil {
		// return nil, apperror.NewBadRequestError("DATABASE_ERROR")
	}

	return userId, nil

}
