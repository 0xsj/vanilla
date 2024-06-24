package main

import (
	"fmt"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ping")
}

func main() {
	http.HandleFunc("/", getHandler)
	fmt.Println("starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
