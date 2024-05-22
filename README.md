# poker

*[English](README.md) ∙ [繁體中文](README_zh-tw.md) ∙ [简体中文](README_zh-cn.md)*

## Introduction
`poker` is a poker card counter that uses Cactus Kev's Poker Hand Evaluator algorithm. It can evaluate the poker card combination in the player's hand and give the corresponding card strength and card type.

## Prerequisites
Make sure you have installed the Go environment.

## Installation
```bash
go get github.com/kinnyhuang529/poker
```

## Configuration
Import this package in Go code:
```bash
import "github.com/kinnyhuang529/poker"
```

## Example
```go
package main

import (
	"fmt"
	"github.com/kinnyhuang529/poker"
)

func main() {
	//Build deck
	deck := poker.NewDeck()
	
	//Shuffle deck
	deck.Shuffle()
	
	//five
	five, err := deck.Draw(5)
	//five: [s3 d2 sK c8 s9]
	
	fiveCardType, fiveCardStrength, fiveBestHand, err := poker.Evaluator(five)
	//fiveCardType: 10, fiveCardStrength: 6952, fiveBestHand: [s3 d2 sK c8 s9]
	
	//six
	six, err := deck.Draw(6)
	//six: [c7 c5 dK h6 cA hK]
	
	sixCardType, sixCardStrength, sixBestHand, err := poker.Evaluator(six)
	//sixCardType: 9, sixCardStrength: 3586, sixBestHand: [c7 dK h6 cA hK]
	
	//seven
	seven, err := deck.Draw(7)
	//seven: [c6 cK h8 dQ s8 cT hT]
	sevenCardType, sevenCardStrength, sevenBestHand, err := poker.Evaluator(seven)
	//sevenCardType: 8, sevenCardStrength: 2942, sevenBestHand: [cK h8 s8 cT hT]
}
```
### Explanation of cardType

- 0: Exceptional Situation
- 1: Royal Flush
- 2: Straight Flush
- 3: Four of a Kind
- 4: Full House
- 5: Flush
- 6: Straight
- 7: Three of a Kind
- 8: Two Pairs
- 9: One Pair
- 10: High Card

### Explanation of cardStrength
The range is from 1 to 7462, the smaller the number, the stronger the card.

## Links
 - [Cactus Kev's Poker Hand Evaluator](https://suffe.cool/poker/evaluator.html)