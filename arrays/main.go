package main

import (
	"fmt"
)

// in go un array è un oggetto, viene passata una copia dell'array alla funzione
// tutte le modifiche dell'array non si propagano sull'array originae
func varia(numbers [5]int) {
	numbers[0] = numbers[0] + 5
	fmt.Println("Altered Numbers: ", numbers)
}

func main() {
	var colors [3]string
	colors[0] = "red" //nota solo = e non := (perché l'array è stato già dichiarato e non serve inferration)
	colors[1] = "green"
	colors[2] = "blue"
	fmt.Println("array:", colors)
	fmt.Println("array first element:", colors[0])

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers: ", numbers)
	fmt.Println("Numbers lenght: ", len(numbers))

	// in go un array è un oggetto, il passaggio di un array ad una funzione ne passa una copia
	varia(numbers)
	fmt.Println("Original Numbers: ", numbers)

	// un array non permette la rimozione di elementi e non è facile gestirne l'ordinamento, da qui l'uso degli slices
}
