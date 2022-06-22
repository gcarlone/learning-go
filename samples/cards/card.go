package main

import (
	"fmt"
	"strings"
)

// struct is a data structure: a collection of properties
type card struct {
	suit  string
	value string
}

func (c card) toString() string {
	return fmt.Sprintf("%v of %v", c.value, c.suit)
}

func newCard(s string) card {
	splitted := strings.Split(s, " of ")
	return card{value: splitted[0], suit: splitted[1]}
}

func (c card) printCard() {
	fmt.Printf("%+v\n", c)
}

// receiver is passed by VALUE (go copies all struct data inside
// a new variable)
func (c card) changeCardValueWrong() {
	c.value = "QUESTO CAMBIA C, MA NON LA CARD DA CUI VIENE INVOCATA LA FUNZIONE"
}

// Turn ADDRESS into VALUE with *ADDRESS
// Turn VALUE into ADDRESS with &VALUE
// -------------------------------------
// (pointerToCard *card) is a type descriptor, means we're
//                       working with a pointer to a card
// *pointerToCard is an operator (means we want to manipulate
//                the value the pointer is referencing)
//
// GO VALUE TYPE: int, float, string, bool, arrays, structs
// GO REFERENCE TYPE: slices, maps, channels, pointers, functions
func (pointerToCard *card) changeCardValue() {
	//(*pointerToCard).value = "King"
	pointerToCard.value = "King" // shortcut
}
