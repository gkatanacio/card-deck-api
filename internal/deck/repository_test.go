//go:build integration

package deck_test

import (
	"testing"
	"time"

	"github.com/gkatanacio/card-deck-api/internal/deck"
	"github.com/gkatanacio/card-deck-api/internal/storage"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var db, _ = storage.NewDb(storage.NewConfig())
var database = deck.NewDatabase(db)

func Test_Database_CreateDeck(t *testing.T) {
	// given
	id := uuid.New()
	shuffled := false
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

	newDeck := &deck.Deck{
		Id:       &id,
		Shuffled: &shuffled,
		Cards:    cards,
	}

	// when
	err := database.CreateDeck(newDeck)

	// then
	assert.NoError(t, err)
}

func Test_Database_GetDeckById(t *testing.T) {
	// given
	id := uuid.New()
	shuffled := false
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

	newDeck := &deck.Deck{
		Id:       &id,
		Shuffled: &shuffled,
		Cards:    cards,
	}

	// when
	err := database.CreateDeck(newDeck)
	assert.NoError(t, err)

	fetched, err := database.GetDeckById(id)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, fetched)
	assert.Equal(t, id, *fetched.Id)
	assert.Equal(t, shuffled, *fetched.Shuffled)
	assert.ElementsMatch(t, cards, fetched.Cards)
	assert.True(t, fetched.CreatedAt.Before(time.Now()))
	assert.Equal(t, fetched.CreatedAt, fetched.UpdatedAt)
}

func Test_Database_UpdateDeck(t *testing.T) {
	// given
	id := uuid.New()
	shuffled := false
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

	newDeck := &deck.Deck{
		Id:       &id,
		Shuffled: &shuffled,
		Cards:    cards,
	}

	// when
	err := database.CreateDeck(newDeck)
	assert.NoError(t, err)

	err = database.UpdateDeck(&deck.Deck{
		Id:       &id,
		Shuffled: &shuffled,
		Cards:    deck.CardList{},
	})
	assert.NoError(t, err)

	updated, err := database.GetDeckById(id)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, updated)
	assert.Equal(t, id, *updated.Id)
	assert.Equal(t, shuffled, *updated.Shuffled)
	assert.Empty(t, updated.Cards)
	assert.True(t, updated.CreatedAt.Before(time.Now()))
	assert.True(t, updated.UpdatedAt.After(updated.CreatedAt))
}
