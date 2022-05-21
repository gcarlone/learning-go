package main

import (
	"fmt"
	"sort"
)

func main() {
	// without an explicit array size, it's a SLICE
	var colors = []string{"red", "blue", "green"}
	fmt.Println("Colors:", colors)

	colors = append(colors, "purple")

	colors = append(colors[1:])
	fmt.Println("Colors:", colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println("Colors:", colors)

	// a slice of length 5 and capacity 10 (0 int inizialization)
	numbers := make([]int, 5, 10)
	numbers[1] = 1
	numbers[2] = 2
	numbers[3] = 3
	fmt.Println("Numbers:", numbers)

	sort.Ints(numbers)
	fmt.Println("Sorted Numbers:", numbers)
}
