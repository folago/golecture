package main

import "fmt"

func main(){
	switch s := "yes"; s{
	case "yes":
		fmt.Println("yes")
	case "no", "noway":
		fmt.Println("no")
	case "maybe":
		fmt.Println("maybe")
	default:
		fmt.Println("whatever")
	}
}
