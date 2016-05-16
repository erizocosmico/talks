package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	})

	fmt.Println("Starting server at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
