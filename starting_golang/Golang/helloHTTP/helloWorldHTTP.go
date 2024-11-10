package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		fmt.Fprintf(w, "My name's Joao Victor")
	})

	http.ListenAndServe(":80", nil)
}
