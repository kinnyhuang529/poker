package poker

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	if deck == nil {
		t.Error("TestNewDeck: new deck error! ")
	}
	t.Log("TestNewDeck: success! ")
}

func TestShuffle(t *testing.T) {
	deck := NewDeck()
	testDesk := NewDeck()
	testDesk.Shuffle()

	if deck == testDesk {
		t.Error("TestShuffle: shuffle deck error! ")
		return
	}

	t.Log("TestShuffleDeck: success! ")
}

func TestDraw(t *testing.T) {
	deck := NewDeck()
	hand, err := deck.Draw(5)
	if err != nil {
		t.Error("Error! ")
		return
	}

	if len(hand) != 5 {
		t.Error("TestDraw: Wrong hand card length! ")
		return
	}
	for _, h := range hand {
		for _, c := range deck.cards {
			if h == c {
				t.Error("TestDraw: Wrong hand card! ")
				return
			}
		}
	}
	t.Log("TestDraw: success! ")
}

func TestNumber(t *testing.T) {
	deck := NewDeck()
	if deck.Number() != 52 {
		t.Error("TestNumber: Wrong number! ")
	}
	t.Log("TestNumber: success! ")
}
