package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new custom type of 'deck'
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

// note how this does not need receiver
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// replace index with _ as it is unused
	// unused variables will throw error
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// note how this function can return 2 separate values
// why did we choose to not add a receiver here
// and add deck as a param
// eg. cards.deal(5) - gives the impression that you modify cards to 5 cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	// type conversion
	// deck -> []string -> string -> []byte
	data := []byte(d.toString())
	// 0666 permission allowing everyone can read write the file
	return ioutil.WriteFile(filename, data, 0666)
}

func (d deck) toString() string {
	// convert deck -> [] string -> string
	return strings.Join([]string(d), ",")
}

func newDeckFromFile(filename string) deck {
	// bs - byte slice
	bs, err := ioutil.ReadFile(filename)

	// error handling
	if err != nil {
		// option 1 - log the error and return a call to newDeck()
		// option 2 - log the error and entirely quit the program
		fmt.Println("Error:", err)

		// quit the program
		// based on documentation
		// any non 0 integer passed indicates theres an error
		os.Exit(1)
	}

	// []byte -> string -> []string
	s := strings.Split(string(bs), ",")
	// we can do this conversion because deck type is a []string
	return deck(s)
}

func (d deck) shuffle() {
	// for each index, card in cards
	// 	generate a random number between 0 and len-1
	//	swap the current card and the card at cards[randomNumber]

	// take current time as the seed value for our random generator
	source := rand.NewSource(time.Now().UnixNano()) // creates new seed
	r := rand.New(source)                           // creates rand

	for i := range d {
		// Go uses the same seed every single time
		// so rerunning will result the same order
		// if you do not create new seed above
		// newPosition := rand.Intn(len(d) - 1)

		newPosition := r.Intn(len(d) - 1)

		// one line code to swap
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
