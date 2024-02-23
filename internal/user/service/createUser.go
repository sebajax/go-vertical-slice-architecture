package service

import (
	"log"

	"github.com/sebajax/go-vertical-slice-architecture/internal/user"
	"github.com/sebajax/go-vertical-slice-architecture/pkg/apperror"
)

// User use cases (port injection)
type CreateUserService struct {
	userRepository user.UserRepository
}

// Create a new user service use case instance
func NewCreateUserService(repository user.UserRepository) *CreateUserService {
	// return the pointer to user service
	return &CreateUserService{
		userRepository: repository,
	}
}

// Create a new user and store the user into the database
func (service *CreateUserService) CreateUser(u *user.User) (int64, error) {
	_, check, err := service.userRepository.GetByEmail(u.Email)
	// check if user does not exist and no database error ocurred
	if err != nil {
		// database error
		log.Println(err)
		err := apperror.InternalServerError()
		return 0, err
	}
	if check {
		// user found
		log.Println(u, user.ErrorEmailExists)
		err := apperror.BadRequest(user.ErrorEmailExists)
		return 0, err
	}

	// create the new user and return the id
	id, err := service.userRepository.Save(u)
	if err != nil {
		// database error
		log.Println(err)
		err := apperror.InternalServerError()
		return 0, err
	}

	// user created successfuly
	return id, nil
}
