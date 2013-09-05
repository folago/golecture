package main

import "errors"
import "fmt"

// START OMIT
//QuotRem calculates the quotient and reminder of the integer division.
//It returns an error in case of division by zero.
func Div(a, b int) (quot, rem int, err error) { // HL
	if b != 0 {
		quot, rem = a/b, a%b // HL

	} else {
		err = errors.New("cannot divide by 0") // HL
	}
	return
}

func main() {
	q, r, err := Div(29, 17)
	fmt.Printf("Quotient: %d Reminder: %d Error: %v\n", q,r,err)
}
//STOP OMIT
