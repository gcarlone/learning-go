package main

import (
	"fmt"
	"net/http"
	"time"
)

// Una GO ROUTINE può essere pensata come l'ENGINE CHE ESEGUE IL CODICE
// (istruzione per istruzione) ogni go run esegue almeno una go routine
// (MAIN ROUTINE).
// GO SCHEDULER controlla l'esecuzione sulla singola CPU CORE delle
// diverse go routine in esecuzione. GO di default usa una sola CPU.
// Quindi lo scheduler non permette realmente di parallelizzare l'esecuzione,
// c'è sempre una sola go routine in esecuzione al momento (CONCORRENZA).
// Si può forzare GO in modo da usare più CPU Core, ricorda che
// CONCURRENCY IS NOT PARALLELISM
// CONCURRENCY: si hanno più thread in esecuzione contemporanea, ma
// per attivarne uno, deve essere stoppato il corrente, quindi di fatto
// c'è un solo reale processo in esecuzione al momento (1 CPU Core)
// PARALLELISM è possibile solo con esecuzione "contemporanea" di molteplici
// thread (+ CPU Core)
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.com",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	// CHANNELS are the olny way to communicate between go routines
	// a CHANNEL has a typed content
	c := make(chan string)

	// Repeating Routines
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		// FUNCTION LITERAL è una funzione anonima, l'equivalente di una lambda
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l) // voglio che l venga passato alla function come valore
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
