package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello from the other side")
}
func about(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello from the about function")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server starting...")
	_ = http.ListenAndServe(":3000", nil)
}
