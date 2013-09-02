package main

import "fmt"

func main() {
	var i interface{} // HL
	x := 127
	i = x
	fmt.Println("i:", i)
	fmt.Println("x:", x)
}
