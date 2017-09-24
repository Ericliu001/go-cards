package main

import "fmt"

var deckSize int

func main() {
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
