package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// create a new type of deck, which is a slice of string
type deck []string

func newDeck() deck {
	// this cards same as var cards = deck{}
	cards := deck{}

	// slice of string
	cardSuits := []string {"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string {"Ace", "Two", "Three", "Four"}

	// _ means that no index needed
	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, suit + " of " + val)
		}
	}

	return cards
}

func (d deck) printDeck() {
	// (d deck) is receiver declaration for create a instance method for a type of data
	// the object type of data can call that method
	// this receiver declaration is like recursive call, but in different way
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func deal (d deck, handSize int) (deck, deck) {
	// d[:handSize] === [1,2,3]
	// d[handSize:] === [3] 
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// strings package to jpin all array to string within coma seperate
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// call ioutil package
	// and then, convert slice of byte to to.String method
	// 0666 is only read for all user, group and other
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}