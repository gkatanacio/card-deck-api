// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	deck "github.com/gkatanacio/card-deck-api/internal/deck"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// Servicer is an autogenerated mock type for the Servicer type
type Servicer struct {
	mock.Mock
}

// CreateDeck provides a mock function with given fields: shuffle, cardCodes
func (_m *Servicer) CreateDeck(shuffle bool, cardCodes []string) (*deck.Deck, error) {
	ret := _m.Called(shuffle, cardCodes)

	var r0 *deck.Deck
	if rf, ok := ret.Get(0).(func(bool, []string) *deck.Deck); ok {
		r0 = rf(shuffle, cardCodes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*deck.Deck)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool, []string) error); ok {
		r1 = rf(shuffle, cardCodes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DrawCards provides a mock function with given fields: deckId, count
func (_m *Servicer) DrawCards(deckId uuid.UUID, count int) (*deck.Deck, error) {
	ret := _m.Called(deckId, count)

	var r0 *deck.Deck
	if rf, ok := ret.Get(0).(func(uuid.UUID, int) *deck.Deck); ok {
		r0 = rf(deckId, count)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*deck.Deck)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID, int) error); ok {
		r1 = rf(deckId, count)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeck provides a mock function with given fields: deckId
func (_m *Servicer) GetDeck(deckId uuid.UUID) (*deck.Deck, error) {
	ret := _m.Called(deckId)

	var r0 *deck.Deck
	if rf, ok := ret.Get(0).(func(uuid.UUID) *deck.Deck); ok {
		r0 = rf(deckId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*deck.Deck)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(deckId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewServicerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewServicer creates a new instance of Servicer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServicer(t NewServicerT) *Servicer {
	mock := &Servicer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
