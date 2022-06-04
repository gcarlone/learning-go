package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.it")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	content := make([]byte, 99999)
	size, err := resp.Body.Read(content)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Readed", size, "bytes")
	fmt.Println(string(content))
}
