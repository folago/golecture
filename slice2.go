package main

import "fmt"

func main() {
	//func make([]T, len, cap) []T
	sl := make([]int, 3, 5)  // HL

    fmt.Println(sl)
	fmt.Println(len(sl), cap(sl))

	nsl := append(sl, 3)
	fmt.Println(nsl)
    fmt.Println(len(nsl), cap(nsl))

}
