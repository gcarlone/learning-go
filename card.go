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
