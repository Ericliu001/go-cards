package main

import "fmt"

var deckSize int

func main() {
	cards := newDeck()
	//cards.print()

	first, last := cards.deal(5)

	first.print()
	fmt.Println()
	last.print()
}

func playAround() {
	var message = "Hello, world!"
	fmt.Println(message)
	var eric = "Eric"
	fmt.Println(eric)
	var amIOld = true
	if amIOld {
		fmt.Println("I'm old.")
	}
	deckSize = 20
	fmt.Println(deckSize)
	sayHiTo("Eric")
	placeCards()
}
