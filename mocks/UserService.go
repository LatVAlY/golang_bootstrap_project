// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	model "golang_bootstrap_project/model"

	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// FindActiveUserWithUserId provides a mock function with given fields: userId
func (_m *UserService) FindActiveUserWithUserId(userId string) (*model.UserDTO, error) {
	ret := _m.Called(userId)

	var r0 *model.UserDTO
	if rf, ok := ret.Get(0).(func(string) *model.UserDTO); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserService(t mockConstructorTestingTNewUserService) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}