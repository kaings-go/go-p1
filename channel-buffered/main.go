package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int, done chan bool) {
	for {
		// print a got data message
		i, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			done <- true
			fmt.Println("Sent signal to done channel")
			return
		}
		fmt.Println("Got", i, "from channel")

		// simulate doing a lot of work
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int, 10)
	done := make(chan bool)

	go listenToChan(ch, done)

	for i := 0; i <= 100; i++ {
		// the first 10 times through this loop, things go quickly; after that, things slow down.
		fmt.Println("sending", i, "to channel...")
		ch <- i
		fmt.Println("sent", i, "to channel!")
	}

	// close ch because ch no longer receiving any data
	fmt.Println("Closing chan ch!")
	close(ch)

	// done channel will be blocking code, it only proceeds when 'done' channel receive value
	// without done chan, the code wont be waiting for all data to be consumed by listenToChan goroutine
	<-done

	fmt.Println("Done!")
	fmt.Println("Closing chan done!")
	close(done)
}
