package main

import "testing"

func TestCardToString(t *testing.T) {
	c := card{suit: "Diamond", value: "Ace"}
	if c.toString() != "Ace of Diamond" {
		t.Errorf("Expected of Ace of Diamond, but got %v", c)
	}
}
