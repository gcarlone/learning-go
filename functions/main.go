package main

import (
	"fmt"
)

func main() {
	total, counter := addValues(1, 4, 3, 4, 5, 7)
	fmt.Println("Sommati", counter, "interi la cui somma è", total)
}

func addValues(values ...int) (int, int) {
	var total int
	for _, v := range values {
		total += v
	}
	return total, len(values)
}

type Guitar struct {
	Strings int
	Type    string
}

// go non supporta override
// i parametri non sono passati per reference, ma per valore
// se si vogliono variare i valori, bisogna usare i puntatori
func (g Guitar) Play() {
	fmt.Println("played string A")
}
