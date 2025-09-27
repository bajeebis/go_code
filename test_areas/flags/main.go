package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	n   = flag.Bool("n", false, "Omit trailing newline")
	sep = flag.String("s", " ", "separator")
)

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	fmt.Printf("%v\n", *sep)
}
