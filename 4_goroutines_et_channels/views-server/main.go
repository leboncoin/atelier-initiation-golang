package main

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const rootPath = "/views/"
const port = ":9091"

func viewsHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(rootPath):]

	views := computeNumberOfViews(id)
	log.Printf("number of views for ad %s is %d\n", id, views)

	str := strconv.Itoa(views)
	w.Write([]byte(str))
}

func computeNumberOfViews(id string) int {
	// crude simulation of a slow server ¯\_(ツ)_/¯

	h := sha1.New()
	h.Write([]byte(id))
	bs := h.Sum(nil)
	data := binary.BigEndian.Uint64(bs)

	delay := time.Duration(rand.Intn(2000))
	time.Sleep(delay * time.Millisecond)

	views := int(data % 10000)
	return views
}

func main() {
	fmt.Println("this returns the number of views for a given ad id")
	fmt.Println("try:")
	fmt.Println("  curl http://localhost" + port + rootPath + "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
	fmt.Println("")

	http.HandleFunc(rootPath, viewsHandler)

	log.Printf("serving on %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
