package main

import (
	"fmt"
	"sort"
)

func main() {
	// map is an unordered collection with key-value pairs (hashtable)
	states := make(map[string]string)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println("maps:", states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "OR")
	fmt.Println("maps:", states)

	// order is not guaranteed
	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	// per l'ordinamento creiamo uno slice delle chiavi ed ordiniamole
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("keys:", keys)

	for i := range keys {
		fmt.Printf("%v:%v ", keys[i], states[keys[i]])
	}
}
