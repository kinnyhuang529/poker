package poker

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	fiveHands = [][]string{
		{"cT", "cJ", "cQ", "cK", "cA"},
		{"c3", "c4", "c5", "c6", "c7"},
		{"dT", "sT", "hT", "cT", "cA"},
		{"cT", "dT", "hT", "c3", "d3"},
		{"h3", "hA", "h8", "hT", "hK"},
		{"c3", "d4", "h5", "s6", "c7"},
		{"c3", "d3", "s3", "cK", "cA"},
		{"c5", "d5", "c3", "d3", "cA"},
		{"c5", "d5", "cT", "d3", "cA"},
		{"c3", "c8", "d2", "sT", "hA"},
	}
	fiveCardTypes     = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fiveCardStrengths = []int32{1, 8, 59, 225, 414, 1607, 2336, 3271, 5338, 6589}
)
var (
	sixHands = [][]string{
		{"cT", "cJ", "cQ", "d3", "cK", "cA"},
		{"c3", "c4", "hK", "c5", "c6", "c7"},
		{"dT", "sT", "hT", "d2", "cT", "cA"},
		{"d2", "cT", "s3", "hT", "c3", "d3"},
		{"h3", "h5", "h8", "s8", "hT", "hK"},
		{"c3", "d4", "h5", "s6", "c7", "h8"},
		{"c3", "d9", "d3", "s3", "cK", "cA"},
		{"c5", "d5", "c3", "cK", "d3", "cA"},
		{"c5", "d5", "c8", "d6", "d3", "cA"},
		{"c3", "c8", "cK", "dJ", "sT", "hA"},
	}
	testSixHands = [][]string{
		{"cT", "cJ", "cQ", "cK", "cA"},
		{"c3", "c4", "c5", "c6", "c7"},
		{"dT", "sT", "hT", "cT", "cA"},
		{"cT", "s3", "hT", "c3", "d3"},
		{"h3", "h5", "h8", "hT", "hK"},
		{"d4", "h5", "s6", "c7", "h8"},
		{"c3", "d3", "s3", "cK", "cA"},
		{"c5", "d5", "c3", "d3", "cA"},
		{"c5", "d5", "c8", "d6", "cA"},
		{"c8", "cK", "dJ", "sT", "hA"},
	}
	sixCardTypes     = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sixCardStrengths = []int32{1, 8, 59, 303, 1050, 1606, 2336, 3271, 5347, 6231}
)

var (
	sevenHands = [][]string{
		{"cT", "cJ", "c2", "cQ", "d3", "cK", "cA"},
		{"c3", "c4", "cT", "hK", "c5", "c6", "c7"},
		{"dT", "sT", "hT", "s3", "d2", "cT", "cA"},
		{"d2", "cT", "dT", "c8", "hT", "c3", "d3"},
		{"h3", "d5", "hA", "h8", "s8", "hT", "hK"},
		{"c3", "d4", "dK", "h5", "s6", "c7", "h8"},
		{"c3", "d9", "h7", "d3", "s3", "cK", "cA"},
		{"c5", "d5", "d8", "c3", "cK", "d3", "cA"},
		{"c5", "d5", "cT", "c8", "d6", "d3", "cA"},
		{"c3", "c8", "d2", "cK", "dJ", "sT", "hA"},
	}
	testSevenHands = [][]string{
		{"cT", "cJ", "cQ", "cK", "cA"},
		{"c3", "c4", "c5", "c6", "c7"},
		{"dT", "sT", "hT", "cT", "cA"},
		{"cT", "dT", "hT", "c3", "d3"},
		{"h3", "hA", "h8", "hT", "hK"},
		{"d4", "h5", "s6", "c7", "h8"},
		{"c3", "d3", "s3", "cK", "cA"},
		{"c5", "d5", "c3", "d3", "cA"},
		{"c5", "d5", "cT", "c8", "cA"},
		{"c8", "cK", "dJ", "sT", "hA"},
	}
	sevenCardTypes     = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sevenCardStrengths = []int32{1, 8, 59, 225, 414, 1606, 2336, 3271, 5334, 6231}
)

