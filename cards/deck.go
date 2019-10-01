package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
type CardDeck []string

func newDeck() CardDeck {
	cards := CardDeck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (deck CardDeck) print() {
	for i, card := range deck {
		fmt.Println(i, card)
	}
}

func (deck CardDeck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range deck {
		np := r.Intn(len(deck) - 1)
		deck[i], deck[np] = deck[np], deck[i]
	}
}

func deal(deck CardDeck, handSize int) (CardDeck, CardDeck) {
	return deck[:handSize], deck[handSize:]
}

func (deck *CardDeck) addCard(newCard string) {
	*deck = append(*deck, newCard)
}

func (deck CardDeck) toString() string {
	return strings.Join(deck, ",")
}

func (deck CardDeck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(deck.toString()), 0666)
}

func newDeckFromFile(filename string) CardDeck {
	deckBuffer, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	}
	return CardDeck(strings.Split(string(deckBuffer), ","))
}
