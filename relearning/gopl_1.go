package main

import "fmt"

func main() {
	vars()
}

func vars() {
	x := 1
	p := &x
	fmt.Println(*p)
}
