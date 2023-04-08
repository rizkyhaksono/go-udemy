package main

import "fmt"

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.printDeck()
	// remainingCards.printDeck()
	cards := newDeck()
	fmt.Println(cards.toString())
}
