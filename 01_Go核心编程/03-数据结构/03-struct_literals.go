package main

import (
	"fmt"
)

type Vertex struct {
	x, y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{x: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, v2, v3, p, *p)

}
