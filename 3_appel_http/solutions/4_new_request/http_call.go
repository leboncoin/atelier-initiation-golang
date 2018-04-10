package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func httpCall(url string, data string) string {
	request, _ := http.NewRequest("POST", url, strings.NewReader(data))
	res, _ := http.DefaultClient.Do(request)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

func main() {
	data := httpCall("http://localhost:8080/bye", "Chris")
	fmt.Print(data)
}
