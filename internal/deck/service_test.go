package deck_test

import (
	"testing"

	"github.com/gkatanacio/card-deck-api/internal/deck"
	"github.com/gkatanacio/card-deck-api/internal/errs"
	"github.com/gkatanacio/card-deck-api/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Service_CreateDeck_NoCardCodes(t *testing.T) {
	// given
	repositoryMock := mocks.NewRepository(t)
	repositoryMock.On("CreateDeck", mock.AnythingOfType("*deck.Deck")).Return(nil)

	service := deck.NewService(repositoryMock)

	// when
	created, err := service.CreateDeck(false, nil)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, created)
	assert.NotEqual(t, uuid.Nil, *created.Id)
	assert.False(t, *created.Shuffled)
	assert.Equal(t, 52, *created.Remaining)
}

func Test_Service_CreateDeck_WithCardCodes(t *testing.T) {
	// given
	repositoryMock := mocks.NewRepository(t)
	repositoryMock.On("CreateDeck", mock.AnythingOfType("*deck.Deck")).Return(nil)

	service := deck.NewService(repositoryMock)

	// when
	created, err := service.CreateDeck(false, []string{"AC", "2C", "10H"})

	// then
	assert.NoError(t, err)
	assert.NotNil(t, created)
	assert.NotEqual(t, uuid.Nil, *created.Id)
	assert.False(t, *created.Shuffled)
	assert.Equal(t, 3, *created.Remaining)
}

func Test_Service_CreateDeck_InvalidCardCodes(t *testing.T) {
	// given
	service := deck.NewService(nil)

	// when
	_, err := service.CreateDeck(false, []string{"AA", "BB"})

	// then
	assert.Error(t, err)
	assert.IsType(t, &errs.BadRequest{}, err)
}

func Test_Service_GetDeck_ReturnsDeck(t *testing.T) {
	// given
	id := uuid.New()
	shuffled := true
	cards := deck.CardList{
		{
			Value: "10",
			Suit:  "CLUBS",
			Code:  "10C",
		},
		{
			Value: "QUEEN",
			Suit:  "HEARTS",
			Code:  "QH",
		},
	}

	repositoryMock := mocks.NewRepository(t)
	repositoryMock.On("GetDeckById", id).Return(&deck.Deck{
		Id:       &id,
		Shuffled: &shuffled,
		Cards:    cards,
	}, nil)

	service := deck.NewService(repositoryMock)

	// when
	fetched, err := service.GetDeck(id)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, fetched)
	assert.Equal(t, id, *fetched.Id)
	assert.Equal(t, shuffled, *fetched.Shuffled)
	assert.Equal(t, len(cards), *fetched.Remaining)
	assert.ElementsMatch(t, cards, fetched.Cards)
}

func Test_Service_GetDeck_NotFound(t *testing.T) {
	// given
	id := uuid.New()

	repositoryMock := mocks.NewRepository(t)
	repositoryMock.On("GetDeckById", id).Return(nil, errs.NewNotFound(""))

	service := deck.NewService(repositoryMock)

	// when
	_, err := service.GetDeck(id)

	// then
	assert.Error(t, err)
	assert.IsType(t, &errs.NotFound{}, err)
}

func Test_Service_DrawCards_ReturnsDrawnCards(t *testing.T) {
	// given
	id := uuid.New()
	count := 2
	cards := deck.CardList{
		{
			Value: "10",
			Suit:  "CLUBS",
			Code:  "10C",
		},
		{
			Value: "QUEEN",
			Suit:  "HEARTS",
			Code:  "QH",
		},
		{
			Value: "ACE",
			Suit:  "HEARTS",
			Code:  "AH",
		},
	}

	repositoryMock := mocks.NewRepository(t)
	repositoryMock.On("GetDeckById", id).Return(&deck.Deck{
		Id:    &id,
		Cards: cards,
	}, nil)
	repositoryMock.On("UpdateDeck", mock.AnythingOfType("*deck.Deck")).Return(nil)

	service := deck.NewService(repositoryMock)

	// when
	drawn, err := service.DrawCards(id, count)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, drawn)
	assert.Len(t, drawn.Cards, count)
	assert.ElementsMatch(t, cards[:count], drawn.Cards)
}

func Test_Service_DrawCards_NotEnoughCards(t *testing.T) {
	// given
	id := uuid.New()
	count := 3
	cards := deck.CardList{
		{
			Value: "10",
			Suit:  "CLUBS",
			Code:  "10C",
		},
		{
			Value: "QUEEN",
			Suit:  "HEARTS",
			Code:  "QH",
		},
	}

	repositoryMock := mocks.NewRepository(t)
	repositoryMock.On("GetDeckById", id).Return(&deck.Deck{
		Id:    &id,
		Cards: cards,
	}, nil)

	service := deck.NewService(repositoryMock)

	// when
	_, err := service.DrawCards(id, count)

	// then
	assert.Error(t, err)
	assert.IsType(t, &errs.BadRequest{}, err)
}
