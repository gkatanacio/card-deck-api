// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	deck "github.com/gkatanacio/card-deck-api/internal/deck"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateDeck provides a mock function with given fields: _a0
func (_m *Repository) CreateDeck(_a0 *deck.Deck) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*deck.Deck) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDeckById provides a mock function with given fields: id
func (_m *Repository) GetDeckById(id uuid.UUID) (*deck.Deck, error) {
	ret := _m.Called(id)

	var r0 *deck.Deck
	if rf, ok := ret.Get(0).(func(uuid.UUID) *deck.Deck); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*deck.Deck)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDeck provides a mock function with given fields: _a0
func (_m *Repository) UpdateDeck(_a0 *deck.Deck) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*deck.Deck) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewRepositoryT interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t NewRepositoryT) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}