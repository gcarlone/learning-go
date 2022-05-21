package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["theAnswer"] = 42
	m["secondAnswer"] = 321
	fmt.Println(m)
}
