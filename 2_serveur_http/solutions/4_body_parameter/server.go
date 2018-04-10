package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func bye(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "Bye, %s!\n", data)
}

func main() {
	http.HandleFunc("/bye", bye)
	http.ListenAndServe(":8080", nil)
}
