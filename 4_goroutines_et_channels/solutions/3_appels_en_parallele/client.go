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
	// let's create a large number of arbitrary ids
	var ids []string
	for i := 0; i < 1000; i++ {
		ids = append(ids, strconv.Itoa(i))
	}

	startTime := time.Now()
	var count int
	for _, id := range ids {
		go func(adID string) {
			nb := getNumberOfViews(adID)
			count += nb
			fmt.Printf("ad %s had %d views\n", adID, nb)
		}(id)
	}

	time.Sleep(5 * time.Second)

	fmt.Printf("total number of views: %d, took: %v\n", count, time.Since(startTime).Truncate(1*time.Second))
}
