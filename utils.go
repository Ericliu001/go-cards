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


func (d deck) print() {
	for index, value := range d {
		fmt.Print(index, ".", value, "; ")
	}
}