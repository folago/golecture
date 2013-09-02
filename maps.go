package main

import "fmt"
// START OMIT
var m = map[string]int{"one": 1, "two": 2, "three": 3}

func main() {
	//m := make(map[string]int)

	m["five"] = 5
	fmt.Println(m)

	delete(m, "three")
	fmt.Println(m)

	b, ok := m["four"]
	if !ok {
		fmt.Println("cannot find four")
	}else{
		fmt.Println(b)
	}

	fmt.Println(m["four"])
}
// STOP OMIT
