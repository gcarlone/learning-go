package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("If-Else statement\n ")
	theAnswer := 42
	var result string

	if theAnswer < 0 {
		result = "less than 0"
	} else if theAnswer == 0 {
		result = "equal than 0"

	} else {
		result = "greater than 0"
	}
	fmt.Println(result)

	// rispetto a java e C, non servono le parentesi, la graffa di apertura
	// deve necessariamente essere sulla stessa riga della keyword ed è
	// possibile dichiarare ed inizializzare contestualmente variabili
	if x := -42; x < 0 {
		fmt.Println("x less than 0")
	}

	fmt.Println("\nSwitch statement\n ")
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1 // genera un numero casuale tra 1 e 7
	fmt.Println("Day Of Week is ", dow)

	switch sun := 7; dow {
	case 6, sun:
		fmt.Println("It's a weekend ")
		// in go il break è implicito; se volessi emululare il comportamento
		// di java senza il break dovrei decommentare la seguente riga
		// fallthrough
	default:
		fmt.Println("It's a working day")
	}

	fmt.Println("\nLoops\n ")
	colors := []string{"red", "green", "blue"}
	fmt.Println(colors)

	fmt.Println("\nlooping using a standard for")
	for i := 0; i < len(colors); i++ {
		fmt.Println(i, "-", colors[i])
	}

	fmt.Println("\nlooping using an index range ")
	for i := range colors {
		fmt.Println(i, "-", colors[i])
	}

	fmt.Println("\nlooping using an index, value range ")
	for i, color := range colors {
		fmt.Println(i, "-", color)
	}

	// break interrompe l'esecuzione del corrente ciclo
	// continue
	// goto

	value := 1
	for value < 10 {
		fmt.Print(value)
		value++
		if value == 4 {
			continue // riprendo dalla successiva iterazione, non arrivo alla virgola
		} else if value == 10 {
			goto theEnd
		}
		fmt.Print(", ")
	}
	fmt.Println("qui non arrivo più")

theEnd:
	fmt.Println(".")

}