func TestEvaluator(t *testing.T) {
	//five
	for i, fiveHand := range fiveHands {
		hand := make([]Card, 0)
		hand = append(hand, NewCard(fiveHand[0]))
		hand = append(hand, NewCard(fiveHand[1]))
		hand = append(hand, NewCard(fiveHand[2]))
		hand = append(hand, NewCard(fiveHand[3]))
		hand = append(hand, NewCard(fiveHand[4]))

		testHand := make([]Card, 0)
		testHand = append(testHand, NewCard(fiveHand[0]))
		testHand = append(testHand, NewCard(fiveHand[1]))
		testHand = append(testHand, NewCard(fiveHand[2]))
		testHand = append(testHand, NewCard(fiveHand[3]))
		testHand = append(testHand, NewCard(fiveHand[4]))

		cardType, cardStrength, bestHand, err := Evaluator(hand)
		if cardType != fiveCardTypes[i] || cardStrength != fiveCardStrengths[i] || err != nil || !reflect.DeepEqual(bestHand, testHand) {
			t.Error(fmt.Sprintf("TestEvaluator: card type-%v Wrong hand card! ", fiveCardTypes[i]))
			return
		}
	}
	//six
	for i, sixHand := range sixHands {
		hand := make([]Card, 0)
		hand = append(hand, NewCard(sixHand[0]))
		hand = append(hand, NewCard(sixHand[1]))
		hand = append(hand, NewCard(sixHand[2]))
		hand = append(hand, NewCard(sixHand[3]))
		hand = append(hand, NewCard(sixHand[4]))
		hand = append(hand, NewCard(sixHand[5]))

		testHand := make([]Card, 0)
		testHand = append(testHand, NewCard(testSixHands[i][0]))
		testHand = append(testHand, NewCard(testSixHands[i][1]))
		testHand = append(testHand, NewCard(testSixHands[i][2]))
		testHand = append(testHand, NewCard(testSixHands[i][3]))
		testHand = append(testHand, NewCard(testSixHands[i][4]))

		cardType, cardStrength, bestHand, err := Evaluator(hand)
		if cardType != sixCardTypes[i] || cardStrength != sixCardStrengths[i] || err != nil || !reflect.DeepEqual(bestHand, testHand) {
			t.Error(fmt.Sprintf("TestEvaluator: six card type-%v Wrong hand card! ", sixCardTypes[i]))
			return
		}
	}
	//seven
	for i, sevenHand := range sevenHands {
		hand := make([]Card, 0)
		hand = append(hand, NewCard(sevenHand[0]))
		hand = append(hand, NewCard(sevenHand[1]))
		hand = append(hand, NewCard(sevenHand[2]))
		hand = append(hand, NewCard(sevenHand[3]))
		hand = append(hand, NewCard(sevenHand[4]))
		hand = append(hand, NewCard(sevenHand[5]))
		hand = append(hand, NewCard(sevenHand[6]))

		testHand := make([]Card, 0)
		testHand = append(testHand, NewCard(testSevenHands[i][0]))
		testHand = append(testHand, NewCard(testSevenHands[i][1]))
		testHand = append(testHand, NewCard(testSevenHands[i][2]))
		testHand = append(testHand, NewCard(testSevenHands[i][3]))
		testHand = append(testHand, NewCard(testSevenHands[i][4]))

		cardType, cardStrength, bestHand, err := Evaluator(hand)
		if cardType != sevenCardTypes[i] || cardStrength != sevenCardStrengths[i] || err != nil || !reflect.DeepEqual(bestHand, testHand) {
			t.Error(fmt.Sprintf("TestEvaluator: card type-%v Wrong hand card! ", sevenCardTypes[i]))
			return
		}
	}
	t.Log("TestEvaluator: success! ")
}
