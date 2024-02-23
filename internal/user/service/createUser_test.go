package service

import (
	"errors"
	"testing"
	"time"

	"github.com/sebajax/go-vertical-slice-architecture/internal/user"
	"github.com/sebajax/go-vertical-slice-architecture/internal/user/mock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserService_CreateUser_Success(t *testing.T) {
	var insertedId int64 = 1
	// Create a new User
	u, _ := user.New("identity_test", "fistname_test", "lastname_test", "email_test", time.Now())

	// Create a mock UserRepository instance
	mockUserRepository := &mock.UserRepository{}

	// Set expectations on the mock
	mockUserRepository.EXPECT().GetByEmail(u.Email).Return(nil, false, nil)
	mockUserRepository.EXPECT().Save(u).Return(insertedId, nil)

	// Pass the mock as a dependency
	userService := NewCreateUserService(mockUserRepository)

	// Call the method being tested
	id, err := userService.CreateUser(u)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, insertedId, id)

	// Assert that the method was called exactly once or not called
	mockUserRepository.AssertCalled(t, "GetByEmail", u.Email)
	mockUserRepository.AssertCalled(t, "Save", u)

	// Assert that the expectations were met
	mockUserRepository.AssertExpectations(t)
}

func TestCreateProductService_CreateProduct_GetByEmail_Failure(t *testing.T) {
	// Create a new User
	u, _ := user.New("identity_test", "fistname_test", "lastname_test", "email_test", time.Now())

	// Create a mock UserRepository instance
	mockUserRepository := &mock.UserRepository{}

	// Set expectations on the mock
	mockUserRepository.EXPECT().GetByEmail(u.Email).Return(u, true, nil)

	// Pass the mock as a dependency
	userService := NewCreateUserService(mockUserRepository)

	// Call the method being tested
	id, err := userService.CreateUser(u)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, int64(0), id)

	// Assert that the method was called exactly once or not called
	mockUserRepository.AssertCalled(t, "GetByEmail", u.Email)
	mockUserRepository.AssertNotCalled(t, "Save")

	// Assert that the expectations were met
	mockUserRepository.AssertExpectations(t)
}

func TestCreateProductService_CreateProduct_Save_Failure(t *testing.T) {
	// Create a new User
	u, _ := user.New("identity_test", "fistname_test", "lastname_test", "email_test", time.Now())

	// Create a mock UserRepository instance
	mockUserRepository := &mock.UserRepository{}

	// Set expectations on the mock
	mockUserRepository.EXPECT().GetByEmail(u.Email).Return(nil, false, nil)
	mockUserRepository.EXPECT().Save(u).Return(0, errors.New("DB ERROR"))

	// Pass the mock as a dependency
	userService := NewCreateUserService(mockUserRepository)

	// Call the method being tested
	id, err := userService.CreateUser(u)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, int64(0), id)

	// Assert that the method was called exactly once or not called
	mockUserRepository.AssertCalled(t, "GetByEmail", u.Email)
	mockUserRepository.AssertCalled(t, "Save", u)

	// Assert that the expectations were met
	mockUserRepository.AssertExpectations(t)
}
