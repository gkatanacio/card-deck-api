package deck

import (
	"database/sql"
	"fmt"

	"github.com/gkatanacio/card-deck-api/internal/errs"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Repository is the interface defining the data store-related operations for Decks.
type Repository interface {
	CreateDeck(deck *Deck) error
	GetDeckById(id uuid.UUID) (*Deck, error)
	UpdateDeck(deck *Deck) error
}

// Database is a concrete implementation of Repository.
type Database struct {
	*sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{db}
}

const sqlCreateDeck = `
INSERT INTO decks (id, shuffled, cards) 
VALUES ($1, $2, $3)
`

// CreateDeck creates a new Deck record in the database.
func (db *Database) CreateDeck(deck *Deck) error {
	_, err := db.Exec(sqlCreateDeck, deck.Id, deck.Shuffled, deck.Cards)

	return err
}

const sqlGetDeck = `
SELECT * FROM decks 
WHERE id = $1
`

// GetDeckById retrieves a particular Deck from the database.
// Returns NotFound error if provided id does not exist.
func (db *Database) GetDeckById(id uuid.UUID) (*Deck, error) {
	deck := &Deck{}
	if err := db.Get(deck, sqlGetDeck, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFound(fmt.Sprintf("no deck found with id: %s", id))
		}
		return nil, err
	}

	return deck, nil
}

const sqlUpdateDeck = `
UPDATE decks SET 
(shuffled, cards) = ($2, $3)
WHERE id = $1 
`

// UpdateDeck saves the Deck record in the database.
func (db *Database) UpdateDeck(deck *Deck) error {
	_, err := db.Exec(sqlUpdateDeck, deck.Id, deck.Shuffled, deck.Cards)

	return err
}
