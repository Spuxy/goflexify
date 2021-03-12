package service

import (
	"testing"

	"github.com/Spuxy/Goflexify/model"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (d *MockRepository) CreateUser(u *model.User) error {
	return nil
}

func (u *MockRepository) GetUserById(id int) (model.User, error) {
	return model.User{}, nil
}

func (d *MockRepository) GetUserByEmailPassword(email, password string) (model.User, error) {
	args := d.Called(email, password)
	return args.Get(0).(model.User), args.Error(1)
}

func (d *MockRepository) UpdateUser(u *model.User) error {
	return nil
}

func (d *MockRepository) DeleteUserById(id int) error {
	return nil
}

func TestGetUserInvalidPassword(t *testing.T) {
	mockl := MockRepository{}
	server := NewUserService(&mockl)
	m, err := server.GetUser("GF", "R")
	assert.NotNil(t, err)
	assert.Empty(t, m, "user dsadawas empty")
}

func TestGetUserInvalidEmail(t *testing.T) {
	var user model.User
	email, password := "test@test.hu", "helloworld"
	mockObject := new(MockRepository)
	mockObject.On("GetUserByEmailPassword", email, password).Return(user, nil)
	server := NewUserService(mockObject)
	m, err := server.GetUser("testtest.com", "agf3")
	assert.NotNil(t, err)
	assert.Empty(t, m, "user dsadawas empty")
}

func TestGetUserRepositoryByEmailPassword(t *testing.T) {
	var user model.User
	email, password := "test@test.hu", "helloworld"
	mockObject := new(MockRepository)
	mockObject.On("GetUserByEmailPassword", email, password).Return(user, nil)
	server := NewUserService(mockObject)
	user, err := server.GetUser(email, password)
	assert.Nil(t, err, "it returns correctly struct of user")
	assert.NotNil(t, user, "the user is not empty")
}
