package poker

import (
	"errors"
	"math/rand"
)

var initDecks *Deck

type Deck struct {
	cards []Card
}

func init() {
	var cards []Card
	for _, suit := range suitKey {
		for _, rank := range rankKey {
			cards = append(cards, NewCard(string(suit)+string(rank)))
		}
	}
	initDecks = &Deck{cards: cards}
}

// NewDeck Generate a deck
func NewDeck() *Deck {
	deck := &Deck{}
	deck.cards = make([]Card, len(initDecks.cards))
	copy(deck.cards, initDecks.cards)

	return deck
}

// Shuffle Shuffle deck
func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

// Draw Draw from deck
func (d *Deck) Draw(count int) ([]Card, error) {
	if count > len(d.cards) {
		return nil, errors.New("not enough cards! ")
	}
	cards := make([]Card, count)
	copy(cards, d.cards[:count])
	d.cards = d.cards[count:]
	return cards, nil
}

// Number Number of remaining cards in deck
func (d *Deck) Number() int {
	return len(d.cards)
}
