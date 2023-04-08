// package  main for read all files that contain in this package
package main

// import package
import "fmt"

// driver method
func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.printDeck()
	// remainingCards.printDeck()
	cards := newDeck()
	fmt.Println(cards.toString())
}
