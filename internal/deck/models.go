package deck

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Deck struct {
	Id         *uuid.UUID `json:"deck_id,omitempty" db:"id"`
	Shuffled   *bool      `json:"shuffled,omitempty" db:"shuffled"`
	Remaining  *int       `json:"remaining,omitempty"`
	Cards      CardList   `json:"cards,omitempty" db:"cards"`
	Timestamps `json:"-"`
}

type CardList []Card

func (cl CardList) Value() (driver.Value, error) {
	return json.Marshal(cl)
}

func (cl *CardList) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &cl)
}

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

type Timestamps struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
