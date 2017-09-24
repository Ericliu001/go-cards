package main

import "fmt"

type deck []string

func (d deck) print() {
	for index, value := range d {
		fmt.Print(index, ".", value, "; ")
	}
}