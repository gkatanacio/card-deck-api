package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gkatanacio/card-deck-api/internal/deck"
	"github.com/gkatanacio/card-deck-api/internal/errs"
	"github.com/gkatanacio/card-deck-api/internal/handler"
	"github.com/gkatanacio/card-deck-api/mocks"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CreateDeckHandler_Handle_NoQueryParams(t *testing.T) {
	// given
	deckServiceMock := mocks.NewServicer(t)
	deckServiceMock.On("CreateDeck", false, mock.AnythingOfType("[]string")).Return(&deck.Deck{}, nil)

	handler := handler.NewCreateDeckHandler(deckServiceMock)

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/decks", nil)

	// when
	handler.Handle(rr, req)

	// then
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func Test_CreateDeckHandler_Handle_WithQueryParams(t *testing.T) {
	// given
	deckServiceMock := mocks.NewServicer(t)
	deckServiceMock.On("CreateDeck", true, []string{"AC", "2C", "KH"}).Return(&deck.Deck{}, nil)

	handler := handler.NewCreateDeckHandler(deckServiceMock)

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/decks?shuffle=true&cards=AC,2C,KH", nil)

	// when
	handler.Handle(rr, req)

	// then
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func Test_CreateDeckHandler_Handle_InvalidShuffle(t *testing.T) {
	// given
	handler := handler.NewCreateDeckHandler(nil)

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/decks?shuffle=zzz", nil)

	// when
	handler.Handle(rr, req)

	// then
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func Test_GetDeckHandler_Handle_ReturnsDeck(t *testing.T) {
	// given
	id := uuid.New()

	deckServiceMock := mocks.NewServicer(t)
	deckServiceMock.On("GetDeck", id).Return(&deck.Deck{}, nil)

	handler := handler.NewGetDeckHandler(deckServiceMock)

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/decks/%s", id.String()), nil)
	req = mux.SetURLVars(req, map[string]string{"id": id.String()})

	// when
	handler.Handle(rr, req)

	// then
	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_GetDeckHandler_Handle_NotFound(t *testing.T) {
	// given
	id := uuid.New()

	deckServiceMock := mocks.NewServicer(t)
	deckServiceMock.On("GetDeck", id).Return(nil, errs.NewNotFound(""))

	handler := handler.NewGetDeckHandler(deckServiceMock)

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/decks/%s", id.String()), nil)
	req = mux.SetURLVars(req, map[string]string{"id": id.String()})

	// when
	handler.Handle(rr, req)

	// then
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func Test_GetDeckHandler_Handle_InvalidId(t *testing.T) {
	// given
	id := "xxx-yyy-zzz"

	handler := handler.NewGetDeckHandler(nil)

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/decks/%s", id), nil)
	req = mux.SetURLVars(req, map[string]string{"id": id})

	// when
	handler.Handle(rr, req)

	// then
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func Test_DeleteCardsHandler_Handle_NoCount(t *testing.T) {
	// given
	id := uuid.New()

	deckServiceMock := mocks.NewServicer(t)
	deckServiceMock.On("DrawCards", id, 1).Return(&deck.Deck{}, nil)

	handler := handler.NewDeleteCardsHandler(deckServiceMock)

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/decks/%s/cards", id.String()), nil)
	req = mux.SetURLVars(req, map[string]string{"id": id.String()})

	// when
	handler.Handle(rr, req)

	// then
	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_DeleteCardsHandler_Handle_WithCount(t *testing.T) {
	// given
	id := uuid.New()
	count := 5

	deckServiceMock := mocks.NewServicer(t)
	deckServiceMock.On("DrawCards", id, count).Return(&deck.Deck{}, nil)

	handler := handler.NewDeleteCardsHandler(deckServiceMock)

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/decks/%s/cards?count=%d", id.String(), count), nil)
	req = mux.SetURLVars(req, map[string]string{"id": id.String()})

	// when
	handler.Handle(rr, req)

	// then
	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_DeleteCardsHandler_Handle_InvalidId(t *testing.T) {
	// given
	id := "xxx-yyy-zzz"

	handler := handler.NewDeleteCardsHandler(nil)

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/decks/%s/cards", id), nil)
	req = mux.SetURLVars(req, map[string]string{"id": id})

	// when
	handler.Handle(rr, req)

	// then
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
