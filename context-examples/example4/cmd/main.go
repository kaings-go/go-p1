package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set a timeout of 5 seconds on the context.
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		fmt.Println("Handler started")
		defer fmt.Println("Handler ended")

		select {
		case <-time.After(10 * time.Second): // Case for long processing time
			// This branch is executed if 10 seconds pass without context cancellation.
			fmt.Fprintln(w, "Hello, World!")
		case <-ctx.Done(): // Case for context cancellation
			// This branch is executed if the context is cancelled before 10 seconds.
			err := ctx.Err()
			fmt.Println("Handler:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", nil)
}
