package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return (s.sideLength * s.sideLength)
}

func printArea(s shape) {
	fmt.Println("Area =", s.getArea())
}

func main() {
	t := triangle{base: 10, height: 5}
	printArea(t)

	s := square{sideLength: 10}
	printArea(s)
}
