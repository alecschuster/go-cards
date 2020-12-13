package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expecting deck length 52 but got %v", len(d))
	}
	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card of Ace of Hearts, got %s", d[0])
	}
	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected last card of King of Diamonds, got %s", d[len(d)-1])
	}
}

func TestSaveToFile(t *testing.T) {
	//remove dirty test files if exists
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expecting deck length 52 but got %v", len(loadedDeck))
	}
	//clean test file
	os.Remove("_decktesting")
}
