package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"

	"mockingtest/domain"
	"mockingtest/repo"
)

var (
	testUserId = "test-user-id"
)

func TestAddUser_InvalidName(t *testing.T) {
	req := domain.AddUserRequest{
		Name:     "",
		Email:    "test@test.com",
		Password: "password",
	}

	userService := NewUserService(nil)
	userId, err := userService.AddUser(req)

	assert.Empty(t, userId)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrInvalidName)
}

func TestAddUser_InvalidEmail(t *testing.T) {
	req := domain.AddUserRequest{
		Name:     "test",
		Email:    "",
		Password: "password",
	}

	userService := NewUserService(nil)
	userId, err := userService.AddUser(req)

	assert.Empty(t, userId)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrInvalidEmail)
}

func TestAddUser_InvalidPassword(t *testing.T) {
	req := domain.AddUserRequest{
		Name:     "test",
		Email:    "test@test.com",
		Password: "pass",
	}

	userService := NewUserService(nil)
	userId, err := userService.AddUser(req)

	assert.Empty(t, userId)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrInvalidPassword)
}

func TestAddUser_ErrorAddToDB(t *testing.T) {

	req := domain.AddUserRequest{
		Name:     "test",
		Email:    "test@test.com",
		Password: "password",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := repo.NewMockUserRepository(ctrl)

	mockUserRepo.
		EXPECT().
		AddUser(gomock.Any(), gomock.Any()).
		Return("", mongo.ErrNilDocument)

	userService := NewUserService(mockUserRepo)
	userId, err := userService.AddUser(req)

	assert.Empty(t, userId)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, mongo.ErrNilDocument)
}

func TestAddUser_Success(t *testing.T) {

	req := domain.AddUserRequest{
		Name:     "test",
		Email:    "test@test.com",
		Password: "password",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := repo.NewMockUserRepository(ctrl)

	mockUserRepo.
		EXPECT().
		AddUser(gomock.Any(), gomock.Any()).
		Return(testUserId, nil)

	userService := NewUserService(mockUserRepo)
	userId, err := userService.AddUser(req)

	assert.NotEmpty(t, userId)
	assert.Nil(t, err)
	assert.Equal(t, userId, testUserId)
}
