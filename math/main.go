package main

import (
	"fmt"
	"math"
)

func main() {
	i1, i2, i3 := 23.5, 65.1, 76.3
	floatSum := i1 + i2 + i3
	fmt.Println("value ", floatSum)

	roundedSum := math.Round(floatSum*100) / 100
	fmt.Println("value ", roundedSum)

	raggio := 15
	circonferenza := float64(raggio) * 2 * math.Pi
	fmt.Printf("circonferenza: %.2f\n", circonferenza)
}
