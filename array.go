package main

import "fmt"

var b [10]int
var f [3]*float64

func main() {
	primes := [...]int{2, 3, 5, 7, 11, 13, 17}
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(primes)

}
