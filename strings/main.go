package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	aString := "45"
	aInt, err := strconv.ParseInt(strings.TrimSpace(aString), 10, 64)
	if err == nil {
		fmt.Printf("My parsed variable's type is %T and value is %v)\n", aInt, aInt)
	} else {
		fmt.Println("not parsable string")
	}
}
