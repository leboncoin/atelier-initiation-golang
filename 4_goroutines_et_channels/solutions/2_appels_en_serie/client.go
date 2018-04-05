package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
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
	ids := []string{"2aae6c35c94fcfb415dbe95f408b9ce91ee846ed", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}

	startTime := time.Now()
	var count int
	for _, id := range ids {
		nb := getNumberOfViews(id)
		count += nb
		fmt.Printf("ad %s had %d views\n", id, nb)
	}

	fmt.Printf("total number of views: %d, took: %v\n", count, time.Since(startTime).Truncate(1*time.Second))
}
