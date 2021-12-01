// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	bookmarks "github.com/rizadwiandhika/miniproject-backend-alterra/features/bookmarks"
	mock "github.com/stretchr/testify/mock"
)

// IData is an autogenerated mock type for the IData type
type IData struct {
	mock.Mock
}

// DeleteBookmark provides a mock function with given fields: bookmark
func (_m *IData) DeleteBookmark(bookmark bookmarks.BookmarkCore) error {
	ret := _m.Called(bookmark)

	var r0 error
	if rf, ok := ret.Get(0).(func(bookmarks.BookmarkCore) error); ok {
		r0 = rf(bookmark)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBookmarksByArticleId provides a mock function with given fields: articleID
func (_m *IData) DeleteBookmarksByArticleId(articleID uint) error {
	ret := _m.Called(articleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(articleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBookmarksByUserId provides a mock function with given fields: userID
func (_m *IData) DeleteBookmarksByUserId(userID uint) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertBookmark provides a mock function with given fields: bookmark
func (_m *IData) InsertBookmark(bookmark bookmarks.BookmarkCore) error {
	ret := _m.Called(bookmark)

	var r0 error
	if rf, ok := ret.Get(0).(func(bookmarks.BookmarkCore) error); ok {
		r0 = rf(bookmark)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectDetailUserBookmark provides a mock function with given fields: userID, articleID
func (_m *IData) SelectDetailUserBookmark(userID uint, articleID uint) (bookmarks.BookmarkCore, error) {
	ret := _m.Called(userID, articleID)

	var r0 bookmarks.BookmarkCore
	if rf, ok := ret.Get(0).(func(uint, uint) bookmarks.BookmarkCore); ok {
		r0 = rf(userID, articleID)
	} else {
		r0 = ret.Get(0).(bookmarks.BookmarkCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(userID, articleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectUserBookmarks provides a mock function with given fields: userID
func (_m *IData) SelectUserBookmarks(userID uint) ([]bookmarks.BookmarkCore, error) {
	ret := _m.Called(userID)

	var r0 []bookmarks.BookmarkCore
	if rf, ok := ret.Get(0).(func(uint) []bookmarks.BookmarkCore); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bookmarks.BookmarkCore)
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
