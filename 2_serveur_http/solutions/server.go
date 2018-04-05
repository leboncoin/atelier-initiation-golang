package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s\n", r.URL.Path)
}

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong\n")
}

func hello(w http.ResponseWriter, r *http.Request) {
	pathParam := r.URL.Path[7:]
	fmt.Fprintf(w, "Hello, %s!\n", pathParam)
}

func bye(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, _ := ioutil.ReadAll(r.Body)

	fmt.Fprintf(w, "Bye, %s!\n", data)
}

func main() {
	http.HandleFunc("/test/", test)
	http.HandleFunc("/ping", pong)
	http.HandleFunc("/hello/", hello)
	http.HandleFunc("/bye", bye)
	http.ListenAndServe(":8080", nil)
}
