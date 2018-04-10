package main

import (
	"net/http"
)

func main() {
	http.Get("http://localhost:8080/ping")
}
