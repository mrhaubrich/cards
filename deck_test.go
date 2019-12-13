package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck lenght of 52, but got %v", len(d))
	}
	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first card Ace of Clubs, but got %v", d[0])
	}
	if d[len(d)-1] != "King of Spades" {
		t.Errorf("Expected first card King of Spades, but got %v", len(d)-1)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()

	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck lenght of 52, but got %v", len(d))
	}
	if loadedDeck[0] != "Ace of Clubs" {
		t.Errorf("Expected first card Ace of Clubs, but got %v", d[0])
	}
	if loadedDeck[len(d)-1] != "King of Spades" {
		t.Errorf("Expected first card King of Spades, but got %v", len(d)-1)
	}

	os.Remove("_decktesting")

}

func TestDeal(t *testing.T) {
	//func deal(d deck, handSize int) (deck, deck) {
	//	return d[:handSize], d[handSize:]
	//}
	d := newDeck()
	dealed, remaining := deal(d, 5)

	if len(dealed) != 5 {
		t.Errorf("Expected size of the dealed deck is 5, but got %v", len(dealed))
	}
	if len(remaining) != 47 {
		t.Errorf("Expected size of the remaining deck is 47, but got %v", len(remaining))
	}
}
