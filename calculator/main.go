package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Value 1: ")
	str1, _ := reader.ReadString('\n')
	value1, err1 := strconv.ParseFloat(strings.TrimSpace(str1), 64)
	if err1 != nil {
		panic("Value1 must be a number")
	}

	fmt.Print("Value 2: ")
	str2, _ := reader.ReadString('\n')
	value2, err2 := strconv.ParseFloat(strings.TrimSpace(str2), 64)
	if err2 != nil {
		panic("Value2 must be a number")
	}

	sum := value1 + value2
	sum = math.Round(sum*100) / 100
	fmt.Printf("The sum of %v and %v is %v\n", value1, value2, sum)
}
