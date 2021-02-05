package main

func main() {
	cards := newDeck()
	cards.saveDeck("my_deck")
}

func newCard() string {
	return "5 of Diamonds"
}
