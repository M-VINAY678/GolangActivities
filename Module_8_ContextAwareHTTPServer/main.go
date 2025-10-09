package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/process", processHandler)
	fmt.Println("Server is listening on port 8085...")
	if err := http.ListenAndServe(":8085", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	done := make(chan string)
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
	case result := <-done:
		fmt.Println(result)
		fmt.Fprintf(w, result)
	}
}
