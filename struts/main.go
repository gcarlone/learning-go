package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)

	poodle.Weight = 9
	fmt.Printf("%+v\n", poodle)
}

// uppercase indica che il tipo è pubblico, il lowercase che è privato
type Dog struct {
	Breed  string
	Weight int
}
