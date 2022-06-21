package deck

import (
	"github.com/gkatanacio/card-deck-api/internal/errs"
	"github.com/google/uuid"
)

// Servicer is the interface for the service layer containing Deck logic.
type Servicer interface {
	CreateDeck(shuffle bool, cardCodes []string) (*Deck, error)
	GetDeck(deckId uuid.UUID) (*Deck, error)
	DrawCards(deckId uuid.UUID, count int) (*Deck, error)
}

// Service is a concrete implementation of Servicer.
type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

// CreateDeck creates a Deck record in the repository then returns a reference to the created Deck.
// A shuffle parameter controls if the Deck to be created should be shuffled or not.
// An optional cardCodes parameter can be passed to create a Deck only with specific Cards.
// If no cardCodes are provided, a standard 52-card Deck is created.
func (s *Service) CreateDeck(shuffle bool, cardCodes []string) (*Deck, error) {
	var cards CardList
	var err error

	if len(cardCodes) > 0 {
		cards, err = buildCards(cardCodes)
		if err != nil {
			return nil, errs.NewBadRequest(err.Error())
		}
	} else {
		cards = standardDeckCards()
	}

	if shuffle {
		cards.Shuffle()
	}

	id := uuid.New()

	deck := &Deck{
		Id:       &id,
		Shuffled: &shuffle,
		Cards:    cards,
	}

	if err := s.repo.CreateDeck(deck); err != nil {
		return nil, err
	}

	remaining := len(deck.Cards)

	return &Deck{
		Id:        deck.Id,
		Shuffled:  deck.Shuffled,
		Remaining: &remaining,
	}, nil
}

// GetDeck returns complete information about a particular Deck given the deckId.
func (s *Service) GetDeck(deckId uuid.UUID) (*Deck, error) {
	deck, err := s.repo.GetDeckById(deckId)
	if err != nil {
		return nil, err
	}

	remaining := len(deck.Cards)

	return &Deck{
		Id:        deck.Id,
		Shuffled:  deck.Shuffled,
		Remaining: &remaining,
		Cards:     deck.Cards,
	}, nil
}

// DrawCards removes and returns a number of Cards from a particular Deck.
// Cards will be removed in the order by which they were specified when the Deck was created.
// The number of Cards to be drawn can be specified using the count parameter.
func (s *Service) DrawCards(deckId uuid.UUID, count int) (*Deck, error) {
	deck, err := s.repo.GetDeckById(deckId)
	if err != nil {
		return nil, err
	}

	if count > len(deck.Cards) {
		return nil, errs.NewBadRequest("not enough cards")
	}

	drawn := deck.Cards[:count]

	deck.Cards = deck.Cards[count:]
	if err := s.repo.UpdateDeck(deck); err != nil {
		return nil, err
	}

	return &Deck{
		Cards: drawn,
	}, nil
}
