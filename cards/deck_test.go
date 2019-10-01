package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// create new deck to be tested
	d := newDeck()
	// run our basic deck assertions
	runBasicDeckAssertions(d, t)
}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	// delete any existing '_decktesting' file in directory
	checkAndDeleteFile("./_decktesting")
	// create new deck
	d := newDeck()
	// save to '_decktesting' file
	d.saveToFile("_decktesting")
	// load from file
	sd := newDeckFromFile("_decktesting")
	// run assertions
	runBasicDeckAssertions(sd, t)
	// delete any existing '_decktesting' file in directory
	checkAndDeleteFile("./_decktesting")
}

func runBasicDeckAssertions(d CardDeck, t *testing.T) {
	// make sure deck is created with n cards
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}
	// make sure first card is "Ace of Spades"
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}
	// make sure the last card is "Four of Clubs"
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but got %v", d[len(d)-1])
	}
}

func checkAndDeleteFile(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return
	}
	os.Remove(path)
}
