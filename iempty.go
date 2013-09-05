package main

import "fmt"

func main() {
	var i interface{} // HL
	x := 127
	i = x
	fmt.Println("i:", i)
	v := struct{ π, e float32 }{3.14159, 2.71828}
	i = v
	fmt.Println("i:", i)
}
