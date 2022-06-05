package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	r := regexp.MustCompile("[_-].")
	result := r.ReplaceAllStringFunc("the_stealth_warrior", func(s string) string {
		return strings.ToUpper(s[1:])
	})

	fmt.Println(result)

	fmt.Println(toWeirdCase("String"))            // => returns "StRiNg"
	fmt.Println(toWeirdCase("Weird string case")) // => returns "WeIrD StRiNg CaSe"
}

func toWeirdCase(s string) string {
	r := regexp.MustCompile("..? ?")

	return r.ReplaceAllStringFunc(s, func(s string) string {
		return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
	})
}
