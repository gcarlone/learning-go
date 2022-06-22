package main

func createAnotherCard() card {
	var c card // a "zero value" card: "" string, 0 int and float, false bool

	c.value = "Ace"
	c.suit = "Flower"

	return c
}
