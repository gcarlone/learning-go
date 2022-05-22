package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	value1 := readInputValue("Value 1")
	value2 := readInputValue("Value 2")
	operator := readOperator()

	calculator := Calculator{value1, value2, operator}
	result := calculator.getResult()
	fmt.Printf("%v %v %v = %v\n", value1, operator, value2, result)
}

func readInputValue(id string) float64 {
	fmt.Print(id, ": ")
	str, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		panic(fmt.Sprint(id, "must be a number"))
	}
	return value
}

func readOperator() string {
	fmt.Print("Operator (+,-,*,/): ")
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSpace(str)
}

type Calculator struct {
	Value1   float64
	Value2   float64
	Operator string
}

func (c Calculator) getResult() float64 {
	var value float64
	switch strings.TrimSpace(c.Operator) {
	case "+":
		value = c.Value1 + c.Value2
	case "-":
		value = c.Value1 - c.Value2
	case "*":
		value = c.Value1 * c.Value2
	case "/":
		value = c.Value1 / c.Value2
	default:
		panic("invalid operation")
	}
	return math.Round(value*100) / 100
}
