package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Register an HTTP handler for the "/process" route
	http.HandleFunc("/process", processHandler)
	// Start the HTTP server on port 8085
	fmt.Println("Server is listening on port 8085...")
	if err := http.ListenAndServe(":8085", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

// processHandler handles requests to the "/process" endpoint
func processHandler(w http.ResponseWriter, r *http.Request) {
	// Create a context with a 5-second timeout
	// This will automatically cancel the context after 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	// Channel to signal when the goroutine has finished work
	done := make(chan string)
	// Start a goroutine to simulate a long-running task
	go func(ctx context.Context) {
		for i := 1; i <= 10; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine stoppped", ctx.Err())
				return
			default:
				fmt.Println("Working... step", i)
				fmt.Fprintf(w, "Working... step %v\n", i)
				time.Sleep(1 * time.Second)
			}
		}
		done <- "work completed successfully"
	}(ctx)

	// Wait for either the work to complete or the context to be canceled/timeout
	select {
	case <-ctx.Done():
		err := ctx.Err()
		switch err {
		case context.DeadlineExceeded:
			fmt.Println("Handler timeout:", err)
		case context.Canceled:
			fmt.Println("Handler canceled by client:", err)
		default:
			fmt.Println("Unknown context error :", err)
		}
	case result := <-done: // Work completed successfully
		fmt.Println(result)
		fmt.Fprintf(w, result) // Send final message to client
	}
}
