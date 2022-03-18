package main

import (
	"os"
	"testing"
)

// file naming convention:
// to test deck.go
// name test file deck_test.go

// function naming convention:
// to test newDeck()
// name test function TestNewDeck()
// note how it starts with capital letter
// go runner will automatically call this t argument which is the test handler
func TestNewDeck(t *testing.T) {
	// note how there are three tests here but
	// go won't be able to differentiate that there are three diff tests
	// just one test

	d := newDeck()

	// test procedure:
	// 1. create a new deck
	// 2. write if stmt to see if the deck has right no of cards
	// 3. if not, tell test handler that smth is wrong

	if len(d) != 16 {
		// notify test handler that something went wrong
		// using formatted string here by using %v
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// test procedure:
	// 1. check first and last card

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spafes, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}

}

// function naming convention to test 2 functions at the same time
// to test saveToFile() and newDeckFromFile()
// name test function TestSaveToFileAndNewDeckFromFile
// its long but it makes it easier for the person who modify the original funciton
// to find the test function by command f
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	// test procedure:
	// 1. create a deck
	// 2. save to file
	// 3. file created!
	// 4. attempt to load file
	// what if it crashes here?
	// we will not be able to remove the crashed file and cleanup

	// hence we need to do take care of the cleanup ourselves:
	// 1. delete any files in current working directory with the name "_decktesting"
	// 2. create a deck
	// 3. save to file "_decktesting"
	// 4. load from file
	// 5. assert deck length
	// 6. delete any files in current working directoy with the name "_decktesting"

	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != len(d) {
		t.Errorf("Expected deck length of %v, but got %v", len(d), len(loadedDeck))
	}

	os.Remove("_decktesting")
}
