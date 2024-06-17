package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i <= 3; i++ {
		go worker(ctx, i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Stopping workers")
	cancel()
	time.Sleep(1 * time.Second) // Give some time for workers to stop
}

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopping\n", id)
			return
		default:
			fmt.Printf("Worker %d working\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}


