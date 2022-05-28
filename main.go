package main

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

	anotherCard := newCard() // interred from function return type

	// every eleement in a slice or in an array have to be of the same type
	// cards := []string{card, anotherCard, createAnotherCard()}
	cards := deck{card, anotherCard, createAnotherCard()}

	// append do not modify existent slice, return a new slice
	cards = append(cards, "Six of Spades", "Four of Spades")
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

func newCard() string {
	return "Three of Diamond"
}
