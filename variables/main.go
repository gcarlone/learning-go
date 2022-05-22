package main

import (
	"fmt"
)

// a costant can be declared only outside a function
const aConst int = 64

func main() {
	// statically typed variables
	var aString string = "My string!"
	fmt.Printf("The variable's type is %T and value is %v\n", aString, aString)

	var aInt int64 = 32
	fmt.Printf("The variable's type is %T and value is %v\n", aInt, aInt)

	// every type has a default value
	var defaultInt int64
	fmt.Println("Default int value is ", defaultInt)

	var defaultString string
	fmt.Println("Default int String is ", defaultString, " ( length ", len(defaultString), ")")

	// for implicit types, the variable types as inferred based on initial value
	var anotherString = aString
	fmt.Printf("The anotherString variable's type is %T\n", anotherString)

	// := Operator (without var keyword) --> only works inside functions
	myString := aString
	fmt.Printf("The variable's type is %T and value is %v\n", myString, myString)

	fmt.Printf("The constant's type is %T and value is %v\n", aConst, aConst)
}
