package user

import (
	"log"

	"github.com/sebajax/go-architecture-angrycoders/pkg/apperror"
)

// UserService interface for DI
type UserService interface {
	CreateUser(user *User) (int64, error)
}

// User use cases (port injection)
type userService struct {
	userRepository UserRepository
}

// Create a new user service use case instance
func NewUserService(repository UserRepository) UserService {
	// return the pointer to user service
	return &userService{
		userRepository: repository,
	}
}

// Create a new user and store the user in the database
func (service *userService) CreateUser(user *User) (int64, error) {
	_, check, err := service.userRepository.GetByEmail(user.Email)
	// check if user does not exist and no database error ocurred
	if err != nil {
		// database error
		log.Fatalln(err)
		err := apperror.InternalServerError()
		return 0, err
	}
	if check {
		// user found
		log.Println(user, ErrorEmailExists)
		err := apperror.BadRequest(ErrorEmailExists)
		return 0, err
	}

	// create the new user and return the id
	userId, err := service.userRepository.Save(user)
	if err != nil {
		// database error
		log.Fatalln(err)
		err := apperror.InternalServerError()
		return 0, err
	}

	// user created successfuly
	return userId, nil
}
