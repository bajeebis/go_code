package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1])
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println(" " + arg)
	}
	// fmt.Println(s)
}
