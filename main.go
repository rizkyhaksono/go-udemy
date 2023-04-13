// package  main for read all files that contain in this package
package main

// import package

// driver method
func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.printDeck()
	// remainingCards.printDeck()
	
	cards := newDeck()
	cards.saveToFile("my_cards.txt")
}
