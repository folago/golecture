package main

import "fmt"

type Vertex struct {
	X, Y int
}

func (v1 Vertex) Add(v2 Vertex) Vertex {
	return Vertex{v1.X + v2.X, v1.Y + v2.Y}
}

func (v *Vertex) Sum(a Vertex) {
	v.X += a.X
	v.Y += a.Y
}

func main() {
	a := Vertex{1, 3}
	b := Vertex{2, 6}
	fmt.Println(a)
	a.Sum(b)
	fmt.Println(a)
}
