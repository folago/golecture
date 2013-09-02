package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

var m = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5}

func main() {
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }

	for k, v := range m{
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}


}
