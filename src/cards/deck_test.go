package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck to be length of 52, intead got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first element to be Ace of Spades, intead got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected first element to be King of Clubs, intead got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	path := "_decktesting"
	os.Remove(path)

	deck := newDeck()
	deck.saveToFile(path)

	loadedDeck := newDeckFromFile(path)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck to be length of 52, intead got %v", len(loadedDeck))
	}

	os.Remove(path)
}
