package main

import "fmt"

var a, b [4]int
var c [3]int

func main() {
	a[0], a[1] = 1, 2
	b = a
	fmt.Printf("a: %v\nb: %v\n\n", a, b)
	a[3] = 3
	fmt.Printf("a: %v\nb: %v\n", a, b)

	// [3]int != [4]int
	//c = a
}
