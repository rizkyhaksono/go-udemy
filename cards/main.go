package main

func main() {
	// cards := newDeckFromFile("my_card.txt")
	// cards.printDeck()

	cards := newDeck()
	cards.shuffle()
	cards.printDeck()
}
