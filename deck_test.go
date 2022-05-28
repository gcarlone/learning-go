package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected length of 16, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first element Ace of Spades, but got %v", string(d[0]))
	}
	if d[15] != "Four of Clubs" {
		t.Errorf("Expected last element Four of Clubs, but got %v", d[15])
	}
}
