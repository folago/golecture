package main

import (
    "fmt"
)

func main() {
    switch x := 10 ; x {
    case 3:
        fmt.Println("3")
    case 4:
        fmt.Println("4")
	case 7,8,9:
        fmt.Println("[7,9]")
    default:
        // something else
        fmt.Println("some number")
    }

}

