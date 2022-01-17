package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// (d deck) - receiver
// d - the instance of the deck (by convention its a 1/2 letter abbreviation)
// deck - type we want to attach the function to
// any variable of type deck has access to this print function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
