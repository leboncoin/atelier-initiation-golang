package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func callService(url string) string {
	res, _ := http.Get(url)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

func main() {
	data := callService("http://localhost:8080/ping")
	fmt.Print(data)
}
