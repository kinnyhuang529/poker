package main

import (
	"fmt"
	"github.com/kinnyhuang529/poker"
)

func main() {
	deck := poker.NewDeck()
	deck.Shuffle()

	//five
	five, _ := deck.Draw(5)
	fiveCardType, fiveCardStrength, fiveBestHand, _ := poker.Evaluator(five)
	fmt.Println("5 cards hand: ", five)
	fmt.Println("5 cards card type: ", fiveCardType)
	fmt.Println("5 cards card strength: ", fiveCardStrength)
	fmt.Println("5 cards best hand: ", fiveBestHand)

	//six
	six, _ := deck.Draw(6)
	sixCardType, sixCardStrength, sixBestHand, _ := poker.Evaluator(six)
	fmt.Println("6 cards hand: ", six)
	fmt.Println("6 cards card type: ", sixCardType)
	fmt.Println("6 cards card strength: ", sixCardStrength)
	fmt.Println("6 cards best hand: ", sixBestHand)

	//seven
	seven, _ := deck.Draw(7)
	sevenCardType, sevenCardStrength, sevenBestHand, _ := poker.Evaluator(seven)
	fmt.Println("7 cards hand: ", seven)
	fmt.Println("7 cards card type: ", sevenCardType)
	fmt.Println("7 cards card strength: ", sevenCardStrength)
	fmt.Println("7 cards best hand: ", sevenBestHand)
}
