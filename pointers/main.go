package main

import (
	"fmt"
)

func main() {
	var anInt = 42
	var p *int = &anInt
	fmt.Println("pointer:", *p)

	*p = *p + 3

	fmt.Println("pointer:", *p)
	fmt.Println("original variable:", anInt)
}
