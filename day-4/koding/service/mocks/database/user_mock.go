package database_mock

import (
	"day2-task1/models"

	"github.com/stretchr/testify/mock"
)

type MockUser struct {
	mock.Mock
}

func (o *MockUser) GetUsers() (interface{}, error) {
	args := o.Called()

	var (
		err      error
		respData interface{}
	)

	if n, ok := args.Get(0).(*models.User); ok {
		respData = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}
	return respData, err
}

func (o *MockUser) GetUserById(id string) (interface{}, error) {
	args := o.Called(id)

	var (
		err      error
		respData interface{}
	)

	if n, ok := args.Get(0).(*models.User); ok {
		respData = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}
	return respData, err
}

func (o *MockUser) SaveUser(user *models.User) error {
	args := o.Called(user)

	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}
	return err
}

func (o *MockUser) UpdateUsers(user *models.User, id int) error {
	args := o.Called(user)

	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}
	return err
}

func (o *MockUser) DeleteUser(id int) error {
	args := o.Called(id)

	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}
	return err
}

func (o *MockUser) LoginUser(user *models.User) (interface{}, error) {
	args := o.Called(user)

	var (
		err      error
		respData interface{}
	)

	if n, ok := args.Get(0).(*models.User); ok {
		respData = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}
	return respData, err
}
