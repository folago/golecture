// START1 OMIT
package main

import (
	"errors"
	"fmt"
)
//A is exported and visible form outside this package
var A int
var (
	B       string = "hello"
	x, y, z float32
	//p is not exported and not visible outside this package
	p       *int
)

// STOP1 OMIT
// STARTF OMIT
const (
	C int = iota //0
	D            //1
	E            //2
	Big = 1 << 100
)


func F(a, b int) (int, float32) {
	return f(a, b), float32(a) / float32(b) // no implicit conversion
}
func f(p, q int) int {
	res := p + q //shorthand variable declaration
	return res
}

// STOPF OMIT

// START2 OMIT
type Q struct { //exported type
	a, b, c int //not exported fields
	D       int //exported field
	buf     [256]byte
}

func (q Q) SumABC() int {
	return q.a + q.b + q.c
}

func (q Q) A() int {
	return q.a
}

func (q *Q) SetFlag(offset int, flag byte) (err error) {
	if offset > 255 || offset < 0 {
		err = errors.New("wrong offset")
		return
	} else {
		q.buf[uint8(offset)] = flag
		return nil
	}
}

// STOP2 OMIT

/*
// START3 OMIT
type error interface {
        Error() string
}
// STOP3 OMIT
*/

// START4 OMIT
type CoolError struct {
	A int
	B float64
	C complex128
}

func (ce CoolError) Error() string {
	return fmt.Sprintf("Error %d, %f, f%", ce.A, ce.B, ce.C)
}

// STOP4 OMIT
