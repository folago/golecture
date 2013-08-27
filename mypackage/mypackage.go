/*
// STARTBASIC OMIT
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
// STOPBASIC OMIT
*/

// STARTIMPORT OMIT
package mypackage

import (
	"errors"
	"fmt"

//	"github.com/golang/glog"
)

//A is exported and visible form outside this package
var A int

//b is not exported and not visible outside this package
var b float32

// STOPIMPORT OMIT

//STARTVAR OMIT
var (
	B       string = "hello"
	x, y, z float32
	p       *int
)

const (
	C   = iota //0
	D          //1
	E          //2
	Big = 1 << 100
)

// STOPVAR OMIT
/*
// STARTIOTA OMIT

type ByteSize float64

const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

// STOPIOTA OMIT
*/

// STARTFUNC OMIT

func APIFun(a int, b int) (int, float32) { // HL
	if b != 0 {
		return privateFun(a, b), float32(a) / float32(b) // no implicit conversion
	} else {
		return 0, 0
	}
}
func privateFun(p, q int) int { // HL
	res := p + q //shorthand variable declaration
	return res
}

// STOPFUNC OMIT

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
