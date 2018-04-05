package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func getNumberOfViews(adID string) int {
	res, err := http.Get(fmt.Sprintf("http://localhost:9091/views/%s", adID))
	if err != nil {
		log.Fatalf("error calling the server: %s", err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	nb, err := strconv.Atoi(string(body))
	if err != nil {
		log.Fatalf("failed to convert to an int: %s", err)
	}
	return nb
}

func main() {
	id := "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed"
	nb := getNumberOfViews(id)
	fmt.Printf("ad %s had %d views\n", id, nb)
}
