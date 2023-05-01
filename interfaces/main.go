package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {}
type indoBot struct {}

func main() {
	eb := englishBot{}
	ib := indoBot{}

	printGreeting(eb)
	printGreeting(ib)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi there!"
}

func (ib indoBot) getGreeting() string {
	return "Hai disana!"
}