package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// deck behaves as a slice of string with additional behaviours
type deck []string

func newDeck() deck {
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	cards := deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Receiver function is like a method in OOP
// d is like this or self in OOP languageas
// d (a receiver argument) in go usually is a
// single letter, go avoid terms like this or self
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
	fmt.Println()
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func parseString(s string) deck {
	return deck(strings.Split(s, ","))
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return parseString(string(bs))
}

// Multiple return values and slice splitting
func deal(d deck, handSize int) (deck, deck) {
	// d[startIndexInclusive:upToNotIncluding]
	// d[:upToNotIncluding] starts from 0
	// d[startIndexInclusive:] up to last element
	// d[:] all the slice
	return d[:handSize], d[handSize:]
}
