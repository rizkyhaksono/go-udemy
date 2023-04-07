package main

import "fmt"

func main() {
	// old way
	// var card string = "Joker"

	// new way
	card := deck{
		newCard(),
		newCard(),
		"Joker",
		"Ace of Diamonds",
	}

	card = append(card, "Six of Spades")

	// card = "Ace of Spades"

	// fmt.Println(card)

	// foreach
	for index, card := range card {
		fmt.Println(index, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}