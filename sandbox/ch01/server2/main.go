package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mutator sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mutator.Lock()
	count++
	mutator.Unlock()
	message := fmt.Sprintf("URL.Path = %q\n", r.URL.Path)
	fmt.Print(message)
	fmt.Fprint(w, message)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mutator.Lock()
	message := fmt.Sprintf("Count %d\n", count)
	fmt.Print(message)
	fmt.Fprint(w, message)
	mutator.Unlock()
}
