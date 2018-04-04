package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func bye(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	fmt.Fprintf(w, "Bye, %s!\n", data)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[7:])
}

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong\n")
}

func main() {
	http.HandleFunc("/ping", pong)
	http.HandleFunc("/hello/", hello)
	http.HandleFunc("/bye", bye)
	http.ListenAndServe(":8080", nil)
}
