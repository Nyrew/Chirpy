package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	fmt.Println("Server running on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
