package main

import "fmt"

var b [10]int
var f [3]*float64

func main() {
	primes := [...]int{2, 3, 5, 7, 11, 13, 17}

	fmt.Println("primes", primes)

	fmt.Println("b", b)

	fmt.Println("f", f)

}
