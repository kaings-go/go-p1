package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	// defer wg.Done()

	fmt.Println(s)

	defer wg.Done()
}

func main() {
	wg := sync.WaitGroup{}

	words := []string{
		"first",
		"second",
		"third",
		"fourth",
		"five",
	}

	wg.Add(len(words))

	for _, word := range words {
		go printSomething(word, &wg)
	}

	wg.Wait()

	wg.Add(1)

	printSomething("End of the program", &wg)
}
