package main

import (
	"fmt"
	"net/http"
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

	// questo è un approccio seriale, nel quale ogni nuovo check
	// inizia solo dopo che il precedente è terminato
	for _, link := range links {
		// con questa istruzione la funzione viene invocata dalla go routine principale
		// checkLink(link)

		// aggiungengo la keyword go, chiediamo che la funzione venga
		// eseguita in un nuovo thread (CHILD ROUTINE)
		go checkLink(link, c)
	}

	// in such a case we need a channel to make the main routine
	// aware of the execution status of the child routines elsewhere
	// the program exit without wait for child termination
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c) // blocking code: finché non arriva un messaggio sul canale
	}
}

func checkLink(link string, c chan string) {
	// la go routine rimane in attesa della risposta del GET
	// bloccando l'esecuzione del programma
	_, err := http.Get(link)
	if err != nil {
		c <- fmt.Sprintln(link, "might be down!")
	} else {
		c <- fmt.Sprintln(link, "is up!")
	}
}
