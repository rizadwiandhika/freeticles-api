// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	users "github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
	mock "github.com/stretchr/testify/mock"
)

// IData is an autogenerated mock type for the IData type
type IData struct {
	mock.Mock
}

// DeleteAllUserFollow provides a mock function with given fields: userID
func (_m *IData) DeleteAllUserFollow(userID uint) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteFollowing provides a mock function with given fields: following
func (_m *IData) DeleteFollowing(following users.FollowerCore) error {
	ret := _m.Called(following)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.FollowerCore) error); ok {
		r0 = rf(following)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: username
func (_m *IData) DeleteUser(username string) error {
	ret := _m.Called(username)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertFollower provides a mock function with given fields: follower
func (_m *IData) InsertFollower(follower users.FollowerCore) error {
	ret := _m.Called(follower)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.FollowerCore) error); ok {
		r0 = rf(follower)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertUser provides a mock function with given fields: user
func (_m *IData) InsertUser(user users.UserCore) (users.UserCore, error) {
	ret := _m.Called(user)

	var r0 users.UserCore
	if rf, ok := ret.Get(0).(func(users.UserCore) users.UserCore); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(users.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(users.UserCore) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectUserByEmail provides a mock function with given fields: email
func (_m *IData) SelectUserByEmail(email string) (users.UserCore, error) {
	ret := _m.Called(email)

	var r0 users.UserCore
	if rf, ok := ret.Get(0).(func(string) users.UserCore); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(users.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectUserById provides a mock function with given fields: id
func (_m *IData) SelectUserById(id uint) (users.UserCore, error) {
	ret := _m.Called(id)

	var r0 users.UserCore
	if rf, ok := ret.Get(0).(func(uint) users.UserCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(users.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectUserByUsername provides a mock function with given fields: username
func (_m *IData) SelectUserByUsername(username string) (users.UserCore, error) {
	ret := _m.Called(username)

	var r0 users.UserCore
	if rf, ok := ret.Get(0).(func(string) users.UserCore); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(users.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectUserFollowers provides a mock function with given fields: userID
func (_m *IData) SelectUserFollowers(userID uint) ([]users.FollowerCore, error) {
	ret := _m.Called(userID)

	var r0 []users.FollowerCore
	if rf, ok := ret.Get(0).(func(uint) []users.FollowerCore); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.FollowerCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectUserFollowings provides a mock function with given fields: userID
func (_m *IData) SelectUserFollowings(userID uint) ([]users.FollowerCore, error) {
	ret := _m.Called(userID)

	var r0 []users.FollowerCore
	if rf, ok := ret.Get(0).(func(uint) []users.FollowerCore); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.FollowerCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectUsers provides a mock function with given fields:
func (_m *IData) SelectUsers() ([]users.UserCore, error) {
	ret := _m.Called()

	var r0 []users.UserCore
	if rf, ok := ret.Get(0).(func() []users.UserCore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.UserCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectUsersByIds provides a mock function with given fields: ids
func (_m *IData) SelectUsersByIds(ids []uint) ([]users.UserCore, error) {
	ret := _m.Called(ids)

	var r0 []users.UserCore
	if rf, ok := ret.Get(0).(func([]uint) []users.UserCore); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.UserCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]uint) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: user
func (_m *IData) UpdateUser(user users.UserCore) (users.UserCore, error) {
	ret := _m.Called(user)

	var r0 users.UserCore
	if rf, ok := ret.Get(0).(func(users.UserCore) users.UserCore); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(users.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(users.UserCore) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
