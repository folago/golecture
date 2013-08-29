package main

import "fmt"

func AddOne(val int) int {
	return val + 1
}

func AddMore(fn func(int) int) int {
	return fn(41)
}

func main() {
	f := AddOne
	x := 41
	fmt.Println("The answer is", f(x))
	fmt.Println("The answer is", AddMore(f))

}

