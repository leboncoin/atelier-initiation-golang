package main

import (
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/test/", test)
	http.ListenAndServe(":8080", nil)
}
