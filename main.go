package main

func main() {
	/* GO is a statically typed language
	string keyword può essere omessa perché c'è
	assegnazione diretta (inferred)
	*/
	// var card string = "Ace of Spades" // long form
	// var card = "Ace of Spades"
	// var card string
	c := card{"Spades", "Ace"} // := is used to infer type

	// better approach is to specify which property you are referring to
	// that approach is more secure on refactoring
	c = card{suit: "Diamond", value: "Five"}

	anotherCard := newCard("Three of Clubs") // inferred from function return type

	// every eleement in a slice or in an array have to be of the same type
	// cards := []string{card, anotherCard, createAnotherCard()}
	cards := deck{c, anotherCard, createAnotherCard()}

	// append do not modify existent slice, return a new slice
	cards = append(cards, newCard("Six of Spades"), card{value: "Four", suit: "Spades"})
	cards.print()

	cards = newDeck()
	//fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	cards = newDeckFromFile("my_cards")

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	cards.shuffle()
	cards.print()
}
