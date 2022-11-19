// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/markwallsgrove/muzz_devops/src/models/domain"

	mock "github.com/stretchr/testify/mock"
)

// Database is an autogenerated mock type for the Database type
type Database struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Database) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *Database) CreateUser(ctx context.Context, user *domain.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindMatches provides a mock function with given fields: ctx, user, gender, minAge, maxAge
func (_m *Database) FindMatches(ctx context.Context, user *domain.User, gender []domain.Gender, minAge int, maxAge int) ([]domain.UserProfile, error) {
	ret := _m.Called(ctx, user, gender, minAge, maxAge)

	var r0 []domain.UserProfile
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User, []domain.Gender, int, int) []domain.UserProfile); ok {
		r0 = rf(ctx, user, gender, minAge, maxAge)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.UserProfile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.User, []domain.Gender, int, int) error); ok {
		r1 = rf(ctx, user, gender, minAge, maxAge)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSwipe provides a mock function with given fields: ctx, firstUserId, secondUserId
func (_m *Database) GetSwipe(ctx context.Context, firstUserId int, secondUserId int) (domain.Swipe, error) {
	ret := _m.Called(ctx, firstUserId, secondUserId)

	var r0 domain.Swipe
	if rf, ok := ret.Get(0).(func(context.Context, int, int) domain.Swipe); ok {
		r0 = rf(ctx, firstUserId, secondUserId)
	} else {
		r0 = ret.Get(0).(domain.Swipe)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, firstUserId, secondUserId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx, id
func (_m *Database) GetUser(ctx context.Context, id int) (domain.User, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(context.Context, int) domain.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *Database) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	ret := _m.Called(ctx, email)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Swipe provides a mock function with given fields: ctx, firstUserId, secondUserId
func (_m *Database) Swipe(ctx context.Context, firstUserId int, secondUserId int) (domain.Swipe, error) {
	ret := _m.Called(ctx, firstUserId, secondUserId)

	var r0 domain.Swipe
	if rf, ok := ret.Get(0).(func(context.Context, int, int) domain.Swipe); ok {
		r0 = rf(ctx, firstUserId, secondUserId)
	} else {
		r0 = ret.Get(0).(domain.Swipe)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, firstUserId, secondUserId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDatabase interface {
	mock.TestingT
	Cleanup(func())
}

// NewDatabase creates a new instance of Database. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDatabase(t mockConstructorTestingTNewDatabase) *Database {
	mock := &Database{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
