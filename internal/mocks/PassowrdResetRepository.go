// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	model "github.com/saadozone/gin-gorm-rest/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// PassowrdResetRepository is an autogenerated mock type for the PassowrdResetRepository type
type PassowrdResetRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: passwordReset
func (_m *PassowrdResetRepository) Delete(passwordReset *model.PasswordReset) (*model.PasswordReset, error) {
	ret := _m.Called(passwordReset)

	var r0 *model.PasswordReset
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.PasswordReset) (*model.PasswordReset, error)); ok {
		return rf(passwordReset)
	}
	if rf, ok := ret.Get(0).(func(*model.PasswordReset) *model.PasswordReset); ok {
		r0 = rf(passwordReset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PasswordReset)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.PasswordReset) error); ok {
		r1 = rf(passwordReset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByToken provides a mock function with given fields: token
func (_m *PassowrdResetRepository) FindByToken(token string) (*model.PasswordReset, error) {
	ret := _m.Called(token)

	var r0 *model.PasswordReset
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.PasswordReset, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) *model.PasswordReset); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PasswordReset)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUserId provides a mock function with given fields: id
func (_m *PassowrdResetRepository) FindByUserId(id int) (*model.PasswordReset, error) {
	ret := _m.Called(id)

	var r0 *model.PasswordReset
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*model.PasswordReset, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *model.PasswordReset); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PasswordReset)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: passwordReset
func (_m *PassowrdResetRepository) Save(passwordReset *model.PasswordReset) (*model.PasswordReset, error) {
	ret := _m.Called(passwordReset)

	var r0 *model.PasswordReset
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.PasswordReset) (*model.PasswordReset, error)); ok {
		return rf(passwordReset)
	}
	if rf, ok := ret.Get(0).(func(*model.PasswordReset) *model.PasswordReset); ok {
		r0 = rf(passwordReset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PasswordReset)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.PasswordReset) error); ok {
		r1 = rf(passwordReset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPassowrdResetRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewPassowrdResetRepository creates a new instance of PassowrdResetRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPassowrdResetRepository(t mockConstructorTestingTNewPassowrdResetRepository) *PassowrdResetRepository {
	mock := &PassowrdResetRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}