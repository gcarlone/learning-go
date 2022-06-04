package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func (logWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	fmt.Println(len(p), "bytes written.")
	return len(p), nil
}

func main() {
	resp, err := http.Get("http://www.google.it")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//PrintResponseBody(resp)

	io.Copy(logWriter{}, resp.Body)
}

// nella realtà viene fatto in modo più conciso
// ed elegante con la helper function
// io.Copy(os.Stdout, resp.Body)
func PrintResponseBody(resp *http.Response) {
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
