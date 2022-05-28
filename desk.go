package main

import "fmt"

type desk []string

func (d desk) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
