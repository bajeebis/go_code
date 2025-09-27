package main

import "fmt"

func main() {
	// sepOne()
	// sepTwo()
	x := 1
	v := 1
	incr(&v)
	incr(x)
	j := g()
	fmt.Println("%v", j)
	fmt.Println(incr(&v))
}

func g() int {
	v := 1
	return &v
}

func incr(p *int) {
	*p++ // increments what p point to; does not change p
	return *p
}

// func sepOne() {
// 	a, b, c := true, 2.3, ""
// 	var s, sep string
// 	for i := 1; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }
//
// func sepTwo() {
// 	s, sep := "", ""
// 	for _, arg := range os.Args[1:] {
// 		s += sep + arg
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }
//
// func dupOne() {
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// func f() *int {
// 	v := 1
// 	return &v
// }
