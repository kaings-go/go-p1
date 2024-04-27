package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wq *sync.WaitGroup) {
	wq.Done()

	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 28, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	// msg = "Hello, world!"

	// updateMessage("Hello, universe!")
	// printMessage()

	// updateMessage("Hello, cosmos!")
	// printMessage()

	// updateMessage("Hello, world!")

	// printMessage()

	msgs := []string{"Hello, universe!", "Hello, cosmos!", "Hello, world!"}

	var wq sync.WaitGroup
	// wg.Add(len(msgs))

	for _, m := range msgs {
		wq.Add(1)
		go updateMessage(m, &wq)
		wq.Wait()
		printMessage()
	}

}
