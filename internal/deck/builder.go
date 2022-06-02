package deck

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type suit struct {
	name string
	code rune
}

var suits = map[rune]suit{
	'C': {"CLUBS", 'C'},
	'S': {"SPADES", 'S'},
	'H': {"HEARTS", 'H'},
	'D': {"DIAMONDS", 'D'},
}

var orderedSuits = []suit{
	suits['S'],
	suits['D'],
	suits['C'],
	suits['H'],
}

type value struct {
	name string
	code string
}

var values = map[string]value{
	"A":  {"ACE", "A"},
	"2":  {"2", "2"},
	"3":  {"3", "3"},
	"4":  {"4", "4"},
	"5":  {"5", "5"},
	"6":  {"6", "6"},
	"7":  {"7", "7"},
	"8":  {"8", "8"},
	"9":  {"9", "9"},
	"10": {"10", "10"},
	"J":  {"JACK", "J"},
	"Q":  {"QUEEN", "Q"},
	"K":  {"KING", "K"},
}

var orderedValues = []value{
	values["A"],
	values["2"],
	values["3"],
	values["4"],
	values["5"],
	values["6"],
	values["7"],
	values["8"],
	values["9"],
	values["10"],
	values["J"],
	values["Q"],
	values["K"],
}

func buildCard(v value, s suit) Card {
	return Card{
		Value: v.name,
		Suit:  s.name,
		Code:  fmt.Sprintf("%s%c", v.code, s.code),
	}
}

func buildCards(codes []string) (CardList, error) {
	var cards CardList

	for _, c := range codes {
		lastIndex := len(c) - 1

		valueCode := c[:lastIndex]

		v, exists := values[valueCode]
		if !exists {
			return nil, errors.New(fmt.Sprintf("invalid value for card code: %s", c))
		}

		suitCode := rune(c[lastIndex])

		s, exists := suits[suitCode]
		if !exists {
			return nil, errors.New(fmt.Sprintf("invalid value for card code: %s", c))
		}

		cards = append(cards, buildCard(v, s))
	}

	return cards, nil
}

func standardDeckCards() CardList {
	var cards CardList

	for _, s := range orderedSuits {
		for _, v := range orderedValues {
			cards = append(cards, buildCard(v, s))
		}
	}

	return cards
}

func shuffleCards(cards CardList) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}