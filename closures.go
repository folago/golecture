package main

import "fmt"

func main() {
	x := 5
	f := func() int {
		x++
		return x
	}
	g := func(y int) {
		x = y
	}
	fmt.Println("x:", x)
	f()
	fmt.Println("x after f():", x)
	g(42)
	fmt.Println("x after g(42):", x)
}
