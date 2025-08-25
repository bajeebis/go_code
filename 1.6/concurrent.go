package main

import (
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	var filename string
	uniqueNo := strconv.Itoa(rand.Int())
	if strings.HasPrefix(url, "https://") {
		filename = strings.TrimPrefix(url, "https://")
	} else if strings.HasPrefix(url, "http://") {
		filename = strings.TrimPrefix(url, "http://")
		// Check strings.cutprefix
	}
	out, fileError := os.Create(filename + uniqueNo + "_body.txt")
	if fileError != nil {
		fmt.Fprintf(os.Stderr, "Error writing to %s: %v\n", filename+uniqueNo+"_body.txt", fileError)
	}
	defer out.Close()
	// io.Copy(out, resp.Body)

	nbytes, err := io.Copy(out, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
