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

	printResponseBody(resp)
}

// nella realtà viene fatto in modo più conciso
// ed elegante con la helper function
// io.Copy(os.Stdout, resp.Body)
func printResponseBody(resp *http.Response) {
	content := make([]byte, 99999)
	size, err := resp.Body.Read(content)
	if err != nil {
		// come da API EOF è l'errore che viene in ogni caso generato
		// al termine del content body: non è una vera condizione di
		// errore o bloccante
		fmt.Println("Error:", err)
	}
	fmt.Println("Readed", size, "bytes")
	fmt.Println(string(content))
}
