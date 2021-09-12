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
	for index, url := range os.Args[1:] {
		// begin go-routine
		fileName := fmt.Sprintf("%v_%d.html", start.Unix(), index)
		go fetch(url, fileName, ch)
	}

	for range os.Args[1:] {
		// receive from `ch` channel
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, destinationFileName string, ch chan<- string) {
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		// send to channel
		ch <- fmt.Sprint(err)
		return
	}

	destinationFile, err := os.Create(destinationFileName)
	if err != nil {
		ch <- fmt.Sprintf("failed to create file %s, cause %v", destinationFileName, err)
		return
	}
	defer destinationFile.Close()

	writtenBytes, err := io.Copy(destinationFile, response.Body)
	// avoid to leek resource
	response.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	elapsedTimeSecond := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", elapsedTimeSecond, writtenBytes, url)
}
