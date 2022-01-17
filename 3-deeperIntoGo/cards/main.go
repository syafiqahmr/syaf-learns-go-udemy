package main

func main() {
	cards := deck{newCard(), "Ace of Diamonds"}
	// this returns a new slice and assign back to cards
	cards = append(cards, "Six of Spades")
	cards.print()
}

// string here specifies what data type the function will return
func newCard() string {
	return "Three of Diamonds"
}
