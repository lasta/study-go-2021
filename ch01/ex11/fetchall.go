package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// begin go-routine
		go fetch("http://" + url, ch)
	}

	for range os.Args[1:] {
		// receive from `ch` channel
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		// send to channel
		ch <- fmt.Sprint(err)
		return
	}

	writtenBytes, err := io.Copy(io.Discard, response.Body)
	// avoid to leek resource
	response.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	elapsedTimeSecond := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", elapsedTimeSecond, writtenBytes, url)
}
