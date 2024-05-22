package poker

import (
	"testing"
)

func TestNewCard(t *testing.T) {
	cardList := []string{"c5", "dT", "sA", "h2"}
	testCardList := []int32{557831, 16795671, 268442665, 73730}
	for i, c := range cardList {
		if int32(NewCard(c)) != testCardList[i] {
			t.Error("TestNewCard: NewCard error! ")
			return
		}
	}
	t.Log("TestNewCard: success! ")
}
func TestPrime(t *testing.T) {
	cardList := []string{"c5", "dT", "sA", "h2"}
	testPrime := []int32{7, 23, 41, 2}

	for i, c := range cardList {
		if NewCard(c).Prime() != testPrime[i] {
			t.Error("TestPrime: Prime error! ")
			return
		}
	}
	t.Log("TestPrime: success! ")
}
func TestRank(t *testing.T) {
	cardList := []string{"c5", "dT", "sA", "h2"}
	testRank := []int32{3, 8, 12, 0}

	for i, c := range cardList {
		if NewCard(c).Rank() != testRank[i] {
			t.Error("TestRank: Rank error! ")
			return
		}
	}
	t.Log("TestRank: success! ")
}
func TestSuit(t *testing.T) {
	cardList := []string{"c5", "dT", "sA", "h2"}
	testSuit := []int32{8, 4, 1, 2}

	for i, c := range cardList {
		if NewCard(c).Suit() != testSuit[i] {
			t.Error("TestSuit: Suit error! ")
			return
		}
	}
	t.Log("TestSuit: success! ")
}
func TestBitRank(t *testing.T) {
	cardList := []string{"c5", "dT", "sA", "h2"}
	testBitRank := []int32{8, 256, 4096, 1}

	for i, c := range cardList {
		if NewCard(c).BitRank() != testBitRank[i] {
			t.Error("TestBitRank: BitRank error! ")
			return
		}
	}
	t.Log("TestBitRank: success! ")
}

func TestString(t *testing.T) {
	cardList := []string{"c5", "dT", "sA", "h2"}
	for i, c := range cardList {
		if NewCard(c).String() != cardList[i] {
			t.Error("TestString: String error! ")
			return
		}
	}
	t.Log("TestString: success! ")
}
