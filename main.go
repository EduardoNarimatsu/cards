package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	//card := newCard()
	//fmt.Println(newCard())

	// cards := newDeck()

	// hand, remainingCards := deal(cards,5)

	// hand.print()
	// remainingCards.print()

	cards := newDeck()
	fmt.Println(cards.toString())
}
