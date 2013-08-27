// START OMIT
package main // HLpackage

import "fmt" // HLimport

func main() { // HLmain
	fmt.Println("Hello, 世界") // HLfunc
}

// STOP OMIT

// STARTF OMIT
func test(a int, b int)(int, int){
	return a + b, a - b
}

// STOPF OMIT
