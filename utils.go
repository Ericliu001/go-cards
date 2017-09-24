package main

import "fmt"

func sayHiTo(name string)  {
	fmt.Println("Hi " + name)
}


func placeCards()  {
	cards := deck{"Eric", "Gurveer", "Goodwill"}
	cards = append(cards, "Josh")
	cards.print()
}