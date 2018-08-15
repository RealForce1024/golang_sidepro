package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 22}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Printf("%v, %T\n", a, a)
	b := a[0:5]
	b[0] = 11
	fmt.Printf("%v, %T\n", b, b)
	c :=a[:]
	fmt.Printf("%v, %T", c, c)
}
