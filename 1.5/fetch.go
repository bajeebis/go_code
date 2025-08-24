// Fetch prints the content found at a URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	fetchExercise18()
}

func fetch1() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", string(b))
	}
}

func fetchExercise17() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
		// b, err := io.Copy(dst io.Writer, src io.Reader)
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}
		fmt.Printf("%d", b)
	}
}

func fetchExercise18() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			// When reassigning, you can use short hand notation
			// https://stackoverflow.com/questions/71303279/go-believes-variables-inside-conditional-statements-are-unused
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}
		fmt.Printf("%d\n", b)
		fmt.Printf("%d", resp.StatusCode)
	}
}
