package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func longRunningHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting long running task")
	time.Sleep(10 * time.Second) // Simulate a long-running task
	log.Println("Finished long running task")
	fmt.Fprintln(w, "Long running task finished")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/long", longRunningHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start the server in a new goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("Server started")

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// select {
	// 	case <- time.After(10*time.Second):
	// 		fmt.Println("after 10s........")
	// }


	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
		os.Exit(1)
	}

	

	log.Println("Server exiting")
}
