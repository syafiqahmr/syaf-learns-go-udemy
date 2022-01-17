package main

import "fmt"

// cannot declare outside function body
// deckSize := 20

// can initialize a variable outside of a function
// we just can't assign a value to it.
var deckSize1 int

func main() {
	// Go - Static type language
	// var - create new variable
	// card - name of var
	// string - only a string will ever be assigned to this var
	// common types - bool, string, int, float64
	var card string = "Ace of Spades"
	fmt.Println(card)

	// same as above
	// can omit var
	// can omit string as it can be inferred
	// only use := for initialisation
	card1 := "Ace of Spades"
	card1 = "Five of Diamonds"
	fmt.Println(card1)

	// intialise first then declare
	var deckSize int
	deckSize = 52
	fmt.Println(deckSize)

	deckSize1 = 52
	fmt.Println(deckSize1)

	card2 := newCard()
	fmt.Println(card2)

	// can do this without importing state.go file
	// to run - go run main.go state.go
	printState()

	// array - fixed length list of things
	// slice - an array that can grow/shrink
	// every element in array and slice must be the same data type

	// initialising a slice
	cards := []string{newCard(), "Ace of Diamonds"}
	// this returns a new slice and assign back to cards
	cards = append(cards, "Six of Spades")
	// iterating
	// i - index
	// card - current item we're iterating over
	// range cards - take the slice 'cards' and loop over it
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// string here specifies what data type the function will return
func newCard() string {
	return "Three of Diamonds"
}
