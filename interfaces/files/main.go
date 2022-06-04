package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 1 {
		fmt.Println("filename, as argument, is needed")
		os.Exit(1)
	}

	filename := args[1]
	fmt.Println("cat", filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
