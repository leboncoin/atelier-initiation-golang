package main

import (
	"fmt"
	"net/http"
)

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong\n")
}

func main() {
	http.HandleFunc("/ping", pong)
	http.ListenAndServe(":8080", nil)
}
