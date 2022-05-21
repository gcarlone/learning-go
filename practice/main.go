package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello from Go!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter text:")
	input, _ := reader.ReadString('\n')
	fmt.Println("you entered:", input)

	fmt.Println("Hello from Go!")
	fmt.Print("enter numer:")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("you entered:", aFloat)
	}

}
