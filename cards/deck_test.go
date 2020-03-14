package main

import (
	"testing"
	"os"
)

//This function is automatically called by go test runner
func TestNewDeck(t *testing.T) {
	d := newDeck()

	//test the number of cards in a deck
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))

	}
	//test if the first card in the deck is Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
	
	//test if the last card in the deck is King of Clubs
	if d[len(d) - 1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d) - 1])
	}
}

func TestDecKFileFunctions(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	//test the number of cards in a deck loaded froma file
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
	
}
