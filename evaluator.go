package poker

import (
	"errors"
)

func Evaluator(cards []Card) (cardType int32, cardStrength int32, bestHand []Card, err error) {
	switch len(cards) {
	case 5:
		return five(cards)
	case 6:
		return six(cards)
	case 7:
		return seven(cards)
	default:
		err = errors.New("the number of cards must be 5, 6, 7! ")
		return
	}

}

func five(cards []Card) (cardType int32, cardStrength int32, bestHand []Card, err error) {
	var q = (cards[0] | cards[1] | cards[2] | cards[3] | cards[4]) >> 16
	var s int32
	if cards[0]&cards[1]&cards[2]&cards[3]&cards[4]&0xF000 != 0 {
		s = flushes[q]
	} else {
		s = unique5[q]
		if s == 0 {
			q = (cards[0] & 0xFF) * (cards[1] & 0xFF) * (cards[2] & 0xFF) * (cards[3] & 0xFF) * (cards[4] & 0xFF)
			s = hashValues[findFast(uint32(q))]
		}
	}
	if s < 1 || s > 7462 {
		err = errors.New("error! ")
		return
	}
	cardType = strengthToType(s)
	cardStrength = s
	bestHand = cards
	return
}

func six(cards []Card) (cardType int32, cardStrength int32, bestHand []Card, err error) {
	var tmpCards []Card
	tmpCards = make([]Card, 5)

	var bestCardType int32 = highCardType
	var bestCardStrength int32 = maxHighCardStrength

	var best []Card
	best = make([]Card, 5)
	for _, item := range t6c5 {
		tmpCards[0] = cards[item[0]]
		tmpCards[1] = cards[item[1]]
		tmpCards[2] = cards[item[2]]
		tmpCards[3] = cards[item[3]]
		tmpCards[4] = cards[item[4]]

		var tmpCardType int32
		var tmpCardStrength int32
		tmpCardType, tmpCardStrength, _, err = five(tmpCards)
		if err != nil {
			return 0, 0, []Card{}, err
		}
		if tmpCardStrength < bestCardStrength {
			copy(best, tmpCards)
			bestCardStrength = tmpCardStrength
			bestCardType = tmpCardType
		}
	}
	cardType = bestCardType
	cardStrength = bestCardStrength
	bestHand = best
	return
}

func seven(cards []Card) (cardType int32, cardStrength int32, bestHand []Card, err error) {
	var tmpCards []Card
	tmpCards = make([]Card, 5)

	var bestCardType int32 = highCardType
	var bestCardStrength int32 = maxHighCardStrength

	var best []Card
	best = make([]Card, 5)
	for _, item := range t7c5 {
		tmpCards[0] = cards[item[0]]
		tmpCards[1] = cards[item[1]]
		tmpCards[2] = cards[item[2]]
		tmpCards[3] = cards[item[3]]
		tmpCards[4] = cards[item[4]]

		var tmpCardType int32
		var tmpCardStrength int32
		tmpCardType, tmpCardStrength, _, err = five(tmpCards)
		if err != nil {
			return 0, 0, []Card{}, err
		}
		if tmpCardStrength < bestCardStrength {
			copy(best, tmpCards)
			bestCardStrength = tmpCardStrength
			bestCardType = tmpCardType
		}
	}
	cardType = bestCardType
	cardStrength = bestCardStrength
	bestHand = best
	return
}

func findFast(u uint32) uint32 {
	u += 0xE91AAA35
	u ^= u >> 16
	u += u << 8
	u ^= u >> 4
	b := (u >> 8) & 0x1FF
	a := (u + (u << 2)) >> 19
	return a ^ hashAdjust[b]
}

var t6c5 = [6][6]uint8{
	{0, 1, 2, 3, 4, 5},
	{0, 1, 2, 3, 5, 4},
	{0, 1, 2, 4, 5, 3},
	{0, 1, 3, 4, 5, 2},
	{0, 2, 3, 4, 5, 1},
	{1, 2, 3, 4, 5, 0},
}

var t7c5 = [21][7]uint8{
	{0, 1, 2, 3, 4, 5, 6},
	{0, 1, 2, 3, 5, 4, 6},
	{0, 1, 2, 3, 6, 4, 5},
	{0, 1, 2, 4, 5, 3, 6},
	{0, 1, 2, 4, 6, 3, 5},
	{0, 1, 2, 5, 6, 3, 4},
	{0, 1, 3, 4, 5, 2, 6},
	{0, 1, 3, 4, 6, 2, 5},
	{0, 1, 3, 5, 6, 2, 4},
	{0, 1, 4, 5, 6, 2, 3},
	{0, 2, 3, 4, 5, 1, 6},
	{0, 2, 3, 4, 6, 1, 5},
	{0, 2, 3, 5, 6, 1, 4},
	{0, 2, 4, 5, 6, 1, 3},
	{0, 3, 4, 5, 6, 1, 2},
	{1, 2, 3, 4, 5, 0, 6},
	{1, 2, 3, 4, 6, 0, 5},
	{1, 2, 3, 5, 6, 0, 4},
	{1, 2, 4, 5, 6, 0, 3},
	{1, 3, 4, 5, 6, 0, 2},
	{2, 3, 4, 5, 6, 0, 1},
}

func strengthToType(strength int32) int32 {
	switch {
	case strength == maxRoyalFlushStrength:
		return royalFlushType
	case strength <= maxStraightFlushStrength:
		return straightFlushType
	case strength <= maxFourOfAKindStrength:
		return fourOfAKindType
	case strength <= maxFullHouseStrength:
		return fullHouseType
	case strength <= maxFlushStrength:
		return flushType
	case strength <= maxStraightStrength:
		return straightType
	case strength <= maxThreeOfAKindStrength:
		return threeOfAKindType
	case strength <= maxTwoPairStrength:
		return twoPairType
	case strength <= maxOnePairStrength:
		return onePairType
	case strength <= maxHighCardStrength:
		return highCardType
	default:
		return 0
	}
}

const (
	maxRoyalFlushStrength    = 1
	maxStraightFlushStrength = 10
	maxFourOfAKindStrength   = 166
	maxFullHouseStrength     = 322
	maxFlushStrength         = 1599
	maxStraightStrength      = 1609
	maxThreeOfAKindStrength  = 2467
	maxTwoPairStrength       = 3325
	maxOnePairStrength       = 6185
	maxHighCardStrength      = 7462
)

const (
	royalFlushType = iota + 1
	straightFlushType
	fourOfAKindType
	fullHouseType
	flushType
	straightType
	threeOfAKindType
	twoPairType
	onePairType
	highCardType
)
