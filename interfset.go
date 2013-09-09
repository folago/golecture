package main

import (
	"fmt"
)
// START OMIT
type Worker interface {
	DoStuff()
}

type Mule struct{}

func (m *Mule) DoStuff() {
	fmt.Println("It's hard work!!")
}

func main() {
	m := Mule{}
	var w Worker
	w = m
	w.DoStuff()
}
// STOP OMIT
