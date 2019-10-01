package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.addCard("SSHHIITTEEPPIILLEEEEEEE")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
