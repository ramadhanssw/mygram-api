// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "mygram-api/domain"

	mock "github.com/stretchr/testify/mock"
)

// CommentRepository is an autogenerated mock type for the CommentRepository type
type CommentRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *CommentRepository) Delete(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: _a0, _a1, _a2
func (_m *CommentRepository) Fetch(_a0 context.Context, _a1 *[]domain.Comment, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *[]domain.Comment, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: _a0, _a1, _a2
func (_m *CommentRepository) GetByID(_a0 context.Context, _a1 *domain.Comment, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Comment, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *CommentRepository) Store(_a0 context.Context, _a1 *domain.Comment) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Comment) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0, _a1, _a2
func (_m *CommentRepository) Update(_a0 context.Context, _a1 domain.Comment, _a2 string) (domain.Photo, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 domain.Photo
	if rf, ok := ret.Get(0).(func(context.Context, domain.Comment, string) domain.Photo); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(domain.Photo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Comment, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCommentRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommentRepository creates a new instance of CommentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentRepository(t mockConstructorTestingTNewCommentRepository) *CommentRepository {
	mock := &CommentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
