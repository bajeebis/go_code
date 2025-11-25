package main

import "fmt"

// var global *int

func main() {
	xy := new(int)
	xz := returnMoney()
	fmt.Printf("%v", *xy)
	fmt.Printf("%f", xz)
}

func returnMoney() string {
	x := 5.0
	return x
}

// func delta(old, new int) int {
// 	return new - old
// }
