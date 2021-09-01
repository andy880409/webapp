package main

import (
	"fmt"
	"net/http"
)

func DefaultHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "test123")
}
func main() {
	http.HandleFunc("/", DefaultHandle)
	http.ListenAndServe(":8080", nil)
}
