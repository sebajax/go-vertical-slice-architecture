package user

import (
	"log"

	"github.com/sebajax/go-architecture-angrycoders/pkg/apperror"
)

// User use cases
type UserService struct {
	userRepository IUserRepository
}

// Create a new user service use case instance
func NewUserService(userRepository IUserRepository) *UserService {
	// return the pointer to user service
	return &UserService{
		userRepository,
	}
}

// Create a new user and store the user in the database
func (userService *UserService) CreateUser(user *User) (int64, error) {
	_, check, err := userService.userRepository.GetByEmail(user.Email)
	// check if user does not exist and no database error ocurred
	if err != nil {
		// database error
		log.Fatalln(err)
		return 0, apperror.InternalServerError()
	}
	if check != false {
		// no user found
		log.Println(user, ErrorEmailExists)
		return 0, apperror.NotFound(ErrorEmailExists)
	}

	// create the new user and return the id
	userId, err := userService.userRepository.Save(user)
	if err != nil {
		// database error
		log.Fatalln(err)
		return 0, apperror.InternalServerError()
	}

	// user created successfuly
	return userId, nil
}
