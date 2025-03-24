package main

import (
	"fmt"
	"net/http"
	"simple-api/handlers"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc("/health", handlers.HealthCheckHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
