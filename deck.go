package main

import (
	"fmt"
	"strings"
)

// create a new type of deck, which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string {"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string {"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, suit + " of " + val)
		}
	}

	return cards
}

func (d deck) printDeck() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func deal (d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}