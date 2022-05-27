package main

import "fmt"

func main() {
	/* GO is a statically typed language
	string keyword può essere omessa perché c'è
	assegnazione diretta (inferred)
	*/
	// var card string = "Ace of Spades" // long form
	// var card = "Ace of Spades"
	// var card string
	card := "Ace of Spades" // := to infer
	card = "Five of Diamond"

	fmt.Println(card)

	anotherCard := newCard() // interred from function return type

	fmt.Println(anotherCard)
	fmt.Println(createAnotherCard())
}

func newCard() string {
	return "Three of Diamond"
}
