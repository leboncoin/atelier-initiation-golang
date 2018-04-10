package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	pathParam := r.URL.Path[7:]
	fmt.Fprintf(w, "Hello, %s!\n", pathParam)
}

func main() {
	http.HandleFunc("/hello/", hello)
	http.ListenAndServe(":8080", nil)
}
